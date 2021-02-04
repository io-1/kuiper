package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lamptoggleevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lamptoggleevents/response"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (client InteractionsClient) CreateLampToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampToggleEventRequest
		res response.CreateLampToggleEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampToggleEvent(ctx, &interactions_pb.CreateLampToggleEventRequest{
		Mac: req.Mac,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampToggleEventRequest
		res           response.GetLampToggleEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_toggle_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampToggleEvent(ctx, &interactions_pb.GetLampToggleEventRequest{ID: id})
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

	res = response.GetLampToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampToggleEventRequest
		res           response.UpdateLampToggleEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_toggle_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateLampToggleEvent(ctx, &interactions_pb.UpdateLampToggleEventRequest{
		ID:  id,
		Mac: req.Mac,
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

	res = response.UpdateLampToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampToggleEventRequest
		res           response.PatchLampToggleEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_toggle_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetLampToggleEvent(ctx, &interactions_pb.GetLampToggleEventRequest{ID: id})

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

	if req.Mac == "" {
		req.Mac = r.Mac
	}

	// save the request difference
	re, err := client.interactionsServiceClient.UpdateLampToggleEvent(ctx, &interactions_pb.UpdateLampToggleEventRequest{
		ID:  id,
		Mac: req.Mac,
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

	res = response.PatchLampToggleEventResponse{
		ID:  re.ID,
		Mac: re.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampToggleEventRequest
		res           response.DeleteLampToggleEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_toggle_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampToggleEvent(ctx, &interactions_pb.DeleteLampToggleEventRequest{
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

	res = response.DeleteLampToggleEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
