package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnesstoggleevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnesstoggleevents/response"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (client InteractionsClient) CreateLampAutoBrightnessToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampAutoBrightnessToggleEventRequest
		res response.CreateLampAutoBrightnessToggleEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampAutoBrightnessToggleEvent(ctx, &interactions_pb.CreateLampAutoBrightnessToggleEventRequest{
		Mac: req.Mac,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampAutoBrightnessToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampAutoBrightnessToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampAutoBrightnessToggleEventRequest
		res           response.GetLampAutoBrightnessToggleEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampAutoBrightnessToggleEvent(ctx, &interactions_pb.GetLampAutoBrightnessToggleEventRequest{ID: id})
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

	res = response.GetLampAutoBrightnessToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampAutoBrightnessToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampAutoBrightnessToggleEventRequest
		res           response.UpdateLampAutoBrightnessToggleEventResponse
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

	r, err := client.interactionsServiceClient.UpdateLampAutoBrightnessToggleEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessToggleEventRequest{
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

	res = response.UpdateLampAutoBrightnessToggleEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampAutoBrightnessToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampAutoBrightnessToggleEventRequest
		res           response.PatchLampAutoBrightnessToggleEventResponse
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
	r, err := client.interactionsServiceClient.GetLampAutoBrightnessToggleEvent(ctx, &interactions_pb.GetLampAutoBrightnessToggleEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampAutoBrightnessToggleEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessToggleEventRequest{
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

	res = response.PatchLampAutoBrightnessToggleEventResponse{
		ID:  re.ID,
		Mac: re.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampAutoBrightnessToggleEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampAutoBrightnessToggleEventRequest
		res           response.DeleteLampAutoBrightnessToggleEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampAutoBrightnessToggleEvent(ctx, &interactions_pb.DeleteLampAutoBrightnessToggleEventRequest{
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

	res = response.DeleteLampAutoBrightnessToggleEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
