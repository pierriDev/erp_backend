package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListStockHandler(c *gin.Context) {
	stocks := []schemas.Stock{}

	if err := db.Preload("Product").Preload("Supplier").Find(&stocks).Error; err != nil {
		logger.ErrorF("Error listing Stocks: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar o seu estoque")
		return
	}

	sendSuccess(c, stocks)

}
