package main

import (
	"fmt"
	"os"
	"serendipity_backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Serendipity backend")

	router := gin.Default()
	router.Use(gin.Logger())
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Authorization", "public-request", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.SerendipityRoute(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060"
	}
	router.Run(":" + port)
}
