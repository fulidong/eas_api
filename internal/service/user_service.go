package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	return s.loginUc.Login(ctx, in)
}

func (s *EasService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return s.userUc.CreateUser(ctx, in)
}

func (s *EasService) GetPageList(ctx context.Context, in *v1.GetPageListRequest) (*v1.GetPageListResponse, error) {
	return s.userUc.GetPageList(ctx, in)
}

func (s *EasService) GetUserDetail(ctx context.Context, in *v1.GetUserDetailRequest) (*v1.GetUserDetailResponse, error) {
	return s.userUc.GetUserDetail(ctx, in)
}

func (s *EasService) GetUserSelfDetail(ctx context.Context, in *v1.GetUserSelfDetailRequest) (*v1.GetUserSelfDetailResponse, error) {
	return s.userUc.GetUserSelfDetail(ctx, in)
}

func (s *EasService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return s.userUc.UpdateUser(ctx, in)
}

func (s *EasService) UpdateUserSelf(ctx context.Context, in *v1.UpdateUserSelfRequest) (*v1.UpdateUserSelfResponse, error) {
	return s.userUc.UpdateUserSelf(ctx, in)
}

func (s *EasService) SetUserStatus(ctx context.Context, in *v1.SetUserStatusRequest) (*v1.SetUserStatusResponse, error) {
	return s.userUc.SetUserStatus(ctx, in)
}

func (s *EasService) ResetUserPassWord(ctx context.Context, in *v1.ResetUserPassWordRequest) (*v1.ResetUserPassWordResponse, error) {
	return s.userUc.ResetUserPassWord(ctx, in)
}

func (s *EasService) UpdateUserPassWord(ctx context.Context, in *v1.UpdateUserPassWordRequest) (*v1.UpdateUserPassWordResponse, error) {
	return s.userUc.UpdateUserPassWord(ctx, in)
}

func (s *EasService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return s.userUc.DeleteUser(ctx, in)
}
