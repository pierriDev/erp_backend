package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
	"golang.org/x/crypto/bcrypt"
)

func generateToken(name string) string {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(name), bcrypt.DefaultCost)

	if hashErr != nil {
		logger.ErrorF("Error generating unique code")
		return "nil"
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateProductHandler(c *gin.Context) {
	request := CreateProductRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	category := schemas.Category{}
	// FIND CATEGORY
	if err := db.First(&category, request.CategoryID).Error; err != nil {
		logger.ErrorF("Category of id: %d not found", request.CategoryID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id %d não foi encontrada", request.CategoryID))
		return
	}

	supplier := schemas.Supplier{}
	// FIND SUPPLIER
	if err := db.First(&supplier, request.SupplierID).Error; err != nil {
		logger.ErrorF("Supplier of id: %d not found", request.SupplierID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O fornecedor de id %d não foi encontrada", request.SupplierID))
		return
	}

	generatedCode := generateToken(request.Title)

	logger.InfoF("generatedCode: %s", generatedCode)

	if generatedCode == "nil" {
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro. Tente novamente mais tarde"))
	}

	product := schemas.Product{
		Title:       request.Title,
		Price:       request.Price,
		Code:        generateToken(request.Title),
		Description: request.Description,
		CategoryID:  category.ID,
		Category:    category,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.ErrorF("Error creating Product: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar o Produto. Tente novamente mais tarde")
		return
	}

	productSupplier := schemas.ProductSupplier{
		BuyPrice: request.BuyPrice,

		ProductID:  product.ID,
		Product:    product,
		SupplierID: supplier.ID,
		Supplier:   supplier,
	}
	if err := db.Create(&productSupplier).Error; err != nil {
		logger.ErrorF("Error creating the relation of Product and Supplier: %v", err.Error())

		if err := db.Delete(&product).Error; err != nil {
			logger.ErrorF("Error deleting the product with id: %d", product.ID)
		}

		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar o Produto. Tente novamente  mais tarde")
		return
	}
	sendSuccess(c, product)
}
