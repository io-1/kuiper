package servers

import (
	"context"
	"encoding/json"
	"fmt"

	lamp_events "github.com/io-1/kuiper/internal/events/lamp"
	devices_pb "github.com/io-1/kuiper/internal/pb/devices"
)

func (s *DevicesServer) SendLampDeviceOn(ctx context.Context, req *devices_pb.SendLampDeviceOnRequest) (*devices_pb.SendLampDeviceOnResponse, error) {
	l := lamp_events.NewLampDeviceOnEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceOnResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceOnResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceOff(ctx context.Context, req *devices_pb.SendLampDeviceOffRequest) (*devices_pb.SendLampDeviceOffResponse, error) {
	l := lamp_events.NewLampDeviceOffEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceOffResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceOffResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceToggle(ctx context.Context, req *devices_pb.SendLampDeviceToggleRequest) (*devices_pb.SendLampDeviceToggleResponse, error) {
	l := lamp_events.NewLampDeviceToggleEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceToggleResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceToggleResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceColor(ctx context.Context, req *devices_pb.SendLampDeviceColorRequest) (*devices_pb.SendLampDeviceColorResponse, error) {
	l := lamp_events.NewLampDeviceColorEvent(req.Red, req.Green, req.Blue)

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceColorResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceColorResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceBrightness(ctx context.Context, req *devices_pb.SendLampDeviceBrightnessRequest) (*devices_pb.SendLampDeviceBrightnessResponse, error) {
	l := lamp_events.NewLampDeviceBrightnessEvent(req.Brightness)

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceBrightnessResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceBrightnessResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceAutoBrightnessOn(ctx context.Context, req *devices_pb.SendLampDeviceAutoBrightnessOnRequest) (*devices_pb.SendLampDeviceAutoBrightnessOnResponse, error) {
	l := lamp_events.NewLampDeviceAutoBrightnessOnEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceAutoBrightnessOnResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceAutoBrightnessOnResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceAutoBrightnessOff(ctx context.Context, req *devices_pb.SendLampDeviceAutoBrightnessOffRequest) (*devices_pb.SendLampDeviceAutoBrightnessOffResponse, error) {
	l := lamp_events.NewLampDeviceAutoBrightnessOffEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceAutoBrightnessOffResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceAutoBrightnessOffResponse{}, nil
}

func (s *DevicesServer) SendLampDeviceAutoBrightnessToggle(ctx context.Context, req *devices_pb.SendLampDeviceAutoBrightnessToggleRequest) (*devices_pb.SendLampDeviceAutoBrightnessToggleResponse, error) {
	l := lamp_events.NewLampDeviceAutoBrightnessToggleEvent()

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDeviceAutoBrightnessToggleResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)
	return &devices_pb.SendLampDeviceAutoBrightnessToggleResponse{}, nil
}

func (s *DevicesServer) SendLampDevicePulse(ctx context.Context, req *devices_pb.SendLampDevicePulseRequest) (*devices_pb.SendLampDevicePulseResponse, error) {
	l := lamp_events.NewLampDevicePulseEvent(req.Red, req.Green, req.Blue)

	j, err := json.Marshal(l)
	if err != nil {
		return &devices_pb.SendLampDevicePulseResponse{}, err
	}

	topic := fmt.Sprintf("devices/%s", req.Mac)

	s.publisher.Publish(topic, j)

	return &devices_pb.SendLampDevicePulseResponse{}, nil
}
