package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/middleware"
	"eas_api/internal/pkg/isecurity"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

//type LoginRepo interface {
//	GetByLoginAccount(ctx context.Context, login_account string) (*model.Administrator, error)
//}

type LoginUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewLoginUseCase(repo UserRepo, logger log.Logger) *LoginUseCase {
	return &LoginUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LoginUseCase) Login(ctx context.Context, req *v1.LoginRequest) (resp *v1.LoginResponse, err error) {
	resp = &v1.LoginResponse{}
	l := uc.log.WithContext(ctx)
	user, err := uc.repo.GetByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("Login.repo.GetByLoginAccount Failed, req:%v", req)
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
	if user.Status != int32(v1.AccountStatus_AccountStatus_Active) {
		err = errors.New("用户未激活")
	}
	// 生成jwt
	accessJWT, err := middleware.JWT.GenerateAccessToken(fmt.Sprintf("%d", user.ID), user.UserName, fmt.Sprintf("%d", user.UserType))
	if err != nil {
		// 处理错误
	}
	resp.UserName = user.UserName
	resp.Token = accessJWT
	return resp, nil
}
