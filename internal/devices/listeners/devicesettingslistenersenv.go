package listeners

import (
	"github.com/n7down/kuiper/internal/devices/persistence"

	logger "github.com/n7down/kuiper/internal/logger"
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
