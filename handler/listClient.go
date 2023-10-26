package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListClientHandler(c *gin.Context) {
	clients := []schemas.Client{}

	if err := db.Preload("User.Address").Find(&clients).Error; err != nil {
		logger.ErrorF("Error listing Clients: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar os clientes")
		return
	}

	sendSuccess(c, clients)
}
