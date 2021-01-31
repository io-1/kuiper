package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnessonevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnessonevents/response"
	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (client InteractionsClient) CreateLampAutoBrightnessOnEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampAutoBrightnessOnEventRequest
		res response.CreateLampAutoBrightnessOnEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampAutoBrightnessOnEvent(ctx, &interactions_pb.CreateLampAutoBrightnessOnEventRequest{
		Mac: req.Mac,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampAutoBrightnessOnEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampAutoBrightnessOnEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampAutoBrightnessOnEventRequest
		res           response.GetLampAutoBrightnessOnEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampAutoBrightnessOnEvent(ctx, &interactions_pb.GetLampAutoBrightnessOnEventRequest{ID: id})
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

	res = response.GetLampAutoBrightnessOnEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampAutoBrightnessOnEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampAutoBrightnessOnEventRequest
		res           response.UpdateLampAutoBrightnessOnEventResponse
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

	r, err := client.interactionsServiceClient.UpdateLampAutoBrightnessOnEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessOnEventRequest{
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

	res = response.UpdateLampAutoBrightnessOnEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampAutoBrightnessOnEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampAutoBrightnessOnEventRequest
		res           response.PatchLampAutoBrightnessOnEventResponse
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
	r, err := client.interactionsServiceClient.GetLampAutoBrightnessOnEvent(ctx, &interactions_pb.GetLampAutoBrightnessOnEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampAutoBrightnessOnEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessOnEventRequest{
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

	res = response.PatchLampAutoBrightnessOnEventResponse{
		ID:  re.ID,
		Mac: re.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampAutoBrightnessOnEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampAutoBrightnessOnEventRequest
		res           response.DeleteLampAutoBrightnessOnEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampAutoBrightnessOnEvent(ctx, &interactions_pb.DeleteLampAutoBrightnessOnEventRequest{
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

	res = response.DeleteLampAutoBrightnessOnEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
