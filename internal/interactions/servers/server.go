package servers

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/interactions/persistence"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

type InteractionsServer struct {
	persistence persistence.Persistence
	interactions_pb.UnimplementedInteractionsServiceServer
}

func NewInteractionsServer(persistence persistence.Persistence) *InteractionsServer {
	return &InteractionsServer{
		persistence: persistence,
	}
}

func (s *InteractionsServer) CreateInteraction(ctx context.Context, req *interactions_pb.CreateInteractionRequest) (*interactions_pb.CreateInteractionResponse, error) {

	// generate uuid
	id := uuid.New().String()

	interaction := persistence.Interaction{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	s.persistence.CreateInteraction(interaction)

	return &interactions_pb.CreateInteractionResponse{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}, nil
}

func (s *InteractionsServer) GetInteraction(ctx context.Context, req *interactions_pb.GetInteractionRequest) (*interactions_pb.GetInteractionResponse, error) {
	recordNotFound, interaction := s.persistence.GetInteraction(req.ID)
	if recordNotFound {
		return &interactions_pb.GetInteractionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &interactions_pb.GetInteractionResponse{
		ID:          interaction.ID,
		Name:        interaction.Name,
		Description: interaction.Description,
	}, nil
}

func (s *InteractionsServer) UpdateInteraction(ctx context.Context, req *interactions_pb.UpdateInteractionRequest) (*interactions_pb.UpdateInteractionResponse, error) {
	interaction := persistence.Interaction{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	recordNotFound, err := s.persistence.UpdateInteraction(interaction)
	if recordNotFound {
		return &interactions_pb.UpdateInteractionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.UpdateInteractionResponse{}, err
	}

	return &interactions_pb.UpdateInteractionResponse{
		ID:          interaction.ID,
		Name:        interaction.Name,
		Description: interaction.Description,
	}, nil
}

func (s *InteractionsServer) DeleteInteraction(ctx context.Context, req *interactions_pb.DeleteInteractionRequest) (*interactions_pb.DeleteInteractionResponse, error) {
	interaction := persistence.Interaction{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteInteraction(interaction)
	if recordNotFound {
		return &interactions_pb.DeleteInteractionResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &interactions_pb.DeleteInteractionResponse{}, err
	}

	return &interactions_pb.DeleteInteractionResponse{
		ID: interaction.ID,
	}, nil
}
