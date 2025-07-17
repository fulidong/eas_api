package _const

const (
	AdministratorPrefix                     = "A"
	SysLoginRecordPrefix                    = "SLR"
	SalesPaperPrefix                        = "SPP"
	SalesPaperCommentPrefix                 = "SPCP"
	SalesPaperDimensionPrefix               = "SPDP"
	SalesPaperDimensionCommentPrefix        = "SPDCP"
	SalesPaperDimensionQuestionPrefix       = "SPDQP"
	SalesPaperDimensionQuestionOptionPrefix = "SPDQOP"
)

var AllowedVars = map[string]interface{}{
	"raw_score":     0.0,
	"average_mark":  0.0,
	"standard_mark": 0.0,
}
