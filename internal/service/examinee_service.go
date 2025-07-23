package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasExamineeService) SaveExaminee(ctx context.Context, in *v1.SaveExamineeRequest) (*v1.SaveExamineeResponse, error) {
	return s.examineeUc.SaveExaminee(ctx, in)
}

func (s *EasExamineeService) GetExamineePageList(ctx context.Context, in *v1.GetExamineePageListRequest) (*v1.GetExamineePageListResponse, error) {
	return s.examineeUc.GetExamineePageList(ctx, in)
}

func (s *EasExamineeService) GetExamineeDetail(ctx context.Context, in *v1.GetExamineeDetailRequest) (*v1.GetExamineeDetailResponse, error) {
	return s.examineeUc.GetExamineeDetail(ctx, in)
}

func (s *EasExamineeService) UpdateExaminee(ctx context.Context, in *v1.UpdateExamineeRequest) (*v1.UpdateExamineeResponse, error) {
	return s.examineeUc.UpdateExaminee(ctx, in)
}

func (s *EasExamineeService) DeleteExaminee(ctx context.Context, in *v1.DeleteExamineeRequest) (*v1.DeleteExamineeResponse, error) {
	return s.examineeUc.DeleteExaminee(ctx, in)
}
