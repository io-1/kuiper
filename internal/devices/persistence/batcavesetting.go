package persistence

import "time"

type BatCaveDeviceSetting struct {
	DeviceID       string `gorm:"primary_key"`
	DeepSleepDelay uint32
	CreatedAt      *time.Time `gorm:"index" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"index" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at"`
}

func (s BatCaveDeviceSetting) Equal(ss BatCaveDeviceSetting) bool {
	if s.DeviceID != ss.DeviceID {
		return false
	}

	if s.DeepSleepDelay != ss.DeepSleepDelay {
		return false
	}

	if s.CreatedAt != ss.CreatedAt {
		return false
	}

	if s.UpdatedAt != ss.UpdatedAt {
		return false
	}

	if s.DeletedAt != ss.DeletedAt {
		return false
	}

	return true
}
