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

		//Clients
		v1.GET("/clients", handler.ListClientHandler)
		v1.GET("/client", handler.GetClientHandler)
		v1.POST("/client", handler.CreateClientHandler)
		v1.PUT("/client", handler.UpdateClientHandler)
		v1.DELETE("/client", handler.DeleteClientHandler)

		//CATEGORY
		v1.GET("/categories", handler.ListCategoryHandler)
		v1.GET("/category", handler.GetCategoryHandler)
		v1.POST("/category", handler.CreateCategoryHandler)
		v1.PUT("/category", handler.UpdateCategoryHandler)
		v1.DELETE("/category", handler.DeleteCategoryHandler)

		// PRODUCT
		v1.GET("/products", handler.ListProductHandler)
		v1.GET("/product", handler.GetProductHandler)
		v1.POST("/product", handler.CreateProductHandler)
		v1.PUT("/product", handler.UpdateProductHandler)
		v1.DELETE("/product", handler.DeleteProductHandler)

		//STOCK
		v1.GET("/stocks", handler.ListStockHandler)
		v1.GET("/stock", handler.GetStockHandler)
		v1.POST("/stock", handler.CreateStockHandler)
		v1.PUT("/stock", handler.UpdateStockHandler)
		v1.DELETE("/stock", handler.DeleteStockHandler)

		//SUPPLIER
		v1.GET("/suppliers", handler.ListSupplierHandler)
		v1.GET("/supplier", handler.GetSupplierHandler)
		v1.POST("/supplier", handler.CreateSupplierHandler)
		v1.PUT("/supplier", handler.UpdateSupplierHandler)
		v1.DELETE("/supplier", handler.DeleteSupplierHandler)

		//PAYMENT METHOD
		v1.GET("/paymentmethods", handler.ListPaymentMethodHandler)
		v1.GET("/paymentmethod", handler.GetPaymentMethodHandler)
		v1.POST("/paymentmethod", handler.CreatePaymentMethodHandler)
		v1.PUT("/paymentmethod", handler.UpdatePaymentMethodHandler)
		v1.DELETE("/paymentmethod", handler.DeletePaymentMethodHandler)

		//SELL
		v1.GET("/sells", handler.ListSellHandler)
		v1.GET("/sell", handler.GetSellHandler)
		v1.POST("/sell", handler.CreateSellHandler)
		v1.PUT("/sell", handler.UpdateSellHandler)
		// v1.DELETE("/sell", handler.DeleteSellHandler)
	}
}
