package interactions

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions/response"
	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (client InteractionsClient) CreateAttach(c *gin.Context) {
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
		ConditionID: req.ConditionID,
		EventID:     req.EventID,
		EventType:   req.EventType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateAttachResponse{
		ID:          r.ID,
		ConditionID: r.ConditionID,
		EventID:     r.EventID,
		EventType:   r.EventType,
	}

	c.JSON(http.StatusOK, res)
}

func (client InteractionsClient) GetAttach(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	// defer cancel()

	// var (
	// 	req           request.GetInteractionRequest
	// 	res           response.GetInteractionResponse
	// 	errorResponse response.ErrorResponse
	// )

	// id := c.Params.ByName("interaction_id")

	// if validationErrors := req.Validate(id); len(validationErrors) > 0 {
	// 	err := map[string]interface{}{"validationError": validationErrors}
	// 	c.JSON(http.StatusMethodNotAllowed, err)
	// 	return
	// }

	// r, err := client.interactionsServiceClient.GetInteraction(ctx, &interactions_pb.GetInteractionRequest{ID: id})
	// if err != nil {
	// 	st, ok := status.FromError(err)

	// 	// unknown error
	// 	if !ok {
	// 		client.logger.Errorf("unknown error: %v", err)
	// 		errorResponse = response.ErrorResponse{
	// 			Message: fmt.Sprintf("an error has occurred"),
	// 		}
	// 		c.JSON(http.StatusInternalServerError, errorResponse)
	// 		return
	// 	}
	// 	errorResponse = response.ErrorResponse{
	// 		Message: st.Message(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, errorResponse)
	// 	return
	// }

	// if r.ID == "" {
	// 	c.JSON(http.StatusNoContent, res)
	// 	return
	// }

	// res = response.GetInteractionResponse{
	// 	ID:          r.ID,
	// 	Name:        r.Name,
	// 	Description: r.Description,
	// }

	// c.JSON(http.StatusOK, res)
	c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented"})
}

func (client InteractionsClient) UpdateAttach(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	// defer cancel()

	// var (
	// 	req           request.UpdateInteractionRequest
	// 	res           response.UpdateInteractionResponse
	// 	errorResponse response.ErrorResponse
	// )

	// if err := c.BindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id := c.Params.ByName("interaction_id")

	// if validationErrors := req.Validate(id); len(validationErrors) > 0 {
	// 	err := map[string]interface{}{"validationError": validationErrors}
	// 	c.JSON(http.StatusMethodNotAllowed, err)
	// 	return
	// }

	// r, err := client.interactionsServiceClient.UpdateInteraction(ctx, &interactions_pb.UpdateInteractionRequest{
	// 	ID:          id,
	// 	Name:        req.Name,
	// 	Description: req.Description,
	// })

	// if err != nil {
	// 	st, ok := status.FromError(err)

	// 	// unknown error
	// 	if !ok {
	// 		client.logger.Errorf("unknown error: %v", err)
	// 		errorResponse = response.ErrorResponse{
	// 			Message: fmt.Sprintf("an error has occurred"),
	// 		}
	// 		c.JSON(http.StatusInternalServerError, errorResponse)
	// 		return
	// 	}
	// 	errorResponse = response.ErrorResponse{
	// 		Message: st.Message(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, errorResponse)
	// 	return
	// }

	// if r.ID == "" {
	// 	c.JSON(http.StatusNoContent, res)
	// 	return
	// }

	// res = response.UpdateInteractionResponse{
	// 	ID:          r.ID,
	// 	Name:        r.Name,
	// 	Description: r.Description,
	// }

	// c.JSON(http.StatusOK, res)
	c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented"})
}

func (client InteractionsClient) PatchAttach(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	// defer cancel()

	// var (
	// 	req           request.PatchInteractionRequest
	// 	res           response.PatchInteractionResponse
	// 	errorResponse response.ErrorResponse
	// )

	// if err := c.BindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id := c.Params.ByName("interaction_id")

	// if validationErrors := req.Validate(id); len(validationErrors) > 0 {
	// 	err := map[string]interface{}{"validationError": validationErrors}
	// 	c.JSON(http.StatusMethodNotAllowed, err)
	// 	return
	// }

	// // get the user
	// r, err := client.interactionsServiceClient.GetInteraction(ctx, &interactions_pb.GetInteractionRequest{ID: id})

	// if err != nil {
	// 	st, ok := status.FromError(err)

	// 	// unknown error
	// 	if !ok {
	// 		client.logger.Errorf("unknown error: %v", err)
	// 		errorResponse = response.ErrorResponse{
	// 			Message: fmt.Sprintf("an error has occurred"),
	// 		}
	// 		c.JSON(http.StatusInternalServerError, errorResponse)
	// 		return
	// 	}
	// 	errorResponse = response.ErrorResponse{
	// 		Message: st.Message(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, errorResponse)
	// 	return
	// }

	// if r.ID == "" {
	// 	c.JSON(http.StatusNoContent, res)
	// 	return
	// }

	// if req.Name == "" {
	// 	req.Name = r.Name
	// }

	// if req.Description == "" {
	// 	req.Description = r.Description
	// }

	// // save the request difference
	// re, err := client.interactionsServiceClient.UpdateInteraction(ctx, &interactions_pb.UpdateInteractionRequest{
	// 	ID:          id,
	// 	Name:        req.Name,
	// 	Description: req.Description,
	// })

	// if err != nil {
	// 	st, _ := status.FromError(err)
	// 	errorResponse = response.ErrorResponse{
	// 		Message: st.Message(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, errorResponse)
	// 	return
	// }

	// if re.ID == "" {
	// 	c.JSON(http.StatusNoContent, res)
	// 	return
	// }

	// res = response.PatchInteractionResponse{
	// 	ID:          re.ID,
	// 	Name:        re.Name,
	// 	Description: re.Description,
	// }

	// c.JSON(http.StatusOK, res)
	c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented"})
}

func (client InteractionsClient) DeleteAttach(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	// defer cancel()

	// var (
	// 	req           request.DeleteInteractionRequest
	// 	res           response.DeleteInteractionResponse
	// 	errorResponse response.ErrorResponse
	// )

	// id := c.Params.ByName("interaction_id")

	// if validationErrors := req.Validate(id); len(validationErrors) > 0 {
	// 	err := map[string]interface{}{"validationError": validationErrors}
	// 	c.JSON(http.StatusMethodNotAllowed, err)
	// 	return
	// }

	// r, err := client.interactionsServiceClient.DeleteInteraction(ctx, &interactions_pb.DeleteInteractionRequest{
	// 	ID: id,
	// })

	// if err != nil {
	// 	st, ok := status.FromError(err)

	// 	// unknown error
	// 	if !ok {
	// 		client.logger.Errorf("unknown error: %v", err)
	// 		errorResponse = response.ErrorResponse{
	// 			Message: fmt.Sprintf("an error has occurred"),
	// 		}
	// 		c.JSON(http.StatusInternalServerError, errorResponse)
	// 		return
	// 	}
	// 	errorResponse = response.ErrorResponse{
	// 		Message: st.Message(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, errorResponse)
	// 	return
	// }

	// if r.ID == "" {
	// 	c.JSON(http.StatusNoContent, res)
	// 	return
	// }

	// res = response.DeleteInteractionResponse{
	// 	ID: r.ID,
	// }

	// c.JSON(http.StatusOK, res)
	c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented"})
}
