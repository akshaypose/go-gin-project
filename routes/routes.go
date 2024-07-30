package routes

import (
	"go-gin-project/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Product routes
	r.GET("/products", handlers.GetAllProducts)
	r.GET("/product/:id", handlers.GetProductByID)

	return r
}
