package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateUserHandler(c *gin.Context) {
	request := CreateUserRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
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
	}

	if err := db.Create(&user).Error; err != nil {
		logger.ErrorF("Error creating User: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar seu usuário. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, user)

}
