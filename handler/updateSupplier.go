package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateSupplierHandler(c *gin.Context) {
	request := UpdateSupplierRequest{}
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

	supplier := schemas.Supplier{
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

	if err := db.Preload("User.Address").First(&supplier, id).Error; err != nil {
		logger.ErrorF("Supplier of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O fornecedor de id: %v não foi encontrado", id))
		return
	}

	if err := db.Preload("Address").First(&user, supplier.UserID).Error; err != nil {
		logger.ErrorF("User of id: %d not found", supplier.UserID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O usuario de id: %v não foi encontrado", supplier.UserID))
		return
	}

	if err := db.First(&address, user.AddressID).Error; err != nil {
		logger.ErrorF("Address of id: %d not found", supplier.User.AddressID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O endereco de id: %v não foi encontrado", supplier.User.AddressID))
		return
	}

	oldUser := user

	if request.Name != "" {
		supplier.User.Name = request.Name
		user.Name = request.Name
	}
	if request.NationalID != "" {
		supplier.User.NationalID = request.NationalID
		user.NationalID = request.NationalID
	}
	if request.Phone != "" {
		supplier.User.Phone = request.Phone
		user.Phone = request.Phone
	}
	if request.Email != "" {
		supplier.User.Email = request.Email
		user.Email = request.Email
	}
	if request.Address != "" {
		supplier.User.Address.Adress = request.Address
		address.Adress = request.Address
	}
	if request.Number > 0 {
		supplier.User.Address.Number = request.Number
		address.Number = request.Number
	}
	if request.Neighborhood != "" {
		supplier.User.Address.Neighborhood = request.Neighborhood
		address.Neighborhood = request.Neighborhood
	}
	if request.CEP != "" {
		supplier.User.Address.CEP = request.CEP
		address.CEP = request.CEP
	}
	if request.City != "" {
		supplier.User.Address.City = request.City
		address.City = request.City
	}
	if request.State != "" {
		supplier.User.Address.State = request.State
		address.State = request.State
	}
	if request.Country != "" {
		supplier.User.Address.Country = request.Country
		address.Country = request.Country
	}

	if err := db.Save(&supplier).Error; err != nil {
		logger.ErrorF("error updating supplier: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&user).Error; err != nil {
		logger.ErrorF("error updating supplier: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&address).Error; err != nil {
		logger.ErrorF("error updating supplier: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")

		if err := db.Save(&oldUser).Error; err != nil {
			logger.ErrorF("error updating supplier: %v", err.Error())
			sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
			return
		}

		return
	}

	sendSuccess(c, supplier)
}
