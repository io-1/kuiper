package persistence

import "time"

type LampOnEvent struct {
	ID        string `grom:"primary_key"`
	Mac       string
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
