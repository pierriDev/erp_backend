package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func DeleteSupplierHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.ErrorF("A Query da requisição está vazia ou mal formada")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "Parâmetro da Query").Error())
		return
	}

	supplier := schemas.Supplier{}
	if err := db.First(&supplier, id).Error; err != nil {
		logger.ErrorF("Supplier of id: %s not found", id)
		sendError(c, http.StatusNotFound, fmt.Sprintf("O fornecedor de id: %v não foi encontrado", id))
		return
	}

	if err := db.Delete(&supplier).Error; err != nil {
		logger.ErrorF("Error deleting the supplier with id: %s", id)
		sendError(c, http.StatusInternalServerError, "Ocorreu um erro, tente novamente mais tarde")
		return
	}

	sendSuccess(c, supplier)
}
