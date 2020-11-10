package mosquitto

import (
	"github.com/io-1/kuiper/internal/logger"
)

type MosquittoPubSub struct {
	logger logger.Logger
}

func NewMosquittoPubSub(logger logger.Logger) *MosquittoPubSub {
	return &MosquittoPubSub{
		logger: logger,
	}
}
