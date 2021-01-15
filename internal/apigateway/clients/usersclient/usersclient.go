package usersclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway/clients/usersclient/request"
	"github.com/io-1/kuiper/internal/apigateway/clients/usersclient/response"
	"github.com/io-1/kuiper/internal/logger"

	users_pb "github.com/io-1/kuiper/internal/pb/users"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type UsersClient struct {
	logger      logger.Logger
	usersClient users_pb.UsersServiceClient
}

func NewUsersClient(serverEnv string, logger logger.Logger) (*UsersClient, error) {
	conn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &UsersClient{
		usersClient: users_pb.NewUsersServiceClient(conn),
	}
	return client, nil
}

func NewUsersClientWithMock(usersClient users_pb.UsersServiceClient, logger logger.Logger) *UsersClient {
	client := &UsersClient{
		logger:      logger,
		usersClient: usersClient,
	}
	return client
}

// Create User
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
		c.JSON(http.StatusMethodNotAllowed, err)
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

// Get a user
func (client *UsersClient) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.GetUserRequest
		res           response.GetUserResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("user_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.usersClient.GetUser(ctx, &users_pb.GetUserRequest{ID: id})
	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.GetUserResponse{
		ID:       r.ID,
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
	}

	c.JSON(http.StatusOK, res)
}

func (client *UsersClient) GetUserByUsername(username string) (*response.GetUserByUsernameResponse, error) {
	var (
		req request.GetUserByUsernameRequest
		res *response.GetUserByUsernameResponse
		ctx = context.Background()
	)

	if validationErrors := req.Validate(username); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		return &response.GetUserByUsernameResponse{}, errors.New(fmt.Sprintf("%v", err))
	}

	r, err := client.usersClient.GetUserByUsername(ctx, &users_pb.GetUserByUsernameRequest{Username: username})
	if err != nil {
		return &response.GetUserByUsernameResponse{}, err
	}

	res = &response.GetUserByUsernameResponse{
		ID:       r.ID,
		Username: r.Username,
		Password: r.Password,
		Name:     r.Name,
		Email:    r.Email,
	}

	return res, nil
}

// Update User
func (client *UsersClient) UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.UpdateUserRequest
		res           response.UpdateUserResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("user_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.usersClient.UpdateUser(ctx, &users_pb.UpdateUserRequest{
		ID:       id,
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
	})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
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

// Patch a user
func (client *UsersClient) PatchUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.PatchUserRequest
		res           response.PatchUserResponse
		errorResponse response.ErrorResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Params.ByName("user_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	// get the user
	r, err := client.usersClient.GetUser(ctx, &users_pb.GetUserRequest{ID: id})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	// use the difference
	if req.Username == "" {
		req.Username = r.Username
	}

	if req.Name == "" {
		req.Name = r.Name
	}

	if req.Email == "" {
		req.Email = r.Email
	}

	// save the request difference
	re, err := client.usersClient.UpdateUser(ctx, &users_pb.UpdateUserRequest{
		ID:       id,
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
	})

	if err != nil {
		st, _ := status.FromError(err)
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if re.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.PatchUserResponse{
		ID:       re.ID,
		Username: re.Username,
		Name:     re.Name,
		Email:    re.Email,
	}

	c.JSON(http.StatusOK, res)
}

// Delete User
func (client *UsersClient) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVE_MINUTES)
	defer cancel()

	var (
		req           request.DeleteUserRequest
		res           response.DeleteUserResponse
		errorResponse response.ErrorResponse
	)

	id := c.Params.ByName("user_id")

	if validationErrors := req.Validate(id); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusMethodNotAllowed, err)
		return
	}

	r, err := client.usersClient.DeleteUser(ctx, &users_pb.DeleteUserRequest{
		ID: id,
	})

	if err != nil {
		st, ok := status.FromError(err)

		// unknown error
		if !ok {
			client.logger.Errorf("unknown error: %v", err)
			errorResponse = response.ErrorResponse{
				Message: fmt.Sprintf("an error has occurred"),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		errorResponse = response.ErrorResponse{
			Message: st.Message(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if r.ID == "" {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.DeleteUserResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
