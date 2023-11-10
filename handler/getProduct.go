package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

type productReturnType struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	BuyPrice    float32 `json:"buyprice"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	CategoryID  int     `json:"categoryId"`

	Supplier schemas.Supplier
}

func GetProductHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	product := schemas.Product{}
	// FIND PRODUCT
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		logger.ErrorF("Product of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O produto de id %s não foi encontrada", id))
		return
	}

	productSupplier := schemas.ProductSupplier{}

	if err := db.Preload("Supplier.User.Address").First(&productSupplier, product.ID).Error; err != nil {
		logger.ErrorF("Relation between product and supplier of product with id: %d not found", product.ID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro. Tente novamente mais tarde"))
		return
	}

	productReturn := productReturnType{
		ID:          product.ID,
		Title:       product.Title,
		Price:       product.Price,
		BuyPrice:    productSupplier.BuyPrice,
		Code:        product.Code,
		Description: product.Description,
		CategoryID:  product.CategoryID,

		Supplier: productSupplier.Supplier,
	}

	sendSuccess(c, productReturn)
}
