package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListCategoryHandler(c *gin.Context) {
	categories := []schemas.Category{}

	if err := db.Find(&categories).Error; err != nil {
		logger.ErrorF("Error listing Categories: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar as categorias")
		return
	}

	sendSuccess(c, categories)

}
