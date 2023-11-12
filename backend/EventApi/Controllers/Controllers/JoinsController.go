package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Join struct {
	gorm.Model
	U_Id int
	// Add other fields as needed
}

var db *gorm.DB

func init() {
	// Open a database connection
	database, err := gorm.Open(sqlite.Open("joins.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create the Joins table
	database.AutoMigrate(&Join{})

	db = database
}

func main() {
	router := gin.Default()

	// GET: /api/joins
	router.GET("/api/joins", GetJoins)

	// GET: /api/joins/:id
	router.GET("/api/joins/:id", GetJoin)

	// PUT: /api/joins/:id
	router.PUT("/api/joins/:id", PutJoin)

	// POST: /api/joins
	router.POST("/api/joins", PostJoin)

	// DELETE: /api/joins/:id
	router.DELETE("/api/joins/:id", DeleteJoin)

	router.Run(":8080")
}

// Handler for GET: /api/joins
func GetJoins(c *gin.Context) {
	var joins []Join
	if result := db.Find(&joins); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No joins found"})
		return
	}
	c.JSON(http.StatusOK, joins)
}

// Handler for GET: /api/joins/:id
func GetJoin(c *gin.Context) {
	var join Join
	id := c.Param("id")
	if result := db.First(&join, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Join not found"})
		return
	}
	c.JSON(http.StatusOK, join)
}

// Handler for PUT: /api/joins/:id
func PutJoin(c *gin.Context) {
	var join Join
	id := c.Param("id")
	if result := db.First(&join, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Join not found"})
		return
	}

	if err := c.ShouldBindJSON(&join); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&join); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update join"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Handler for POST: /api/joins
func PostJoin(c *gin.Context) {
	var join Join
	if err := c.ShouldBindJSON(&join); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&join); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create join"})
		return
	}

	c.JSON(http.StatusCreated, join)
}

// Handler for DELETE: /api/joins/:id
func DeleteJoin(c *gin.Context) {
	id := c.Param("id")
	var join Join
	if result := db.First(&join, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Join not found"})
		return
	}

	if result := db.Delete(&join); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete join"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
