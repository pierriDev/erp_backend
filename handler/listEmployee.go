package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListEmployeeHandler(c *gin.Context) {
	employees := []schemas.Employee{}

	if err := db.Preload("User.Address").Find(&employees).Error; err != nil {
		logger.ErrorF("Error listing employees: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar os funcionários")
		return
	}

	sendSuccess(c, employees)
}
