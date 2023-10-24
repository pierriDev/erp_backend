package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func GetOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	opening := schemas.Opening{}
	// FIND OPENING
	if err := db.First(&opening, id).Error; err != nil {
		logger.ErrorF("Opening de id: %s não foi encontrada", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("A opening de id %s não foi encontrada", id))
		return
	}

	sendSuccess(c, opening)
}
