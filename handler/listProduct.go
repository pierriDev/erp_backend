package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListProductHandler(c *gin.Context) {
	products := []schemas.Product{}

	if err := db.Preload("Category").Find(&products).Error; err != nil {
		logger.ErrorF("Error listing Products: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar os Produtos")
		return
	}

	sendSuccess(c, products)

}
