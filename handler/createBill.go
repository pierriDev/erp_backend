package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateBillHandler(c *gin.Context) {
	request := CreateBillRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	day := request.BillingDay

	currentTime := time.Now()

	year := currentTime.Year()
	month := currentTime.Month()

	billingDueDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	bill := schemas.Bill{
		Title:          request.Title,
		Value:          request.Value,
		Description:    request.Description,
		BillingDueDate: billingDueDate,
		IsPaid:         *request.IsPaid,
	}

	if err := db.Create(&bill).Error; err != nil {
		logger.ErrorF("Error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao criar a sua Conta")
		return
	}

	sendSuccess(c, bill)
}
