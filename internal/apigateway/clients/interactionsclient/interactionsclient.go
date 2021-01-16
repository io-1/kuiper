package interactionsclient

import (
	"time"

	"github.com/io-1/kuiper/internal/logger"
	"google.golang.org/grpc"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

const (
	FIVE_MINUTES = time.Minute * 5
)

type InteractionsClient struct {
	logger                    logger.Logger
	interactionsServiceClient interactions_pb.InteractionsServiceClient
}

func NewInteractionsClient(serverEnv string, logger logger.Logger) (*InteractionsClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &InteractionsClient{
		interactionsServiceClient: interactions_pb.NewInteractionsServiceClient(conn),
	}
	return client, nil
}

func NewInteractionsClientWithMock(interactionsServiceClient interactions_pb.InteractionsServiceClient, logger logger.Logger) *InteractionsClient {
	client := &InteractionsClient{
		logger:                    logger,
		interactionsServiceClient: interactionsServiceClient,
	}
	return client
}
