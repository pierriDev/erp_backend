package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/schemas"
)

func ListSupplierHandler(c *gin.Context) {
	supplier := []schemas.Supplier{}

	if err := db.Preload("User.Address").Find(&supplier).Error; err != nil {
		logger.ErrorF("Error listing Suppliers: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Erro ao buscar os fornecedores")
		return
	}

	sendSuccess(c, supplier)
}
