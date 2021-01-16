package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (s *InteractionsServer) CreateLampToggleEvent(ctx context.Context, req *interactions_pb.CreateLampToggleEventRequest) (*interactions_pb.CreateLampToggleEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampToggleEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampToggleEvent(lampEvent)

	return &interactions_pb.CreateLampToggleEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampToggleEvent(ctx context.Context, req *interactions_pb.GetLampToggleEventRequest) (*interactions_pb.GetLampToggleEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampToggleEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampToggleEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampToggleEvent(ctx context.Context, req *interactions_pb.UpdateLampToggleEventRequest) (*interactions_pb.UpdateLampToggleEventResponse, error) {
	lampEvent := persistence.LampToggleEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampToggleEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampToggleEventResponse{}, err
	}

	return &interactions_pb.UpdateLampToggleEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampToggleEvent(ctx context.Context, req *interactions_pb.DeleteLampToggleEventRequest) (*interactions_pb.DeleteLampToggleEventResponse, error) {
	lampEvent := persistence.LampToggleEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampToggleEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampToggleEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampToggleEventResponse{}, err
	}

	return &interactions_pb.DeleteLampToggleEventResponse{
		ID: lampEvent.ID,
	}, nil
}
