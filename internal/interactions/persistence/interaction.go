package persistence

import "time"

type Interaction struct {
	ID          string `grom:"primary_key"`
	Name        string
	Description string
	CreatedAt   *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}
