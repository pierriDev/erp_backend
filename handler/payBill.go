package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func PayBillHandler(c *gin.Context) {
	request := PayBillRequest{}
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

	bill := schemas.Bill{}

	if err := db.First(&bill, id).Error; err != nil {
		logger.ErrorF("Bill of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A conta de id %s não foi encontrada", id))
		return
	}

	if request.IsPaid != nil {
		bill.IsPaid = *request.IsPaid
	}

	if err := db.Save(&bill).Error; err != nil {
		logger.ErrorF("error updating bill: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao atualizar a conta")
		return
	}

	sendSuccess(c, bill)
}
