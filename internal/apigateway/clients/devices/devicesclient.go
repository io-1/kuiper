package devices

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/devices/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/devices/response"
	"google.golang.org/grpc"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type DevicesClient struct {
	deviceSettingsClient devices_pb.DevicesServiceClient
}

func NewDevicesClient(serverEnv string) (*DevicesClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &DevicesClient{
		deviceSettingsClient: devices_pb.NewDevicesServiceClient(conn),
	}
	return client, nil
}

func NewDevicesClientWithMock(mockSettingsServiceClient devices_pb.DevicesServiceClient) *DevicesClient {
	client := &DevicesClient{
		deviceSettingsClient: mockSettingsServiceClient,
	}
	return client
}

// Create Bat Cave Device Setting
func (client *DevicesClient) CreateBatCaveDeviceSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateBatCaveDeviceSettingRequest
		res response.CreateBatCaveDeviceSettingResponse
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

	r, err := client.deviceSettingsClient.CreateBatCaveDeviceSetting(ctx, &devices_pb.CreateBatCaveDeviceSettingRequest{Mac: req.Mac, DeepSleepDelay: req.DeepSleepDelay})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		req request.GetBatCaveDeviceSettingRequest
		res response.GetBatCaveDeviceSettingResponse
	)

	id := c.Params.ByName("id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.deviceSettingsClient.GetBatCaveDeviceSetting(ctx, &devices_pb.GetBatCaveDeviceSettingRequest{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		req request.UpdateBatCaveDeviceSettingRequest
		res response.UpdateBatCaveDeviceSettingResponse
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

	r, err := client.deviceSettingsClient.UpdateBatCaveDeviceSetting(ctx, &devices_pb.UpdateBatCaveDeviceSettingRequest{
		ID:             id,
		DeepSleepDelay: req.DeepSleepDelay,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
