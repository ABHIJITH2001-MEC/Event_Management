package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Event struct {
	ID            uint   `gorm:"primaryKey"`
	EventName     string `gorm:"not null"`
	EventDate     time.Time
	EventUser     User `gorm:"foreignKey:UserNameU_Id"`
	UserNameU_Id  uint
	EventDuration string
	EventLocation string
	Type          string
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

	// Auto Migrate the Event and User structs to the database
	err = db.AutoMigrate(&Event{}, &User{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	// Add foreign key relationship
	db.Model(&Event{}).AddForeignKey("user_name_u_id", "users(u_id)", "CASCADE", "CASCADE")

	fmt.Println("Tables migrated successfully.")
}
