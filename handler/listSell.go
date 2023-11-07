package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListSellHandler(c *gin.Context) {
	sells := []schemas.Sell{}

	if err := db.Preload("PaymentMethod").Preload("Employee.User.Address").Preload("Client.User.Address").Find(&sells).Error; err != nil {
		logger.ErrorF("Error listing Sells: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar historico de vendas")
		return
	}

	sendSuccess(c, sells)
}
