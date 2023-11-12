// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jinzhu/gorm"
	"github.com/rubenv/sql-migrate"
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
	db, err := gorm.Open("pgx", "user=youruser dbname=yourdb sslmode=disable") // Update with your database connection details
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migrations
	err = runMigrations(db.DB())
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin
	router := gin.Default()

	// Your Gin routes and handlers go here...

	// Start the Gin server
	router.Run(":8080")
}

func runMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "path/to/migrations", // Update with the path to your migration files
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
