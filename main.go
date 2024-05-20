package main

import (
	routes "go-jwt-mongo/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoute(router)
	routes.UserRoute(router)
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Hello",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Success 2"})
	})

	router.Run(":" + port)

}
