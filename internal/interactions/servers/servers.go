package servers

import (
	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

type InteractionsServer struct {
	persistence persistence.Persistence
	interactions_pb.UnimplementedInteractionsServiceServer
}

func NewInteractionsServer(persistence persistence.Persistence) *InteractionsServer {
	return &InteractionsServer{
		persistence: persistence,
	}
}
