package persistence

import "time"

type KeypadConditionsToLampEvents struct {
	ID            string `grom:"primary_key"`
	InteractionID string
	ConditionID   string
	EventID       string
	CreatedAt     *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at"`
}
