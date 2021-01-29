package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (s *InteractionsServer) CreateLampAutoBrightnessOnEvent(ctx context.Context, req *interactions_pb.CreateLampAutoBrightnessOnEventRequest) (*interactions_pb.CreateLampAutoBrightnessOnEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampAutoBrightnessOnEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampAutoBrightnessOnEvent(lampEvent)

	return &interactions_pb.CreateLampAutoBrightnessOnEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampAutoBrightnessOnEvent(ctx context.Context, req *interactions_pb.GetLampAutoBrightnessOnEventRequest) (*interactions_pb.GetLampAutoBrightnessOnEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampAutoBrightnessOnEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampAutoBrightnessOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampAutoBrightnessOnEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampAutoBrightnessOnEvent(ctx context.Context, req *interactions_pb.UpdateLampAutoBrightnessOnEventRequest) (*interactions_pb.UpdateLampAutoBrightnessOnEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessOnEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampAutoBrightnessOnEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampAutoBrightnessOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampAutoBrightnessOnEventResponse{}, err
	}

	return &interactions_pb.UpdateLampAutoBrightnessOnEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampAutoBrightnessOnEvent(ctx context.Context, req *interactions_pb.DeleteLampAutoBrightnessOnEventRequest) (*interactions_pb.DeleteLampAutoBrightnessOnEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessOnEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampAutoBrightnessOnEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampAutoBrightnessOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampAutoBrightnessOnEventResponse{}, err
	}

	return &interactions_pb.DeleteLampAutoBrightnessOnEventResponse{
		ID: lampEvent.ID,
	}, nil
}
