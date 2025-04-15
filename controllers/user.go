package controllers

import (
	"context"

	models "github.com/rinard84/auth-service/models/mysql"
	"github.com/rinard84/auth-service/utils"
	api "github.com/rinard84/proto/api/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Controller) Register(ctx context.Context, in *api.RegisterRequest) (*api.RegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(in.Password)
	if err != nil {
		s.logger.Error(err, "HashPassword")
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	err = s.userRepo.Create(ctx, &models.User{
		Username: in.Username,
		Password: hashedPassword,
		Name:     in.Name,
		Phone:    in.Phone,
		Email:    in.Email,
	})
	if err != nil {
		s.logger.Error(err, "Create", "username", in.Username)
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &api.RegisterResponse{
		Code:    int32(codes.OK),
		Message: codes.OK.String(),
	}, nil
}

func (s *Controller) Login(ctx context.Context, in *api.LoginRequest) (*api.LoginResponse, error) {
	user, err := s.userRepo.GetOne(ctx, map[string]interface{}{"username": in.Username}, nil)
	if err != nil {
		s.logger.Error(err, "GetUser", "username", in.Username)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err = utils.ComparePassword(user.Password, in.Password); err != nil {
		s.logger.Error(err, "ComparePassword", "username", in.Username)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	accessToken := ""
	refreshToken := ""

	return &api.LoginResponse{
		Code:    int32(codes.OK),
		Message: codes.OK.String(),
		Data: &api.LoginResponse_Data{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
func (s *Controller) RefreshToken(ctx context.Context, in *api.RefreshTokenRequest) (*api.RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (s *Controller) ResetPassword(ctx context.Context, in *api.ResetPasswordRequest) (*api.ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (s *Controller) ChangePassword(ctx context.Context, in *api.ChangePasswordRequest) (*api.ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (s *Controller) UpdateUserProfile(ctx context.Context, in *api.UpdateUserProfileRequest) (*api.UpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (s *Controller) GetUserProfile(ctx context.Context, in *api.GetUserProfileRequest) (*api.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (s *Controller) VerifyAccount(ctx context.Context, in *api.VerifyAccountRequest) (*api.VerifyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
func (s *Controller) DeactivateAccount(ctx context.Context, in *api.DeactivateAccountRequest) (*api.DeactivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateAccount not implemented")
}
func (s *Controller) CreatePermission(ctx context.Context, in *api.CreatePermissionRequest) (*api.CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (s *Controller) ListPermission(ctx context.Context, in *api.ListPermissionRequest) (*api.ListPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermission not implemented")
}
func (s *Controller) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (s *Controller) ListUser(ctx context.Context, in *api.ListUserRequest) (*api.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (s *Controller) RegisterDevice(ctx context.Context, in *api.RegisterDeviceRequest) (*api.RegisterDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (s *Controller) UnregisterDevice(ctx context.Context, in *api.UnregisterDeviceRequest) (*api.UnregisterDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterDevice not implemented")
}
func (s *Controller) ListDevice(ctx context.Context, in *api.ListDeviceRequest) (*api.ListDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevice not implemented")
}
