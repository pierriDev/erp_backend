package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateBillHandler(c *gin.Context) {
	request := UpdateBillRequest{}
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

	if request.Title != "" {
		bill.Title = request.Title
	}

	if request.Value > 0 {
		bill.Value = request.Value
	}

	if request.Description != "" {
		bill.Description = request.Description
	}

	if request.BillingDay > 0 {
		day := request.BillingDay

		currentTime := time.Now()

		year := currentTime.Year()
		month := currentTime.Month()

		billingDueDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		bill.BillingDueDate = billingDueDate

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
