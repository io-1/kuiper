package devicesclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/apigateway/clients/devicesclient/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/devicesclient/response"

	devices_pb "github.com/io-1/kuiper/pkg/pb/devices"
)

// Create Bat Cave Device Setting
func (client *DevicesClient) CreateBatCaveDeviceSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.CreateBatCaveDeviceSettingRequest
		res           response.CreateBatCaveDeviceSettingResponse
		errorResponse response.ErrorResponse
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

	r, err := client.devicesClient.CreateBatCaveDeviceSetting(ctx, &devices_pb.CreateBatCaveDeviceSettingRequest{Mac: req.Mac, DeepSleepDelay: req.DeepSleepDelay})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	res = response.CreateBatCaveDeviceSettingResponse{
		ID:             r.ID,
		Mac:            r.Mac,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

// Get Bat Cave Device Setting
func (client *DevicesClient) GetBatCaveDeviceSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetBatCaveDeviceSettingRequest
		res           response.GetBatCaveDeviceSettingResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.devicesClient.GetBatCaveDeviceSetting(ctx, &devices_pb.GetBatCaveDeviceSettingRequest{ID: id})
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

	res = response.GetBatCaveDeviceSettingResponse{
		ID:             r.ID,
		Mac:            r.Mac,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

// Update Bat Cave Device Setting
func (client *DevicesClient) UpdateBatCaveDeviceSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateBatCaveDeviceSettingRequest
		res           response.UpdateBatCaveDeviceSettingResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.devicesClient.UpdateBatCaveDeviceSetting(ctx, &devices_pb.UpdateBatCaveDeviceSettingRequest{
		ID:             id,
		DeepSleepDelay: req.DeepSleepDelay,
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

	res = response.UpdateBatCaveDeviceSettingResponse{
		ID:             r.ID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}
