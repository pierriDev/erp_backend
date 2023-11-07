package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

type sellReturnType struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	DeletedAt       time.Time `json:"deletedAt,omitempty"`
	TotalValue      float32   `json:"totalValue"`
	LiquidValue     float32   `json:"liquidValue"`
	PaymentMethodID int       `json:"paymentMethodId"`
	ClientID        int       `json:"clientId"`
	EmployeeID      int       `json:"employeeId"`

	SellProduct []schemas.SellProduct
}

func convertProductReturn(sellProducts []schemas.SellProduct) []schemas.Product {
	products := make([]schemas.Product, len(sellProducts))
	for i, sp := range sellProducts {
		// Assuming there's a conversion function from SellProduct to Product, e.g., ToProduct
		products[i] = sp.Product
	}
	return products
}

func GetSellHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	sell := schemas.Sell{}
	if err := db.Preload("PaymentMethod").Preload("Employee.User.Address").Preload("Client.User.Address").First(&sell, id).Error; err != nil {
		logger.ErrorF("Sell of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A venda de id %s não foi encontrado", id))
		return
	}

	sellProducts := []schemas.SellProduct{}
	if err := db.Where("sell_id = ?", sell.ID).Preload("Product").Find(&sellProducts).Error; err != nil {
		logger.ErrorF("Sell Products of sell id %d not found", sell.ID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Ocorreu um erro ao carregar os produtos da sua venda"))
		return
	}

	sellReturn := sellReturnType{
		ID:              sell.ID,
		CreatedAt:       sell.CreatedAt,
		UpdatedAt:       sell.UpdatedAt,
		TotalValue:      sell.TotalValue,
		LiquidValue:     sell.LiquidValue,
		PaymentMethodID: sell.PaymentMethodID,
		ClientID:        sell.ClientID,
		EmployeeID:      sell.EmployeeID,

		SellProduct: sellProducts,
	}

	sendSuccess(c, sellReturn)
}
