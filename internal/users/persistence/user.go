package persistence

import "time"

type User struct {
	ID        string `grom:"primary_key"`
	Username  string `grom:"unique"`
	Password  string
	Name      string
	Email     string     `grom:"unique"`
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
