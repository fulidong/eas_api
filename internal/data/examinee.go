package data

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ExamineeRepo struct {
	data *Data
	log  *log.Helper
}

func ExamineRepo(data *Data, logger log.Logger) biz.ExamineeRepo {
	return &ExamineeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ExamineeRepo) GetPageList(ctx context.Context, in *v1.GetExamineePageListRequest) (res []*entity.Examinee, total int64, err error) {
	session := r.data.db.WithContext(ctx)
	session = session.Table((&entity.Examinee{}).TableName())
	q, v := r.buildConditions(in)
	if q != "" {
		session.Where(q, v...)
	}
	session.Count(&total)
	err = session.
		Order("updated_at desc,created_at desc").
		Offset(int((in.PageIndex - 1) * in.PageSize)).
		Limit(int(in.PageSize)).
		Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

func (r *ExamineeRepo) GetByEmails(ctx context.Context, email []string) (res []*entity.Examinee, err error) {
	err = r.data.db.WithContext(ctx).Model(res).Where(" email in ? ", email).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ExamineeRepo) GetByID(ctx context.Context, examineeId string) (resEntity *entity.Examinee, err error) {
	resEntity, err = getSingleRecordByScope[entity.Examinee](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" id = ? ", examineeId),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

func (r *ExamineeRepo) GetByIDs(ctx context.Context, examineeIds []string) (list []*entity.Examinee, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.Examinee{}).Where(" id in ? ", examineeIds).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 创建方法
func (r *ExamineeRepo) Save(ctx context.Context, addExaminees []*entity.Examinee, updateExaminees []*entity.Examinee, updatedBy string) error {
	if err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(addExaminees) > 0 {
			for _, examinee := range addExaminees {
				examinee.CreatedBy = updatedBy
				examinee.UpdatedBy = updatedBy
			}
			if e := tx.Create(addExaminees).Error; e != nil {
				return e
			}
		}
		if len(updateExaminees) > 0 {
			for _, comment := range updateExaminees {
				comment.UpdatedBy = updatedBy
			}
			if e := r.update(tx, updateExaminees); e != nil {
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
func (r *ExamineeRepo) update(tx *gorm.DB, updateExaminees []*entity.Examinee) error {
	ids := make([]string, 0, len(updateExaminees))
	for _, examinee := range updateExaminees {
		ids = append(ids, examinee.ID)
	}
	// 执行更新
	err := tx.Model(&entity.Examinee{}).
		Where(" id in ? ", ids).
		Updates(map[string]interface{}{
			"user_name": buildCaseExpr[*entity.Examinee](updateExaminees, func(comment *entity.Examinee) interface{} {
				return comment.ID
			}, func(comment *entity.Examinee) interface{} {
				return comment.UserName
			}),
			"phone": buildCaseExpr[*entity.Examinee](updateExaminees, func(comment *entity.Examinee) interface{} {
				return comment.ID
			}, func(comment *entity.Examinee) interface{} {
				return comment.Phone
			}),
			"updated_by": buildCaseExpr[*entity.Examinee](updateExaminees, func(comment *entity.Examinee) interface{} {
				return comment.ID
			}, func(comment *entity.Examinee) interface{} {
				return comment.UpdatedBy
			}),
		}).Error

	return err
}

// 更新方法
func (r *ExamineeRepo) Update(ctx context.Context, examinee *entity.Examinee) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"user_name":  examinee.UserName,
		"phone":      examinee.Phone,
		"updated_by": examinee.UpdatedBy,
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Examinee{}).
		Where(" id = ? ", examinee.ID).
		Updates(updates).Error

	return err
}

// 修改密码
func (r *ExamineeRepo) UpdateExamineePassWord(ctx context.Context, examineeId, passWord, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"hash_password": passWord,
		"updated_by":    updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Examinee{}).
		Where(" id = ? ", examineeId).
		Updates(updates).Error

	return err
}

func (r *ExamineeRepo) Delete(ctx context.Context, examineeId, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updatedBy,
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Examinee{}).
		Where(" id = ? ", examineeId).
		Updates(updates).Error

	return err
}

func (r *ExamineeRepo) buildConditions(in *v1.GetExamineePageListRequest) (string, []interface{}) {
	var (
		query strings.Builder
		value []interface{}
		q     string
	)
	if in.KeyWord != "" {
		query.WriteString(" email like ? ")
		keyWord := fmt.Sprintf("%%%s%%", strings.TrimSpace(in.KeyWord))
		value = append(value, keyWord)
		query.WriteString(" AND")
	}
	//if in.UserStatus >= 0 && in.UserStatus <= 1 {
	//	query.WriteString("  status=?")
	//	value = append(value, in.UserStatus)
	//	query.WriteString(" AND")
	//}

	// 过滤已删除
	query.WriteString(" deleted_at is NULL ")
	query.WriteString(" AND")

	if query.String() != "" {
		q = strings.TrimSuffix(query.String(), "AND")
	} else {
		q = query.String()
	}

	return q, value
}
