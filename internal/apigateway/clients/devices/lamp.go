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

func (client *DevicesClient) SendLampDeviceOn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDeviceOnRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceOn(ctx, &devices_pb.SendLampDeviceOnRequest{
		Mac: mac,
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
		req           request.SendLampDeviceOffRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceToggle(ctx, &devices_pb.SendLampDeviceToggleRequest{
		Mac: mac,
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

func (client *DevicesClient) SendLampDeviceToggle(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDeviceToggleRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceToggle(ctx, &devices_pb.SendLampDeviceToggleRequest{
		Mac: mac,
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
		req           request.SendLampDeviceColorRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceColor(ctx, &devices_pb.SendLampDeviceColorRequest{
		Mac:   mac,
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
		req           request.SendLampDeviceBrightnessRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceBrightness(ctx, &devices_pb.SendLampDeviceBrightnessRequest{
		Mac:        mac,
		Brightness: *req.Brightness,
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

func (client *DevicesClient) SendLampDeviceAutoBrightness(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDeviceAutoBrightnessRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDeviceAutoBrightness(ctx, &devices_pb.SendLampDeviceAutoBrightnessRequest{
		Mac: mac,
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

func (client *DevicesClient) SendLampDevicePulse(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.SendLampDevicePulseRequest
		errorResponse response.ErrorResponse
	)

	mac := c.Params.ByName("send_lamp_mac")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(mac); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	_, err := client.devicesClient.SendLampDevicePulse(ctx, &devices_pb.SendLampDevicePulseRequest{
		Mac:   mac,
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
