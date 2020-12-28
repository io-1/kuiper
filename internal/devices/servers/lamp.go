package servers

import (
	"context"
	"encoding/json"
	"fmt"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

type LampDevicePulseResponse struct {
	EventType string `json:"e"`
	Red       int32  `json:"r"`
	Green     int32  `json:"g"`
	Blue      int32  `json:"b"`
}

func (s *DevicesServer) SendLampDevicePulse(ctx context.Context, req *devices_pb.SendLampDevicePulseRequest) (*devices_pb.SendLampDevicePulseResponse, error) {

	l := LampDevicePulseResponse{
		EventType: "pulse",
		Red:       req.Red,
		Green:     req.Green,
		Blue:      req.Blue,
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDevicePulseResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)

	return &devices_pb.SendLampDevicePulseResponse{}, nil
}
