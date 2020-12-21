package servers

import (
	"github.com/io-1/kuiper/internal/interactions/persistence"
	"github.com/io-1/kuiper/internal/logger"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

type InteractionsServer struct {
	persistence persistence.Persistence
	logger      logger.Logger
	interactions_pb.UnimplementedInteractionsServiceServer
}

func NewInteractionsServer(persistence persistence.Persistence, logger logger.Logger) *InteractionsServer {
	return &InteractionsServer{
		persistence: persistence,
		logger:      logger,
	}
}
