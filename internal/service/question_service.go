package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasSalesPaperService) SaveSalesPaperDimensionQuestion(ctx context.Context, in *v1.SaveSalesPaperDimensionQuestionRequest) (*v1.SaveSalesPaperDimensionQuestionResponse, error) {
	return s.questionUc.SaveSalesPaperDimensionQuestion(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionQuestionList(ctx context.Context, in *v1.GetSalesPaperDimensionQuestionListRequest) (*v1.GetSalesPaperDimensionQuestionListResponse, error) {
	return s.questionUc.GetSalesPaperDimensionQuestionList(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionQuestionDetail(ctx context.Context, in *v1.GetSalesPaperDimensionQuestionDetailRequest) (*v1.GetSalesPaperDimensionQuestionDetailResponse, error) {
	return s.questionUc.GetSalesPaperDimensionQuestionDetail(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionQuestionPreView(ctx context.Context, in *v1.GetSalesPaperDimensionQuestionPreViewRequest) (*v1.GetSalesPaperDimensionQuestionPreViewResponse, error) {
	return s.questionUc.GetSalesPaperDimensionQuestionPreView(ctx, in)
}

func (s *EasSalesPaperService) DeleteSalesPaperDimensionQuestion(ctx context.Context, in *v1.DeleteSalesPaperDimensionQuestionRequest) (*v1.DeleteSalesPaperDimensionQuestionResponse, error) {
	return s.questionUc.DeleteSalesPaperDimensionQuestion(ctx, in)
}
