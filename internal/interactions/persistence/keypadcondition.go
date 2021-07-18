package persistence

import "time"

type KeypadCondition struct {
	ID        *string `grom:"primary_key"`
	Mac       *string
	ButtonID  *int32
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
