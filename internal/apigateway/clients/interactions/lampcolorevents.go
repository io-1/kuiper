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

func (client InteractionsClient) CreateLampColorEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateLampColorEventRequest
		res response.CreateLampColorEventResponse
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

	r, err := client.interactionsServiceClient.CreateLampColorEvent(ctx, &interactions_pb.CreateLampColorEventRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateLampColorEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetLampColorEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetLampColorEventRequest
		res           response.GetLampColorEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetLampColorEvent(ctx, &interactions_pb.GetLampColorEventRequest{ID: id})
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

	res = response.GetLampColorEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateLampColorEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateLampColorEventRequest
		res           response.UpdateLampColorEventResponse
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

	r, err := client.interactionsServiceClient.UpdateLampColorEvent(ctx, &interactions_pb.UpdateLampColorEventRequest{
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

	res = response.UpdateLampColorEventResponse{
		ID:    r.ID,
		Mac:   r.Mac,
		Red:   r.Red,
		Green: r.Green,
		Blue:  r.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchLampColorEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchLampColorEventRequest
		res           response.PatchLampColorEventResponse
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
	r, err := client.interactionsServiceClient.GetLampColorEvent(ctx, &interactions_pb.GetLampColorEventRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateLampColorEvent(ctx, &interactions_pb.UpdateLampColorEventRequest{
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

	res = response.PatchLampColorEventResponse{
		ID:    re.ID,
		Mac:   re.Mac,
		Red:   re.Red,
		Green: re.Green,
		Blue:  re.Blue,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteLampColorEvent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteLampColorEventRequest
		res           response.DeleteLampColorEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("lamp_event_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteLampColorEvent(ctx, &interactions_pb.DeleteLampColorEventRequest{
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

	res = response.DeleteLampColorEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
