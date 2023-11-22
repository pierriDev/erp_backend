package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreatePromotionHandler(c *gin.Context) {
	request := CreatePromotionRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dateStart, dateStartErr := time.Parse("2006-01-02", request.DateStart)
	dateEnd, dateEndErr := time.Parse("2006-01-02", request.DateEnd)

	if dateStartErr != nil {
		logger.ErrorF("Error creating promotion: %v", dateStartErr.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao criar a Promoção")
		return
	}

	if dateEndErr != nil {
		logger.ErrorF("Error creating promotion: %v", dateStartErr.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao criar a Promoção")
		return
	}

	promotion := schemas.Promotion{
		Title:     request.Title,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		IsActive:  *request.IsActive,
	}

	if err := db.Create(&promotion).Error; err != nil {
		logger.ErrorF("Error creating promotion: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao criar a Promoção")
		return
	}

	arrayOfProduct := request.Products

	for index, element := range arrayOfProduct {
		product := schemas.Product{}
		if err := db.First(&product, element.ProductID).Error; err != nil {
			logger.ErrorF("Product of id: %d not found", element.ProductID)
			sendError(c, http.StatusNotFound, fmt.Sprintf("O produto de id %d não foi encontrado", element.ProductID))

			if err := db.Delete(&promotion).Error; err != nil {
				logger.ErrorF("Error deleting the Promotion")
				sendError(c, http.StatusInternalServerError, "Erro ao deletar a Promoção")
			}

			return
		}

		promotionProduct := schemas.PromotionProduct{
			ProductID:   product.ID,
			PromotionID: promotion.ID,
			Value:       element.Value,
		}

		if err := db.Create(&promotionProduct).Error; err != nil {
			logger.ErrorF("Error relating Promotion and Product: %v", err.Error())
			sendError(c, http.StatusInternalServerError, "Erro ao criar relação entre Promoção e Produto")

			if err := db.Delete(&promotion).Error; err != nil {
				logger.ErrorF("Error deleting the Promotion")
				sendError(c, http.StatusInternalServerError, "Erro ao deletar a Promoção")
			}
			return
		}
		logger.InfoF("Index: %v", index)
	}

	sendSuccess(c, promotion)
}
