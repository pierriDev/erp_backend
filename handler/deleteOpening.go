package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteOpeningHandler(c *gin.Context) {
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

	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		logger.ErrorF("Erro ao deletar a opening de id: %s", id)
		sendError(c, http.StatusInternalServerError, "Erro ao deletar a opening")
		return
	}

	sendSuccess(c, opening)
}
