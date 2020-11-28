package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (s *InteractionsServer) CreateLampEvent(ctx context.Context, req *interactions_pb.CreateLampEventRequest) (*interactions_pb.CreateLampEventResponse, error) {

	// generate uuid
	id := uuid.New().String()

	lampEvent := persistence.LampEvent{
		ID:        id,
		Mac:       req.Mac,
		EventType: req.EventType,
		Color:     req.Color,
	}

	s.persistence.CreateLampEvent(lampEvent)

	return &interactions_pb.CreateLampEventResponse{
		ID:        id,
		Mac:       req.Mac,
		EventType: req.EventType,
		Color:     req.Color,
	}, nil
}

func (s *InteractionsServer) GetLampEvent(ctx context.Context, req *interactions_pb.GetLampEventRequest) (*interactions_pb.GetLampEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampEventResponse{
		ID:        lampEvent.ID,
		Mac:       lampEvent.Mac,
		EventType: lampEvent.EventType,
		Color:     lampEvent.Color,
	}, nil
}

func (s *InteractionsServer) UpdateLampEvent(ctx context.Context, req *interactions_pb.UpdateLampEventRequest) (*interactions_pb.UpdateLampEventResponse, error) {
	lampEvent := persistence.LampEvent{
		ID:        req.ID,
		Mac:       req.Mac,
		EventType: req.EventType,
		Color:     req.Color,
	}

	recordNotFound, err := s.persistence.UpdateLampEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampEventResponse{}, err
	}

	return &interactions_pb.UpdateLampEventResponse{
		ID:        lampEvent.ID,
		Mac:       lampEvent.Mac,
		EventType: lampEvent.EventType,
		Color:     lampEvent.Color,
	}, nil
}

func (s *InteractionsServer) DeleteLampEvent(ctx context.Context, req *interactions_pb.DeleteLampEventRequest) (*interactions_pb.DeleteLampEventResponse, error) {
	lampEvent := persistence.LampEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampEventResponse{}, err
	}

	return &interactions_pb.DeleteLampEventResponse{
		ID: lampEvent.ID,
	}, nil
}
