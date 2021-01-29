package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (s *InteractionsServer) CreateLampBrightnessEvent(ctx context.Context, req *interactions_pb.CreateLampBrightnessEventRequest) (*interactions_pb.CreateLampBrightnessEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampBrightnessEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampBrightnessEvent(lampEvent)

	return &interactions_pb.CreateLampBrightnessEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampBrightnessEvent(ctx context.Context, req *interactions_pb.GetLampBrightnessEventRequest) (*interactions_pb.GetLampBrightnessEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampBrightnessEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampBrightnessEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampBrightnessEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampBrightnessEvent(ctx context.Context, req *interactions_pb.UpdateLampBrightnessEventRequest) (*interactions_pb.UpdateLampBrightnessEventResponse, error) {
	lampEvent := persistence.LampBrightnessEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampBrightnessEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampBrightnessEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampBrightnessEventResponse{}, err
	}

	return &interactions_pb.UpdateLampBrightnessEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampBrightnessEvent(ctx context.Context, req *interactions_pb.DeleteLampBrightnessEventRequest) (*interactions_pb.DeleteLampBrightnessEventResponse, error) {
	lampEvent := persistence.LampBrightnessEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampBrightnessEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampBrightnessEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampBrightnessEventResponse{}, err
	}

	return &interactions_pb.DeleteLampBrightnessEventResponse{
		ID: lampEvent.ID,
	}, nil
}
