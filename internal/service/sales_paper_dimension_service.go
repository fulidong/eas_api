package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasSalesPaperService) CreateSalesPaperDimension(ctx context.Context, in *v1.CreateSalesPaperDimensionRequest) (*v1.CreateSalesPaperDimensionResponse, error) {
	return s.salesPaperDimensionUc.CreateSalesPaperDimension(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionList(ctx context.Context, in *v1.GetSalesPaperDimensionListRequest) (*v1.GetSalesPaperDimensionListResponse, error) {
	return s.salesPaperDimensionUc.GetSalesPaperDimensionList(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDimensionDetail(ctx context.Context, in *v1.GetSalesPaperDimensionDetailRequest) (*v1.GetSalesPaperDimensionDetailResponse, error) {
	return s.salesPaperDimensionUc.GetSalesPaperDimensionDetail(ctx, in)
}

func (s *EasSalesPaperService) UpdateSalesPaperDimension(ctx context.Context, in *v1.UpdateSalesPaperDimensionRequest) (*v1.UpdateSalesPaperDimensionResponse, error) {
	return s.salesPaperDimensionUc.UpdateSalesPaperDimension(ctx, in)
}

func (s *EasSalesPaperService) DeleteSalesPaperDimension(ctx context.Context, in *v1.DeleteSalesPaperDimensionRequest) (*v1.DeleteSalesPaperDimensionResponse, error) {
	return s.salesPaperDimensionUc.DeleteSalesPaperDimension(ctx, in)
}
