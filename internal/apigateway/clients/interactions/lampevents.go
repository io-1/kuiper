package interactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/response"
	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
	"google.golang.org/grpc/status"
)

func (client InteractionsClient) CreateLampEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampEventRequest
		res response.CreateLampEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampEvent(ctx, &interactions_pb.CreateLampEventRequest{
		Mac:       req.Mac,
		EventType: req.EventType,
		Red:       *req.Red,
		Green:     *req.Green,
		Blue:      *req.Blue,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampEventResponse{
		ID:        r.ID,
		Mac:       r.Mac,
		EventType: r.EventType,
		Red:       r.Red,
		Green:     r.Green,
		Blue:      r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampEventRequest
		res           response.GetLampEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampEvent(ctx, &interactions_pb.GetLampEventRequest{ID: id})
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

	res = response.GetLampEventResponse{
		ID:        r.ID,
		Mac:       r.Mac,
		EventType: r.EventType,
		Red:       r.Red,
		Green:     r.Green,
		Blue:      r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampEventRequest
		res           response.UpdateLampEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateLampEvent(ctx, &interactions_pb.UpdateLampEventRequest{
		ID:        id,
		Mac:       req.Mac,
		EventType: req.EventType,
		Red:       *req.Red,
		Green:     *req.Green,
		Blue:      *req.Blue,
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

	res = response.UpdateLampEventResponse{
		ID:        r.ID,
		Mac:       r.Mac,
		EventType: r.EventType,
		Red:       r.Red,
		Green:     r.Green,
		Blue:      r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampEventRequest
		res           response.PatchLampEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetLampEvent(ctx, &interactions_pb.GetLampEventRequest{ID: id})

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

	if req.EventType == "" {
		req.EventType = r.EventType
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
	re, err := client.interactionsServiceClient.UpdateLampEvent(ctx, &interactions_pb.UpdateLampEventRequest{
		ID:        id,
		Mac:       req.Mac,
		EventType: req.EventType,
		Red:       *req.Red,
		Green:     *req.Green,
		Blue:      *req.Blue,
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

	res = response.PatchLampEventResponse{
		ID:        re.ID,
		Mac:       re.Mac,
		EventType: re.EventType,
		Red:       re.Red,
		Green:     re.Green,
		Blue:      re.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampEventRequest
		res           response.DeleteLampEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampEvent(ctx, &interactions_pb.DeleteLampEventRequest{
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

	res = response.DeleteLampEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
