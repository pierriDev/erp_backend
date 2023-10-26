package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func GetClientHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	client := schemas.Client{}
	// FIND OPENING
	if err := db.Preload("User.Address").First(&client, id).Error; err != nil {
		logger.ErrorF("Client of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O cliente de id %s não foi encontrado", id))
		return
	}

	sendSuccess(c, client)
}
