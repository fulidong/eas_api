package data

import (
	"context"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type QuestionRepo struct {
	data *Data
	log  *log.Helper
}

func NewQuestionRepo(data *Data, logger log.Logger) biz.QuestionRepo {
	return &QuestionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *QuestionRepo) GetList(ctx context.Context, salesPaperId, dimensionId string) (res []*entity.Question, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.Question{}).
		Where(" sales_paper_id = ? and dimension_id = ? ", salesPaperId, dimensionId).
		Order(" `order` asc").
		Find(&res).Error
	if err != nil {
		return
	}

	return
}

func (r *QuestionRepo) GetListBySalesPaperId(ctx context.Context, salesPaperId string) (res []*entity.Question, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.Question{}).
		Where(" sales_paper_id = ? ", salesPaperId).
		Order(" `order` asc").
		Find(&res).Error
	if err != nil {
		return
	}

	return
}

func (r *QuestionRepo) GetOptionList(ctx context.Context, questionId string) (res []*entity.QuestionOption, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.QuestionOption{}).
		Where(" question_id = ?", questionId).
		Order(" `order` asc").
		Find(&res).Error
	if err != nil {
		return
	}

	return
}

func (r *QuestionRepo) GetOptionListByQuestionIds(ctx context.Context, questionIds []string) (res map[string][]*entity.QuestionOption, err error) {
	qOptions := make([]*entity.QuestionOption, 0, 10)
	res = make(map[string][]*entity.QuestionOption)
	err = r.data.db.WithContext(ctx).Model(&entity.QuestionOption{}).
		Where(" question_id in ?", questionIds).
		Order(" `order` asc").
		Find(&qOptions).Error
	if err != nil {
		return
	}
	for _, option := range qOptions {
		value, ok := res[option.QuestionID]
		if !ok {
			res[option.QuestionID] = []*entity.QuestionOption{option}
		} else {
			res[option.QuestionID] = append(value, option)
		}
	}
	return
}
func (r *QuestionRepo) GetById(ctx context.Context, questionId string) (qEntity *entity.Question, qOptionsEntities []*entity.QuestionOption, err error) {
	d := r.data.db.WithContext(ctx)
	qEntity, err = getSingleRecordByScope[entity.Question](
		d.Model(qEntity).Where(" id = ? ", questionId),
	)
	err = d.Model(qOptionsEntities).Where(" question_id = ? ", questionId).Find(&qOptionsEntities).Error
	if err != nil {
		return
	}
	return
}

// 创建方法
func (r *QuestionRepo) Save(ctx context.Context, question *entity.Question, addOptions, updateOptions []*entity.QuestionOption, delOptions []string, updatedBy string) error {
	if err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if e := tx.Save(question).Error; e != nil {
			return e
		}

		if len(addOptions) > 0 {
			for _, comment := range addOptions {
				comment.CreatedBy = updatedBy
				comment.UpdatedBy = updatedBy
			}
			if e := r.createQuestionOptions(tx, addOptions); e != nil {
				return e
			}
		}
		if len(updateOptions) > 0 {
			for _, comment := range updateOptions {
				comment.UpdatedBy = updatedBy
			}
			if e := r.updateQuestionOptions(tx, updateOptions); e != nil {
				return e
			}
		}
		if len(delOptions) > 0 {
			if e := r.deleteQuestionOptions(tx, delOptions, updatedBy); e != nil {
				return e
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// 创建
func (r *QuestionRepo) createQuestionOptions(tx *gorm.DB, QuestionOptions []*entity.QuestionOption) error {
	// 执行更新
	err := tx.Create(QuestionOptions).Error
	return err
}

// 更新方法
func (r *QuestionRepo) updateQuestionOptions(tx *gorm.DB, updateQuestions []*entity.QuestionOption) error {
	ids := make([]string, 0, len(updateQuestions))
	for _, comment := range updateQuestions {
		ids = append(ids, comment.ID)
	}
	// 执行更新
	err := tx.Model(&entity.QuestionOption{}).
		Where(" id in ? ", ids).
		Updates(map[string]interface{}{
			"order": buildCaseExpr[*entity.QuestionOption](updateQuestions, func(comment *entity.QuestionOption) interface{} {
				return comment.ID
			}, func(comment *entity.QuestionOption) interface{} {
				return comment.Order_
			}),
			"score": buildCaseExpr[*entity.QuestionOption](updateQuestions, func(comment *entity.QuestionOption) interface{} {
				return comment.ID
			}, func(comment *entity.QuestionOption) interface{} {
				return comment.Score
			}),
			"description": buildCaseExpr[*entity.QuestionOption](updateQuestions, func(comment *entity.QuestionOption) interface{} {
				return comment.ID
			}, func(comment *entity.QuestionOption) interface{} {
				return comment.Description
			}),
			"updated_by": buildCaseExpr[*entity.QuestionOption](updateQuestions, func(comment *entity.QuestionOption) interface{} {
				return comment.ID
			}, func(comment *entity.QuestionOption) interface{} {
				return comment.UpdatedBy
			}),
		}).Error

	return err
}

func (r *QuestionRepo) Delete(ctx context.Context, questionId string, updatedBy string) error {
	if err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if e := r.deleteQuestion(tx, questionId, updatedBy); e != nil {
			return e
		}
		if e := r.deleteQuestionOptionsByQuestionId(tx, questionId, updatedBy); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// 删除题目
func (r *QuestionRepo) deleteQuestion(tx *gorm.DB, QuestionId string, updateBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updateBy,
	}

	// 执行更新
	err := tx.Model(&entity.Question{}).
		Where(" id = ? ", QuestionId).
		Updates(updates).Error

	return err
}

// 删除题目选项
func (r *QuestionRepo) deleteQuestionOptionsByQuestionId(tx *gorm.DB, QuestionId string, updateBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updateBy,
	}

	// 执行更新
	err := tx.Model(&entity.QuestionOption{}).
		Where(" question_id = ? ", QuestionId).
		Updates(updates).Error

	return err
}

// 删除题目选项
func (r *QuestionRepo) deleteQuestionOptions(tx *gorm.DB, QuestionOptionId []string, updateBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updateBy,
	}

	// 执行更新
	err := tx.Model(&entity.QuestionOption{}).
		Where(" id in ? ", QuestionOptionId).
		Updates(updates).Error

	return err
}
