package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateStockHandler(c *gin.Context) {
	request := CreateStockRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// FIND Product
	product := schemas.Product{}
	if err := db.First(&product, request.ProductID).Error; err != nil {
		logger.ErrorF("Product of id: %d not found", request.ProductID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O Produto de id %d não foi encontrado", request.ProductID))
		return
	}

	// FIND Supplier
	supplier := schemas.Supplier{}
	if err := db.First(&supplier, request.SupplierID).Error; err != nil {
		logger.ErrorF("Product of id: %d not found", request.SupplierID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O Fornecedor de id %d não foi encontrado", request.SupplierID))
		return
	}

	stock := schemas.Stock{
		Quantity:   request.Quantity,
		PriceOfBuy: request.PriceOfBuy,
		ProductID:  product.ID,
		SupplierID: supplier.ID,
	}

	if err := db.Create(&stock).Error; err != nil {
		logger.Error("Error inserting the Product on stock", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao atualizar seu estoque. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, stock)

}
