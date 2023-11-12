package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Join struct {
	U_Id          uint   `gorm:"primaryKey"`
	UserName      string `gorm:"not null"`
	EventName     string
	ContactNumber string
}

type Event struct {
	ID            uint `gorm:"primaryKey"`
	EventName     string
	EventDate     time.Time
	EventUser     User `gorm:"foreignKey:UserU_Id"`
	UserU_Id      uint
	EventDuration string
	EventLocation string
	Type          string
}

type User struct {
	U_Id          uint   `gorm:"primaryKey"`
	UserName      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	UserEmail     string
	City          string
	ContactNumber string
}

func main() {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the Join, Event, and User structs to the database
	err = db.AutoMigrate(&Join{}, &Event{}, &User{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	// Add the new columns and relationship in your existing data
	db.Exec("ALTER TABLE joins ADD COLUMN contact_number TEXT;")
	db.Exec("ALTER TABLE events ADD COLUMN u_id INTEGER;")
	db.Exec("ALTER TABLE events ADD COLUMN user_u_id INTEGER;")
	db.Exec("CREATE INDEX idx_events_user_u_id ON events (user_u_id);")
	db.Exec("UPDATE events SET user_u_id = (SELECT u_id FROM users WHERE users.u_id = events.u_id);")

	fmt.Println("Tables migrated successfully.")
}
