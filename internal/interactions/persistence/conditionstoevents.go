package persistence

import "time"

type ConditionsToEvents struct {
	ID          string `grom:"primary_key"`
	ConditionID string
	EventID     string
	EventType   string
	CreatedAt   *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}
