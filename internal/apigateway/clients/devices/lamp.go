package devices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/devices/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/devices/response"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

func (client *DevicesClient) SendLampDevicePulse(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
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

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (client *DevicesClient) SendLampDeviceOn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
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

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (client *DevicesClient) SendLampDeviceOff(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
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

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (client *DevicesClient) SendLampDeviceColor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
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

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (client *DevicesClient) SendLampDeviceBrightness(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
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

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   req.Mac,
		Red:   *req.Red,
		Green: *req.Green,
		Blue:  *req.Blue,
	})
	if err != nil {
		client.logger.Errorf("unknown error: %v", err)
		errorResponse = response.ErrorResponse{
			Message: fmt.Sprintf("an error has occurred"),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}
