package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func GetStockHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	stock := schemas.Stock{}
	// FIND PRODUCT
	if err := db.Preload("Product").Preload("Supplier").First(&stock, id).Error; err != nil {
		logger.ErrorF("Stock of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro ao carregar o estoque do seu produto. Tente novamente mais tarde"))
		return
	}

	sendSuccess(c, stock)
}
