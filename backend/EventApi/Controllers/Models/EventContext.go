// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import the SQLite driver
)

var db *gorm.DB

func main() {
	// Initialize Gin
	router := gin.Default()

	// Initialize the database
	initDatabase()

	// Define routes
	router.GET("/events", getEvents)

	// Run the server
	router.Run(":8080")
}

func initDatabase() {
	// Connect to the SQLite database
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate the models
	db.AutoMigrate(&Event{}, &User{}, &Join{})
}

// Event model
type Event struct {
	gorm.Model
	EventName        string
	EventDescription string
	Type             string
	EventLocation    string
	EventDate        string
	EventDuration    string
}

// User model
type User struct {
	gorm.Model
	UserName      string
	Password      string
	UserEmail     string
	City          string
	IsOrganiser   bool
	ContactNumber string
}

// Join model
type Join struct {
	gorm.Model
	UserName  string
	EventName string
}

// Handler to get events
func getEvents(c *gin.Context) {
	var events []Event
	db.Find(&events)
	c.JSON(200, events)
}
