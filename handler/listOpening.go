package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.ErrorF("Error listing openings: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar as openings")
	}

	sendSuccess(c, openings)
}
