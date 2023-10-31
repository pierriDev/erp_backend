package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateSupplierHandler(c *gin.Context) {
	request := CreateSupplierRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	supplier := schemas.Supplier{
		User: schemas.User{
			Name:       request.Name,
			NationalID: request.NationalID,
			Phone:      request.Phone,
			Email:      request.Email,
			Address: schemas.Address{
				Adress:       request.Adress,
				Number:       request.Number,
				Neighborhood: request.Neighborhood,
				CEP:          request.CEP,
				City:         request.City,
				State:        request.State,
				Country:      request.Country,
			},
		},
	}

	if err := db.Create(&supplier).Error; err != nil {
		logger.ErrorF("Error creating Supplier: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar seu fornecedor. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, supplier)
}
