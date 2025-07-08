package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/data/entity"
	"eas_api/internal/pkg/icontext"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/isecurity"
	"eas_api/internal/pkg/isnowflake"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type UserRepo interface {
	Create(ctx context.Context, administrator *entity.Administrator) error
	GetByLoginAccount(ctx context.Context, login_account string) (*entity.Administrator, error)
	GetByUserName(ctx context.Context, user_name string) (*entity.Administrator, error)
	GetPageList(ctx context.Context, in *v1.GetPageListRequest) (res []*entity.Administrator, total int64, err error)
	GetByID(ctx context.Context, userId int64) (resEntity *entity.Administrator, err error)
	GetListByLoginAccount(ctx context.Context, loginAccount string) (list []*entity.Administrator, err error)
	Update(ctx context.Context, user *entity.Administrator, isOwn bool) error
	SetUserStatus(ctx context.Context, userId int64, userStatus v1.AccountStatus) error
	UpdateUserPassWord(ctx context.Context, userId int64, passWord string) error
	DeleteUser(ctx context.Context, userId int64) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (resp *v1.CreateUserResponse, err error) {
	resp = &v1.CreateUserResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	//判断账号是否存在
	user, err := uc.repo.GetByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("CreateUser.repo.GetByLoginAccount Failed, req:%v", req)
		return resp, err
	}
	if user != nil {
		return resp, errors.New("该账号已存在！")
	}
	//判断用户名是否存在
	user, err = uc.repo.GetByUserName(ctx, req.UserName)
	if err != nil {
		l.Errorf("CreateUser.repo.GetByUserName Failed, req:%v", req)
		return resp, err
	}
	if user != nil {
		return resp, errors.New("该用户名已存在！")
	}

	id, err := isnowflake.SnowFlake.NextID()
	if err != nil {
		return resp, err
	}
	hashed, err := isecurity.HashPassword(req.PassWord)
	if err != nil {
		l.Errorf("CreateUser.isecurity.HashPassword Failed, req:%v", req)
		return resp, err
	}

	administrator := &entity.Administrator{
		ID:           id,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		HashPassword: hashed,
		Status:       int32(req.UserStatus),
		Email:        req.Email,
		UserType:     int32(req.UserType),
	}
	err = uc.repo.Create(ctx, administrator)
	if err != nil {
		l.Errorf("CreateUser.repo.Create Failed, req:%v", req)
		return resp, err
	}
	return resp, nil
}

