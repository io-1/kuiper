package interactions

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/response"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

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
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetInteractionRequest
		res           response.GetInteractionResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("interaction_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetInteraction(ctx, &interactions_pb.GetInteractionRequest{ID: id})
	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetInteractionResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetInteractionDetails(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetInteractionDetailsRequest
		res           response.GetInteractionDetailsResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("interaction_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	stream, err := client.interactionsServiceClient.GetInteractionDetails(ctx, &interactions_pb.GetInteractionDetailsRequest{ID: id})
	if err != nil {
		client.logger.Errorf("error with interaction service: %v", err)
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	i := []response.KeypadConditionsToLampEventsInteraction{}
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}

		re := response.KeypadConditionsToLampEventsInteraction{
			KeypadCondition: response.KeypadCondition{
				ID:       r.KeypadConditionID,
				Mac:      r.KeypadConditionMac,
				ButtonID: r.KeypadConditionButtonID,
			},
			LampEvent: response.LampEvent{
				ID:        r.LampEventID,
				Mac:       r.LampEventMac,
				EventType: r.LampEventType,
				Red:       r.LampEventRed,
				Green:     r.LampEventGreen,
				Blue:      r.LampEventBlue,
			},
		}
		i = append(i, re)
	}

	res = response.GetInteractionDetailsResponse{
		ID:           id,
		Interactions: i,
	}
	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateInteractionRequest
		res           response.UpdateInteractionResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("interaction_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateInteraction(ctx, &interactions_pb.UpdateInteractionRequest{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.UpdateInteractionResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchInteractionRequest
		res           response.PatchInteractionResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("interaction_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetInteraction(ctx, &interactions_pb.GetInteractionRequest{ID: id})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	if req.Name == "" {
		req.Name = r.Name
	}

	if req.Description == "" {
		req.Description = r.Description
	}

	// save the request difference
	re, err := client.interactionsServiceClient.UpdateInteraction(ctx, &interactions_pb.UpdateInteractionRequest{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		st, _ := status.FromError(err)
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if re.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.PatchInteractionResponse{
		ID:          re.ID,
		Name:        re.Name,
		Description: re.Description,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteInteractionRequest
		res           response.DeleteInteractionResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("interaction_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteInteraction(ctx, &interactions_pb.DeleteInteractionRequest{
		ID: id,
	})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.DeleteInteractionResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
