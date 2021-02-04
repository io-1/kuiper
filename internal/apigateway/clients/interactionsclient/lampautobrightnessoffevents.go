package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnessoffevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lampautobrightnessoffevents/response"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (client InteractionsClient) CreateLampAutoBrightnessOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampAutoBrightnessOffEventRequest
		res response.CreateLampAutoBrightnessOffEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampAutoBrightnessOffEvent(ctx, &interactions_pb.CreateLampAutoBrightnessOffEventRequest{
		Mac: req.Mac,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampAutoBrightnessOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampAutoBrightnessOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampAutoBrightnessOffEventRequest
		res           response.GetLampAutoBrightnessOffEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampAutoBrightnessOffEvent(ctx, &interactions_pb.GetLampAutoBrightnessOffEventRequest{ID: id})
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

	res = response.GetLampAutoBrightnessOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampAutoBrightnessOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampAutoBrightnessOffEventRequest
		res           response.UpdateLampAutoBrightnessOffEventResponse
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

	r, err := client.interactionsServiceClient.UpdateLampAutoBrightnessOffEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessOffEventRequest{
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

	res = response.UpdateLampAutoBrightnessOffEventResponse{
		ID:  r.ID,
		Mac: r.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampAutoBrightnessOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampAutoBrightnessOffEventRequest
		res           response.PatchLampAutoBrightnessOffEventResponse
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
	r, err := client.interactionsServiceClient.GetLampAutoBrightnessOffEvent(ctx, &interactions_pb.GetLampAutoBrightnessOffEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampAutoBrightnessOffEvent(ctx, &interactions_pb.UpdateLampAutoBrightnessOffEventRequest{
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

	res = response.PatchLampAutoBrightnessOffEventResponse{
		ID:  re.ID,
		Mac: re.Mac,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampAutoBrightnessOffEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampAutoBrightnessOffEventRequest
		res           response.DeleteLampAutoBrightnessOffEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_color_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampAutoBrightnessOffEvent(ctx, &interactions_pb.DeleteLampAutoBrightnessOffEventRequest{
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

	res = response.DeleteLampAutoBrightnessOffEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
