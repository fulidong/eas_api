package data

import (
	"context"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type SalesPaperDimensionCommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewSalesPaperDimensionCommentRepo(data *Data, logger log.Logger) biz.SalesPaperDimensionCommentRepo {
	return &SalesPaperDimensionCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SalesPaperDimensionCommentRepo) GetList(ctx context.Context, salesPaperDimensionId string) (res []*entity.SalesPaperDimensionComment, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.SalesPaperDimensionComment{}).
		Where(" sales_paper_dimension_id = ?", salesPaperDimensionId).
		Order("up_score asc").
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 创建方法
func (r *SalesPaperDimensionCommentRepo) Save(ctx context.Context, addComments []*entity.SalesPaperDimensionComment, updateComments []*entity.SalesPaperDimensionComment, delComments []string, updatedBy string) error {
	if err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(addComments) > 0 {
			for _, comment := range addComments {
				comment.CreatedBy = updatedBy
				comment.UpdatedBy = updatedBy
			}
			if e := tx.Create(addComments).Error; e != nil {
				return e
			}
		}
		if len(updateComments) > 0 {
			for _, comment := range updateComments {
				comment.UpdatedBy = updatedBy
			}
			if e := r.update(tx, updateComments); e != nil {
				return e
			}
		}
		if len(delComments) > 0 {
			if e := r.delete(tx, delComments, updatedBy); e != nil {
				return e
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// 更新方法
func (r *SalesPaperDimensionCommentRepo) update(tx *gorm.DB, updateSalesPaperDimensionComments []*entity.SalesPaperDimensionComment) error {
	ids := make([]string, 0, len(updateSalesPaperDimensionComments))
	for _, comment := range updateSalesPaperDimensionComments {
		ids = append(ids, comment.ID)
	}
	// 执行更新
	err := tx.Model(&entity.SalesPaperDimensionComment{}).
		Where(" id in ? ", ids).
		Updates(map[string]interface{}{
			"content": buildCaseExpr[*entity.SalesPaperDimensionComment](updateSalesPaperDimensionComments, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.Content
			}),
			"up_score": buildCaseExpr[*entity.SalesPaperDimensionComment](updateSalesPaperDimensionComments, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.UpScore
			}),
			"low_score": buildCaseExpr[*entity.SalesPaperDimensionComment](updateSalesPaperDimensionComments, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.LowScore
			}),
			"updated_by": buildCaseExpr[*entity.SalesPaperDimensionComment](updateSalesPaperDimensionComments, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperDimensionComment) interface{} {
				return comment.UpdatedBy
			}),
		}).Error

	return err
}

// 删除
func (r *SalesPaperDimensionCommentRepo) delete(tx *gorm.DB, SalesPaperDimensionCommentId []string, updateBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updateBy,
	}

	// 执行更新
	err := tx.Model(&entity.SalesPaperDimensionComment{}).
		Where(" id in ? ", SalesPaperDimensionCommentId).
		Updates(updates).Error

	return err
}
