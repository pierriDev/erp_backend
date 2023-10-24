package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateEmployeeHandler(c *gin.Context) {
	request := CreateUserWorkerRequest{}

	c.BindJSON(&request)

	employee := schemas.Employee{
		Password: request.Password,
		Picture:  request.Picture,
		User: schemas.User{
			Name:  request.Name,
			CPF:   request.CPF,
			Phone: request.Phone,
			Email: request.Email,
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

	if err := db.Create(&employee).Error; err != nil {
		logger.ErrorF("Error creating Employee: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar seu funcionário. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, employee)

}
