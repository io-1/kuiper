package servers

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/io-1/kuiper/internal/interactions/persistence"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

func (s *InteractionsServer) CreateKeypadCondition(ctx context.Context, req *interactions_pb.CreateKeypadConditionRequest) (*interactions_pb.CreateKeypadConditionResponse, error) {

	// generate uuid
	id := uuid.New().String()

	keypadCondition := persistence.KeypadCondition{
		ID:            id,
		InteractionID: req.InteractionID,
		Mac:           req.Mac,
		ButtonID:      req.ButtonID,
	}

	s.persistence.CreateKeypadCondition(keypadCondition)

	return &interactions_pb.CreateKeypadConditionResponse{
		ID:            id,
		InteractionID: req.InteractionID,
		Mac:           req.Mac,
		ButtonID:      req.ButtonID,
	}, nil
}

func (s *InteractionsServer) GetKeypadCondition(ctx context.Context, req *interactions_pb.GetKeypadConditionRequest) (*interactions_pb.GetKeypadConditionResponse, error) {
	recordNotFound, interaction := s.persistence.GetKeypadCondition(req.ID)
	if recordNotFound {
		return &interactions_pb.GetKeypadConditionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetKeypadConditionResponse{
		ID:            interaction.ID,
		InteractionID: interaction.InteractionID,
		Mac:           interaction.Mac,
		ButtonID:      interaction.ButtonID,
	}, nil
}

func (s *InteractionsServer) UpdateKeypadCondition(ctx context.Context, req *interactions_pb.UpdateKeypadConditionRequest) (*interactions_pb.UpdateKeypadConditionResponse, error) {
	interaction := persistence.KeypadCondition{
		ID:            req.ID,
		InteractionID: req.InteractionID,
		Mac:           req.Mac,
		ButtonID:      req.ButtonID,
	}

	recordNotFound, err := s.persistence.UpdateKeypadCondition(interaction)
	if recordNotFound {
		return &interactions_pb.UpdateKeypadConditionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateKeypadConditionResponse{}, err
	}

	return &interactions_pb.UpdateKeypadConditionResponse{
		ID:            interaction.ID,
		InteractionID: req.InteractionID,
		Mac:           req.Mac,
		ButtonID:      req.ButtonID,
	}, nil
}

func (s *InteractionsServer) DeleteKeypadCondition(ctx context.Context, req *interactions_pb.DeleteKeypadConditionRequest) (*interactions_pb.DeleteKeypadConditionResponse, error) {
	interaction := persistence.KeypadCondition{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteKeypadCondition(interaction)
	if recordNotFound {
		return &interactions_pb.DeleteKeypadConditionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteKeypadConditionResponse{}, err
	}

	return &interactions_pb.DeleteKeypadConditionResponse{
		ID: interaction.ID,
	}, nil
}
