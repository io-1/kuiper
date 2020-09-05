// +build unit,!integration

package devices

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/n7down/kuiper/internal/logger/blanklogger"
	"github.com/n7down/kuiper/internal/mock"
	"github.com/stretchr/testify/assert"

	devices_pb "github.com/n7down/kuiper/internal/pb/devices"
)

func Test_CreateBatCaveDeviceSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
	var (
		deviceIDUpperCase string = "0011001100FF"
		deviceIDLowerCase string = "0011001100ff"
		deepSleepDelay    uint32 = 15
		expectedCode             = http.StatusOK
		reqParam                 = fmt.Sprintf(`{"deviceID":"%s","deepSleepDelay":%d}`, deviceIDUpperCase, deepSleepDelay)
		expectedRes              = fmt.Sprintf(`{"deviceID":"%s","deepSleepDelay":%d}`, deviceIDLowerCase, deepSleepDelay)
		err               error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockDevicesServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	logger := blanklogger.NewBlankLogger()
	devicesClient := NewDevicesClientWithMock(mockDevicesServiceClient, logger)

	mockDevicesServiceClient.EXPECT().CreateBatCaveDeviceSetting(
		gomock.Any(),
		&devices_pb.CreateBatCaveDeviceSettingRequest{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.CreateBatCaveDeviceSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.POST("/bc", devicesClient.CreateBatCaveDeviceSetting)

	c.Request, err = http.NewRequest("POST", "/bc", strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_GetBatCaveDeviceSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
	var (
		deviceIDUpperCase string = "0011001100FF"
		deviceIDLowerCase string = "0011001100ff"
		deepSleepDelay    uint32 = 15
		expectedCode             = http.StatusOK
		expectedRes              = fmt.Sprintf(`{"deviceID":"%s","deepSleepDelay":%d}`, deviceIDLowerCase, deepSleepDelay)
		err               error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockSettingsServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	logger := blanklogger.NewBlankLogger()
	devicesClient := NewDevicesClientWithMock(mockSettingsServiceClient, logger)

	mockSettingsServiceClient.EXPECT().GetBatCaveDeviceSetting(
		gomock.Any(),
		&devices_pb.GetBatCaveDeviceSettingRequest{
			DeviceID: deviceIDLowerCase,
		},
	).Return(
		&devices_pb.GetBatCaveDeviceSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.GET("/bc/:device_id", devicesClient.GetBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", deviceIDUpperCase)
	c.Request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_GetBatCaveDeviceSetting_Should_Return_StatusNoContent_When_DeviceID_Is_Empty(t *testing.T) {
	var (
		deviceID     string = "0011001100ff"
		expectedCode        = http.StatusNoContent
		expectedRes         = ""
		err          error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockDevicesServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	logger := blanklogger.NewBlankLogger()
	devicesClient := NewDevicesClientWithMock(mockDevicesServiceClient, logger)

	mockDevicesServiceClient.EXPECT().GetBatCaveDeviceSetting(
		gomock.Any(),
		&devices_pb.GetBatCaveDeviceSettingRequest{
			DeviceID: deviceID,
		},
	).Return(
		&devices_pb.GetBatCaveDeviceSettingResponse{}, nil,
	)

	r.GET("/bc/:device_id", devicesClient.GetBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", deviceID)
	c.Request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_UpdateBatCaveDeviceSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
	var (
		deviceIDUpperCase string = "0011001100FF"
		deviceIDLowerCase string = "0011001100ff"
		deepSleepDelay    uint32 = 15
		reqParam                 = fmt.Sprintf(`{"deepSleepDelay":%d}`, deepSleepDelay)
		expectedCode             = http.StatusOK
		expectedRes              = fmt.Sprintf(`{"deviceID":"%s","deepSleepDelay":%d}`, deviceIDLowerCase, deepSleepDelay)
		err               error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockSettingsServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	logger := blanklogger.NewBlankLogger()
	devicesClient := NewDevicesClientWithMock(mockSettingsServiceClient, logger)

	mockSettingsServiceClient.EXPECT().UpdateBatCaveDeviceSetting(
		gomock.Any(),
		&devices_pb.UpdateBatCaveDeviceSettingRequest{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.UpdateBatCaveDeviceSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.PUT("/bc/:device_id", devicesClient.UpdateBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", deviceIDUpperCase)
	c.Request, err = http.NewRequest("PUT", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_UpdateBatCaveDeviceSetting_Should_Return_StatusNoContent_When_DeviceID_Is_Empty(t *testing.T) {
	var (
		deviceID       string = "0011001100ff"
		deepSleepDelay uint32 = 15
		reqParam              = fmt.Sprintf(`{"deepSleepDelay":%d}`, deepSleepDelay)
		expectedCode          = http.StatusNoContent
		expectedRes           = ""
		err            error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockSettingsServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	logger := blanklogger.NewBlankLogger()
	devicesClient := NewDevicesClientWithMock(mockSettingsServiceClient, logger)

	mockSettingsServiceClient.EXPECT().UpdateBatCaveDeviceSetting(
		gomock.Any(),
		&devices_pb.UpdateBatCaveDeviceSettingRequest{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.UpdateBatCaveDeviceSettingResponse{}, nil,
	)

	r.PUT("/bc/:device_id", devicesClient.UpdateBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", deviceID)
	c.Request, err = http.NewRequest("PUT", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}
