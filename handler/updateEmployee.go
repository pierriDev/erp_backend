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

	if err := db.Preload("User.Address").First(&employee, id).Error; err != nil {
		logger.ErrorF("Employee of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O funcionário de id: %v não foi encontrado", id))
		return
	}

	// Update opening
	if request.Name != "" {
		employee.User.Name = request.Name
	}
	if request.CPF != "" {
		employee.User.CPF = request.CPF
	}
	if request.Phone != "" {
		employee.User.Phone = request.Phone
	}
	if request.Email != "" {
		employee.User.Email = request.Email
	}
	if request.Address != "" {
		employee.User.Address.Adress = request.Address
	}
	if request.Number > 0 {
		employee.User.Address.Number = request.Number
	}
	if request.Neighborhood != "" {
		employee.User.Address.Neighborhood = request.Neighborhood
	}
	if request.CEP != "" {
		employee.User.Address.CEP = request.CEP
	}
	if request.City != "" {
		employee.User.Address.City = request.City
	}
	if request.State != "" {
		employee.User.Address.State = request.State
	}
	if request.Country != "" {
		employee.User.Address.Country = request.Country
	}
	if request.Password != "" {
		employee.Password = request.Password
	}
	if request.Picture != "" {
		employee.Picture = request.Picture
	}

	//Save Opening

	if err := db.Save(&employee).Error; err != nil {
		logger.ErrorF("error updating employee: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, employee)
}
