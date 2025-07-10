package data

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
	"time"
)

type SalesPaperRepo struct {
	data *Data
	log  *log.Helper
}

func NewSalesPaperRepo(data *Data, logger log.Logger) biz.SalesPaperRepo {
	return &SalesPaperRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SalesPaperRepo) GetPageList(ctx context.Context, in *v1.GetSalesPaperPageListRequest) (res []*entity.SalesPaper, total int64, err error) {
	session := r.data.db.WithContext(ctx)
	session = session.Table((&entity.SalesPaper{}).TableName())
	q, v := r.buildConditions(in.KeyWord, in.SalesPaperStatus)
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

func (r *SalesPaperRepo) GetUsablePageList(ctx context.Context, in *v1.GetUsableSalesPaperPageListRequest) (res []*entity.SalesPaper, total int64, err error) {
	session := r.data.db.WithContext(ctx)
	session = session.Table((&entity.SalesPaper{}).TableName())
	q, v := r.buildConditions(in.KeyWord, int64(v1.SalesPaperStatus_Enable))
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
func (r *SalesPaperRepo) GetBySalesPaperName(ctx context.Context, salesPaperName string) (list []*entity.SalesPaper, err error) {
	err = r.data.db.WithContext(ctx).Model(list).Where(" name = ? ", salesPaperName).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *SalesPaperRepo) GetByID(ctx context.Context, salesPaperId int64) (resEntity *entity.SalesPaper, err error) {
	resEntity, err = getSingleRecordByScope[entity.SalesPaper](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" id = ? ", salesPaperId),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

// 创建方法
func (r *SalesPaperRepo) Create(ctx context.Context, salesPaper *entity.SalesPaper) error {
	return r.data.db.WithContext(ctx).Create(salesPaper).Error
}

// 更新方法
func (r *SalesPaperRepo) Update(ctx context.Context, salesPaper *entity.SalesPaper) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"name":               salesPaper.Name,
		"recommend_time_lim": salesPaper.RecommendTimeLim,
		"max_score":          salesPaper.MaxScore,
		"min_score":          salesPaper.MinScore,
		"is_enabled":         salesPaper.IsEnabled,
		"mark":               salesPaper.Mark,
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.SalesPaper{}).
		Where(" id = ? and is_used = 0 ", salesPaper.ID).
		Updates(updates).Error

	return err
}

// 更新激活状态
func (r *SalesPaperRepo) SetSalesPaperStatus(ctx context.Context, salesPaperId int64, salesPaperStatus v1.SalesPaperStatus, updatedBy int64) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"is_enabled": salesPaperStatus,
		"updated_by": updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.SalesPaper{}).
		Where(" id = ? and is_used = 0 ", salesPaperId).
		Updates(updates).Error

	return err
}

// 删除用户
func (r *SalesPaperRepo) DeleteSalesPaper(ctx context.Context, salesPaperId, updatedBy int64) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.SalesPaper{}).
		Where(" id = ? and is_used = 0 ", salesPaperId).
		Updates(updates).Error

	return err
}

func (r *SalesPaperRepo) buildConditions(keyWord string, salesPagerStatus int64) (string, []interface{}) {
	var (
		query strings.Builder
		value []interface{}
		q     string
	)
	if keyWord != "" {
		query.WriteString(" name like ? ")
		keyWord = fmt.Sprintf("%%%s%%", strings.TrimSpace(keyWord))
		value = append(value, keyWord)
		query.WriteString(" AND")
	}
	if salesPagerStatus >= 0 && salesPagerStatus <= 1 {
		query.WriteString("  is_enabled=?")
		value = append(value, salesPagerStatus)
		query.WriteString(" AND")
	}

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
