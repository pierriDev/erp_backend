package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	category := schemas.Category{}
	// FIND CATEGORY
	if err := db.First(&category, id).Error; err != nil {
		logger.ErrorF("Category of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id: %v não foi encontrada", id))
		return
	}

	//Delete CAategory
	if err := db.Delete(&category).Error; err != nil {
		logger.ErrorF("Error deleting the category with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, category)
}
