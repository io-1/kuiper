package users

import (
	"time"

	"google.golang.org/grpc"

	users_pb "github.com/n7down/kuiper/internal/pb/users"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type UsersClient struct {
	usersClient users_pb.UsersServiceClient
}

func NewUsersClient(serverEnv string) (*UsersClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &UsersClient{
		usersClient: users_pb.NewUsersServiceClient(conn),
	}
	return client, nil
}

func NewDevicesClientWithMock(usersClient users_pb.UsersServiceClient) *UsersClient {
	client := &UsersClient{
		usersClient: usersClient,
	}
	return client
}

// func (client *DevicesClient) CreateBatCaveSetting(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
// 	defer cancel()

// 	var (
// 		req request.CreateBatCaveSettingRequest
// 		res response.CreateBatCaveSettingResponse
// 	)

// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	req.DeviceID = strings.ToLower(req.DeviceID)

// 	if validationErrors := req.Validate(); len(validationErrors) > 0 {
// 		err := map[string]interface{}{"validationError": validationErrors}
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	r, err := client.settingsClient.CreateBatCaveSetting(ctx, &devices_pb.CreateBatCaveSettingRequest{DeviceID: req.DeviceID, DeepSleepDelay: req.DeepSleepDelay})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	res = response.CreateBatCaveSettingResponse{
// 		DeviceID:       r.DeviceID,
// 		DeepSleepDelay: r.DeepSleepDelay,
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (client *DevicesClient) GetBatCaveSetting(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
// 	defer cancel()

// 	var (
// 		req request.GetBatCaveSettingRequest
// 		res response.GetBatCaveSettingResponse
// 	)

// 	deviceID := c.Params.ByName("device_id")

// 	req = request.GetBatCaveSettingRequest{
// 		DeviceID: deviceID,
// 	}

// 	req.DeviceID = strings.ToLower(req.DeviceID)

// 	if validationErrors := req.Validate(); len(validationErrors) > 0 {
// 		err := map[string]interface{}{"validationError": validationErrors}
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	r, err := client.settingsClient.GetBatCaveSetting(ctx, &devices_pb.GetBatCaveSettingRequest{DeviceID: req.DeviceID})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	if r.DeviceID == "" {
// 		c.JSON(http.StatusNoContent, res)
// 		return
// 	}

// 	res = response.GetBatCaveSettingResponse{
// 		DeviceID:       r.DeviceID,
// 		DeepSleepDelay: r.DeepSleepDelay,
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (client *DevicesClient) UpdateBatCaveSetting(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
// 	defer cancel()

// 	var (
// 		req request.UpdateBatCaveSettingRequest
// 		res response.UpdateBatCaveSettingResponse
// 	)

// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	deviceID := c.Params.ByName("device_id")

// 	req = request.UpdateBatCaveSettingRequest{
// 		DeviceID:       deviceID,
// 		DeepSleepDelay: req.DeepSleepDelay,
// 	}

// 	req.DeviceID = strings.ToLower(req.DeviceID)

// 	if validationErrors := req.Validate(); len(validationErrors) > 0 {
// 		err := map[string]interface{}{"validationError": validationErrors}
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	r, err := client.settingsClient.UpdateBatCaveSetting(ctx, &devices_pb.UpdateBatCaveSettingRequest{
// 		DeviceID:       req.DeviceID,
// 		DeepSleepDelay: req.DeepSleepDelay,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	if r.DeviceID == "" {
// 		c.JSON(http.StatusNoContent, res)
// 		return
// 	}

// 	res = response.UpdateBatCaveSettingResponse{
// 		DeviceID:       r.DeviceID,
// 		DeepSleepDelay: r.DeepSleepDelay,
// 	}

// 	c.JSON(http.StatusOK, res)
// }
