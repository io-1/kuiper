package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (s *InteractionsServer) CreateLampOnEvent(ctx context.Context, req *interactions_pb.CreateLampOnEventRequest) (*interactions_pb.CreateLampOnEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampOnEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampOnEvent(lampEvent)

	return &interactions_pb.CreateLampOnEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampOnEvent(ctx context.Context, req *interactions_pb.GetLampOnEventRequest) (*interactions_pb.GetLampOnEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampOnEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampOnEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampOnEvent(ctx context.Context, req *interactions_pb.UpdateLampOnEventRequest) (*interactions_pb.UpdateLampOnEventResponse, error) {
	lampEvent := persistence.LampOnEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampOnEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampOnEventResponse{}, err
	}

	return &interactions_pb.UpdateLampOnEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampOnEvent(ctx context.Context, req *interactions_pb.DeleteLampOnEventRequest) (*interactions_pb.DeleteLampOnEventResponse, error) {
	lampEvent := persistence.LampOnEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampOnEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampOnEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampOnEventResponse{}, err
	}

	return &interactions_pb.DeleteLampOnEventResponse{
		ID: lampEvent.ID,
	}, nil
}
