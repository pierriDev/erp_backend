package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

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

	sendSuccess(c, product)
}
