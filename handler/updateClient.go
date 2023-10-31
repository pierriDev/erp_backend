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
			Name:       request.Name,
			NationalID: request.NationalID,
			Phone:      request.Phone,
			Email:      request.Email,
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

	user := schemas.User{
		Name:       request.Name,
		NationalID: request.NationalID,
		Phone:      request.Phone,
		Email:      request.Email,
		Address: schemas.Address{
			Adress:       request.Address,
			Number:       request.Number,
			Neighborhood: request.Neighborhood,
			CEP:          request.CEP,
			City:         request.City,
			State:        request.State,
			Country:      request.Country,
		},
	}

	address := schemas.Address{
		Adress:       request.Address,
		Number:       request.Number,
		Neighborhood: request.Neighborhood,
		CEP:          request.CEP,
		City:         request.City,
		State:        request.State,
		Country:      request.Country,
	}

	if err := db.Preload("User.Address").First(&client, id).Error; err != nil {
		logger.ErrorF("Client of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O cliente de id: %v não foi encontrado", id))
		return
	}

	if err := db.Preload("Address").First(&user, client.UserID).Error; err != nil {
		logger.ErrorF("User of id: %d not found", client.UserID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O usuario de id: %v não foi encontrado", client.UserID))
		return
	}

	if err := db.First(&address, user.AddressID).Error; err != nil {
		logger.ErrorF("Address of id: %d not found", client.User.AddressID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O endereco de id: %v não foi encontrado", client.User.AddressID))
		return
	}

	oldUser := user

	// Update Client
	if request.Name != "" {
		client.User.Name = request.Name
		user.Name = request.Name
	}
	if request.NationalID != "" {
		client.User.NationalID = request.NationalID
		user.NationalID = request.NationalID
	}
	if request.Phone != "" {
		client.User.Phone = request.Phone
		user.Phone = request.Phone
	}
	if request.Email != "" {
		client.User.Email = request.Email
		user.Email = request.Email
	}
	if request.Address != "" {
		client.User.Address.Adress = request.Address
		address.Adress = request.Address
	}
	if request.Number > 0 {
		client.User.Address.Number = request.Number
		address.Number = request.Number
	}
	if request.Neighborhood != "" {
		client.User.Address.Neighborhood = request.Neighborhood
		address.Neighborhood = request.Neighborhood
	}
	if request.CEP != "" {
		client.User.Address.CEP = request.CEP
		address.CEP = request.CEP
	}
	if request.City != "" {
		client.User.Address.City = request.City
		address.City = request.City
	}
	if request.State != "" {
		client.User.Address.State = request.State
		address.State = request.State
	}
	if request.Country != "" {
		client.User.Address.Country = request.Country
		address.Country = request.Country
	}

	if err := db.Save(&client).Error; err != nil {
		logger.ErrorF("error updating client: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&user).Error; err != nil {
		logger.ErrorF("error updating user client: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&address).Error; err != nil {
		logger.ErrorF("error updating address client: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")

		if err := db.Save(&oldUser).Error; err != nil {
			logger.ErrorF("error updating user client: %v", err.Error())
			sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
			return
		}

		return
	}

	sendSuccess(c, client)
}
