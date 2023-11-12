// event.go
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Event represents the Event model
type Event struct {
	gorm.Model
	EventName        string    `gorm:"column:event_name"`
	EventDescription string    `gorm:"column:event_description"`
	Type             string    `gorm:"column:type"`
	EventLocation    string    `gorm:"column:event_location"`
	EventDate        time.Time `gorm:"column:event_date"`
	EventDuration    string    `gorm:"column:event_duration"`
}

// TableName specifies the table name for the Event model
func (Event) TableName() string {
	return "events"
}
