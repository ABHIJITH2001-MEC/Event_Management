package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName        string
	EventDescription string
	Type             string
	EventLocation    string
	EventDate        time.Time
	EventDuration    string
}

type User struct {
	gorm.Model
	UserName      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	UserEmail     string
	ContactNumber string
	IsOrganiser   *bool
}

type Join struct {
	gorm.Model
	UserName      string
	EventName     string
	ContactNumber string
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/events", func(c *gin.Context) {
		// Handle GET request for events
		c.JSON(http.StatusOK, gin.H{"message": "Get Events"})
	})

	r.POST("/events", func(c *gin.Context) {
		// Handle POST request for events
		c.JSON(http.StatusOK, gin.H{"message": "Create Event"})
	})

	// Add more routes as needed

	return r
}

func main() {
	dsn := "your_sql_server_connection_string"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create tables based on the provided models
	db.AutoMigrate(&Event{}, &User{}, &Join{})

	r := SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
