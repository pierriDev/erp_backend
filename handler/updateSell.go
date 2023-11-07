package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateSellHandler(c *gin.Context) {
	request := UpdateSellRequest{}
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

	sell := schemas.Sell{}

	// }
	if err := db.Preload("PaymentMethod").Preload("Employee.User.Address").Preload("Client.User.Address").First(&sell, id).Error; err != nil {
		logger.ErrorF("Sell of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A venda de id: %v não foi encontrada", id))
		return
	}

	// Update Category
	if request.Status != "" {
		sell.Status = request.Status
	}

	if err := db.Save(&sell).Error; err != nil {

		logger.ErrorF("error updating status: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, sell)
}
