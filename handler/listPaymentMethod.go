package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListPaymentMethodHandler(c *gin.Context) {
	paymentMethods := []schemas.PaymentMethod{}

	if err := db.Find(&paymentMethods).Error; err != nil {
		logger.ErrorF("Error listing Payment Methods: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar os Metodos de Pagamentos")
		return
	}

	sendSuccess(c, paymentMethods)

}
