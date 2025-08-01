package service

import (
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
)

type EasExamineeService struct {
	v1.UnimplementedEasExamineeServiceServer
	userUc                          *biz.UserUseCase
	examineeUc                      *biz.ExamineeUseCase
	examineeSalesPaperAssociationUc *biz.ExamineeSalesPaperAssociationUseCase
}

func NewEasExamineeService(
	userUc *biz.UserUseCase,
	examineeUc *biz.ExamineeUseCase,
	examineeSalesPaperAssociationUc *biz.ExamineeSalesPaperAssociationUseCase) *EasExamineeService {
	return &EasExamineeService{
		userUc:                          userUc,
		examineeUc:                      examineeUc,
		examineeSalesPaperAssociationUc: examineeSalesPaperAssociationUc,
	}
}
