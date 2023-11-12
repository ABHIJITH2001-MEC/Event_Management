package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Join struct {
	U_Id      uint `gorm:"primaryKey"`
	UserName  string
	EventName string
}

func main() {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("joins.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the Join struct to the database
	err = db.AutoMigrate(&Join{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	fmt.Println("Tables migrated successfully.")
}
