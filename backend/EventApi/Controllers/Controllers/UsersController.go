package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	U_Id int
	// Add other fields as needed
}

var db *gorm.DB

func init() {
	// Open a database connection
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create the Users table
	database.AutoMigrate(&User{})

	db = database
}

func main() {
	router := gin.Default()

	// GET: /api/users
	router.GET("/api/users", GetUsers)

	// GET: /api/users/:id
	router.GET("/api/users/:id", GetUser)

	// PUT: /api/users/:id
	router.PUT("/api/users/:id", PutUser)

	// POST: /api/users
	router.POST("/api/users", PostUser)

	// DELETE: /api/users/:id
	router.DELETE("/api/users/:id", DeleteUser)

	router.Run(":8080")
}

// Handler for GET: /api/users
func GetUsers(c *gin.Context) {
	var users []User
	if result := db.Find(&users); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Handler for GET: /api/users/:id
func GetUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Handler for PUT: /api/users/:id
func PutUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Handler for POST: /api/users
func PostUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Handler for DELETE: /api/users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if result := db.Delete(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
