package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (s *InteractionsServer) CreateLampPulseEvent(ctx context.Context, req *interactions_pb.CreateLampPulseEventRequest) (*interactions_pb.CreateLampPulseEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampPulseEvent{
		ID:    id,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}

	s.persistence.CreateLampPulseEvent(lampEvent)

	return &interactions_pb.CreateLampPulseEventResponse{
		ID:    id,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}, nil
}

func (s *InteractionsServer) GetLampPulseEvent(ctx context.Context, req *interactions_pb.GetLampPulseEventRequest) (*interactions_pb.GetLampPulseEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampPulseEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampPulseEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampPulseEventResponse{
		ID:    lampEvent.ID,
		Mac:   lampEvent.Mac,
		Red:   lampEvent.Red,
		Green: lampEvent.Green,
		Blue:  lampEvent.Blue,
	}, nil
}

func (s *InteractionsServer) UpdateLampPulseEvent(ctx context.Context, req *interactions_pb.UpdateLampPulseEventRequest) (*interactions_pb.UpdateLampPulseEventResponse, error) {
	lampEvent := persistence.LampPulseEvent{
		ID:    req.ID,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}

	recordNotFound, err := s.persistence.UpdateLampPulseEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampPulseEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampPulseEventResponse{}, err
	}

	return &interactions_pb.UpdateLampPulseEventResponse{
		ID:    lampEvent.ID,
		Mac:   lampEvent.Mac,
		Red:   lampEvent.Red,
		Green: lampEvent.Green,
		Blue:  lampEvent.Blue,
	}, nil
}

func (s *InteractionsServer) DeleteLampPulseEvent(ctx context.Context, req *interactions_pb.DeleteLampPulseEventRequest) (*interactions_pb.DeleteLampPulseEventResponse, error) {
	lampEvent := persistence.LampPulseEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampPulseEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampPulseEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampPulseEventResponse{}, err
	}

	return &interactions_pb.DeleteLampPulseEventResponse{
		ID: lampEvent.ID,
	}, nil
}
