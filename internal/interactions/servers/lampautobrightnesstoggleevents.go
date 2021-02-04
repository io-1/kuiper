package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (s *InteractionsServer) CreateLampAutoBrightnessToggleEvent(ctx context.Context, req *interactions_pb.CreateLampAutoBrightnessToggleEventRequest) (*interactions_pb.CreateLampAutoBrightnessToggleEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampAutoBrightnessToggleEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampAutoBrightnessToggleEvent(lampEvent)

	return &interactions_pb.CreateLampAutoBrightnessToggleEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampAutoBrightnessToggleEvent(ctx context.Context, req *interactions_pb.GetLampAutoBrightnessToggleEventRequest) (*interactions_pb.GetLampAutoBrightnessToggleEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampAutoBrightnessToggleEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampAutoBrightnessToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampAutoBrightnessToggleEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampAutoBrightnessToggleEvent(ctx context.Context, req *interactions_pb.UpdateLampAutoBrightnessToggleEventRequest) (*interactions_pb.UpdateLampAutoBrightnessToggleEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessToggleEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampAutoBrightnessToggleEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampAutoBrightnessToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampAutoBrightnessToggleEventResponse{}, err
	}

	return &interactions_pb.UpdateLampAutoBrightnessToggleEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampAutoBrightnessToggleEvent(ctx context.Context, req *interactions_pb.DeleteLampAutoBrightnessToggleEventRequest) (*interactions_pb.DeleteLampAutoBrightnessToggleEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessToggleEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampAutoBrightnessToggleEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampAutoBrightnessToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampAutoBrightnessToggleEventResponse{}, err
	}

	return &interactions_pb.DeleteLampAutoBrightnessToggleEventResponse{
		ID: lampEvent.ID,
	}, nil
}
