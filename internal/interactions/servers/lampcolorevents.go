package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions"
)

func (s *InteractionsServer) CreateLampColorEvent(ctx context.Context, req *interactions_pb.CreateLampColorEventRequest) (*interactions_pb.CreateLampColorEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampColorEvent{
		ID:    id,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}

	s.persistence.CreateLampColorEvent(lampEvent)

	return &interactions_pb.CreateLampColorEventResponse{
		ID:    id,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}, nil
}

func (s *InteractionsServer) GetLampColorEvent(ctx context.Context, req *interactions_pb.GetLampColorEventRequest) (*interactions_pb.GetLampColorEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampColorEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampColorEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampColorEventResponse{
		ID:    lampEvent.ID,
		Mac:   lampEvent.Mac,
		Red:   lampEvent.Red,
		Green: lampEvent.Green,
		Blue:  lampEvent.Blue,
	}, nil
}

func (s *InteractionsServer) UpdateLampColorEvent(ctx context.Context, req *interactions_pb.UpdateLampColorEventRequest) (*interactions_pb.UpdateLampColorEventResponse, error) {
	lampEvent := persistence.LampColorEvent{
		ID:    req.ID,
		Mac:   req.Mac,
		Red:   req.Red,
		Green: req.Green,
		Blue:  req.Blue,
	}

	recordNotFound, err := s.persistence.UpdateLampColorEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampColorEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampColorEventResponse{}, err
	}

	return &interactions_pb.UpdateLampColorEventResponse{
		ID:    lampEvent.ID,
		Mac:   lampEvent.Mac,
		Red:   lampEvent.Red,
		Green: lampEvent.Green,
		Blue:  lampEvent.Blue,
	}, nil
}

func (s *InteractionsServer) DeleteLampColorEvent(ctx context.Context, req *interactions_pb.DeleteLampColorEventRequest) (*interactions_pb.DeleteLampColorEventResponse, error) {
	lampEvent := persistence.LampColorEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampColorEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampColorEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampColorEventResponse{}, err
	}

	return &interactions_pb.DeleteLampColorEventResponse{
		ID: lampEvent.ID,
	}, nil
}
