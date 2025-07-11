package service

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
)

func (s *EasSalesPaperService) CreateSalesPaper(ctx context.Context, in *v1.CreateSalesPaperRequest) (*v1.CreateSalesPaperResponse, error) {
	return s.salesPaperUc.CreateSalesPaper(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperPageList(ctx context.Context, in *v1.GetSalesPaperPageListRequest) (*v1.GetSalesPaperPageListResponse, error) {
	return s.salesPaperUc.GetSalesPaperPageList(ctx, in)
}

func (s *EasSalesPaperService) GetUsableSalesPaperPageList(ctx context.Context, in *v1.GetUsableSalesPaperPageListRequest) (*v1.GetUsableSalesPaperPageListResponse, error) {
	return s.salesPaperUc.GetUsableSalesPaperPageList(ctx, in)
}

func (s *EasSalesPaperService) GetSalesPaperDetail(ctx context.Context, in *v1.GetSalesPaperDetailRequest) (*v1.GetSalesPaperDetailResponse, error) {
	return s.salesPaperUc.GetSalesPaperDetail(ctx, in)
}

func (s *EasSalesPaperService) UpdateSalesPaper(ctx context.Context, in *v1.UpdateSalesPaperRequest) (*v1.UpdateSalesPaperResponse, error) {
	return s.salesPaperUc.UpdateSalesPaper(ctx, in)
}

func (s *EasSalesPaperService) SetSalesPaperStatus(ctx context.Context, in *v1.SetSalesPaperStatusRequest) (*v1.SetSalesPaperStatusResponse, error) {
	return s.salesPaperUc.SetSalesPaperStatus(ctx, in)
}

func (s *EasSalesPaperService) DeleteSalesPaper(ctx context.Context, in *v1.DeleteSalesPaperRequest) (*v1.DeleteSalesPaperResponse, error) {
	return s.salesPaperUc.DeleteSalesPaper(ctx, in)
}
