package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateClientHandler(c *gin.Context) {
	request := UpdateClientRequest{}
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

	client := schemas.Client{
		User: schemas.User{
			Name:  request.Name,
			CPF:   request.CPF,
			Phone: request.Phone,
			Email: request.Email,
			Address: schemas.Address{
				Adress:       request.Address,
				Number:       request.Number,
				Neighborhood: request.Neighborhood,
				CEP:          request.CEP,
				City:         request.City,
				State:        request.State,
				Country:      request.Country,
			},
		},
	}

	if err := db.Preload("User.Address").First(&client, id).Error; err != nil {
		logger.ErrorF("Client of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O cliente de id: %v não foi encontrado", id))
		return
	}

	logger.InfoF("Selected Client %+v", client)

	// Update Client
	if request.Name != "" {
		client.User.Name = request.Name
	}
	if request.CPF != "" {
		client.User.CPF = request.CPF
	}
	if request.Phone != "" {
		client.User.Phone = request.Phone
	}
	if request.Email != "" {
		client.User.Email = request.Email
	}
	if request.Address != "" {
		client.User.Address.Adress = request.Address
	}
	if request.Number > 0 {
		client.User.Address.Number = request.Number
	}
	if request.Neighborhood != "" {
		client.User.Address.Neighborhood = request.Neighborhood
	}
	if request.CEP != "" {
		client.User.Address.CEP = request.CEP
	}
	if request.City != "" {
		client.User.Address.City = request.City
	}
	if request.State != "" {
		client.User.Address.State = request.State
	}
	if request.Country != "" {
		client.User.Address.Country = request.Country
	}

	if err := db.Save(&client).Error; err != nil {
		logger.ErrorF("error updating client: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, client)
}
