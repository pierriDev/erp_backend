package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}
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

	opening := schemas.Opening{}

	if err := db.Preload("User.Address").First(&opening, id).Error; err != nil {
		logger.ErrorF("Opening de id: %s não foi encontrada", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A opening de id %s não foi encontrada", id))
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	//Save Opening

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("error updating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao atualizar a opening")
		return
	}

	sendSuccess(c, opening)
}
