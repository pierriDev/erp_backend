package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateProductHandler(c *gin.Context) {
	request := UpdateProductRequest{}
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

	category := schemas.Category{}
	// FIND CATEGORY
	if err := db.First(&category, request.CategoryID).Error; err != nil {
		logger.ErrorF("Category of id: %d not found", request.CategoryID)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id %d não foi encontrada", request.CategoryID))
		return
	}

	product := schemas.Product{}
	if err := db.First(&product, id).Error; err != nil {
		logger.ErrorF("Product of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O produto de id: %v não foi encontrado", id))
		return
	}

	// Update Category
	if request.Title != "" {
		product.Title = request.Title
		product.Price = request.Price
		product.Code = request.Code
		product.Description = request.Description
		product.CategoryID = category.ID
	}

	if err := db.Save(&product).Error; err != nil {

		logger.ErrorF("error updating product: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, product)
}
