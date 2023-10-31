package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreatePaymentMethodHandler(c *gin.Context) {
	request := CreatePaymentMethodRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	paymentMethod := schemas.PaymentMethod{
		Title: request.Title,
		Tax:   request.Tax,
	}

	if err := db.Create(&paymentMethod).Error; err != nil {
		logger.ErrorF("Error creating Payment Method: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao cadastrar o Metodo de Pagamento. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, paymentMethod)
}
