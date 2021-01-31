// +build unit,!integration

package devicesclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/io-1/kuiper/internal/logger/blanklogger"
	"github.com/io-1/kuiper/pkg/mock"
	"github.com/stretchr/testify/assert"

	devices_pb "github.com/io-1/kuiper/pkg/pb/devices"
)

func Test_GetBatCaveDeviceSetting_Should_Return_StatusNoContent_When_ID_Is_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		expectedCode        = http.StatusNoContent
		expectedRes         = ""
		err          error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	logger := blanklogger.NewBlankLogger()

	mockCtrl := gomock.NewController(t)
	mockDevicesServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	devicesClient := NewDevicesClientWithMock(mockDevicesServiceClient, logger)

	mockDevicesServiceClient.EXPECT().GetBatCaveDeviceSetting(
		gomock.Any(),

		// FIXME: not sure what is going on here
		// &devices_pb.GetBatCaveDeviceSettingRequest{
		// 	ID: id,
		// },
		gomock.Any(),
	).Return(
		&devices_pb.GetBatCaveDeviceSettingResponse{}, nil,
	)

	r.GET("/bc/:id", devicesClient.GetBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", id)
	c.Request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_UpdateBatCaveDeviceSetting_Should_Return_StatusNoContent_When_ID_Is_Empty(t *testing.T) {
	var (
		id             string = "00000000-1111-2222-3333-444444444444"
		deepSleepDelay uint32 = 15
		reqParam              = fmt.Sprintf(`{"deepSleepDelay":%d}`, deepSleepDelay)
		expectedCode          = http.StatusNoContent
		expectedRes           = ""
		err            error
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	logger := blanklogger.NewBlankLogger()

	mockCtrl := gomock.NewController(t)
	mockSettingsServiceClient := mock.NewMockDevicesServiceClient(mockCtrl)

	devicesClient := NewDevicesClientWithMock(mockSettingsServiceClient, logger)

	mockSettingsServiceClient.EXPECT().UpdateBatCaveDeviceSetting(
		gomock.Any(),

		// FIXME: not sure what is going on here
		// &devices_pb.UpdateBatCaveDeviceSettingRequest{
		// 	ID:             id,
		// 	DeepSleepDelay: deepSleepDelay,
		// },
		gomock.Any(),
	).Return(
		&devices_pb.UpdateBatCaveDeviceSettingResponse{}, nil,
	)

	r.PUT("/bc/:id", devicesClient.UpdateBatCaveDeviceSetting)

	url := fmt.Sprintf("/bc/%s", id)
	c.Request, err = http.NewRequest("PUT", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}
