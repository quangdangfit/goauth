package controllers

import (
	"context"
	"testing"

	"github.com/quangdangfit/goauth/config"
	models "github.com/quangdangfit/goauth/models/mysql"
	"github.com/quangdangfit/goauth/models/types"
	api "github.com/quangdangfit/goauth/proto/api/auth"
	"github.com/quangdangfit/goauth/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateAsync(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) Cancel(ctx context.Context, uuid string, note string) error {
	args := m.Called(ctx, uuid, note)
	return args.Error(0)
}

func (m *MockUserService) DeleteByUUID(ctx context.Context, uuid string) error {
	args := m.Called(ctx, uuid)
	return args.Error(0)
}

func (m *MockUserService) GetById(ctx context.Context, uuid string) (*models.User, error) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) List(ctx context.Context, search *string, filter map[string]interface{}, page, limit int) ([]models.User, *utils.Pagination, error) {
	args := m.Called(ctx, search, filter, page, limit)
	return args.Get(0).([]models.User), args.Get(1).(*utils.Pagination), args.Error(2)
}

func (m *MockUserService) Create(ctx context.Context, doc *models.User) error {
	args := m.Called(ctx, doc)
	return args.Error(0)
}

func (m *MockUserService) Update(ctx context.Context, uuid string, doc *models.User) error {
	args := m.Called(ctx, uuid, doc)
	return args.Error(0)
}

func (m *MockUserService) Delete(ctx context.Context, uuid string) error {
	args := m.Called(ctx, uuid)
	return args.Error(0)
}

type MockStream struct {
	mock.Mock
}

func (m *MockStream) Send(msg *api.ListUserResponse) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockStream) SetHeader(md metadata.MD) error {
	args := m.Called(md)
	return args.Error(0)
}

func (m *MockStream) SendHeader(md metadata.MD) error {
	args := m.Called(md)
	return args.Error(0)
}

func (m *MockStream) SetTrailer(md metadata.MD) {
	m.Called(md)
}

func (m *MockStream) Context() context.Context {
	args := m.Called()
	return args.Get(0).(context.Context)
}

func (m *MockStream) SendMsg(msg any) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockStream) RecvMsg(msg any) error {
	args := m.Called(msg)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}

	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)

	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.CreateUserRequest{
		OrderCode:   "order_123",
		Amount:      100.50,
		FromAccount: "account_1",
		ToAccount:   "account_2",
		Description: "Payment for order 123",
	}

	mockUserSvc.On("CreateAsync", ctx, mock.Anything).Return(nil)

	resp, err := server.CreateUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotNil(t, resp.Data)
}

func TestCancelUser(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}
	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)
	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.CancelUserRequest{
		Uuid: "user_123",
		Note: "Cancelled by user",
	}

	mockUserSvc.On("Cancel", ctx, req.Uuid, req.Note).Return(nil)

	resp, err := server.CancelUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}
	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)
	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.DeleteUserRequest{
		Uuid: "user_123",
	}

	mockUserSvc.On("DeleteByUUID", ctx, req.Uuid).Return(nil)

	resp, err := server.DeleteUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestGetUserDetail(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}
	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)
	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.GetUserDetailRequest{
		Uuid: "user_123",
	}

	user := &models.User{
		OrderCode:   "order_123",
		Amount:      100.50,
		FromAccount: "account_1",
		ToAccount:   "account_2",
		Description: "Payment for order 123",
		Status:      types.UserStatusInit,
	}

	mockUserSvc.On("GetById", ctx, req.Uuid).Return(user, nil)

	resp, err := server.GetUserDetail(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotNil(t, resp.Data)
}

func TestListUser(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}
	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)
	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.ListUserRequest{
		Search:   "order_123",
		Status:   api.UserStatus_INIT.Enum(),
		Page:     1,
		Limit:    10,
		FromDate: "2023-01-01",
		ToDate:   "2023-12-31",
	}

	users := []models.User{
		{OrderCode: "order_123", Amount: 100.50, FromAccount: "account_1", ToAccount: "account_2", Status: types.UserStatusInit},
	}
	pagination := utils.NewPagination(1, 10)
	mockUserSvc.On("List", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(users, pagination, nil)

	resp, err := server.ListUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotNil(t, resp.Data)
	assert.Len(t, resp.Data.Items, 1)
}

func TestStreamUser(t *testing.T) {
	ctx := context.Background()
	mockUserSvc := &MockUserService{}
	mockStream := &MockStream{}
	cfg, err := config.Load()
	assert.NoError(t, err)
	logger := cfg.Log.Build().WithName(cfg.ServiceName)
	server := &Controller{
		userSvc: mockUserSvc,
		logger:  logger,
	}

	req := &api.ListUserRequest{
		Search: "order_123",
		Status: api.UserStatus_INIT.Enum(),
		Page:   1,
		Limit:  10,
	}

	users := []models.User{
		{OrderCode: "order_123", Amount: 100.50, FromAccount: "account_1", ToAccount: "account_2", Status: types.UserStatusInit},
	}
	pagination := utils.NewPagination(1, 10)
	mockUserSvc.On("List", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(users, pagination, nil)

	// Mock stream.Send to return nil for each chunk
	mockStream.On("Send", mock.Anything).Return(nil).Once()

	err = server.StreamUser(req, mockStream)
	assert.NoError(t, err)
	mockStream.AssertExpectations(t)
}
