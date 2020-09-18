// +build !unit,integration

package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/devices/persistence"
	"github.com/stretchr/testify/assert"
)

var (
	db *MysqlPersistence
)

func Test_CreateBatCaveSetting(t *testing.T) {
	var (
		setting = persistence.BatCaveDeviceSetting{
			ID:             uuid.New().String(),
			Mac:            "000000000000",
			DeepSleepDelay: 30,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		rowsAffectedExpected int64 = 1
	)

	rowsAffectedActual := db.CreateBatCaveDeviceSetting(setting)
	assert.Equal(t, rowsAffectedExpected, rowsAffectedActual)
}

func Test_GetBatCaveSetting(t *testing.T) {
	var (
		id              = "11111111-2222-3333-4444-555555555555"
		mac             = "000000000011"
		settingExpected = persistence.BatCaveDeviceSetting{
			ID:             id,
			Mac:            mac,
			DeepSleepDelay: 30,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		recordNotFoundExpected = false
	)

	recordNotFoundActual, settingActual := db.GetBatCaveDeviceSetting(id)
	assert.Equal(t, recordNotFoundExpected, recordNotFoundActual)
	assert.True(t, settingExpected.Equal(settingActual))
}

func Test_UpdateBatCaveSetting(t *testing.T) {
	var (
		id                    = "22222222-3333-4444-5555-777777777777"
		mac                   = "000000001111"
		deepSleepDelay uint32 = 32

		setting = persistence.BatCaveDeviceSetting{
			ID:             id,
			Mac:            mac,
			DeepSleepDelay: deepSleepDelay,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}

		settingExpected = persistence.BatCaveDeviceSetting{
			ID:             id,
			Mac:            mac,
			DeepSleepDelay: deepSleepDelay,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		rowsAffectedExpected   int64 = 1
		recordNotFoundExpected       = false
	)

	rowsAffectedActual := db.UpdateBatCaveDeviceSetting(setting)
	assert.Equal(t, rowsAffectedExpected, rowsAffectedActual)

	recordNotFoundActual, settingActual := db.GetBatCaveDeviceSetting(id)
	assert.Equal(t, recordNotFoundExpected, recordNotFoundActual)
	assert.Equal(t, settingExpected.ID, settingActual.ID)
	assert.Equal(t, settingExpected.DeepSleepDelay, settingActual.DeepSleepDelay)
}

func TestMain(m *testing.M) {
	var (
		err error
	)

	dbConn := os.Getenv("DB_CONN")
	db, err = NewMysqlPersistence(dbConn)
	if err != nil {
		log.Fatal(err.Error())
	}
	code := m.Run()
	os.Exit(code)
}
