package routes

import (
	"go-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
