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

type QuestionRepo interface {
	GetList(ctx context.Context, salesPaperId string) (res []*entity.Question, err error)
	GetOptionList(ctx context.Context, questionId string) (res []*entity.QuestionOption, err error)
	GetOptionListByQuestionIds(ctx context.Context, questionIds []string) (res map[string][]*entity.QuestionOption, err error)
	GetById(ctx context.Context, questionId string) (qEntity *entity.Question, qOptionsEntities []*entity.QuestionOption, err error)
	Save(ctx context.Context, question *entity.Question, addOptions, updateOptions []*entity.QuestionOption, delOptions []string, updatedBy string) error
	Delete(ctx context.Context, questionId string, updatedBy string) error
}

type QuestionUseCase struct {
	repo              QuestionRepo
	salesPaperUseCase *SalesPaperUseCase
	userUseCase       *UserUseCase
	log               *log.Helper
}

func NewQuestionUseCase(repo QuestionRepo, salesPaperUseCase *SalesPaperUseCase, userUseCase *UserUseCase, logger log.Logger) *QuestionUseCase {
	return &QuestionUseCase{repo: repo, salesPaperUseCase: salesPaperUseCase, userUseCase: userUseCase, log: log.NewHelper(logger)}
}

func (uc *QuestionUseCase) SaveSalesPaperDimensionQuestion(ctx context.Context, req *v1.SaveSalesPaperDimensionQuestionRequest) (resp *v1.SaveSalesPaperDimensionQuestionResponse, err error) {
	resp = &v1.SaveSalesPaperDimensionQuestionResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.salesPaperUseCase.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	qOptionsEntities, err := uc.repo.GetOptionList(ctx, req.QuestionData.QuestionId)
	if err != nil {
		l.Errorf("SaveSalesPaperDimensionQuestion.repo.GetOptionList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	id := req.QuestionData.QuestionId
	if strings.Trim(id, " ") == "" {
		id, err = isnowflake.SnowFlake.NextID(_const.SalesPaperDimensionQuestionPrefix)
		if err != nil {
			l.Errorf("SaveSalesPaperDimensionQuestion.SnowFlake.NextID Failed, req:%v, err:%v", req, err.Error())
			err = innErr.ErrInternalServer
		}
	}
	qEntity := &entity.Question{
		ID:             id,
		DimensionID:    req.SalesPaperDimensionId,
		SalesPaperID:   req.SalesPaperId,
		Title:          req.QuestionData.Title,
		Remark:         req.QuestionData.Remark,
		QuestionTypeID: int32(req.QuestionData.QuestionTypeId),
		Order:          req.QuestionData.Order,
		CreatedBy:      userId,
		UpdatedBy:      userId,
	}
	addOptions, updateOptions, delOptions := iutils.DiffEntities[*entity.QuestionOption, *v1.SaveQuestionOptionData, string](
		qOptionsEntities,
		req.QuestionData.Options,
		func(qOption *entity.QuestionOption) string {
			return qOption.ID
		},
		func(data *v1.SaveQuestionOptionData) string {
			return data.QuestionOptionId
		},
		func(qOption *entity.QuestionOption, data *v1.SaveQuestionOptionData) *entity.QuestionOption {
			if qOption == nil {
				qOption = &entity.QuestionOption{}
			}
			qOption.DimensionID = req.SalesPaperDimensionId
			qOption.QuestionID = id
			qOption.Score = data.Score
			qOption.Description = data.Description
			qOption.Order = data.Order
			qOption.UpdatedBy = userId
			return qOption
		})

	if len(addOptions) > 0 {
		for _, option := range addOptions {
			id, err := isnowflake.SnowFlake.NextID(_const.SalesPaperDimensionQuestionOptionPrefix)
			if err != nil {
				return resp, err
			}
			option.ID = id
		}
	}

	err = uc.repo.Save(ctx, qEntity, addOptions, updateOptions, delOptions, userId)
	if err != nil {
		l.Errorf("SaveSalesPaperDimensionQuestion.repo.Save Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *QuestionUseCase) GetSalesPaperDimensionQuestionList(ctx context.Context, req *v1.GetSalesPaperDimensionQuestionListRequest) (resp *v1.GetSalesPaperDimensionQuestionListResponse, err error) {
	resp = &v1.GetSalesPaperDimensionQuestionListResponse{QuestionData: make([]*v1.QuestionData, 0, 10)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}

	res, err := uc.repo.GetList(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionQuestionList.repo.GetList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.Question, string](res, func(question *entity.Question) string {
		return question.UpdatedBy
	})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, err := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if err != nil {
			l.Errorf("GetSalesPaperDimensionQuestionList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
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
		cur := &v1.QuestionData{
			QuestionId:           re.ID,
			Title:                re.Title,
			Remark:               re.Remark,
			QuestionTypeId:       v1.QuestionType(re.QuestionTypeID),
			Order:                re.Order,
			SalePaperId:          re.SalesPaperID,
			SalePaperDimensionId: re.DimensionID,
			UpdatedAt:            re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:            updatedBy,
		}

		resp.QuestionData = append(resp.QuestionData, cur)
	}
	return
}

func (uc *QuestionUseCase) GetSalesPaperDimensionQuestionDetail(ctx context.Context, req *v1.GetSalesPaperDimensionQuestionDetailRequest) (resp *v1.GetSalesPaperDimensionQuestionDetailResponse, err error) {
	resp = &v1.GetSalesPaperDimensionQuestionDetailResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}

	qEntity, qOptionEntities, err := uc.repo.GetById(ctx, req.QuestionId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionQuestionDetail.repo.GetById Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	if qEntity == nil {
		err = errors.New("该题目已被删除")
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*entity.QuestionOption, string](qOptionEntities, func(questionOptions *entity.QuestionOption) string {
		return questionOptions.UpdatedBy
	})
	updatedIds = append(updatedIds, qEntity.UpdatedBy)
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, err := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if err != nil {
			l.Errorf("GetSalesPaperDimensionQuestionDetail.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
			err = innErr.ErrInternalServer
			return resp, err
		}
		for _, administrator := range userList {
			userMap[administrator.ID] = administrator
		}
	}
	updatedBy := ""
	if _, ok := userMap[qEntity.UpdatedBy]; ok {
		updatedBy = userMap[qEntity.UpdatedBy].UserName
	}
	resp.QuestionData = &v1.QuestionData{
		QuestionId:           qEntity.ID,
		Title:                qEntity.Title,
		Remark:               qEntity.Remark,
		QuestionTypeId:       v1.QuestionType(qEntity.QuestionTypeID),
		Order:                qEntity.Order,
		SalePaperId:          qEntity.SalesPaperID,
		SalePaperDimensionId: qEntity.DimensionID,
		UpdatedAt:            qEntity.UpdatedAt.Format(time.DateTime),
		UpdatedBy:            updatedBy,
		QuestionOptionsData:  make([]*v1.QuestionOptionData, 0, 5),
	}
	for _, re := range qOptionEntities {
		updatedBy := ""
		if _, ok := userMap[re.UpdatedBy]; ok {
			updatedBy = userMap[re.UpdatedBy].UserName
		}
		cur := &v1.QuestionOptionData{
			QuestionOptionId: re.ID,
			Description:      re.Description,
			Score:            re.Score,
			Order:            re.Order,
			UpdatedAt:        re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:        updatedBy,
		}
		resp.QuestionData.QuestionOptionsData = append(resp.QuestionData.QuestionOptionsData, cur)
	}
	return
}

func (uc *QuestionUseCase) GetSalesPaperDimensionQuestionPreView(ctx context.Context, req *v1.GetSalesPaperDimensionQuestionPreViewRequest) (resp *v1.GetSalesPaperDimensionQuestionPreViewResponse, err error) {
	resp = &v1.GetSalesPaperDimensionQuestionPreViewResponse{QuestionData: make([]*v1.QuestionData, 0, 10)}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}

	res, err := uc.repo.GetList(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionQuestionPreView.repo.GetList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	questionIds := make([]string, 0, len(res))
	for _, re := range res {
		questionIds = append(questionIds, re.ID)
	}
	mQuestionOptions, err := uc.repo.GetOptionListByQuestionIds(ctx, questionIds)
	if err != nil {
		l.Errorf("GetSalesPaperDimensionQuestionPreView.repo.GetOptionListByQuestionIds Failed, questionIds:%v, err:%v", questionIds, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	for _, re := range res {
		cur := &v1.QuestionData{
			QuestionId:           re.ID,
			Title:                re.Title,
			Remark:               re.Remark,
			QuestionTypeId:       v1.QuestionType(re.QuestionTypeID),
			Order:                re.Order,
			SalePaperId:          re.SalesPaperID,
			SalePaperDimensionId: re.DimensionID,
		}
		if v, ok := mQuestionOptions[cur.QuestionId]; ok {
			for _, option := range v {
				cur.QuestionOptionsData = append(cur.QuestionOptionsData, &v1.QuestionOptionData{
					QuestionOptionId: option.ID,
					Description:      option.Description,
					Score:            option.Score,
					Order:            option.Order,
					SerialNumber:     iutils.OrderToLetter(option.Order),
				})
			}
		}
		resp.QuestionData = append(resp.QuestionData, cur)
	}
	return
}

func (uc *QuestionUseCase) DeleteSalesPaperDimensionQuestion(ctx context.Context, req *v1.DeleteSalesPaperDimensionQuestionRequest) (resp *v1.DeleteSalesPaperDimensionQuestionResponse, err error) {
	resp = &v1.DeleteSalesPaperDimensionQuestionResponse{}
	l := uc.log.WithContext(ctx)
	if _, err = adminPermission(ctx); err != nil {
		return
	}
	err = uc.salesPaperUseCase.CheckSalesPaper(ctx, req.SalesPaperId, l)
	if err != nil {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	err = uc.repo.Delete(ctx, req.QuestionId, userId)
	if err != nil {
		l.Errorf("DeleteSalesPaperDimensionQuestion.repo.Delete Failed, req:%v, err:%v", req, err.Error())
		return
	}
	return resp, nil
}
