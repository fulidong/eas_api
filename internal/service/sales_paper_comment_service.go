package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasService) SaveSalesPaperComment(ctx context.Context, in *v1.SaveSalesPaperCommentRequest) (*v1.SaveSalesPaperCommentResponse, error) {
	return s.salesPaperCommentUc.SaveSalesPaperComment(ctx, in)
}

func (s *EasService) GetSalesPaperCommentList(ctx context.Context, in *v1.GetSalesPaperCommentListRequest) (*v1.GetSalesPaperCommentListResponse, error) {
	return s.salesPaperCommentUc.GetSalesPaperCommentList(ctx, in)
}
