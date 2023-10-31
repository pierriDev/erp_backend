package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateClientHandler(c *gin.Context) {
	request := CreateClientRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	client := schemas.Client{
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

	if err := db.Create(&client).Error; err != nil {
		logger.ErrorF("Error creating Client: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar seu cliente. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, client)
}
