package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeletePaymentMethodHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	paymentMethod := schemas.PaymentMethod{}
	if err := db.First(&paymentMethod, id).Error; err != nil {
		logger.ErrorF("Payment Method of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O metodo de pagamento de id: %v não foi encontrada", id))
		return
	}

	if err := db.Delete(&paymentMethod).Error; err != nil {
		logger.ErrorF("Error deleting the payment method with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, paymentMethod)
}
