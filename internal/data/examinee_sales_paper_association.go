package data

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"eas_api/internal/model"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ExamineeSalesPaperAssociationRepo struct {
	data *Data
	log  *log.Helper
}

func NewExamineeSalesPaperAssociationRepo(data *Data, logger log.Logger) biz.ExamineeSalesPaperAssociationRepo {
	return &ExamineeSalesPaperAssociationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 获取发放试卷列表
func (r *ExamineeSalesPaperAssociationRepo) GetProvidePageList(ctx context.Context, in *v1.GetProvidePageListRequest) (res []*model.ExamineeSalesPaperAssociation, total int64, err error) {
	session := r.data.db.WithContext(ctx)
	session = session.Table((&entity.ExamineeSalesPaperAssociation{}).TableName() + " as assoc").
		Select("assoc.*, examinee.user_name as examinee_name, examinee.email as examinee_email, examinee.phone as examinee_phone ").
		Joins(" inner join examinee on examinee.id = assoc.examinee_id ")
	q, v := r.buildConditions(in)
	if q != "" {
		session.Where(q, v...)
	}
	session.Count(&total)
	err = session.
		Order("assoc.updated_at desc,assoc.created_at desc").
		Offset(int((in.PageIndex - 1) * in.PageSize)).
		Limit(int(in.PageSize)).
		Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

func (r *ExamineeSalesPaperAssociationRepo) GetBySalesPaperIds(ctx context.Context, salesPaperIds []string) (list []*entity.ExamineeSalesPaperAssociation, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.ExamineeSalesPaperAssociation{}).Where(" sales_paper_id in ? ", salesPaperIds).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ExamineeSalesPaperAssociationRepo) GetByExamineeIds(ctx context.Context, examineeIds []string) (list []*entity.ExamineeSalesPaperAssociation, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.ExamineeSalesPaperAssociation{}).Where(" examinee_id in ? ", examineeIds).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 发放试卷
func (r *ExamineeSalesPaperAssociationRepo) Provide(ctx context.Context, associationEntities []*entity.ExamineeSalesPaperAssociation,
	examineeEmailRecords []*entity.ExamineeEmailRecord) (err error) {
	if len(associationEntities) == 0 || len(examineeEmailRecords) == 0 {
		return
	}
	return r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(associationEntities) > 0 {
			if e := tx.Create(associationEntities).Error; e != nil {
				return e
			}
		}
		if len(examineeEmailRecords) > 0 {
			if e := tx.Create(examineeEmailRecords).Error; e != nil {
				return e
			}
		}
		return nil
	})
}

// 更新邮件状态
func (r *ExamineeSalesPaperAssociationRepo) UpdateEmailStatus(ctx context.Context, examineeSalesPaperAssociationIds []string) error {
	// 执行更新
	return r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		e := tx.Model(&entity.ExamineeEmailRecord{}).
			Where(" examinee_sales_paper_association_id in ? ", examineeSalesPaperAssociationIds).
			Updates(map[string]interface{}{
				"is_sended": 1,
				"send_time": time.Now(),
			}).Error
		if e != nil {
			return e
		}
		e = tx.Model(&entity.ExamineeSalesPaperAssociation{}).
			Where(" id in ? ", examineeSalesPaperAssociationIds).
			Updates(map[string]interface{}{
				"email_status": v1.EmailStatus_Send,
			}).Error
		if e != nil {
			return e
		}
		return nil
	})
}

// 更新方法
func (r *ExamineeSalesPaperAssociationRepo) UpdateReport(tx *gorm.DB, reportPath, examineeSalesPaperAssociationId string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"report_path": reportPath,
	}
	// 执行更新
	err := tx.Model(&entity.ExamineeSalesPaperAssociation{}).
		Where(" id = ? ", examineeSalesPaperAssociationId).
		Updates(updates).Error

	return err
}

func (r *ExamineeSalesPaperAssociationRepo) Delete(ctx context.Context, examineeSalesPaperAssociationId, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updatedBy,
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.ExamineeSalesPaperAssociation{}).
		Where(" id = ? ", examineeSalesPaperAssociationId).
		Updates(updates).Error

	return err
}

func (r *ExamineeSalesPaperAssociationRepo) buildConditions(in *v1.GetProvidePageListRequest) (string, []interface{}) {
	var (
		query strings.Builder
		value []interface{}
		q     string
	)
	if in.SalesPaperId != "" {
		query.WriteString(" assoc.sales_paper_id = ? ")
		value = append(value, in.SalesPaperId)
		query.WriteString(" AND")
	}
	if in.StageNumber >= 0 && in.StageNumber <= int32(v1.StageNumber_Report) {
		query.WriteString("  assoc.stage_number = ?")
		value = append(value, in.StageNumber)
		query.WriteString(" AND")
	}
	if in.KeyWord != "" {
		query.WriteString(" (examinee.user_name like ? or examinee.email like ? or examinee.phone like ?)")
		keyWord := fmt.Sprintf("%%%s%%", strings.TrimSpace(in.KeyWord))
		value = append(value, keyWord, keyWord, keyWord, keyWord)
		query.WriteString(" AND")
	}
	// 过滤已删除
	query.WriteString(" assoc.deleted_at is NULL ")
	query.WriteString(" AND")

	if query.String() != "" {
		q = strings.TrimSuffix(query.String(), "AND")
	} else {
		q = query.String()
	}

	return q, value
}
