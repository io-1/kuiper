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
	"github.com/n7down/kuiper/internal/mock"
	"github.com/stretchr/testify/assert"

	devices_pb "github.com/n7down/kuiper/internal/pb/devices"
)

func Test_CreateBatCaveSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
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
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	devicesClient := NewDevicesClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().CreateBatCaveSetting(
		gomock.Any(),
		&devices_pb.CreateBatCaveSettingRequest{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.CreateBatCaveSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.POST("/bc", devicesClient.CreateBatCaveSetting)

	c.Request, err = http.NewRequest("POST", "/bc", strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_GetBatCaveSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
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
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	settingsClient := NewDevicesClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().GetBatCaveSetting(
		gomock.Any(),
		&devices_pb.GetBatCaveSettingRequest{
			DeviceID: deviceIDLowerCase,
		},
	).Return(
		&devices_pb.GetBatCaveSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.GET("/bc/:device_id", settingsClient.GetBatCaveSetting)

	url := fmt.Sprintf("/bc/%s", deviceIDUpperCase)
	c.Request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_GetBatCaveSetting_Should_Return_StatusNoContent_When_DeviceID_Is_Empty(t *testing.T) {
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
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	settingsClient := NewDevicesClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().GetBatCaveSetting(
		gomock.Any(),
		&devices_pb.GetBatCaveSettingRequest{
			DeviceID: deviceID,
		},
	).Return(
		&devices_pb.GetBatCaveSettingResponse{}, nil,
	)

	r.GET("/bc/:device_id", settingsClient.GetBatCaveSetting)

	url := fmt.Sprintf("/bc/%s", deviceID)
	c.Request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_UpdateBatCaveSetting_Should_Change_DeviceID_To_Lower_Case_When_DeviceID_Has_Upper_Case_Characters_In_Request(t *testing.T) {
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
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	settingsClient := NewDevicesClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().UpdateBatCaveSetting(
		gomock.Any(),
		&devices_pb.UpdateBatCaveSettingRequest{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.UpdateBatCaveSettingResponse{
			DeviceID:       deviceIDLowerCase,
			DeepSleepDelay: deepSleepDelay,
		}, nil,
	)

	r.PUT("/bc/:device_id", settingsClient.UpdateBatCaveSetting)

	url := fmt.Sprintf("/bc/%s", deviceIDUpperCase)
	c.Request, err = http.NewRequest("PUT", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_UpdateBatCaveSetting_Should_Return_StatusNoContent_When_DeviceID_Is_Empty(t *testing.T) {
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
	mockSettingsServiceClient := mock.NewMockSettingsServiceClient(mockCtrl)

	settingsClient := NewDevicesClientWithMock(mockSettingsServiceClient)

	mockSettingsServiceClient.EXPECT().UpdateBatCaveSetting(
		gomock.Any(),
		&devices_pb.UpdateBatCaveSettingRequest{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
		},
	).Return(
		&devices_pb.UpdateBatCaveSettingResponse{}, nil,
	)

	r.PUT("/bc/:device_id", settingsClient.UpdateBatCaveSetting)

	url := fmt.Sprintf("/bc/%s", deviceID)
	c.Request, err = http.NewRequest("PUT", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}
