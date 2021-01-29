package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampoffevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampoffevents/response"
	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (client InteractionsClient) CreateLampOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampOffEventRequest
		res response.CreateLampOffEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampOffEvent(ctx, &interactions_pb.CreateLampOffEventRequest{
		Mac: req.Mac,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampOffEventRequest
		res           response.GetLampOffEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_off_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampOffEvent(ctx, &interactions_pb.GetLampOffEventRequest{ID: id})
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

	res = response.GetLampOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampOffEventRequest
		res           response.UpdateLampOffEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_off_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateLampOffEvent(ctx, &interactions_pb.UpdateLampOffEventRequest{
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

	res = response.UpdateLampOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampOffEventRequest
		res           response.PatchLampOffEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_off_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetLampOffEvent(ctx, &interactions_pb.GetLampOffEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampOffEvent(ctx, &interactions_pb.UpdateLampOffEventRequest{
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

	res = response.PatchLampOffEventResponse{
		ID:  re.ID,
		Mac: re.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampOffEventRequest
		res           response.DeleteLampOffEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_off_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampOffEvent(ctx, &interactions_pb.DeleteLampOffEventRequest{
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

	res = response.DeleteLampOffEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
