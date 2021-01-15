package interactionsclient

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/keypadconditionstolampevents/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient/keypadconditionstolampevents/response"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (client InteractionsClient) CreateKeypadConditionToLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (

		// FIXME: change to CreateKeypadConditionToLampEventInteractionRequest
		req request.CreateKeypadConditionToLampEventRequest
		res response.CreateKeypadConditionToLampEventResponse
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

	r, err := client.interactionsServiceClient.CreateKeypadConditionToLampEvent(ctx, &interactions_pb.CreateKeypadConditionToLampEventRequest{
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateKeypadConditionToLampEventResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetKeypadConditionToLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetKeypadConditionToLampEventRequest
		res           response.GetKeypadConditionToLampEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetKeypadConditionToLampEvent(ctx, &interactions_pb.GetKeypadConditionToLampEventRequest{ID: id})
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

	res = response.GetKeypadConditionToLampEventResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateKeypadConditionToLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateKeypadConditionToLampEventRequest
		res           response.UpdateKeypadConditionToLampEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.UpdateKeypadConditionToLampEvent(ctx, &interactions_pb.UpdateKeypadConditionToLampEventRequest{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
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

	res = response.UpdateKeypadConditionToLampEventResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchKeypadConditionToLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchKeypadConditionToLampEventRequest
		res           response.PatchKeypadConditionToLampEventResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.interactionsServiceClient.GetKeypadConditionToLampEvent(ctx, &interactions_pb.GetKeypadConditionToLampEventRequest{ID: id})

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

	if req.InteractionID == "" {
		req.InteractionID = r.InteractionID
	}

	if req.ConditionID == "" {
		req.ConditionID = r.ConditionID
	}

	if req.EventID == "" {
		req.EventID = r.EventID
	}

	// save the request difference
	re, err := client.interactionsServiceClient.UpdateKeypadConditionToLampEvent(ctx, &interactions_pb.UpdateKeypadConditionToLampEventRequest{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
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

	res = response.PatchKeypadConditionToLampEventResponse{
		ID:            re.ID,
		InteractionID: re.InteractionID,
		ConditionID:   re.ConditionID,
		EventID:       re.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteKeypadConditionToLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteKeypadConditionToLampEventRequest
		res           response.DeleteKeypadConditionToLampEventResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteKeypadConditionToLampEvent(ctx, &interactions_pb.DeleteKeypadConditionToLampEventRequest{
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

	res = response.DeleteKeypadConditionToLampEventResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
