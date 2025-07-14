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
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type SalesPaperDimensionCommentRepo interface {
	GetList(ctx context.Context, salesPaperDimensionId string) (res []*entity.SalesPaperDimensionComment, err error)
	Save(ctx context.Context, addComments []*entity.SalesPaperDimensionComment, updateComments []*entity.SalesPaperDimensionComment, delComments []string, updatedBy string) error
}

type SalesPaperDimensionCommentUseCase struct {
	repo        SalesPaperDimensionCommentRepo
	userUseCase *UserUseCase
	log         *log.Helper
}

func NewSalesPaperDimensionCommentUseCase(repo SalesPaperDimensionCommentRepo, userUseCase *UserUseCase, logger log.Logger) *SalesPaperDimensionCommentUseCase {
	return &SalesPaperDimensionCommentUseCase{repo: repo, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *SalesPaperDimensionCommentUseCase) SaveSalesPaperDimensionComment(ctx context.Context, req *v1.SaveSalesPaperDimensionCommentRequest) (resp *v1.SaveSalesPaperDimensionCommentResponse, err error) {
	resp = &v1.SaveSalesPaperDimensionCommentResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	salesPaperDimensionComments, err := uc.repo.GetList(ctx, req.SalesPaperDimensionId)
	if err != nil {
		l.Errorf("SaveSalesPaperDimensionComment.repo.GetList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}

	addComments, updateComments, delComments := iutils.DiffEntities[*entity.SalesPaperDimensionComment, *v1.SaveSalesPaperDimensionCommentData, string](
		salesPaperDimensionComments,
		req.DimensionCommentData,
		func(comment *entity.SalesPaperDimensionComment) string {
			return comment.ID
		},
		func(data *v1.SaveSalesPaperDimensionCommentData) string {
			return data.SalesPaperDimensionCommentId
		},
		func() *entity.SalesPaperDimensionComment {
			return &entity.SalesPaperDimensionComment{}
		})

	if len(addComments) > 0 {
		for _, comment := range addComments {
			id, err := isnowflake.SnowFlake.NextID(_const.SalesPaperDimensionCommentPrefix)
			if err != nil {
				return resp, err
			}
			comment.ID = id
			comment.SalesPaperDimensionID = req.SalesPaperDimensionId
		}
	}

	err = uc.repo.Save(ctx, addComments, updateComments, delComments, userId)
	if err != nil {
		l.Errorf("SaveSalesPaperDimensionComment.repo.Save Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *SalesPaperDimensionCommentUseCase) GetSalesPaperDimensionCommentList(ctx context.Context, req *v1.GetSalesPaperDimensionCommentListRequest) (resp *v1.GetSalesPaperDimensionCommentListResponse, err error) {
	resp = &v1.GetSalesPaperDimensionCommentListResponse{DimensionCommentData: make([]*v1.SalesPaperDimensionCommentData, 0, 5)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}

	res, err := uc.repo.GetList(ctx, req.SalesPaperDimensionId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionCommentList.repo.GetSalesPaperDimensionCommentList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.SalesPaperDimensionComment, string](res, func(salesPaper *entity.SalesPaperDimensionComment) string {
		return salesPaper.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, err := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if err != nil {
			l.Errorf("GetSalesPaperDimensionCommentList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
			err = innErr.ErrInternalServer
			return resp, err
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
		cur := &v1.SalesPaperDimensionCommentData{
			SalesPaperDimensionCommentId: re.ID,
			Content:                      re.Content,
			UpScore:                      re.UpScore,
			LowScore:                     re.LowScore,
			UpdatedAt:                    re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:                    updatedBy,
		}
		resp.DimensionCommentData = append(resp.DimensionCommentData, cur)
	}
	return
}
