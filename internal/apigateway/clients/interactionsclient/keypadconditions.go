package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/keypadconditions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/keypadconditions/response"
	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (client InteractionsClient) CreateKeypadCondition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateKeypadConditionRequest
		res response.CreateKeypadConditionResponse
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

	// FIXME: remove interactionID from the service
	r, err := client.interactionsServiceClient.CreateKeypadCondition(ctx, &interactions_pb.CreateKeypadConditionRequest{
		Mac:      req.Mac,
		ButtonID: *req.ButtonID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateKeypadConditionResponse{
		ID:       r.ID,
		Mac:      r.Mac,
		ButtonID: r.ButtonID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetKeypadCondition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetKeypadConditionRequest
		res           response.GetKeypadConditionResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_condition_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetKeypadCondition(ctx, &interactions_pb.GetKeypadConditionRequest{ID: id})
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

	res = response.GetKeypadConditionResponse{
		ID:       r.ID,
		Mac:      r.Mac,
		ButtonID: r.ButtonID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateKeypadCondition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateKeypadConditionRequest
		res           response.UpdateKeypadConditionResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("keypad_condition_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// FIXME: remove the interactionID from the service
	r, err := client.interactionsServiceClient.UpdateKeypadCondition(ctx, &interactions_pb.UpdateKeypadConditionRequest{
		ID:       id,
		Mac:      req.Mac,
		ButtonID: req.ButtonID,
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

	res = response.UpdateKeypadConditionResponse{
		ID:       r.ID,
		Mac:      r.Mac,
		ButtonID: r.ButtonID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchKeypadCondition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchKeypadConditionRequest
		res           response.PatchKeypadConditionResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("keypad_condition_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// FIXME: remove the interactionID from the service
	r, err := client.interactionsServiceClient.GetKeypadCondition(ctx, &interactions_pb.GetKeypadConditionRequest{ID: id})

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

	// FIXME: remove the interactionID from the service
	// save the request difference
	re, err := client.interactionsServiceClient.UpdateKeypadCondition(ctx, &interactions_pb.UpdateKeypadConditionRequest{
		ID:       id,
		Mac:      req.Mac,
		ButtonID: req.ButtonID,
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

	res = response.PatchKeypadConditionResponse{
		ID:       re.ID,
		Mac:      r.Mac,
		ButtonID: r.ButtonID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteKeypadCondition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteKeypadConditionRequest
		res           response.DeleteKeypadConditionResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_condition_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteKeypadCondition(ctx, &interactions_pb.DeleteKeypadConditionRequest{
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

	res = response.DeleteKeypadConditionResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
