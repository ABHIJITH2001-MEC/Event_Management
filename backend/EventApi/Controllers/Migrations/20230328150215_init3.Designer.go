// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jinzhu/gorm"
)

// Event represents the Event model
type Event struct {
	gorm.Model
	EventName        string
	EventDate        time.Time
	EventDescription string
	EventDuration    string
	EventLocation    string
	Type             string
}

// Join represents the Join model
type Join struct {
	gorm.Model
	ContactNumber string
	EventName     string
	UserName      string
}

// User represents the User model
type User struct {
	gorm.Model
	ContactNumber string
	IsOrganiser   *bool
	Password      string
	UserEmail     string
	UserName      string `gorm:"unique"`
}

func main() {
	// Initialize Gin
	router := gin.Default()

	// Your Gin routes and handlers go here...

	// Database connection
	db, err := gorm.Open("postgres", "user=youruser dbname=yourdb sslmode=disable") // Update with your database connection details
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migrations
	err = runMigrations(db)
	if err != nil {
		log.Fatal(err)
	}

	// Start the Gin server
	router.Run(":8080")
}

func runMigrations(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Define your migrations here
		{
			ID: "20230328150215_init3",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Event{}, &Join{}, &User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTableIfExists("Events", "Join", "Users").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	fmt.Println("Migration completed successfully!")
	return nil
}
