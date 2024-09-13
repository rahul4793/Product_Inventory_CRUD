package main

import (
	"go-inventory/routes"
	"go-inventory/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Your frontend URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Initialize MongoDB
	utils.InitMongo()

	// Set up routes
	routes.ProductRoutes(router)

	// Start server
	router.Run(":8080")
}
