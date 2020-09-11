package listeners

import (
	"github.com/io-1/kuiper/internal/devices/persistence"

	logger "github.com/io-1/kuiper/internal/logger"
)

type DeviceSettingsListenersEnv struct {
	persistence persistence.Persistence
	logger      logger.Logger
}

func NewDeviceSettingsListenersEnv(persistence persistence.Persistence, logger logger.Logger) *DeviceSettingsListenersEnv {
	return &DeviceSettingsListenersEnv{
		persistence: persistence,
		logger:      logger,
	}
}
