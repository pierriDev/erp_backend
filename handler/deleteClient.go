package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteClientHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	client := schemas.Client{}
	// FIND OPENING
	if err := db.First(&client, id).Error; err != nil {
		logger.ErrorF("Client of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O funcionário de id: %v não foi encontrado", id))
		return
	}

	//Delete Opening
	if err := db.Delete(&client).Error; err != nil {
		logger.ErrorF("Error deleting the client with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, client)
}
