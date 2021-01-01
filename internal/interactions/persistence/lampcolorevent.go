package persistence

import "time"

type LampColorEvent struct {
	ID        string `grom:"primary_key"`
	Mac       string
	Red       int32
	Green     int32
	Blue      int32
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
