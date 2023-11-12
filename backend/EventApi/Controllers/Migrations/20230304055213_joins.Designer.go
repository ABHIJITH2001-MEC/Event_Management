package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Event struct {
	ID               uint `gorm:"primaryKey"`
	EventName        string
	EventDescription string
	Type             string
	EventLocation    string
	EventDate        time.Time
	EventDuration    string
}

type Join struct {
	U_Id      uint `gorm:"primaryKey"`
	EventName string
	UserName  string
}

type User struct {
	U_Id      uint   `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	UserEmail string
	City      string
}

func main() {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the Event, Join, and User structs to the database
	err = db.AutoMigrate(&Event{}, &Join{}, &User{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	fmt.Println("Tables migrated successfully.")
}
