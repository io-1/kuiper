package mosquitto

import (
	"github.com/n7down/kuiper/internal/devices/persistence"
	"github.com/n7down/kuiper/internal/logger"
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
