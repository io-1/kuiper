package mosquitto

import (
	"github.com/io-1/kuiper/internal/devices/persistence"
	"github.com/io-1/kuiper/internal/logger"
)

type MosquittoPubSub struct {
	persistence persistence.Persistence
	logger      logger.Logger
}

func NewMosquittoPubSub(persistence persistence.Persistence, logger logger.Logger) *MosquittoPubSub {
	return &MosquittoPubSub{
		persistence: persistence,
		logger:      logger,
	}
}