func (uc *UserUseCase) GetPageList(ctx context.Context, req *v1.GetPageListRequest) (resp *v1.GetPageListResponse, err error) {
	resp = &v1.GetPageListResponse{UserList: make([]*v1.UserData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, total, err := uc.repo.GetPageList(ctx, req)
	if err != nil {
		l.Errorf("GetPageList.repo.GetPageList Failed, req:%v", req)
		err = innErr.ErrInternalServer
		return
	}
	resp.Total = total
	for _, re := range res {
		cur := &v1.UserData{
			UserId:       strconv.Itoa(int(re.ID)),
			UserName:     re.UserName,
			LoginAccount: re.LoginAccount,
			Email:        re.Email,
			UserStatus:   v1.AccountStatus(re.Status),
			UserType:     v1.UserType(re.UserType),
		}
		resp.UserList = append(resp.UserList, cur)
	}
	return
}

func (uc *UserUseCase) GetUserDetail(ctx context.Context, req *v1.GetUserDetailRequest) (resp *v1.GetUserDetailResponse, err error) {
	resp = &v1.GetUserDetailResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		err = errors.New("参数无效")
		return
	}
	res, err := uc.repo.GetByID(ctx, userId)
	if err != nil {
		l.Errorf("GetUserDetail.repo.GetByID Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("用户不存在")
		return
	}
	resp.User = &v1.UserData{
		UserId:       strconv.Itoa(int(res.ID)),
		UserName:     res.UserName,
		LoginAccount: res.LoginAccount,
		Email:        res.Email,
		UserStatus:   v1.AccountStatus(res.Status),
		UserType:     v1.UserType(res.UserType),
	}
	return
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (resp *v1.UpdateUserResponse, err error) {
	resp = &v1.UpdateUserResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	if req.UserId == "" {
		err = errors.New("参数无效")
		return
	}
	iUserId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		err = errors.New("参数无效")
		return
	}
	list, err := uc.repo.GetListByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("UpdateUser.repo.GetListByLoginAccount Failed, err:%v ", err)
		err = innErr.ErrInternalServer
		return
	}
	for _, administrator := range list {
		if administrator.ID != iUserId {
			err = errors.New("账户名已存在")
			return
		}
	}
	err = uc.repo.Update(ctx, &entity.Administrator{
		ID:           iUserId,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		Status:       int32(req.UserStatus),
		Email:        req.Email,
		UserType:     int32(req.UserType),
	}, false)
	if err != nil {
		l.Errorf("UpdateUser.repo.Update Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) SetUserStatus(ctx context.Context, req *v1.SetUserStatusRequest) (resp *v1.SetUserStatusResponse, err error) {
	resp = &v1.SetUserStatusResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		err = errors.New("参数无效")
		return
	}
	err = uc.repo.SetUserStatus(ctx, userId, req.UserStatus)
	if err != nil {
		l.Errorf("SetUserStatus.repo.SetUserStatus Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) ResetUserPassWord(ctx context.Context, req *v1.ResetUserPassWordRequest) (resp *v1.ResetUserPassWordResponse, err error) {
	resp = &v1.ResetUserPassWordResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		err = errors.New("参数无效")
		return
	}
	hashPassWord, err := isecurity.HashPassword(req.PassWord)
	if err != nil {
		l.Errorf("ResetUserPassWord.isecurity.HashPassword Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	err = uc.repo.UpdateUserPassWord(ctx, userId, hashPassWord)
	if err != nil {
		l.Errorf("ResetUserPassWord.repo.UpdateUserPassWord Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (resp *v1.DeleteUserResponse, err error) {
	resp = &v1.DeleteUserResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	iuserId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		err = errors.New("参数无效")
		return
	}

	err = uc.repo.DeleteUser(ctx, iuserId)
	if err != nil {
		l.Errorf("DeleteUser.repo.DeleteUser Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) GetUserSelfDetail(ctx context.Context, req *v1.GetUserSelfDetailRequest) (resp *v1.GetUserSelfDetailResponse, err error) {
	resp = &v1.GetUserSelfDetailResponse{}
	l := uc.log.WithContext(ctx)
	userId, ok := icontext.UserIdFrom(ctx)
	if !ok {
		err = innErr.ErrLogin
		return
	}
	iUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		err = innErr.ErrLogin
		return
	}
	res, err := uc.repo.GetByID(ctx, iUserId)
	if err != nil {
		l.Errorf("GetUserSelfDetail.repo.GetByID Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("用户不存在")
		return
	}
	resp.User = &v1.UserData{
		UserId:       strconv.Itoa(int(res.ID)),
		UserName:     res.UserName,
		LoginAccount: res.LoginAccount,
		Email:        res.Email,
		UserStatus:   v1.AccountStatus(res.Status),
		UserType:     v1.UserType(res.UserType),
	}
	return
}

func (uc *UserUseCase) UpdateUserSelf(ctx context.Context, req *v1.UpdateUserSelfRequest) (resp *v1.UpdateUserSelfResponse, err error) {
	resp = &v1.UpdateUserSelfResponse{}
	l := uc.log.WithContext(ctx)

	userId, ok := icontext.UserIdFrom(ctx)
	if !ok {
		err = innErr.ErrLogin
		return
	}
	iUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		err = innErr.ErrLogin
		return
	}
	list, err := uc.repo.GetListByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("UpdateUserSelf.repo.GetListByLoginAccount Failed, err:%v ", err)
		err = innErr.ErrInternalServer
		return
	}
	for _, administrator := range list {
		if administrator.ID != iUserId {
			err = errors.New("账户名已存在")
			return
		}
	}
	err = uc.repo.Update(ctx, &entity.Administrator{
		ID:           iUserId,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		Email:        req.Email,
	}, true)
	if err != nil {
		l.Errorf("UpdateUserSelf.repo.Update Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) UpdateUserPassWord(ctx context.Context, req *v1.UpdateUserPassWordRequest) (resp *v1.UpdateUserPassWordResponse, err error) {
	resp = &v1.UpdateUserPassWordResponse{}
	l := uc.log.WithContext(ctx)
	userId, _ := icontext.UserIdFrom(ctx)
	if userId == "" {
		err = innErr.ErrLogin
		return
	}
	iuserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		err = innErr.ErrLogin
		return
	}
	user, err := uc.repo.GetByID(ctx, iuserId)
	if err != nil {
		err = errors.New("用户不存在")
		return
	}
	if ok := isecurity.CheckPassword(req.OldPassWord, user.HashPassword); !ok {
		err = errors.New("旧密码不正确")
		return
	}
	hashPassWord, err := isecurity.HashPassword(req.NewPassWord)
	if err != nil {
		l.Errorf("UpdateUserPassWord.isecurity.HashPassword Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	err = uc.repo.UpdateUserPassWord(ctx, iuserId, hashPassWord)
	if err != nil {
		l.Errorf("UpdateUserPassWord.repo.UpdateUserPassWord Failed, err:%v", err)
		err = innErr.ErrInternalServer
		return
	}
	return
}

func adminPermission(ctx context.Context) (rule string, err error) {
	rule, ok := icontext.UserRuleFrom(ctx)
	if !ok || rule != fmt.Sprintf("%d", v1.UserType_Admin) {
		err = errors.New("无权限")
		return
	}
	return
}
