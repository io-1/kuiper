package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lamppulseevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/lamppulseevents/response"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (client InteractionsClient) CreateLampPulseEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampPulseEventRequest
		res response.CreateLampPulseEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampPulseEvent(ctx, &interactions_pb.CreateLampPulseEventRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampPulseEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampPulseEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampPulseEventRequest
		res           response.GetLampPulseEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_pulse_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampPulseEvent(ctx, &interactions_pb.GetLampPulseEventRequest{ID: id})
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

	res = response.GetLampPulseEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampPulseEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampPulseEventRequest
		res           response.UpdateLampPulseEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_pulse_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateLampPulseEvent(ctx, &interactions_pb.UpdateLampPulseEventRequest{
		ID:    id,
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
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

	res = response.UpdateLampPulseEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampPulseEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampPulseEventRequest
		res           response.PatchLampPulseEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_pulse_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetLampPulseEvent(ctx, &interactions_pb.GetLampPulseEventRequest{ID: id})

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

	// FIXME: not sure how to fix these
	// if req.Red == nil {
	// 	req.Red = r.Red
	// }

	// if req.Green == "" {
	// 	req.Green = r.Green
	// }

	// if req.Blue == "" {
	// 	req.Blue = r.Blue
	// }

	// save the request difference
	re, err := client.interactionsServiceClient.UpdateLampPulseEvent(ctx, &interactions_pb.UpdateLampPulseEventRequest{
		ID:    id,
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
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

	res = response.PatchLampPulseEventResponse{
		ID:    re.ID,
		Mac:   re.Mac,
		Red:   re.Red,
		Green: re.Green,
		Blue:  re.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampPulseEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampPulseEventRequest
		res           response.DeleteLampPulseEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_pulse_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampPulseEvent(ctx, &interactions_pb.DeleteLampPulseEventRequest{
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

	res = response.DeleteLampPulseEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
