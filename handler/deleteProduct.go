package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteProductHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	product := schemas.Product{}
	// FIND Product
	if err := db.First(&product, id).Error; err != nil {
		logger.ErrorF("Product of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O Produto de id: %v não foi encontrado", id))
		return
	}

	//Delete Product
	if err := db.Delete(&product).Error; err != nil {
		logger.ErrorF("Error deleting the product with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, product)
}
