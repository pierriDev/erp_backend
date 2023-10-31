package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdatePaymentMethodHandler(c *gin.Context) {
	request := UpdatePaymentMethodRequest{}
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

	paymentMethod := schemas.PaymentMethod{}

	// }
	if err := db.First(&paymentMethod, id).Error; err != nil {
		logger.ErrorF("Payment Method of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Metodo de Pagamento de id: %v não foi encontrada", id))
		return
	}

	// Update Category
	if request.Title != "" {
		paymentMethod.Title = request.Title
		paymentMethod.Tax = request.Tax
	}

	if err := db.Save(&paymentMethod).Error; err != nil {

		logger.ErrorF("error updating payment method: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, paymentMethod)
}
