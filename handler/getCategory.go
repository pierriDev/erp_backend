package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func GetCategoryHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	category := schemas.Category{}
	// FIND OPENING
	if err := db.First(&category, id).Error; err != nil {
		logger.ErrorF("Category of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A categoria de id %s não foi encontrada", id))
		return
	}

	sendSuccess(c, category)
}
