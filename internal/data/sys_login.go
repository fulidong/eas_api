package data

import (
	"context"
	"eas_api/internal/biz"
	"eas_api/internal/data/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type SysLoginRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysLoginRepo(data *Data, logger log.Logger) biz.SysLoginRepo {
	return &SysLoginRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建方法
func (r *SysLoginRepo) Create(ctx context.Context, admin *entity.SysLoginRecord) error {
	return r.data.db.WithContext(ctx).Create(admin).Error
}
