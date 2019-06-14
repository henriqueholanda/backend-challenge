package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
)

func SetupRouter(checkoutHandlers *handlers.CheckoutHandlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.POST("/checkout/basket", checkoutHandlers.Create)
		v1.DELETE("/checkout/basket/:id", checkoutHandlers.Delete)
		v1.GET("/checkout/basket/:id/amount", checkoutHandlers.FetchAmount)
		v1.POST("/checkout/basket/:id/products", checkoutHandlers.AddProduct)
	}

	return router
}
