package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateCategoryHandler(c *gin.Context) {
	request := CreateCategoryRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	category := schemas.Category{
		Title: request.Title,
	}

	if err := db.Create(&category).Error; err != nil {
		logger.ErrorF("Error creating Category: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar a categoria. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, category)
}
