package models

import (
	"github.com/jinzhu/gorm"
)

// Join model
type Join struct {
	gorm.Model
	UserName      string
	EventName     string
	ContactNumber string
}
