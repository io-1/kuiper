package servers

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/interactions/persistence"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	interactions_pb "github.com/io-1/kuiper/pkg/pb/interactions/go"
)

func (s *InteractionsServer) CreateAttach(ctx context.Context, req *interactions_pb.CreateAttachRequest) (*interactions_pb.CreateAttachResponse, error) {

	// generate uuid
	id := uuid.New().String()

	conditionsToEvents := persistence.ConditionsToEvents{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}

	// FIXME: change this to keypad_conditions_to_lamp_events
	s.persistence.CreateConditionsToEvents(conditionsToEvents)

	return &interactions_pb.CreateAttachResponse{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}, nil
}

func (s *InteractionsServer) GetAttach(ctx context.Context, req *interactions_pb.GetAttachRequest) (*interactions_pb.GetAttachResponse, error) {
	recordNotFound, conditionsToEvents := s.persistence.GetConditionsToEvents(req.ID)
	if recordNotFound {
		return &interactions_pb.GetAttachResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetAttachResponse{
		ID:            conditionsToEvents.ID,
		InteractionID: conditionsToEvents.InteractionID,
		ConditionID:   conditionsToEvents.ConditionID,
		EventID:       conditionsToEvents.EventID,
	}, nil
}

func (s *InteractionsServer) UpdateAttach(ctx context.Context, req *interactions_pb.UpdateAttachRequest) (*interactions_pb.UpdateAttachResponse, error) {
	conditionsToEvents := persistence.ConditionsToEvents{
		ID:            req.ID,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}

	recordNotFound, err := s.persistence.UpdateConditionsToEvents(conditionsToEvents)
	if recordNotFound {
		return &interactions_pb.UpdateAttachResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateAttachResponse{}, err
	}

	return &interactions_pb.UpdateAttachResponse{
		ID:            conditionsToEvents.ID,
		InteractionID: conditionsToEvents.InteractionID,
		ConditionID:   conditionsToEvents.ConditionID,
		EventID:       conditionsToEvents.EventID,
	}, nil
}

func (s *InteractionsServer) DeleteAttach(ctx context.Context, req *interactions_pb.DeleteAttachRequest) (*interactions_pb.DeleteAttachResponse, error) {
	conditionsToEvents := persistence.ConditionsToEvents{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteConditionsToEvents(conditionsToEvents)
	if recordNotFound {
		return &interactions_pb.DeleteAttachResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteAttachResponse{}, err
	}

	return &interactions_pb.DeleteAttachResponse{
		ID: conditionsToEvents.ID,
	}, nil
}
