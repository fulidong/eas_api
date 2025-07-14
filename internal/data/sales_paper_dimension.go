package data

import (
	"context"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type SalesPaperDimensionRepo struct {
	data *Data
	log  *log.Helper
}

func NewSalesPaperDimensionRepo(data *Data, logger log.Logger) biz.SalesPaperDimensionRepo {
	return &SalesPaperDimensionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SalesPaperDimensionRepo) GetList(ctx context.Context, salesPaperId string) (res []*entity.SalesPaperDimension, err error) {
	err = r.data.db.WithContext(ctx).Model(res).Where(" sales_paper_id = ? ", salesPaperId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 创建
func (r *SalesPaperDimensionRepo) Create(ctx context.Context, dimensions []*entity.SalesPaperDimension) error {
	return r.data.db.WithContext(ctx).Create(dimensions).Error
}

func (r *SalesPaperDimensionRepo) GetById(ctx context.Context, dimensionId string) (resEntity *entity.SalesPaperDimension, err error) {
	resEntity, err = getSingleRecordByScope[entity.SalesPaperDimension](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" id = ? ", dimensionId),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

// 更新方法
func (r *SalesPaperDimensionRepo) Update(ctx context.Context, updateSalesPaperDimensions []*entity.SalesPaperDimension) error {
	ids := make([]string, 0, len(updateSalesPaperDimensions))
	for _, dimension := range updateSalesPaperDimensions {
		ids = append(ids, dimension.ID)
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.SalesPaperDimension{}).
		Where(" id in ? ", ids).
		Updates(map[string]interface{}{
			"name": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.Name
			}),
			"average_mark": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.AverageMark
			}),
			"standard_mark": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.StandardMark
			}),
			"description": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.Description
			}),
			"max_score": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.MaxScore
			}),
			"min_score": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.MinScore
			}),
			"is_choose": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.IsChoose
			}),
			"updated_by": buildCaseExpr[*entity.SalesPaperDimension](updateSalesPaperDimensions, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.ID
			}, func(dimension *entity.SalesPaperDimension) interface{} {
				return dimension.UpdatedBy
			}),
		}).Error

	return err
}

// 删除
func (r *SalesPaperDimensionRepo) Delete(ctx context.Context, salesPaperDimensionId, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.SalesPaperDimension{}).
		Where(" id = ? ", salesPaperDimensionId).
		Updates(updates).Error

	return err
}
