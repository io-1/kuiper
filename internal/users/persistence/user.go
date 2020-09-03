package persistence

import "time"

type User struct {
	Username  string `gorm:"primary_key"`
	Password  string
	Name      string
	Email     string
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
