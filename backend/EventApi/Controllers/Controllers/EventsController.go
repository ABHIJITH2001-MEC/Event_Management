package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName string
	// Add other fields as needed
}

var db *gorm.DB

func init() {
	// Open a database connection
	database, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create the Events table
	database.AutoMigrate(&Event{})

	db = database
}

func main() {
	router := gin.Default()

	// GET: /api/events
	router.GET("/api/events", GetEvents)

	// GET: /api/events/:id
	router.GET("/api/events/:id", GetEvent)

	// PUT: /api/events/:id
	router.PUT("/api/events/:id", PutEvent)

	// POST: /api/events
	router.POST("/api/events", PostEvent)

	// DELETE: /api/events/:id
	router.DELETE("/api/events/:id", DeleteEvent)

	router.Run(":8080")
}

// Handler for GET: /api/events
func GetEvents(c *gin.Context) {
	var events []Event
	if result := db.Find(&events); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No events found"})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Handler for GET: /api/events/:id
func GetEvent(c *gin.Context) {
	var event Event
	id := c.Param("id")
	if result := db.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Handler for PUT: /api/events/:id
func PutEvent(c *gin.Context) {
	var event Event
	id := c.Param("id")
	if result := db.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Handler for POST: /api/events
func PostEvent(c *gin.Context) {
	var event Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// Handler for DELETE: /api/events/:id
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	var event Event
	if result := db.First(&event, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if result := db.Delete(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
