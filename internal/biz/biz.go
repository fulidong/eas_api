package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewLoginUseCase,
	NewUserUseCase,
	NewSalesPaperUseCase,
	NewSalesPaperCommentUseCase,
	NewSalesPaperDimensionUseCase,
	NewSalesPaperDimensionCommentUseCase,
	NewQuestionUseCase,
	NewExamineeUseCase,
	NewExamineeSalesPaperAssociationUseCase)
