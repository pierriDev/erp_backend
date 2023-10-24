package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pierriDev/erp_backend.git/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	//ROTAS DE EXEMPLO (CRIAS AS PROPRIAS DEPOIS)
	v1 := router.Group("/api/v1/")
	{
		//OPENINGS
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		//USER
		v1.POST("/user", handler.CreateUserHandler)

		//Employee
		v1.GET("/employees", handler.ListEmployeeHandler)
		v1.GET("/employee", handler.GetEmployeeHandler)
		v1.POST("/employee", handler.CreateEmployeeHandler)
		v1.PUT("/employee", handler.UpdateEmployeeHandler)
		v1.DELETE("/employee", handler.DeleteEmployeeHandler)
	}
}
