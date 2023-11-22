package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

type promotionReturnType struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	DateStart time.Time `json:"dateStart"`
	DateEnd   time.Time `json:"dateEnd"`
	IsActive  bool      `json:"isActive"`

	PromotionProducts []schemas.PromotionProduct
}

// func convertProductReturn(sellProducts []schemas.SellProduct) []schemas.Product {
// 	products := make([]schemas.Product, len(sellProducts))
// 	for i, sp := range sellProducts {
// 		// Assuming there's a conversion function from SellProduct to Product, e.g., ToProduct
// 		products[i] = sp.Product
// 	}
// 	return products
// }

func GetPromotionHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	promotion := schemas.Promotion{}
	if err := db.First(&promotion, id).Error; err != nil {
		logger.ErrorF("Promotion of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A Promoção de id %s não foi encontrado", id))
		return
	}

	promotionProducts := []schemas.PromotionProduct{}
	if err := db.Where("promotion_id = ?", promotion.ID).Preload("Product").Find(&promotionProducts).Error; err != nil {
		logger.ErrorF("Products of promotion id %d not found", promotion.ID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro ao carregar os produtos da sua Promoção"))
		return
	}

	promotionReturn := promotionReturnType{
		ID:        promotion.ID,
		CreatedAt: promotion.CreatedAt,
		UpdatedAt: promotion.UpdatedAt,
		Title:     promotion.Title,
		DateStart: promotion.DateStart,
		DateEnd:   promotion.DateEnd,
		IsActive:  promotion.IsActive,

		PromotionProducts: promotionProducts,
	}

	sendSuccess(c, promotionReturn)
}
