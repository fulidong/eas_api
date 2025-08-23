package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/middleware"
	"eas_api/internal/pkg/isecurity"
	"eas_api/internal/pkg/isnowflake"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type SysLoginRepo interface {
	Create(ctx context.Context, entity *entity.SysLoginRecord) error
}

type LoginUseCase struct {
	repo     UserRepo
	sysLogin SysLoginRepo
	log      *log.Helper
}

func NewLoginUseCase(repo UserRepo, sysLogin SysLoginRepo, logger log.Logger) *LoginUseCase {
	return &LoginUseCase{repo: repo, sysLogin: sysLogin, log: log.NewHelper(logger)}
}

func (uc *LoginUseCase) Login(ctx context.Context, req *v1.LoginRequest) (resp *v1.LoginResponse, err error) {
	resp = &v1.LoginResponse{}
	l := uc.log.WithContext(ctx)
	user, err := uc.repo.GetByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("Login.repo.GetByLoginAccount Failed, req:%v, err:%v", req, err.Error())
		return nil, err
	}
	if user == nil {
		return resp, errors.New("用户不存在")
	}
	// 验证密码
	if req.PassWord != "" {
		ok := isecurity.CheckPassword(req.PassWord, user.HashPassword)
		if !ok {
			err = errors.New("密码错误")
			return
		}
	}
	if user.Status != int32(v1.AccountStatus_Active) {
		err = errors.New("用户未激活")
		return
	}
	id, _ := isnowflake.SnowFlake.NextID(_const.SysLoginRecordPrefix)
	uc.sysLogin.Create(ctx, &entity.SysLoginRecord{
		ID:            id,
		UserID:        user.ID,
		LoginPlatform: int32(v1.LoginPlatform_Management),
	})
	// 生成jwt
	accessJWT, err := middleware.JWT.GenerateAccessToken(user.ID, user.UserName, fmt.Sprintf("%d", user.UserType))
	if err != nil {
		// 处理错误
		err = errors.New("登录失败")
		return
	}
	resp.UserName = user.UserName
	resp.UserType = v1.UserType(user.UserType)
	resp.Token = accessJWT
	return resp, nil
}
