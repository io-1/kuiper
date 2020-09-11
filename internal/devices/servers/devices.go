package servers

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/devices/persistence"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

type DevicesServer struct {
	persistence persistence.Persistence
	devices_pb.UnimplementedDevicesServiceServer
}

func NewDevicesServer(persistence persistence.Persistence) *DevicesServer {
	return &DevicesServer{
		persistence: persistence,
	}
}

func (s *DevicesServer) CreateBatCaveDeviceSetting(ctx context.Context, req *devices_pb.CreateBatCaveDeviceSettingRequest) (*devices_pb.CreateBatCaveDeviceSettingResponse, error) {

	id := uuid.New().String()

	setting := persistence.BatCaveDeviceSetting{
		ID:             id,
		Mac:            req.Mac,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.CreateBatCaveDeviceSetting(setting)

	return &devices_pb.CreateBatCaveDeviceSettingResponse{
		ID:             setting.ID,
		Mac:            setting.Mac,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) GetBatCaveDeviceSetting(ctx context.Context, req *devices_pb.GetBatCaveDeviceSettingRequest) (*devices_pb.GetBatCaveDeviceSettingResponse, error) {
	_, setting := s.persistence.GetBatCaveDeviceSetting(req.ID)

	return &devices_pb.GetBatCaveDeviceSettingResponse{
		ID:             setting.ID,
		Mac:            setting.Mac,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) UpdateBatCaveDeviceSetting(ctx context.Context, req *devices_pb.UpdateBatCaveDeviceSettingRequest) (*devices_pb.UpdateBatCaveDeviceSettingResponse, error) {
	setting := persistence.BatCaveDeviceSetting{
		ID:             req.ID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.UpdateBatCaveDeviceSetting(setting)

	return &devices_pb.UpdateBatCaveDeviceSettingResponse{
		ID:             setting.ID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}
