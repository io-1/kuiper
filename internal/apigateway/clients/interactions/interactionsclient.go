package interactions

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/response"
	"github.com/io-1/kuiper/internal/logger"
	"google.golang.org/grpc"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
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

func (client InteractionsClient) CreateInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateInteractionRequest
		res response.CreateInteractionResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.CreateInteraction(ctx, &interactions_pb.CreateInteractionRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateInteractionResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetInteraction(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, nil)
}

func (client InteractionsClient) UpdateInteraction(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, nil)
}

func (client InteractionsClient) PatchInteraction(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, nil)
}

func (client InteractionsClient) DeleteInteraction(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, nil)
}
