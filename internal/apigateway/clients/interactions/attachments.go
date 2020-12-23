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

func (client InteractionsClient) CreateKeypadConditionLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateAttachRequest
		res response.CreateAttachResponse
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

	r, err := client.interactionsServiceClient.CreateAttach(ctx, &interactions_pb.CreateAttachRequest{
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateAttachResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetKeypadConditionLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetAttachRequest
		res           response.GetAttachResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.GetAttach(ctx, &interactions_pb.GetAttachRequest{ID: id})
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

	res = response.GetAttachResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) UpdateKeypadConditionLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateAttachRequest
		res           response.UpdateAttachResponse
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

	r, err := client.interactionsServiceClient.UpdateAttach(ctx, &interactions_pb.UpdateAttachRequest{
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

	res = response.UpdateAttachResponse{
		ID:            r.ID,
		InteractionID: r.InteractionID,
		ConditionID:   r.ConditionID,
		EventID:       r.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) PatchKeypadConditionLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchAttachRequest
		res           response.PatchAttachResponse
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
	r, err := client.interactionsServiceClient.GetAttach(ctx, &interactions_pb.GetAttachRequest{ID: id})

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
	re, err := client.interactionsServiceClient.UpdateAttach(ctx, &interactions_pb.UpdateAttachRequest{
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

	res = response.PatchAttachResponse{
		ID:            re.ID,
		InteractionID: re.InteractionID,
		ConditionID:   re.ConditionID,
		EventID:       re.EventID,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) DeleteKeypadConditionLampEventInteraction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteInteractionRequest
		res           response.DeleteInteractionResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("keypad_to_lamp_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.interactionsServiceClient.DeleteInteraction(ctx, &interactions_pb.DeleteInteractionRequest{
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

	res = response.DeleteInteractionResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
