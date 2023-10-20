package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	//ROTAS DE EXEMPLO (CRIAS AS PROPRIAS DEPOIS)
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handler.ListOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.PUT("/opening", handler.DeleteOpeningHandler)
	}
}
