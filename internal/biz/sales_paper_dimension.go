package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/pkg/icontext"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/isnowflake"
	"eas_api/internal/pkg/iutils"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
	"time"
)

type SalesPaperDimensionRepo interface {
	GetList(ctx context.Context, salesPaperId string) (res []*entity.SalesPaperDimension, err error)
	Create(ctx context.Context, dimensions []*entity.SalesPaperDimension) error
	GetById(ctx context.Context, dimensionId string) (resEntity *entity.SalesPaperDimension, err error)
	Update(ctx context.Context, updateSalesPaperDimensions []*entity.SalesPaperDimension) error
	Delete(ctx context.Context, salesPaperDimensionId, updatedBy string) error
}

type SalesPaperDimensionUseCase struct {
	repo              SalesPaperDimensionRepo
	salesPaperUseCase *SalesPaperUseCase
	userUseCase       *UserUseCase
	log               *log.Helper
}

func NewSalesPaperDimensionUseCase(repo SalesPaperDimensionRepo, salesPaperUseCase *SalesPaperUseCase, userUseCase *UserUseCase, logger log.Logger) *SalesPaperDimensionUseCase {
	return &SalesPaperDimensionUseCase{repo: repo, salesPaperUseCase: salesPaperUseCase, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *SalesPaperDimensionUseCase) CreateSalesPaperDimension(ctx context.Context, req *v1.CreateSalesPaperDimensionRequest) (resp *v1.CreateSalesPaperDimensionResponse, err error) {
	resp = &v1.CreateSalesPaperDimensionResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	curUserId, _ := icontext.UserIdFrom(ctx)
	if len(req.DimensionData) == 0 {
		err = innErr.ErrBadRequest
		return
	}

	err = uc.salesPaperUseCase.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	entities := make([]*entity.SalesPaperDimension, 0, len(req.DimensionData))
	for _, datum := range req.DimensionData {
		if strings.Trim(datum.DimensionName, " ") == "" {
			err = errors.New("维度名称不能为空")
			return
		}
		id, e := isnowflake.SnowFlake.NextID(_const.SalesPaperDimensionPrefix)
		if e != nil {
			err = e
			l.Errorf("CreateSalesPaperDimension.isnowflake.SnowFlake.NextID Failed, req:%v, err:%v", req, err.Error())
			return
		}
		salesPaperDimension := &entity.SalesPaperDimension{
			ID:           id,
			SalesPaperID: req.SalesPaperId,
			Name:         datum.DimensionName,
			AverageMark:  datum.AverageMark,
			StandardMark: datum.StandardMark,
			Description:  datum.Description,
			MaxScore:     datum.MaxScore,
			MinScore:     datum.MinScore,
			IsChoose:     datum.IsChoose == 1,
			CreatedBy:    curUserId,
			UpdatedBy:    curUserId,
		}
		entities = append(entities, salesPaperDimension)
	}

	err = uc.repo.Create(ctx, entities)
	if err != nil {
		l.Errorf("CreateSalesPaperDimension.repo.Create Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *SalesPaperDimensionUseCase) GetSalesPaperDimensionList(ctx context.Context, req *v1.GetSalesPaperDimensionListRequest) (resp *v1.GetSalesPaperDimensionListResponse, err error) {
	resp = &v1.GetSalesPaperDimensionListResponse{DimensionData: make([]*v1.SalesPaperDimensionData, 0, 10)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, err := uc.repo.GetList(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionList.repo.GetList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.SalesPaperDimension, string](res, func(salesPaper *entity.SalesPaperDimension) string {
		return salesPaper.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, e := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if e != nil {
			l.Errorf("GetSalesPaperDimensionList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
			err = innErr.ErrInternalServer
			return
		}
		for _, administrator := range userList {
			userMap[administrator.ID] = administrator
		}
	}
	for _, re := range res {
		updatedBy := ""
		if _, ok := userMap[re.UpdatedBy]; ok {
			updatedBy = userMap[re.UpdatedBy].UserName
		}
		cur := &v1.SalesPaperDimensionData{
			DimensionId:   re.ID,
			DimensionName: re.Name,
			AverageMark:   re.AverageMark,
			StandardMark:  re.StandardMark,
			Description:   re.Description,
			MaxScore:      re.MaxScore,
			MinScore:      re.MinScore,
			IsChoose:      iutils.ConvInt(re.IsChoose),
			UpdatedAt:     re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:     updatedBy,
		}
		resp.DimensionData = append(resp.DimensionData, cur)
	}
	return
}

func (uc *SalesPaperDimensionUseCase) GetSalesPaperDimensionDetail(ctx context.Context, req *v1.GetSalesPaperDimensionDetailRequest) (resp *v1.GetSalesPaperDimensionDetailResponse, err error) {
	resp = &v1.GetSalesPaperDimensionDetailResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	res, err := uc.repo.GetById(ctx, req.DimensionId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionDetail.repo.GetById Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if res == nil {
		err = errors.New("试卷维度不存在")
		return
	}
	resp.SalesPaperDimension = &v1.SalesPaperDimensionData{
		DimensionId:   res.ID,
		DimensionName: res.Name,
		AverageMark:   res.AverageMark,
		StandardMark:  res.StandardMark,
		Description:   res.Description,
		MaxScore:      res.MaxScore,
		MinScore:      res.MinScore,
		IsChoose:      iutils.ConvInt(res.IsChoose),
		UpdatedAt:     res.UpdatedAt.Format(time.DateTime),
	}
	return
}

func (uc *SalesPaperDimensionUseCase) UpdateSalesPaperDimension(ctx context.Context, req *v1.UpdateSalesPaperDimensionRequest) (resp *v1.UpdateSalesPaperDimensionResponse, err error) {
	resp = &v1.UpdateSalesPaperDimensionResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	if len(req.DimensionData) == 0 {
		err = innErr.ErrBadRequest
		return
	}
	err = uc.salesPaperUseCase.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	dimensions := make([]*entity.SalesPaperDimension, 0, len(req.DimensionData))
	for _, datum := range req.DimensionData {
		dimensions = append(dimensions, &entity.SalesPaperDimension{
			ID:           datum.DimensionId,
			SalesPaperID: req.SalesPaperId,
			Name:         datum.DimensionName,
			AverageMark:  datum.AverageMark,
			StandardMark: datum.StandardMark,
			Description:  datum.Description,
			MaxScore:     datum.MaxScore,
			MinScore:     datum.MinScore,
			IsChoose:     datum.IsChoose == 1,
			UpdatedBy:    userId,
		})
	}
	err = uc.repo.Update(ctx, dimensions)
	if err != nil {
		l.Errorf("UpdateSalesPaper.repo.Update Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *SalesPaperDimensionUseCase) DeleteSalesPaperDimension(ctx context.Context, req *v1.DeleteSalesPaperDimensionRequest) (resp *v1.DeleteSalesPaperDimensionResponse, err error) {
	resp = &v1.DeleteSalesPaperDimensionResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.salesPaperUseCase.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	err = uc.repo.Delete(ctx, req.SalesPaperDimensionId, userId)
	if err != nil {
		l.Errorf("DeleteSalesPaperDimension.repo.Delete Failed, req:%v, err:%v", err, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}
