package devices

import (
	"context"
	"net/http"
	"strings"
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

// swagger:route POST /api/v1/devices/bc Devices createBatCaveDeviceSetting
//
// Create Bat Cave Device Setting
//
// Allows a Bat Cave device setting to be created.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Responses:
//  200: CreateBatCaveDeviceSettingResponse
//  400: description: Unable to bind request.
//  405: description: Validation error.
//  500: description: Error with the service.
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

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.deviceSettingsClient.CreateBatCaveDeviceSetting(ctx, &devices_pb.CreateBatCaveDeviceSettingRequest{DeviceID: req.DeviceID, DeepSleepDelay: req.DeepSleepDelay})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateBatCaveDeviceSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/devices/bc/:device_id Devices getBatCaveDeviceSetting
//
// Get Bat Cave Device Setting
//
// Get a Bat Cave device setting.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Responses:
//  200: GetBatCaveDeviceSettingResponse
//  204: description: No content.
//  405: description: Validation error.
//  500: description: Error with the service.
func (client *DevicesClient) GetBatCaveDeviceSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.GetBatCaveDeviceSettingRequest
		res response.GetBatCaveDeviceSettingResponse
	)

	deviceID := c.Params.ByName("device_id")

	req = request.GetBatCaveDeviceSettingRequest{
		DeviceID: deviceID,
	}

	req.DeviceID = strings.ToLower(req.DeviceID)

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.deviceSettingsClient.GetBatCaveDeviceSetting(ctx, &devices_pb.GetBatCaveDeviceSettingRequest{DeviceID: req.DeviceID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if r.DeviceID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetBatCaveDeviceSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/devices/bc/:device_id Devices updateBatCaveDeviceSetting
//
// Update Bat Cave Device Setting
//
// Update a Bat Cave device setting.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Responses:
//  200: UpdateBatCaveDeviceSettingResponse
//  204: description: No content.
//  400: description: Unable to bind request.
//  405: description: Validation error.
//  500: description: Error with the service.
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

	deviceID := c.Params.ByName("device_id")

	req = request.UpdateBatCaveDeviceSettingRequest{
		DeepSleepDelay: req.DeepSleepDelay,
	}

	deviceID = strings.ToLower(deviceID)

	if validationErrors := req.Validate(deviceID); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.deviceSettingsClient.UpdateBatCaveDeviceSetting(ctx, &devices_pb.UpdateBatCaveDeviceSettingRequest{
		DeviceID:       deviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if r.DeviceID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.UpdateBatCaveDeviceSettingResponse{
		DeviceID:       r.DeviceID,
		DeepSleepDelay: r.DeepSleepDelay,
	}

	c.JSON(http.StatusOK, res)
}
