package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasSalesPaperService) SaveSalesPaperDimensionComment(ctx context.Context, in *v1.SaveSalesPaperDimensionCommentRequest) (*v1.SaveSalesPaperDimensionCommentResponse, error) {
	return s.salesPaperDimensionCommentUc.SaveSalesPaperDimensionComment(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionCommentList(ctx context.Context, in *v1.GetSalesPaperDimensionCommentListRequest) (*v1.GetSalesPaperDimensionCommentListResponse, error) {
	return s.salesPaperDimensionCommentUc.GetSalesPaperDimensionCommentList(ctx, in)
}
