package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListPromotionHandler(c *gin.Context) {
	promotions := []schemas.Promotion{}

	if err := db.Find(&promotions).Error; err != nil {
		logger.ErrorF("Error listing Promotions: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar as Promoções")
		return
	}

	sendSuccess(c, promotions)
}
