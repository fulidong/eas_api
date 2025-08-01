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

type SalesPaperCommentRepo interface {
	GetList(ctx context.Context, salesPaperId string) (res []*entity.SalesPaperComment, err error)
	Save(ctx context.Context, addComments []*entity.SalesPaperComment, updateComments []*entity.SalesPaperComment, delComments []string, updatedBy string) error
}

type SalesPaperCommentUseCase struct {
	repo        SalesPaperCommentRepo
	userUseCase *UserUseCase
	log         *log.Helper
}

func NewSalesPaperCommentUseCase(repo SalesPaperCommentRepo, userUseCase *UserUseCase, logger log.Logger) *SalesPaperCommentUseCase {
	return &SalesPaperCommentUseCase{repo: repo, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *SalesPaperCommentUseCase) SaveSalesPaperComment(ctx context.Context, req *v1.SaveSalesPaperCommentRequest) (resp *v1.SaveSalesPaperCommentResponse, err error) {
	resp = &v1.SaveSalesPaperCommentResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	if len(req.CommentData) == 0 {
		err = innErr.ErrBadRequest
		return
	}
	salesPaperComments, err := uc.repo.GetList(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperCommentList.repo.GetSalesPaperCommentList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}

	addComments, updateComments, delComments := iutils.DiffEntities[*entity.SalesPaperComment, *v1.SaveSalesPaperCommentData, string](
		salesPaperComments,
		req.CommentData,
		func(comment *entity.SalesPaperComment) string {
			return comment.ID
		},
		func(data *v1.SaveSalesPaperCommentData) string {
			return data.SalesPaperCommentId
		},
		func(comment *entity.SalesPaperComment, data *v1.SaveSalesPaperCommentData) *entity.SalesPaperComment {
			if comment == nil {
				comment = &entity.SalesPaperComment{}
			}
			comment.Content = data.Content
			comment.SalesPaperID = req.SalesPaperId
			comment.UpScore = data.UpScore
			comment.LowScore = data.LowScore
			comment.CreatedBy = userId
			comment.UpdatedBy = userId
			return comment
		})

	if len(addComments) > 0 {
		for _, comment := range addComments {
			id, err := isnowflake.SnowFlake.NextID(_const.SalesPaperCommentPrefix)
			if err != nil {
				return resp, err
			}
			comment.ID = id
			comment.SalesPaperID = req.SalesPaperId
		}
	}

	err = uc.repo.Save(ctx, addComments, updateComments, delComments, userId)
	if err != nil {
		l.Errorf("SaveSalesPaperComment.repo.Save Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *SalesPaperCommentUseCase) GetSalesPaperCommentList(ctx context.Context, req *v1.GetSalesPaperCommentListRequest) (resp *v1.GetSalesPaperCommentListResponse, err error) {
	resp = &v1.GetSalesPaperCommentListResponse{CommentData: make([]*v1.SalesPaperCommentData, 0, 5)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}

	res, err := uc.repo.GetList(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperCommentList.repo.GetList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.SalesPaperComment, string](res, func(salesPaper *entity.SalesPaperComment) string {
		return salesPaper.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, err := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if err != nil {
			l.Errorf("GetSalesPaperCommentList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
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
		cur := &v1.SalesPaperCommentData{
			SalesPaperCommentId: re.ID,
			Content:             re.Content,
			UpScore:             re.UpScore,
			LowScore:            re.LowScore,
			UpdatedAt:           re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:           updatedBy,
		}
		resp.CommentData = append(resp.CommentData, cur)
	}
	return
}
