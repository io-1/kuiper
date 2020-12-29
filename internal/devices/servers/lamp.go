package servers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/io-1/kuiper/internal/devices/servers/response"

	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

func (s *DevicesServer) SendLampDeviceOn(ctx context.Context, req *devices_pb.SendLampDeviceOnRequest) (*devices_pb.SendLampDeviceOnResponse, error) {

	l := response.LampDeviceOnResponse{
		EventType: "on",
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceOnResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceOnResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceOff(ctx context.Context, req *devices_pb.SendLampDeviceOffRequest) (*devices_pb.SendLampDeviceOffResponse, error) {

	l := response.LampDeviceOffResponse{
		EventType: "off",
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceOffResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceOffResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceToggle(ctx context.Context, req *devices_pb.SendLampDeviceToggleRequest) (*devices_pb.SendLampDeviceToggleResponse, error) {

	l := response.LampDeviceToggleResponse{
		EventType: "toggle",
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceToggleResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceToggleResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceColor(ctx context.Context, req *devices_pb.SendLampDeviceColorRequest) (*devices_pb.SendLampDeviceColorResponse, error) {

	l := response.LampDeviceColorResponse{
		EventType: "color",
		Red:       req.Red,
		Green:     req.Green,
		Blue:      req.Blue,
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceColorResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceColorResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceBrightness(ctx context.Context, req *devices_pb.SendLampDeviceBrightnessRequest) (*devices_pb.SendLampDeviceBrightnessResponse, error) {

	l := response.LampDeviceBrightnessResponse{
		EventType:  "brightness",
		Brightness: req.Brightness,
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceBrightnessResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceBrightnessResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceAutoBrightness(ctx context.Context, req *devices_pb.SendLampDeviceAutoBrightnessRequest) (*devices_pb.SendLampDeviceAutoBrightnessResponse, error) {

	l := response.LampDeviceAutoBrightnessResponse{
		EventType: "auto-brightness",
	}

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceAutoBrightnessResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceAutoBrightnessResponse{}, nil
}

func (s *DevicesServer) SendLampDevicePulse(ctx context.Context, req *devices_pb.SendLampDevicePulseRequest) (*devices_pb.SendLampDevicePulseResponse, error) {

	l := response.LampDevicePulseResponse{
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
