package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampbrightnessevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampbrightnessevents/response"
	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (client InteractionsClient) CreateLampBrightnessEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampBrightnessEventRequest
		res response.CreateLampBrightnessEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampBrightnessEvent(ctx, &interactions_pb.CreateLampBrightnessEventRequest{
		Mac:        req.Mac,
		Brightness: *req.Brightness,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampBrightnessEventResponse{
		ID:         r.ID,
		Mac:        r.Mac,
		Brightness: r.Brightness,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampBrightnessEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampBrightnessEventRequest
		res           response.GetLampBrightnessEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampBrightnessEvent(ctx, &interactions_pb.GetLampBrightnessEventRequest{ID: id})
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

	res = response.GetLampBrightnessEventResponse{
		ID:         r.ID,
		Mac:        r.Mac,
		Brightness: r.Brightness,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampBrightnessEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampBrightnessEventRequest
		res           response.UpdateLampBrightnessEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateLampBrightnessEvent(ctx, &interactions_pb.UpdateLampBrightnessEventRequest{
		ID:         id,
		Mac:        req.Mac,
		Brightness: *req.Brightness,
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

	res = response.UpdateLampBrightnessEventResponse{
		ID:         r.ID,
		Mac:        r.Mac,
		Brightness: r.Brightness,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampBrightnessEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampBrightnessEventRequest
		res           response.PatchLampBrightnessEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetLampBrightnessEvent(ctx, &interactions_pb.GetLampBrightnessEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampBrightnessEvent(ctx, &interactions_pb.UpdateLampBrightnessEventRequest{
		ID:         id,
		Mac:        req.Mac,
		Brightness: *req.Brightness,
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

	res = response.PatchLampBrightnessEventResponse{
		ID:         re.ID,
		Mac:        re.Mac,
		Brightness: re.Brightness,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampBrightnessEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampBrightnessEventRequest
		res           response.DeleteLampBrightnessEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampBrightnessEvent(ctx, &interactions_pb.DeleteLampBrightnessEventRequest{
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

	res = response.DeleteLampBrightnessEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
