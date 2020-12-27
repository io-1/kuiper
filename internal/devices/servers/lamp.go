package servers

import (
	"context"
	"encoding/json"
	"fmt"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

type LampPulseSetting struct {
	EventType string `json:"e"`
	Red       int32  `json:"r"`
	Green     int32  `json:"g"`
	Blue      int32  `json:'b'`
}

func (s *DevicesServer) SendLampDevicePulseSetting(ctx context.Context, req *devices_pb.SendLampDevicePulseSettingRequest) (*devices_pb.SendLampDevicePulseSettingResponse, error) {

	l := LampPulseSetting{
		EventType: "pulse",
		Red:       req.Red,
		Green:     req.Green,
		Blue:      req.Blue,
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDevicePulseSettingResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)

	return &devices_pb.SendLampDevicePulseSettingResponse{}, nil
}
