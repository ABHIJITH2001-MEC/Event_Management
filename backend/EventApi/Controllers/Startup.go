package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// JWT middleware (example, you need to customize this based on your authentication mechanism)
	r.Use(authMiddleware)

	// Your routes and handlers go here
	r.GET("/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Get Events"})
	})

	// Add more routes as needed

	port := "8080" // Set your desired port
	err := r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}

func authMiddleware(c *gin.Context) {
	// Your JWT authentication logic goes here
	// Example: check the Authorization header for a valid JWT token

	// For simplicity, we just check if the Authorization header is present
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Your JWT validation logic goes here
	// Example: validate the token using your JWT library
	// Note: Be sure to handle token expiration, signature verification, etc.

	// If the token is valid, you can proceed to the next middleware or handler
}
