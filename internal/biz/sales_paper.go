package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/pkg/icontext"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/iformula"
	"eas_api/internal/pkg/isnowflake"
	"eas_api/internal/pkg/iutils"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
	"time"
)

type SalesPaperRepo interface {
	Create(ctx context.Context, salesPaper *entity.SalesPaper) error
	GetPageList(ctx context.Context, in *v1.GetSalesPaperPageListRequest) (res []*entity.SalesPaper, total int64, err error)
	GetUsablePageList(ctx context.Context, in *v1.GetUsableSalesPaperPageListRequest) (res []*entity.SalesPaper, total int64, err error)
	GetBySalesPaperName(ctx context.Context, salesPaperName string) (list []*entity.SalesPaper, err error)
	GetByID(ctx context.Context, salesPaperId string) (resEntity *entity.SalesPaper, err error)
	Update(ctx context.Context, salesPaper *entity.SalesPaper) error
	SetSalesPaperStatus(ctx context.Context, salesPaperId string, salesPaperStatus v1.SalesPaperStatus, updatedBy string) error
	DeleteSalesPaper(ctx context.Context, salesPaperId, updatedBy string) error
}

type SalesPaperUseCase struct {
	repo        SalesPaperRepo
	userUseCase *UserUseCase
	log         *log.Helper
}

