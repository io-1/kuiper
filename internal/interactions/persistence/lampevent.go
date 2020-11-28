package persistence

import "time"

type LampEvent struct {
	ID        string `grom:"primary_key"`
	Mac       string
	EventType string
	Color     string
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
