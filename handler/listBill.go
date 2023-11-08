package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListBillHandler(c *gin.Context) {
	bills := []schemas.Bill{}

	if err := db.Find(&bills).Error; err != nil {
		logger.ErrorF("Error listing bills: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar as contas")
	}

	sendSuccess(c, bills)
}
