package servers

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/users/persistence"
	"github.com/io-1/kuiper/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	users_pb "github.com/io-1/kuiper/internal/pb/users"
)

type UsersServer struct {
	persistence persistence.Persistence
	users_pb.UnimplementedUsersServiceServer
}

func NewUsersServer(persistence persistence.Persistence) *UsersServer {
	return &UsersServer{
		persistence: persistence,
	}
}

func (s *UsersServer) CreateUser(ctx context.Context, req *users_pb.CreateUserRequest) (*users_pb.CreateUserResponse, error) {

	// generate uuid
	id := uuid.New().String()

	// bcrypt
	encryptedPassword, err := utils.CreateBcryptHashString(req.Password)
	if err != nil {
		return &users_pb.CreateUserResponse{}, status.Error(codes.Internal, "encryption error")
	}

	user := persistence.User{
		ID:       id,
		Username: req.Username,
		Password: encryptedPassword,
		Name:     req.Name,
		Email:    req.Email,
	}

	s.persistence.CreateUser(user)

	return &users_pb.CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) GetUser(ctx context.Context, req *users_pb.GetUserRequest) (*users_pb.GetUserResponse, error) {
	recordNotFound, user := s.persistence.GetUser(req.ID)
	if recordNotFound {
		return &users_pb.GetUserResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	return &users_pb.GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) GetUserByUsername(ctx context.Context, req *users_pb.GetUserByUsernameRequest) (*users_pb.GetUserByUsernameResponse, error) {
	recordNotFound, user := s.persistence.GetUserByUsername(req.Username)
	if recordNotFound {
		return &users_pb.GetUserByUsernameResponse{}, status.Error(codes.NotFound, "username does not exist")
	}

	return &users_pb.GetUserByUsernameResponse{
		ID:       user.ID,
		Password: user.Password,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, req *users_pb.UpdateUserRequest) (*users_pb.UpdateUserResponse, error) {
	user := persistence.User{
		ID:       req.ID,
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
	}

	recordNotFound, err := s.persistence.UpdateUser(user)
	if recordNotFound {
		return &users_pb.UpdateUserResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &users_pb.UpdateUserResponse{}, err
	}

	return &users_pb.UpdateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, req *users_pb.DeleteUserRequest) (*users_pb.DeleteUserResponse, error) {
	user := persistence.User{
		ID: req.ID,
	}

	recordNotFound, err := s.persistence.DeleteUser(user)
	if recordNotFound {
		return &users_pb.DeleteUserResponse{}, status.Error(codes.NotFound, "id was not found")
	}

	if err != nil {
		return &users_pb.DeleteUserResponse{}, err
	}

	return &users_pb.DeleteUserResponse{
		ID: user.ID,
	}, nil
}
