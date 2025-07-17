package service

import (
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
)

type EasSalesPaperService struct {
	v1.UnimplementedEasSalesPaperServiceServer
	userUc                       *biz.UserUseCase
	salesPaperUc                 *biz.SalesPaperUseCase
	salesPaperCommentUc          *biz.SalesPaperCommentUseCase
	salesPaperDimensionUc        *biz.SalesPaperDimensionUseCase
	salesPaperDimensionCommentUc *biz.SalesPaperDimensionCommentUseCase
	questionUc                   *biz.QuestionUseCase
}

func NewEasSalesPaperService(
	userUc *biz.UserUseCase,
	salesPaperUc *biz.SalesPaperUseCase,
	salesPaperCommentUc *biz.SalesPaperCommentUseCase,
	salesPaperDimensionUc *biz.SalesPaperDimensionUseCase,
	salesPaperDimensionCommentUc *biz.SalesPaperDimensionCommentUseCase,
	questionUc *biz.QuestionUseCase) *EasSalesPaperService {
	return &EasSalesPaperService{
		userUc:                       userUc,
		salesPaperUc:                 salesPaperUc,
		salesPaperCommentUc:          salesPaperCommentUc,
		salesPaperDimensionUc:        salesPaperDimensionUc,
		salesPaperDimensionCommentUc: salesPaperDimensionCommentUc,
		questionUc:                   questionUc,
	}
}
