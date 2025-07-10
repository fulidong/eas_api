package data

import (
	"context"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type SalesPaperCommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewSalesPaperCommentRepo(data *Data, logger log.Logger) biz.SalesPaperCommentRepo {
	return &SalesPaperCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SalesPaperCommentRepo) GetSalesPaperCommentList(ctx context.Context, salesPaperId int64) (res []*entity.SalesPaperComment, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.SalesPaperComment{}).
		Where(" sales_paper_id = ?", salesPaperId).
		Order("up_score asc").
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 创建方法
func (r *SalesPaperCommentRepo) SaveSalesPaperComment(ctx context.Context, addComments []*entity.SalesPaperComment, updateComments []*entity.SalesPaperComment, delComments []int64, updatedBy int64) error {
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
func (r *SalesPaperCommentRepo) update(tx *gorm.DB, updateSalesPaperComments []*entity.SalesPaperComment) error {
	ids := make([]int64, 0, len(updateSalesPaperComments))
	for _, comment := range updateSalesPaperComments {
		ids = append(ids, comment.ID)
	}
	// 执行更新
	err := tx.Model(&entity.SalesPaperComment{}).
		Where(" id in ? ", ids).
		Updates(map[string]interface{}{
			"content": buildCaseExpr[*entity.SalesPaperComment](updateSalesPaperComments, func(comment *entity.SalesPaperComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperComment) interface{} {
				return comment.Content
			}),
			"up_score": buildCaseExpr[*entity.SalesPaperComment](updateSalesPaperComments, func(comment *entity.SalesPaperComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperComment) interface{} {
				return comment.UpScore
			}),
			"low_score": buildCaseExpr[*entity.SalesPaperComment](updateSalesPaperComments, func(comment *entity.SalesPaperComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperComment) interface{} {
				return comment.LowScore
			}),
			"updated_by": buildCaseExpr[*entity.SalesPaperComment](updateSalesPaperComments, func(comment *entity.SalesPaperComment) interface{} {
				return comment.ID
			}, func(comment *entity.SalesPaperComment) interface{} {
				return comment.UpdatedBy
			}),
		}).Error

	return err
}

// 删除
func (r *SalesPaperCommentRepo) delete(tx *gorm.DB, salesPaperCommentId []int64, updateBy int64) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updateBy,
	}

	// 执行更新
	err := tx.Model(&entity.SalesPaperComment{}).
		Where(" id in ? ", salesPaperCommentId).
		Updates(updates).Error

	return err
}