func NewSalesPaperUseCase(repo SalesPaperRepo, userUseCase *UserUseCase, logger log.Logger) *SalesPaperUseCase {
	return &SalesPaperUseCase{repo: repo, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *SalesPaperUseCase) CreateSalesPaper(ctx context.Context, req *v1.CreateSalesPaperRequest) (resp *v1.CreateSalesPaperResponse, err error) {
	resp = &v1.CreateSalesPaperResponse{}
	l := uc.log.WithContext(ctx)
	// 验证是否是管理员权限
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	curUserId, _ := icontext.UserIdFrom(ctx)
	if req.SalesPaperData == nil {
		err = innErr.ErrBadRequest
		return
	}
	//判断试卷是否存在
	salesPaperList, err := uc.repo.GetBySalesPaperName(ctx, req.SalesPaperData.SalesPaperName)
	if err != nil {
		l.Errorf("CreateSalesPaper.repo.GetBySalesPaperName Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	if len(salesPaperList) > 0 {
		return resp, errors.New("该试卷已存在！")
	}

	if len(req.SalesPaperData.Expression) > 0 {
		if e := iformula.ValidateExpression(req.SalesPaperData.Expression, _const.AllowedVars); e != nil {
			err = e
			return resp, err
		}
	}

	id, err := isnowflake.SnowFlake.NextID(_const.SalesPaperPrefix)
	if err != nil {
		return resp, err
	}

	salesPaper := &entity.SalesPaper{
		ID:               id,
		Name:             req.SalesPaperData.SalesPaperName,
		RecommendTimeLim: int32(req.SalesPaperData.RecommendTimeLim),
		MaxScore:         req.SalesPaperData.MaxScore,
		MinScore:         req.SalesPaperData.MinScore,
		IsEnabled:        req.SalesPaperData.IsEnabled == 1,
		Mark:             req.SalesPaperData.Mark,
		Expression:       req.SalesPaperData.Expression,
		Rounding:         req.SalesPaperData.Rounding,
		IsSumScore:       req.SalesPaperData.IsSumScore == 1,
		CreatedBy:        curUserId,
		UpdatedBy:        curUserId,
	}
	err = uc.repo.Create(ctx, salesPaper)
	if err != nil {
		l.Errorf("CreateSalesPaper.repo.Create Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *SalesPaperUseCase) GetSalesPaperPageList(ctx context.Context, req *v1.GetSalesPaperPageListRequest) (resp *v1.GetSalesPaperPageListResponse, err error) {
	resp = &v1.GetSalesPaperPageListResponse{SalesPaperList: make([]*v1.SalesPaperData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, total, err := uc.repo.GetPageList(ctx, req)
	if err != nil {
		l.Errorf("GetSalesPaperPageList.repo.GetPageList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.SalesPaper, string](res, func(salesPaper *entity.SalesPaper) string {
		return salesPaper.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, e := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if e != nil {
			l.Errorf("GetSalesPaperPageList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
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
		cur := &v1.SalesPaperData{
			SalesPaperId:     re.ID,
			SalesPaperName:   re.Name,
			RecommendTimeLim: int64(re.RecommendTimeLim),
			MaxScore:         re.MaxScore,
			MinScore:         re.MinScore,
			IsEnabled:        iutils.ConvInt(re.IsEnabled),
			IsUsed:           iutils.ConvInt(re.IsUsed),
			Mark:             re.Mark,
			Expression:       re.Expression,
			Rounding:         re.Rounding,
			IsSumScore:       iutils.ConvInt(re.IsSumScore),
			UpdatedAt:        re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:        updatedBy,
		}
		resp.SalesPaperList = append(resp.SalesPaperList, cur)
	}
	return
}

func (uc *SalesPaperUseCase) GetUsableSalesPaperPageList(ctx context.Context, req *v1.GetUsableSalesPaperPageListRequest) (resp *v1.GetUsableSalesPaperPageListResponse, err error) {
	if req.PageIndex == 0 {
		req.PageIndex = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp = &v1.GetUsableSalesPaperPageListResponse{SalesPaperList: make([]*v1.SalesPaperData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	//if _, err = adminPermission(ctx); err != nil {
	//	return
	//}
	res, total, err := uc.repo.GetUsablePageList(ctx, req)
	if err != nil {
		l.Errorf("GetSalesPaperPageList.repo.GetPageList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	resp.Total = total
	for _, re := range res {
		cur := &v1.SalesPaperData{
			SalesPaperId:     re.ID,
			SalesPaperName:   re.Name,
			RecommendTimeLim: int64(re.RecommendTimeLim),
			MaxScore:         re.MaxScore,
			MinScore:         re.MinScore,
			IsEnabled:        iutils.ConvInt(re.IsEnabled),
			IsUsed:           iutils.ConvInt(re.IsUsed),
			Mark:             re.Mark,
			IsSumScore:       iutils.ConvInt(re.IsSumScore),
			UpdatedAt:        re.UpdatedAt.Format(time.DateTime),
		}
		resp.SalesPaperList = append(resp.SalesPaperList, cur)
	}
	return
}

func (uc *SalesPaperUseCase) GetSalesPaperDetail(ctx context.Context, req *v1.GetSalesPaperDetailRequest) (resp *v1.GetSalesPaperDetailResponse, err error) {
	resp = &v1.GetSalesPaperDetailResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, err := uc.repo.GetByID(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperDetail.repo.GetByID Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("试卷不存在")
		return
	}
	resp.SalesPaper = &v1.SalesPaperData{
		SalesPaperId:     res.ID,
		SalesPaperName:   res.Name,
		RecommendTimeLim: int64(res.RecommendTimeLim),
		MaxScore:         res.MaxScore,
		MinScore:         res.MinScore,
		IsEnabled:        iutils.ConvInt(res.IsEnabled),
		IsUsed:           iutils.ConvInt(res.IsUsed),
		Mark:             res.Mark,
		Expression:       res.Expression,
		Rounding:         res.Rounding,
		IsSumScore:       iutils.ConvInt(res.IsSumScore),
		UpdatedAt:        res.UpdatedAt.Format(time.DateTime),
	}
	return
}

func (uc *SalesPaperUseCase) UpdateSalesPaper(ctx context.Context, req *v1.UpdateSalesPaperRequest) (resp *v1.UpdateSalesPaperResponse, err error) {
	resp = &v1.UpdateSalesPaperResponse{}
	l := uc.log.WithContext(ctx)
	if strings.Trim(req.SalesPaperData.SalesPaperId, " ") == "" {
		err = errors.New("参数无效")
		return
	}
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	if req.SalesPaperData == nil {
		err = innErr.ErrBadRequest
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.CheckSalesPaper(ctx, req.SalesPaperData.SalesPaperId, l)
	if err != nil {
		return
	}
	list, err := uc.repo.GetBySalesPaperName(ctx, req.SalesPaperData.SalesPaperName)
	if err != nil {
		l.Errorf("UpdateSalesPaper.repo.GetBySalesPaperName Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	for _, salesPaper := range list {
		if salesPaper.ID != req.SalesPaperData.SalesPaperId {
			err = errors.New("试卷名已存在")
			return
		}
	}
	if len(req.SalesPaperData.Expression) > 0 {
		if err = iformula.ValidateExpression(req.SalesPaperData.Expression, _const.AllowedVars); err != nil {
			return
		}
	}
	err = uc.repo.Update(ctx, &entity.SalesPaper{
		ID:               req.SalesPaperData.SalesPaperId,
		Name:             req.SalesPaperData.SalesPaperName,
		RecommendTimeLim: int32(req.SalesPaperData.RecommendTimeLim),
		MaxScore:         req.SalesPaperData.MaxScore,
		MinScore:         req.SalesPaperData.MinScore,
		IsEnabled:        req.SalesPaperData.IsEnabled == 1,
		Mark:             req.SalesPaperData.Mark,
		Expression:       req.SalesPaperData.Expression,
		Rounding:         req.SalesPaperData.Rounding,
		IsSumScore:       req.SalesPaperData.IsSumScore == 1,
		UpdatedBy:        userId,
	})
	if err != nil {
		l.Errorf("UpdateSalesPaper.repo.Update Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *SalesPaperUseCase) SetSalesPaperStatus(ctx context.Context, req *v1.SetSalesPaperStatusRequest) (resp *v1.SetSalesPaperStatusResponse, err error) {
	resp = &v1.SetSalesPaperStatusResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	err = uc.repo.SetSalesPaperStatus(ctx, req.SalesPaperId, req.SalesPaperStatus, userId)
	if err != nil {
		l.Errorf("SetUserStatus.repo.SetUserStatus Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *SalesPaperUseCase) DeleteSalesPaper(ctx context.Context, req *v1.DeleteSalesPaperRequest) (resp *v1.DeleteSalesPaperResponse, err error) {
	resp = &v1.DeleteSalesPaperResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	err = uc.repo.DeleteSalesPaper(ctx, req.SalesPaperId, userId)
	if err != nil {
		l.Errorf("DeleteSalesPaper.repo.DeleteSalesPaper Failed, req:%v, err:%v", err, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *SalesPaperUseCase) CheckSalesPaper(ctx context.Context, iSalesPaperId string, l *log.Helper) (err error) {
	salesPaper, err := uc.repo.GetByID(ctx, iSalesPaperId)
	if err != nil {
		l.Errorf("CheckSalesPaper.repo.GetByID Failed, req:%v, err:%v", err, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if salesPaper == nil {
		err = errors.New("试卷不存在")
		return
	}
	if salesPaper.IsUsed {
		err = errors.New("试卷已被使用，不可更新")
		return
	}
	return
}
