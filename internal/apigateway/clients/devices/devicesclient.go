package devices

import (
	"time"

	"github.com/io-1/kuiper/internal/logger"
	"google.golang.org/grpc"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type DevicesClient struct {
	logger        logger.Logger
	devicesClient devices_pb.DevicesServiceClient
}

func NewDevicesClient(serverEnv string, logger logger.Logger) (*DevicesClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &DevicesClient{
		logger:        logger,
		devicesClient: devices_pb.NewDevicesServiceClient(conn),
	}
	return client, nil
}

func NewDevicesClientWithMock(mockSettingsServiceClient devices_pb.DevicesServiceClient, logger logger.Logger) *DevicesClient {
	client := &DevicesClient{
		logger:        logger,
		devicesClient: mockSettingsServiceClient,
	}
	return client
}
