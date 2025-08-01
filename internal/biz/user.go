package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/pkg/icontext"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/isecurity"
	"eas_api/internal/pkg/isnowflake"
	"eas_api/internal/pkg/iutils"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, administrator *entity.Administrator) error
	GetByLoginAccount(ctx context.Context, loginAccount string) (*entity.Administrator, error)
	GetByUserName(ctx context.Context, userName string) (*entity.Administrator, error)
	GetPageList(ctx context.Context, in *v1.GetPageListRequest) (res []*entity.Administrator, total int64, err error)
	GetByID(ctx context.Context, userId string) (resEntity *entity.Administrator, err error)
	GetByIDs(ctx context.Context, userId []string) (list []*entity.Administrator, err error)
	GetListByLoginAccount(ctx context.Context, loginAccount string) (list []*entity.Administrator, err error)
	Update(ctx context.Context, user *entity.Administrator, isOwn bool) error
	SetUserStatus(ctx context.Context, userId string, userStatus v1.AccountStatus, updatedBy string) error
	UpdateUserPassWord(ctx context.Context, userId, passWord, updatedBy string) error
	DeleteUser(ctx context.Context, userId, updatedBy string) error
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
	curUserId, _ := icontext.UserIdFrom(ctx)
	//判断账号是否存在
	user, err := uc.repo.GetByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("CreateUser.repo.GetByLoginAccount Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	if user != nil {
		return resp, errors.New("该账号已存在！")
	}
	//判断用户名是否存在
	user, err = uc.repo.GetByUserName(ctx, req.UserName)
	if err != nil {
		l.Errorf("CreateUser.repo.GetByUserName Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	if user != nil {
		return resp, errors.New("该用户名已存在！")
	}

	id, err := isnowflake.SnowFlake.NextID(_const.AdministratorPrefix)
	if err != nil {
		return resp, err
	}
	hashed, err := isecurity.HashPassword(req.PassWord)
	if err != nil {
		l.Errorf("CreateUser.isecurity.HashPassword Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}

	administrator := &entity.Administrator{
		ID:           id,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		HashPassword: hashed,
		Status:       int32(req.UserStatus),
		Email:        req.Email,
		//UserType:     int32(req.UserType),
		UserType:  1,
		CreatedBy: curUserId,
		UpdatedBy: curUserId,
	}
	err = uc.repo.Create(ctx, administrator)
	if err != nil {
		l.Errorf("CreateUser.repo.Create Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *UserUseCase) GetPageList(ctx context.Context, req *v1.GetPageListRequest) (resp *v1.GetPageListResponse, err error) {
	if req.PageIndex == 0 {
		req.PageIndex = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp = &v1.GetPageListResponse{UserList: make([]*v1.UserData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, total, err := uc.repo.GetPageList(ctx, req)
	if err != nil {
		l.Errorf("GetPageList.repo.GetPageList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.Administrator, string](res, func(administrator *entity.Administrator) string {
		return administrator.ID
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, e := uc.repo.GetByIDs(ctx, updatedIds)
		if e != nil {
			l.Errorf("GetPageList.repo.GetByIDs Failed, updatedIds:%v, err:%v", updatedIds, e.Error())
			err = innErr.ErrInternalServer
			return resp, err
		}
		for _, administrator := range userList {
			userMap[administrator.ID] = administrator
		}
	}

	resp.Total = total
	for _, re := range res {
		updatedBy := ""
		if _, ok := userMap[re.UpdatedBy]; ok {
			updatedBy = userMap[re.UpdatedBy].UserName
		}
		cur := &v1.UserData{
			UserId:       re.ID,
			UserName:     re.UserName,
			LoginAccount: re.LoginAccount,
			Email:        re.Email,
			UserStatus:   v1.AccountStatus(re.Status),
			UserType:     v1.UserType(re.UserType),
			UpdatedAt:    re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:    updatedBy,
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
	res, err := uc.repo.GetByID(ctx, req.UserId)
	if err != nil {
		l.Errorf("GetUserDetail.repo.GetByID Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("用户不存在")
		return
	}
	resp.User = &v1.UserData{
		UserId:       res.ID,
		UserName:     res.UserName,
		LoginAccount: res.LoginAccount,
		Email:        res.Email,
		UserStatus:   v1.AccountStatus(res.Status),
		UserType:     v1.UserType(res.UserType),
		UpdatedAt:    res.CreatedAt.Format(time.DateTime),
	}
	return
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (resp *v1.UpdateUserResponse, err error) {
	resp = &v1.UpdateUserResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	curUserId, _ := icontext.UserIdFrom(ctx)
	if strings.Trim(req.UserId, " ") == "" {
		err = errors.New("参数无效")
		return
	}
	list, err := uc.repo.GetListByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("UpdateUser.repo.GetListByLoginAccount Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	for _, administrator := range list {
		if administrator.ID != req.UserId {
			err = errors.New("账户名已存在")
			return
		}
	}
	err = uc.repo.Update(ctx, &entity.Administrator{
		ID:           req.UserId,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		Status:       int32(req.UserStatus),
		Email:        req.Email,
		UpdatedBy:    curUserId,
	}, false)
	if err != nil {
		l.Errorf("UpdateUser.repo.Update Failed, req:%v, err:%v", req, err.Error())
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
	curUserId, _ := icontext.UserIdFrom(ctx)
	err = uc.repo.SetUserStatus(ctx, req.UserId, req.UserStatus, curUserId)
	if err != nil {
		l.Errorf("SetUserStatus.repo.SetUserStatus Failed, req:%v, err:%v", req, err.Error())
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
	curUserId, _ := icontext.UserIdFrom(ctx)
	hashPassWord, err := isecurity.HashPassword(req.PassWord)
	if err != nil {
		l.Errorf("ResetUserPassWord.isecurity.HashPassword Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	err = uc.repo.UpdateUserPassWord(ctx, req.UserId, hashPassWord, curUserId)
	if err != nil {
		l.Errorf("ResetUserPassWord.repo.UpdateUserPassWord Failed, req:%v, err:%v", req, err.Error())
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
	curUserId, _ := icontext.UserIdFrom(ctx)
	err = uc.repo.DeleteUser(ctx, req.UserId, curUserId)
	if err != nil {
		l.Errorf("DeleteUser.repo.DeleteUser Failed, req:%v, err:%v", req, err.Error())
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
	res, err := uc.repo.GetByID(ctx, userId)
	if err != nil {
		l.Errorf("GetUserSelfDetail.repo.GetByID Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("用户不存在")
		return
	}
	resp.User = &v1.UserData{
		UserId:       res.ID,
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
	list, err := uc.repo.GetListByLoginAccount(ctx, req.LoginAccount)
	if err != nil {
		l.Errorf("UpdateUserSelf.repo.GetListByLoginAccount Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	for _, administrator := range list {
		if administrator.ID != userId {
			err = errors.New("账户名已存在")
			return
		}
	}
	err = uc.repo.Update(ctx, &entity.Administrator{
		ID:           userId,
		UserName:     req.UserName,
		LoginAccount: req.LoginAccount,
		Email:        req.Email,
		UpdatedBy:    userId,
	}, true)
	if err != nil {
		l.Errorf("UpdateUserSelf.repo.Update Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) UpdateUserPassWord(ctx context.Context, req *v1.UpdateUserPassWordRequest) (resp *v1.UpdateUserPassWordResponse, err error) {
	resp = &v1.UpdateUserPassWordResponse{}
	l := uc.log.WithContext(ctx)
	userId, _ := icontext.UserIdFrom(ctx)
	if strings.Trim(userId, " ") == "" {
		err = innErr.ErrLogin
		return
	}
	user, err := uc.repo.GetByID(ctx, userId)
	if err != nil {
		l.Errorf("UpdateUserPassWord.repo.GetByID Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if user == nil {
		err = errors.New("用户不存在")
		return
	}
	if ok := isecurity.CheckPassword(req.OldPassWord, user.HashPassword); !ok {
		err = errors.New("旧密码不正确")
		return
	}
	hashPassWord, err := isecurity.HashPassword(req.NewPassWord)
	if err != nil {
		l.Errorf("UpdateUserPassWord.isecurity.HashPassword Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	err = uc.repo.UpdateUserPassWord(ctx, userId, hashPassWord, userId)
	if err != nil {
		l.Errorf("UpdateUserPassWord.repo.UpdateUserPassWord Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *UserUseCase) GetUserListByIds(ctx context.Context, userIds []string) (resp []*entity.Administrator, err error) {
	resp = make([]*entity.Administrator, 0, len(userIds))
	l := uc.log.WithContext(ctx)
	userId, _ := icontext.UserIdFrom(ctx)
	if strings.Trim(userId, " ") == "" {
		err = innErr.ErrLogin
		return
	}
	resp, err = uc.repo.GetByIDs(ctx, userIds)
	if err != nil {
		l.Errorf("GetUserListByIds.repo.GetByIDs Failed, userIds:%v, err:%v", userIds, err.Error())
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
