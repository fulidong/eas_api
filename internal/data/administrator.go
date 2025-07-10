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

type AdministratorRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &AdministratorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AdministratorRepo) GetPageList(ctx context.Context, in *v1.GetPageListRequest) (res []*entity.Administrator, total int64, err error) {
	session := r.data.db.WithContext(ctx)
	session = session.Table((&entity.Administrator{}).TableName())
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

func (r *AdministratorRepo) GetByLoginAccount(ctx context.Context, loginAccount string) (resEntity *entity.Administrator, err error) {

	resEntity, err = getSingleRecordByScope[entity.Administrator](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" login_account = ? ", loginAccount),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

func (r *AdministratorRepo) GetListByLoginAccount(ctx context.Context, loginAccount string) (list []*entity.Administrator, err error) {
	err = r.data.db.WithContext(ctx).Model(list).Where(" login_account = ? ", loginAccount).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *AdministratorRepo) GetByUserName(ctx context.Context, userName string) (resEntity *entity.Administrator, err error) {
	resEntity, err = getSingleRecordByScope[entity.Administrator](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" user_name = ? ", userName),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

func (r *AdministratorRepo) GetByID(ctx context.Context, userId string) (resEntity *entity.Administrator, err error) {
	resEntity, err = getSingleRecordByScope[entity.Administrator](
		r.data.db.WithContext(ctx).Model(resEntity).Where(" id = ? ", userId),
	)
	if err != nil {
		return nil, err
	}
	return resEntity, nil
}

func (r *AdministratorRepo) GetByIDs(ctx context.Context, userId []string) (list []*entity.Administrator, err error) {
	err = r.data.db.WithContext(ctx).Model(&entity.Administrator{}).Where(" id in ? ", userId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *AdministratorRepo) Create(ctx context.Context, admin *entity.Administrator) error {
	return r.data.db.WithContext(ctx).Create(admin).Error
}

// 更新方法
func (r *AdministratorRepo) Update(ctx context.Context, user *entity.Administrator, isOwn bool) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"user_name":     user.UserName,
		"login_account": user.LoginAccount,
		"email":         user.Email,
		"updated_by":    user.UpdatedBy,
	}
	if !isOwn {
		updates["status"] = user.Status
		updates["user_type"] = user.UserType
	}
	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Administrator{}).
		Where(" id = ? ", user.ID).
		Updates(updates).Error

	return err
}

// 更新激活状态
func (r *AdministratorRepo) SetUserStatus(ctx context.Context, userId string, userStatus v1.AccountStatus, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"status":     userStatus,
		"updated_by": updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Administrator{}).
		Where(" id = ? ", userId).
		Updates(updates).Error

	return err
}

// 修改密码
func (r *AdministratorRepo) UpdateUserPassWord(ctx context.Context, userId, passWord, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"hash_password": passWord,
		"updated_by":    updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Administrator{}).
		Where(" id = ? ", userId).
		Updates(updates).Error

	return err
}

// 删除用户
func (r *AdministratorRepo) DeleteUser(ctx context.Context, userId, updatedBy string) error {
	// 准备更新字段
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_by": updatedBy,
	}

	// 执行更新
	err := r.data.db.WithContext(ctx).Model(&entity.Administrator{}).
		Where(" id = ? ", userId).
		Updates(updates).Error

	return err
}

func (r *AdministratorRepo) buildConditions(in *v1.GetPageListRequest) (string, []interface{}) {
	var (
		query strings.Builder
		value []interface{}
		q     string
	)
	if in.KeyWord != "" {
		query.WriteString("  (user_name like ? OR login_account like ? )")
		keyWord := fmt.Sprintf("%%%s%%", strings.TrimSpace(in.KeyWord))
		value = append(value, keyWord, keyWord)
		query.WriteString(" AND")
	}
	if in.UserStatus >= 0 && in.UserStatus <= 1 {
		query.WriteString("  status=?")
		value = append(value, in.UserStatus)
		query.WriteString(" AND")
	}
	if in.UserType >= 0 && in.UserType <= 1 {
		query.WriteString("  user_type=?")
		value = append(value, in.UserType)
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
