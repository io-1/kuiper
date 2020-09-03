package servers

import (
	"context"
	"errors"

	"github.com/n7down/kuiper/internal/devices/persistence"

	devices_pb "github.com/n7down/kuiper/internal/pb/devices"
)

type DevicesServer struct {
	persistence persistence.Persistence
	devices_pb.UnimplementedSettingsServiceServer
}

func NewDevicesServer(persistence persistence.Persistence) *DevicesServer {
	return &DevicesServer{
		persistence: persistence,
	}
}

func (s *DevicesServer) CreateBatCaveSetting(ctx context.Context, req *devices_pb.CreateBatCaveSettingRequest) (*devices_pb.CreateBatCaveSettingResponse, error) {
	setting := persistence.BatCaveSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.CreateBatCaveSetting(setting)

	return &devices_pb.CreateBatCaveSettingResponse{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) UpdateBatCaveSetting(ctx context.Context, req *devices_pb.UpdateBatCaveSettingRequest) (*devices_pb.UpdateBatCaveSettingResponse, error) {
	setting := persistence.BatCaveSetting{
		DeviceID:       req.DeviceID,
		DeepSleepDelay: req.DeepSleepDelay,
	}

	s.persistence.UpdateBatCaveSetting(setting)

	return &devices_pb.UpdateBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}

func (s *DevicesServer) GetBatCaveSetting(ctx context.Context, req *devices_pb.GetBatCaveSettingRequest) (*devices_pb.GetBatCaveSettingResponse, error) {
	recordNotFound, setting := s.persistence.GetBatCaveSetting(req.DeviceID)
	if recordNotFound {
		return &devices_pb.GetBatCaveSettingResponse{}, errors.New("record not found")
	}

	return &devices_pb.GetBatCaveSettingResponse{
		DeviceID:       setting.DeviceID,
		DeepSleepDelay: setting.DeepSleepDelay,
	}, nil
}
