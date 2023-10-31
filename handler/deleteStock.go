package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteStockHandler(c *gin.Context) {
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

	//Delete Product
	if err := db.Delete(&stock).Error; err != nil {
		logger.ErrorF("Error deleting the stock with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, stock)
}
