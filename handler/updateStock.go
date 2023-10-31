package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateStockHandler(c *gin.Context) {
	request := UpdateStockRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	stock := schemas.Stock{}
	if err := db.First(&stock, id).Error; err != nil {
		logger.ErrorF("Stock of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro ao carregar o estoque do seu produto. Tente novamente mais tarde"))
		return
	}

	stock.Quantity = request.Quantity

	if err := db.Save(&stock).Error; err != nil {

		logger.ErrorF("error updating stock: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao atualizar seu estoque. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, stock)
}
