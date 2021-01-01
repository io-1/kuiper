package servers

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/interactions/persistence"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (s *InteractionsServer) CreateKeypadConditionToLampEvent(ctx context.Context, req *interactions_pb.CreateKeypadConditionToLampEventRequest) (*interactions_pb.CreateKeypadConditionToLampEventResponse, error) {

	// generate uuid
	id := uuid.New().String()

	conditionsToEvents := persistence.KeypadConditionsToLampEvents{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}

	// FIXME: change this to keypad_conditions_to_lamp_events
	s.persistence.CreateKeypadConditionsToLampEvents(conditionsToEvents)

	return &interactions_pb.CreateKeypadConditionToLampEventResponse{
		ID:            id,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}, nil
}

func (s *InteractionsServer) GetKeypadConditionToLampEvent(ctx context.Context, req *interactions_pb.GetKeypadConditionToLampEventRequest) (*interactions_pb.GetKeypadConditionToLampEventResponse, error) {
	recordNotFound, conditionsToEvents := s.persistence.GetKeypadConditionsToLampEvents(req.ID)
	if recordNotFound {
		return &interactions_pb.GetKeypadConditionToLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetKeypadConditionToLampEventResponse{
		ID:            conditionsToEvents.ID,
		InteractionID: conditionsToEvents.InteractionID,
		ConditionID:   conditionsToEvents.ConditionID,
		EventID:       conditionsToEvents.EventID,
	}, nil
}

func (s *InteractionsServer) UpdateKeypadConditionToLampEvent(ctx context.Context, req *interactions_pb.UpdateKeypadConditionToLampEventRequest) (*interactions_pb.UpdateKeypadConditionToLampEventResponse, error) {
	conditionsToEvents := persistence.KeypadConditionsToLampEvents{
		ID:            req.ID,
		InteractionID: req.InteractionID,
		ConditionID:   req.ConditionID,
		EventID:       req.EventID,
	}

	recordNotFound, err := s.persistence.UpdateKeypadConditionsToLampEvents(conditionsToEvents)
	if recordNotFound {
		return &interactions_pb.UpdateKeypadConditionToLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateKeypadConditionToLampEventResponse{}, err
	}

	return &interactions_pb.UpdateKeypadConditionToLampEventResponse{
		ID:            conditionsToEvents.ID,
		InteractionID: conditionsToEvents.InteractionID,
		ConditionID:   conditionsToEvents.ConditionID,
		EventID:       conditionsToEvents.EventID,
	}, nil
}

func (s *InteractionsServer) DeleteKeypadConditionToLampEvent(ctx context.Context, req *interactions_pb.DeleteKeypadConditionToLampEventRequest) (*interactions_pb.DeleteKeypadConditionToLampEventResponse, error) {
	conditionsToEvents := persistence.KeypadConditionsToLampEvents{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteKeypadConditionsToLampEvents(conditionsToEvents)
	if recordNotFound {
		return &interactions_pb.DeleteKeypadConditionToLampEventResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteKeypadConditionToLampEventResponse{}, err
	}

	return &interactions_pb.DeleteKeypadConditionToLampEventResponse{
		ID: conditionsToEvents.ID,
	}, nil
}
