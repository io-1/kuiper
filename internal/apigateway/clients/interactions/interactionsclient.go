package interactions

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/response"
)

const (
	FIVE_MINUTES = time.Minute * 5
)

type InteractionsClient struct {
}

func (client InteractionsClient) CreateInteraction(c *gin.Context) {

	_, cancel := context.WithTimeout(c, FIVE_MINUTES)
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

	c.JSON(http.StatusOK, res)
}
