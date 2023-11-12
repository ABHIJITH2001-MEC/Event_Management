package Models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	U_Id          uint   `gorm:"primaryKey"`
	UserName      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	UserEmail     string
	ContactNumber string
	IsOrganiser   *bool
}
