package listeners

import (
	"github.com/io-1/kuiper/internal/logger"
	"github.com/io-1/kuiper/internal/sensors/persistence"
)

type SensorsListenersEnv struct {
	persistence persistence.Persistence
	logger      logger.Logger
}

func NewSensorsListenersEnv(persistence persistence.Persistence, logger logger.Logger) *SensorsListenersEnv {
	return &SensorsListenersEnv{
		persistence: persistence,
		logger:      logger,
	}
}
