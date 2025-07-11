package service

import (
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
)

type EasService struct {
	v1.UnimplementedEasServiceServer
	loginUc *biz.LoginUseCase
	userUc  *biz.UserUseCase
}

func NewEasService(loginUc *biz.LoginUseCase,
	userUc *biz.UserUseCase) *EasService {
	return &EasService{
		loginUc: loginUc,
		userUc:  userUc,
	}
}
