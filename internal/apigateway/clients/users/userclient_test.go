// +build unit,!integration

package users

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/io-1/kuiper/internal/logger/blanklogger"
	"github.com/io-1/kuiper/internal/mock"
	"github.com/stretchr/testify/assert"

	users_pb "github.com/io-1/kuiper/internal/pb/users"
)

func Test_PatchUser_Should_Update_Username_When_Username_Is_Not_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		err          error
		username     = "test1"
		name         = "test1"
		email        = "test1@io1.com"
		newUsername  = "test2"
		reqParam     = fmt.Sprintf(`{"username":"%s"}`, newUsername)
		expectedCode = http.StatusOK
		expectedRes  = fmt.Sprintf(`{"id":"%s","username":"%s","name":"%s","email":"%s"}`, id, newUsername, name, email)
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockUsersServiceClient := mock.NewMockUsersServiceClient(mockCtrl)

	blankLogger := blanklogger.NewBlankLogger()
	usersClient := NewUsersClientWithMock(mockUsersServiceClient, blankLogger)

	mockUsersServiceClient.EXPECT().GetUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.GetUserResponse{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    email,
		}, nil,
	)

	mockUsersServiceClient.EXPECT().UpdateUser(
		gomock.Any(),
		&users_pb.UpdateUserRequest{
			ID:       id,
			Username: newUsername,
			Name:     name,
			Email:    email,
		},
	).Return(
		&users_pb.UpdateUserResponse{
			ID:       id,
			Username: newUsername,
			Name:     name,
			Email:    email,
		}, nil,
	)

	r.PATCH("/users/:id", usersClient.PatchUser)

	url := fmt.Sprintf("/users/%s", id)
	c.Request, err = http.NewRequest("PATCH", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_PatchUser_Should_Update_Name_When_Name_Is_Not_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		err          error
		username     = "test1"
		name         = "test1"
		email        = "test1@io1.com"
		newName      = "test2"
		reqParam     = fmt.Sprintf(`{"name":"%s"}`, newName)
		expectedCode = http.StatusOK
		expectedRes  = fmt.Sprintf(`{"id":"%s","username":"%s","name":"%s","email":"%s"}`, id, username, newName, email)
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockUsersServiceClient := mock.NewMockUsersServiceClient(mockCtrl)

	blankLogger := blanklogger.NewBlankLogger()
	usersClient := NewUsersClientWithMock(mockUsersServiceClient, blankLogger)

	mockUsersServiceClient.EXPECT().GetUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.GetUserResponse{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    email,
		}, nil,
	)

	mockUsersServiceClient.EXPECT().UpdateUser(
		gomock.Any(),
		&users_pb.UpdateUserRequest{
			ID:       id,
			Username: username,
			Name:     newName,
			Email:    email,
		},
	).Return(
		&users_pb.UpdateUserResponse{
			ID:       id,
			Username: username,
			Name:     newName,
			Email:    email,
		}, nil,
	)

	r.PATCH("/users/:id", usersClient.PatchUser)

	url := fmt.Sprintf("/users/%s", id)
	c.Request, err = http.NewRequest("PATCH", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_PatchUser_Should_Update_Email_When_Email_Is_Not_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		err          error
		username     = "test1"
		name         = "test1"
		email        = "test1@io1.com"
		newEmail     = "test2@io1.com"
		reqParam     = fmt.Sprintf(`{"email":"%s"}`, newEmail)
		expectedCode = http.StatusOK
		expectedRes  = fmt.Sprintf(`{"id":"%s","username":"%s","name":"%s","email":"%s"}`, id, username, name, newEmail)
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockUsersServiceClient := mock.NewMockUsersServiceClient(mockCtrl)

	blankLogger := blanklogger.NewBlankLogger()
	usersClient := NewUsersClientWithMock(mockUsersServiceClient, blankLogger)

	mockUsersServiceClient.EXPECT().GetUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.GetUserResponse{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    email,
		}, nil,
	)

	mockUsersServiceClient.EXPECT().UpdateUser(
		gomock.Any(),
		&users_pb.UpdateUserRequest{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    newEmail,
		},
	).Return(
		&users_pb.UpdateUserResponse{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    newEmail,
		}, nil,
	)

	r.PATCH("/users/:id", usersClient.PatchUser)

	url := fmt.Sprintf("/users/%s", id)
	c.Request, err = http.NewRequest("PATCH", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_PatchUser_Should_Return_NoContent_When_GetUser_Returns_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		expectedCode        = http.StatusNoContent
		expectedRes         = ""
		err          error
		username     = "test"
		reqParam     = fmt.Sprintf(`{"username":"%s"}`, username)
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockUsersServiceClient := mock.NewMockUsersServiceClient(mockCtrl)

	blankLogger := blanklogger.NewBlankLogger()
	usersClient := NewUsersClientWithMock(mockUsersServiceClient, blankLogger)

	mockUsersServiceClient.EXPECT().GetUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.GetUserResponse{}, nil,
	)

	mockUsersServiceClient.EXPECT().UpdateUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.UpdateUserResponse{}, nil,
	)

	r.PATCH("/users/:id", usersClient.PatchUser)

	url := fmt.Sprintf("/users/%s", id)
	c.Request, err = http.NewRequest("PATCH", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}

func Test_PatchUser_Should_Return_NoContent_When_UpdateUser_Returns_Empty(t *testing.T) {
	var (
		id           string = "00000000-1111-2222-3333-444444444444"
		expectedCode        = http.StatusNoContent
		expectedRes         = ""
		err          error
		username     = "test"
		email        = "test@io1.com"
		name         = "test"
		reqParam     = fmt.Sprintf(`{"username":"%s"}`, username)
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	mockCtrl := gomock.NewController(t)
	mockUsersServiceClient := mock.NewMockUsersServiceClient(mockCtrl)

	blankLogger := blanklogger.NewBlankLogger()
	usersClient := NewUsersClientWithMock(mockUsersServiceClient, blankLogger)

	mockUsersServiceClient.EXPECT().GetUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.GetUserResponse{
			ID:       id,
			Username: username,
			Name:     name,
			Email:    email,
		}, nil,
	)

	mockUsersServiceClient.EXPECT().UpdateUser(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&users_pb.UpdateUserResponse{}, nil,
	)

	r.PATCH("/users/:id", usersClient.PatchUser)

	url := fmt.Sprintf("/users/%s", id)
	c.Request, err = http.NewRequest("PATCH", url, strings.NewReader(string(reqParam)))
	assert.NoError(t, err)

	r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	assert.Equal(t, expectedCode, actualCode)

	actualRes := w.Body.String()
	assert.Equal(t, expectedRes, actualRes)
}
