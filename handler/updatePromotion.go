package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdatePromotionHandler(c *gin.Context) {
	request := UpdatePromotionRequest{}
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

	promotion := schemas.Promotion{}

	if err := db.First(&promotion, id).Error; err != nil {
		logger.ErrorF("Promotion of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A Promoção de id %v não foi encontrada", id))
		return
	}

	if request.IsActive != nil {
		promotion.IsActive = *request.IsActive
	}

	if err := db.Save(&promotion).Error; err != nil {
		logger.ErrorF("error updating promotion: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao atualizar a Promoção")
		return
	}

	sendSuccess(c, promotion)
}
