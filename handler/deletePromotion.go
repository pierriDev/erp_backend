package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeletePromotionHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	promotion := schemas.Promotion{}
	if err := db.First(&promotion, id).Error; err != nil {
		logger.ErrorF("Promotion of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro. Tente novamente mais tarde"))
		return
	}

	if err := db.Delete(&promotion).Error; err != nil {
		logger.ErrorF("Error deleting the promotion of id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao deletar a Promoção. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, promotion)
}
