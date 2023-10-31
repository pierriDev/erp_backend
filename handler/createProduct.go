package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func CreateProductHandler(c *gin.Context) {
	request := CreateProductRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	category := schemas.Category{}
	// FIND CATEGORY
	if err := db.First(&category, request.CategoryID).Error; err != nil {
		logger.ErrorF("Category of id: %d not found", request.CategoryID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id %d não foi encontrada", request.CategoryID))
		return
	}

	product := schemas.Product{
		Title:       request.Title,
		Price:       request.Price,
		Code:        request.Code,
		Description: request.Description,
		CategoryID:  category.ID,
		Category:    category,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.ErrorF("Error creating Product: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro ao criar o Produto. Tente novamente mais tarde")
		return
	}
	sendSuccess(c, product)
}
