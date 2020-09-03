package servers

import (
	"context"

	"github.com/n7down/kuiper/internal/devices/persistence"

	devices_pb "github.com/n7down/kuiper/internal/pb/devices"
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
	setting := persistence.BatCaveDeviceSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.CreateBatCaveDeviceSetting(setting)

	return &devices_pb.CreateBatCaveDeviceSettingResponse{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) UpdateBatCaveSetting(ctx context.Context, req *devices_pb.UpdateBatCaveDeviceSettingRequest) (*devices_pb.UpdateBatCaveDeviceSettingResponse, error) {
	setting := persistence.BatCaveDeviceSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.UpdateBatCaveDeviceSetting(setting)

	return &devices_pb.UpdateBatCaveDeviceSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) GetBatCaveSetting(ctx context.Context, req *devices_pb.GetBatCaveDeviceSettingRequest) (*devices_pb.GetBatCaveDeviceSettingResponse, error) {
	_, setting := s.persistence.GetBatCaveDeviceSetting(req.DeviceID)
	// recordNotFound, setting := s.persistence.GetBatCaveDeviceSetting(req.DeviceID)
	// if recordNotFound {
	// 	return &devices_pb.GetBatCaveDeviceSettingResponse{}, errors.New("record not found")
	// }

	return &devices_pb.GetBatCaveDeviceSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}
