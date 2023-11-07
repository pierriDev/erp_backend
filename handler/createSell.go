package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateSellHandler(c *gin.Context) {
	request := CreateSellRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	employee := schemas.Employee{}
	if err := db.Preload("User.Address").First(&employee, request.EmployeeID).Error; err != nil {
		logger.ErrorF("Employee of id: %d not found", request.EmployeeID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O funcionário de id: %d não foi encontrado", request.EmployeeID))
		return
	}

	client := schemas.Client{}
	if request.ClientID > 0 {
		if err := db.Preload("User.Address").First(&client, request.ClientID).Error; err != nil {
			logger.ErrorF("Client of id: %d not found", request.ClientID)
			sendError(c, http.StatusNotFound, fmt.Sprintf("O cliente de id %d não foi encontrado", request.ClientID))
			return
		}
	}

	paymentMethod := schemas.PaymentMethod{}
	if err := db.First(&paymentMethod, request.PaymentMethodID).Error; err != nil {
		logger.ErrorF("Payment Method of id: %d not found", request.PaymentMethodID)
		sendError(c, http.StatusNotFound, fmt.Sprint("Ocorreu um erro. Tente novamente mais tarde"))
		return
	}

	taxValue := (paymentMethod.Tax / 100) * request.TotalValue

	liquidValue := request.TotalValue - taxValue

	sell := schemas.Sell{
		TotalValue:      request.TotalValue,
		LiquidValue:     liquidValue,
		Status:          request.Status,
		PaymentMethodID: request.PaymentMethodID,
		ClientID:        request.EmployeeID,
		EmployeeID:      request.EmployeeID,
		Client:          client,
		Employee:        employee,
		PaymentMethod:   paymentMethod,
	}

	if err := db.Create(&sell).Error; err != nil {
		logger.ErrorF("Error creating Sell: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente")
		return
	}

	for i := 0; i < len(request.Products); i++ {
		product := schemas.Product{}
		if err := db.Preload("Category").First(&product, request.Products[i].ProductID).Error; err != nil {
			logger.ErrorF("Product of id: %d not found", request.Products[i].ProductID)
			sendError(c, http.StatusNotFound, fmt.Sprintf("O produto de id %d não foi encontrada", request.Products[i].ProductID))

			if err := db.Delete(&sell).Error; err != nil {
				logger.ErrorF("Error deleting the sell with id: %d", sell.ID)
				sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
				return
			}

			return
		}

		sellProduct := schemas.SellProduct{
			Quantity:  request.Products[i].Quantity,
			ProductID: request.Products[i].ProductID,
			SellID:    sell.ID,

			Product: product,
			Sell:    sell,
		}

		if err := db.Create(&sellProduct).Error; err != nil {
			logger.ErrorF("Error creating Sell Product: %v", err.Error())
			sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente")

			if err := db.Delete(&sell).Error; err != nil {
				logger.ErrorF("Error deleting the sell with id: %d", sell.ID)
				sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
				return
			}

			return
		}
	}

	sendSuccess(c, sell)
}
