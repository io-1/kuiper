package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (s *InteractionsServer) CreateLampAutoBrightnessOffEvent(ctx context.Context, req *interactions_pb.CreateLampAutoBrightnessOffEventRequest) (*interactions_pb.CreateLampAutoBrightnessOffEventResponse, error) {
	id := uuid.New().String()

	lampEvent := persistence.LampAutoBrightnessOffEvent{
		ID:  id,
		Mac: req.Mac,
	}

	s.persistence.CreateLampAutoBrightnessOffEvent(lampEvent)

	return &interactions_pb.CreateLampAutoBrightnessOffEventResponse{
		ID:  id,
		Mac: req.Mac,
	}, nil
}

func (s *InteractionsServer) GetLampAutoBrightnessOffEvent(ctx context.Context, req *interactions_pb.GetLampAutoBrightnessOffEventRequest) (*interactions_pb.GetLampAutoBrightnessOffEventResponse, error) {
	recordNotFound, lampEvent := s.persistence.GetLampAutoBrightnessOffEvent(req.ID)
	if recordNotFound {
		return &interactions_pb.GetLampAutoBrightnessOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetLampAutoBrightnessOffEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) UpdateLampAutoBrightnessOffEvent(ctx context.Context, req *interactions_pb.UpdateLampAutoBrightnessOffEventRequest) (*interactions_pb.UpdateLampAutoBrightnessOffEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessOffEvent{
		ID:  req.ID,
		Mac: req.Mac,
	}

	recordNotFound, err := s.persistence.UpdateLampAutoBrightnessOffEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.UpdateLampAutoBrightnessOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateLampAutoBrightnessOffEventResponse{}, err
	}

	return &interactions_pb.UpdateLampAutoBrightnessOffEventResponse{
		ID:  lampEvent.ID,
		Mac: lampEvent.Mac,
	}, nil
}

func (s *InteractionsServer) DeleteLampAutoBrightnessOffEvent(ctx context.Context, req *interactions_pb.DeleteLampAutoBrightnessOffEventRequest) (*interactions_pb.DeleteLampAutoBrightnessOffEventResponse, error) {
	lampEvent := persistence.LampAutoBrightnessOffEvent{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteLampAutoBrightnessOffEvent(lampEvent)
	if recordNotFound {
		return &interactions_pb.DeleteLampAutoBrightnessOffEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteLampAutoBrightnessOffEventResponse{}, err
	}

	return &interactions_pb.DeleteLampAutoBrightnessOffEventResponse{
		ID: lampEvent.ID,
	}, nil
}
