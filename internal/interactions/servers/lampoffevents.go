package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (s *InteractionsServer) CreateLampOffEvent(ctx context.Context, req *interactions_pb.CreateLampOffEventRequest) (*interactions_pb.CreateLampOffEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampOffEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampOffEvent(lampEvent)

	return &interactions_pb.CreateLampOffEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampOffEvent(ctx context.Context, req *interactions_pb.GetLampOffEventRequest) (*interactions_pb.GetLampOffEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampOffEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampOffEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampOffEvent(ctx context.Context, req *interactions_pb.UpdateLampOffEventRequest) (*interactions_pb.UpdateLampOffEventResponse, error) {
	lampEvent := persistence.LampOffEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampOffEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampOffEventResponse{}, err
	}

	return &interactions_pb.UpdateLampOffEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampOffEvent(ctx context.Context, req *interactions_pb.DeleteLampOffEventRequest) (*interactions_pb.DeleteLampOffEventResponse, error) {
	lampEvent := persistence.LampOffEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampOffEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampOffEventResponse{}, err
	}

	return &interactions_pb.DeleteLampOffEventResponse{
		ID: lampEvent.ID,
	}, nil
}
