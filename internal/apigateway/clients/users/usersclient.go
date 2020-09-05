package users

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/apigateway/clients/users/request"
	"github.com/n7down/kuiper/internal/apigateway/clients/users/response"

	users_pb "github.com/n7down/kuiper/internal/pb/users"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type UsersClient struct {
	usersClient users_pb.UsersServiceClient
}

func NewUsersClient(serverEnv string) (*UsersClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &UsersClient{
		usersClient: users_pb.NewUsersServiceClient(conn),
	}
	return client, nil
}

func NewUsersClientWithMock(usersClient users_pb.UsersServiceClient) *UsersClient {
	client := &UsersClient{
		usersClient: usersClient,
	}
	return client
}

// FIXME: switch to ID from username
func (client *UsersClient) CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.CreateUserRequest
		res response.CreateUserResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.usersClient.CreateUser(ctx, &users_pb.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res = response.CreateUserResponse{
		ID:       r.ID,
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
	}

	c.JSON(http.StatusOK, res)
}

// FIXME: switch to ID from username
func (client *UsersClient) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.GetUserRequest
		res response.GetUserResponse
	)

	username := c.Params.ByName("username")

	req = request.GetUserRequest{
		Username: username,
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.usersClient.GetUser(ctx, &users_pb.GetUserRequest{Username: req.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if r.Username == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetUserResponse{
		ID:       r.ID,
		Username: r.Username,
		Password: r.Password,
		Name:     r.Name,
		Email:    r.Email,
	}

	c.JSON(http.StatusOK, res)
}

func (client *UsersClient) GetUserLogin(username string) (response.UserLoginResponse, error) {
	var (
		res response.UserLoginResponse
		ctx = context.Background()
	)

	// req = request.GetUserRequest{
	// 	Username: username,
	// }

	// if validationErrors := req.Validate(); len(validationErrors) > 0 {
	// 	err := map[string]interface{}{"validationError": validationErrors}
	// return "", validationErrors
	// }

	r, err := client.usersClient.GetUser(ctx, &users_pb.GetUserRequest{Username: username})
	if err != nil {
		return response.UserLoginResponse{}, err
	}

	res = response.UserLoginResponse{
		ID:       r.ID,
		Username: r.Username,
		Password: r.Password,
		Name:     r.Name,
		Email:    r.Email,
	}

	return res, nil
}

// FIXME: switch to ID from username
func (client *UsersClient) UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.UpdateUserRequest
		res response.UpdateUserResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	username := c.Params.ByName("username")
	req.Username = username

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.usersClient.UpdateUser(ctx, &users_pb.UpdateUserRequest{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if r.Username == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.UpdateUserResponse{
		ID:       r.ID,
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
	}

	c.JSON(http.StatusOK, res)
}

// FIXME: switch to ID from username
func (client *UsersClient) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req request.DeleteUserRequest
		res response.DeleteUserResponse
	)

	username := c.Params.ByName("username")

	req = request.DeleteUserRequest{
		Username: username,
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := client.usersClient.DeleteUser(ctx, &users_pb.DeleteUserRequest{
		Username: username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if r.Username == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.DeleteUserResponse{
		Username: r.Username,
	}

	c.JSON(http.StatusOK, res)
}
