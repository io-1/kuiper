package servers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/users/persistence"
	"github.com/io-1/kuiper/internal/utils"

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
		return &users_pb.CreateUserResponse{}, errors.New("error encrypting password")
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
	_, user := s.persistence.GetUser(req.Username)

	return &users_pb.GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, req *users_pb.UpdateUserRequest) (*users_pb.UpdateUserResponse, error) {
	user := persistence.User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
	}

	s.persistence.UpdateUser(user)

	return &users_pb.UpdateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, req *users_pb.DeleteUserRequest) (*users_pb.DeleteUserResponse, error) {
	user := persistence.User{
		ID:       req.ID,
		Username: req.Username,
	}

	s.persistence.DeleteUser(user)

	return &users_pb.DeleteUserResponse{
		ID:       req.ID,
		Username: req.Username,
	}, nil
}
