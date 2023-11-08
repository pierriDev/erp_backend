package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteBillHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	bill := schemas.Bill{}
	// FIND OPENING
	if err := db.First(&bill, id).Error; err != nil {
		logger.ErrorF("Bill of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A opening de id %s não foi encontrada", id))
		return
	}

	//Delete Opening
	if err := db.Delete(&bill).Error; err != nil {
		logger.ErrorF("Error deleting the bill with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, bill)
}
