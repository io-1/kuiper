package servers

import (
	"context"

	"github.com/n7down/kuiper/internal/users/persistence"

	users_pb "github.com/n7down/kuiper/internal/pb/users"
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
	user := persistence.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
	}

	_, newUser := s.persistence.CreateUser(user)

	return &users_pb.CreateUserResponse{
		Username: newUser.Username,
		Name:     newUser.Name,
		Email:    newUser.Email,
	}, nil
}

func (s *UsersServer) GetUser(ctx context.Context, req *users_pb.GetUserRequest) (*users_pb.GetUserResponse, error) {
	_, user := s.persistence.GetUser(req.Username)
	// recordNotFound, user := s.persistence.GetUser(req.Username)
	// if recordNotFound {
	// 	return &users_pb.GetUserResponse{}, errors.New("record not found")
	// }

	return &users_pb.GetUserResponse{
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, req *users_pb.UpdateUserRequest) (*users_pb.UpdateUserResponse, error) {
	user := persistence.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
	}

	s.persistence.UpdateUser(user)
	// rowsAffected, updatedUser := s.persistence.UpdateUser(user)
	// if rowsAffected == 0 {
	// 	return &users_pb.UpdateUserResponse{}, errors.New("record not found")
	// }

	return &users_pb.UpdateUserResponse{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, req *users_pb.DeleteUserRequest) (*users_pb.DeleteUserResponse, error) {
	user := persistence.User{
		Username: req.Username,
	}

	s.persistence.DeleteUser(user)
	// rowsAffected := s.persistence.DeleteUser(user)
	// if rowsAffected == 0 {
	// 	return &users_pb.DeleteUserResponse{}, errors.New("record not found")
	// }

	return &users_pb.DeleteUserResponse{
		Username: req.Username,
	}, nil
}
