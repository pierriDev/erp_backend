package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateEmployeeHandler(c *gin.Context) {
	request := UpdateEmployeeRequest{}
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

	employee := schemas.Employee{
		Password: request.Password,
		Picture:  request.Picture,
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

	if err := db.Preload("User.Address").First(&employee, id).Error; err != nil {
		logger.ErrorF("Employee of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O funcionário de id: %v não foi encontrado", id))
		return
	}

	if err := db.Preload("Address").First(&user, employee.UserID).Error; err != nil {
		logger.ErrorF("User of id: %d not found", employee.UserID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O usuario de id: %v não foi encontrado", employee.UserID))
		return
	}

	if err := db.First(&address, user.AddressID).Error; err != nil {
		logger.ErrorF("Address of id: %d not found", employee.User.AddressID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O endereco de id: %v não foi encontrado", employee.User.AddressID))
		return
	}

	oldUser := user

	// Update Employee
	if request.Name != "" {
		employee.User.Name = request.Name
		user.Name = request.Name
	}
	if request.NationalID != "" {
		employee.User.NationalID = request.NationalID
		user.NationalID = request.NationalID
	}
	if request.Phone != "" {
		employee.User.Phone = request.Phone
		user.Phone = request.Phone
	}
	if request.Email != "" {
		employee.User.Email = request.Email
		user.Email = request.Email
	}
	if request.Address != "" {
		employee.User.Address.Adress = request.Address
		address.Adress = request.Address
	}
	if request.Number > 0 {
		employee.User.Address.Number = request.Number
		address.Number = request.Number
	}
	if request.Neighborhood != "" {
		employee.User.Address.Neighborhood = request.Neighborhood
		address.Neighborhood = request.Neighborhood
	}
	if request.CEP != "" {
		employee.User.Address.CEP = request.CEP
		address.CEP = request.CEP
	}
	if request.City != "" {
		employee.User.Address.City = request.City
		address.City = request.City
	}
	if request.State != "" {
		employee.User.Address.State = request.State
		address.State = request.State
	}
	if request.Country != "" {
		employee.User.Address.Country = request.Country
		address.Country = request.Country
	}
	if request.Password != "" {
		employee.Password = request.Password
	}
	if request.Picture != "" {
		employee.Picture = request.Picture
	}

	if err := db.Save(&employee).Error; err != nil {
		logger.ErrorF("error updating employee: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&user).Error; err != nil {
		logger.ErrorF("error updating user employee: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	if err := db.Save(&address).Error; err != nil {
		logger.ErrorF("error updating address employee: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")

		if err := db.Save(&oldUser).Error; err != nil {
			logger.ErrorF("error updating user employee: %v", err.Error())
			sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
			return
		}

		return
	}

	sendSuccess(c, employee)
}
