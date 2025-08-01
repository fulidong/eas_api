package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/pkg/icontext"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/iregexp"
	"eas_api/internal/pkg/isecurity"
	"eas_api/internal/pkg/isnowflake"
	"eas_api/internal/pkg/iutils"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
	"time"
)

type ExamineeRepo interface {
	GetPageList(ctx context.Context, in *v1.GetExamineePageListRequest) (res []*entity.Examinee, total int64, err error)
	GetByEmails(ctx context.Context, email []string) (res []*entity.Examinee, err error)
	GetByID(ctx context.Context, examineeId string) (resEntity *entity.Examinee, err error)
	GetByIDs(ctx context.Context, examineeIds []string) (list []*entity.Examinee, err error)
	Save(ctx context.Context, addExaminees []*entity.Examinee, updateExaminees []*entity.Examinee, updatedBy string) error
	Update(ctx context.Context, examinee *entity.Examinee) error
	UpdateExamineePassWord(ctx context.Context, examineeId, passWord, updatedBy string) error
	Delete(ctx context.Context, examineeId, updatedBy string) error
}

type ExamineeUseCase struct {
	repo        ExamineeRepo
	userUseCase *UserUseCase
	log         *log.Helper
}

func NewExamineeUseCase(repo ExamineeRepo, userUseCase *UserUseCase, logger log.Logger) *ExamineeUseCase {
	return &ExamineeUseCase{repo: repo, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *ExamineeUseCase) SaveExaminee(ctx context.Context, req *v1.SaveExamineeRequest) (resp *v1.SaveExamineeResponse, err error) {
	resp = &v1.SaveExamineeResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	if len(req.ExamineeData) == 0 {
		err = innErr.ErrBadRequest
		return
	}
	emails := make([]string, 0, len(req.ExamineeData))
	for _, datum := range req.ExamineeData {
		if !iregexp.IsValidEmail(datum.Email) {
			err = innErr.WithMessage(innErr.ErrInternalServer, "邮箱格式不正确")
			return
		}
		if !iregexp.IsValidPhoneNumberWithCountryCode(datum.Phone) {
			err = innErr.WithMessage(innErr.ErrInternalServer, "手机号格式不正确")
			return
		}
		emails = append(emails, datum.Email)
	}
	examinees, err := uc.repo.GetByEmails(ctx, emails)
	if err != nil {
		l.Errorf("SaveExaminee.repo.GetByEmails Failed, req:%v, err:%v", req, err.Error())
		return
	}
	addExaminees, updateExaminees := iutils.FindCreateAndUpdate[*entity.Examinee, *v1.SaveExamineeData](
		examinees,
		req.ExamineeData,
		func(examinee *entity.Examinee, data *v1.SaveExamineeData) bool {
			return examinee.Email == data.Email
		},
		func(examinee *entity.Examinee, data *v1.SaveExamineeData) *entity.Examinee {
			if examinee == nil {
				examinee = &entity.Examinee{}
			}
			password, hashed := "", ""
			password = iregexp.GetEmailPrefix(data.Email)
			if len(password) >= 6 {
				password = password[(len(password) - 6):]
			} else {
				password = fmt.Sprintf("%06s", password)
			}
			hashed, err := isecurity.HashPassword(password)
			if err != nil {
				l.Errorf("SaveExaminee.isecurity.HashPassword Failed, req:%v, err:%v", req, err.Error())
				return nil
			}
			examinee.UserName = data.UserName
			examinee.Phone = data.Phone
			examinee.HashPassword = hashed
			examinee.Email = data.Email
			examinee.Status = 1
			examinee.CreatedBy = userId
			examinee.UpdatedBy = userId
			return examinee
		})

	if len(addExaminees) > 0 {
		for _, examinee := range addExaminees {
			id, err := isnowflake.SnowFlake.NextID(_const.ExamineePrefix)
			if err != nil {
				return resp, err
			}
			examinee.ID = id
		}
	}

	err = uc.repo.Save(ctx, addExaminees, updateExaminees, userId)
	if err != nil {
		l.Errorf("CreateExaminee.repo.Save Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *ExamineeUseCase) GetExamineePageList(ctx context.Context, req *v1.GetExamineePageListRequest) (resp *v1.GetExamineePageListResponse, err error) {
	if req.PageIndex == 0 {
		req.PageIndex = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp = &v1.GetExamineePageListResponse{ExamineeData: make([]*v1.ExamineeData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, total, err := uc.repo.GetPageList(ctx, req)
	if err != nil {
		l.Errorf("GetExamineePageList.repo.GetPageList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.Examinee, string](res, func(salesPaper *entity.Examinee) string {
		return salesPaper.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, e := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if e != nil {
			l.Errorf("GetExamineePageList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
			err = innErr.ErrInternalServer
			return
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
		cur := &v1.ExamineeData{
			ExamineeId: re.ID,
			UserName:   re.UserName,
			Status:     1,
			Email:      re.Email,
			Phone:      re.Phone,
			UpdatedAt:  re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:  updatedBy,
		}
		resp.ExamineeData = append(resp.ExamineeData, cur)
	}
	return
}

func (uc *ExamineeUseCase) GetExamineeDetail(ctx context.Context, req *v1.GetExamineeDetailRequest) (resp *v1.GetExamineeDetailResponse, err error) {
	resp = &v1.GetExamineeDetailResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, err := uc.repo.GetByID(ctx, req.ExamineeId)
	if err != nil {
		l.Errorf("GetExamineeDetail.repo.GetByID Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("考生不存在")
		return
	}
	resp.ExamineeData = &v1.ExamineeData{
		ExamineeId: res.ID,
		UserName:   res.UserName,
		Status:     1,
		Email:      res.Email,
		Phone:      res.Phone,
		UpdatedAt:  res.UpdatedAt.Format(time.DateTime),
	}

	return
}

func (uc *ExamineeUseCase) GetExamineeByIds(ctx context.Context, examineeIds []string) (resp map[string]*entity.Examinee, err error) {
	resp = make(map[string]*entity.Examinee, len(examineeIds))
	list := make([]*entity.Examinee, 0, len(examineeIds))
	l := uc.log.WithContext(ctx)
	list, err = uc.repo.GetByIDs(ctx, examineeIds)
	if err != nil {
		l.Errorf("GetExamineeByIds.repo.GetByIDs Failed, examineeIds:%v, err:%v", examineeIds, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	for _, examinee := range list {
		resp[examinee.ID] = examinee
	}
	return
}

func (uc *ExamineeUseCase) UpdateExaminee(ctx context.Context, req *v1.UpdateExamineeRequest) (resp *v1.UpdateExamineeResponse, err error) {
	resp = &v1.UpdateExamineeResponse{}
	l := uc.log.WithContext(ctx)
	if req.ExamineeData == nil {
		err = innErr.ErrBadRequest
		return
	}
	if strings.Trim(req.ExamineeData.ExamineeId, " ") == "" {
		err = errors.New("参数无效")
		return
	}
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	if !iregexp.IsValidPhoneNumberWithCountryCode(req.ExamineeData.Phone) {
		err = innErr.WithMessage(innErr.ErrInternalServer, "手机号格式不正确")
		return
	}
	examinee, err := uc.repo.GetByID(ctx, req.ExamineeData.ExamineeId)
	if err != nil {
		l.Errorf("UpdateExaminee.repo.GetByID Failed, req:%v, err:%v", err, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if examinee == nil {
		err = errors.New("考生不存在")
		return
	}
	if examinee.Email != req.ExamineeData.Email {
		err = errors.New("邮箱不可更改")
		return
	}
	err = uc.repo.Update(ctx, &entity.Examinee{
		ID:        req.ExamineeData.ExamineeId,
		UserName:  req.ExamineeData.UserName,
		Status:    1,
		Phone:     req.ExamineeData.Phone,
		UpdatedBy: userId,
	})
	if err != nil {
		l.Errorf("UpdateExaminee.repo.Update Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *ExamineeUseCase) DeleteExaminee(ctx context.Context, req *v1.DeleteExamineeRequest) (resp *v1.DeleteExamineeResponse, err error) {
	resp = &v1.DeleteExamineeResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.repo.Delete(ctx, req.ExamineeId, userId)
	if err != nil {
		l.Errorf("DeleteExaminee.repo.Delete Failed, req:%v, err:%v", err, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}
