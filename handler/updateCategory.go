package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func UpdateCategoryHandler(c *gin.Context) {
	request := UpdateCategoryRequest{}
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

	// }
	if err := db.First(&category, id).Error; err != nil {
		logger.ErrorF("Category of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id: %v não foi encontrada", id))
		return
	}

	// Update Category
	if request.Title != "" {
		category.Title = request.Title
	}

	if err := db.Save(&category).Error; err != nil {

		logger.ErrorF("error updating category: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro. Tente novamente mais tarde")
		return
	}

	sendSuccess(c, category)
}
