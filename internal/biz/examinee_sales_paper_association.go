package biz

import (
	"context"
	v1 "eas_api/api/eas_api/v1"
	_const "eas_api/internal/const"
	"eas_api/internal/data/entity"
	"eas_api/internal/model"
	"eas_api/internal/pkg/icontext"
	"eas_api/internal/pkg/iemail"
	innErr "eas_api/internal/pkg/ierrors"
	"eas_api/internal/pkg/isnowflake"
	"eas_api/internal/pkg/iutils"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type ExamineeSalesPaperAssociationRepo interface {
	GetProvidePageList(ctx context.Context, in *v1.GetProvidePageListRequest, createdBy string) (res []*model.ExamineeSalesPaperAssociation, total int64, err error)
	GetBySalesPaperIds(ctx context.Context, salesPaperIds []string) (list []*entity.ExamineeSalesPaperAssociation, err error)
	GetByExamineeIds(ctx context.Context, examineeIds []string) (list []*entity.ExamineeSalesPaperAssociation, err error)
	Provide(ctx context.Context, associationEntities []*entity.ExamineeSalesPaperAssociation, examineeEmailRecords []*entity.ExamineeEmailRecord) (err error)
	UpdateEmailStatus(ctx context.Context, examineeSalesPaperAssociationIds []string) error
	UpdateReport(tx *gorm.DB, reportPath, examineeSalesPaperAssociationId string) error
	Delete(ctx context.Context, examineeSalesPaperAssociationId, updatedBy string) error
}

type ExamineeSalesPaperAssociationUseCase struct {
	repo             ExamineeSalesPaperAssociationRepo
	userUseCase      *UserUseCase
	salesPaperCase   *SalesPaperUseCase
	examineeUserCase *ExamineeUseCase
	log              *log.Helper
}

func NewExamineeSalesPaperAssociationUseCase(repo ExamineeSalesPaperAssociationRepo,
	userUseCase *UserUseCase,
	salesPaperCase *SalesPaperUseCase,
	examineeUserCase *ExamineeUseCase,
	logger log.Logger) *ExamineeSalesPaperAssociationUseCase {
	return &ExamineeSalesPaperAssociationUseCase{
		repo:             repo,
		userUseCase:      userUseCase,
		salesPaperCase:   salesPaperCase,
		examineeUserCase: examineeUserCase,
		log:              log.NewHelper(logger),
	}
}

func (uc *ExamineeSalesPaperAssociationUseCase) Provide(ctx context.Context, req *v1.ProvideRequest) (resp *v1.ProvideResponse, err error) {
	resp = &v1.ProvideResponse{}
	l := uc.log.WithContext(ctx)
	if len(req.ExamineeIds) == 0 {
		return
	}
	userId, _ := icontext.UserIdFrom(ctx)
	salesPaper, err := uc.salesPaperCase.GetSalesPaperDetail(ctx, &v1.GetSalesPaperDetailRequest{SalesPaperId: req.SalesPaperId})
	if err != nil {
		l.Errorf("Provide.salesPaperCase.GetSalesPaperDetail Failed, req:%v, err:%v", req, err.Error())
		return
	}
	if salesPaper == nil || salesPaper.SalesPaper == nil || salesPaper.SalesPaper.IsEnabled != 1 {
		return resp, errors.New("该试卷不存在或未启用！")
	}
	examineeMap, err := uc.examineeUserCase.GetExamineeByIds(ctx, req.ExamineeIds)
	if err != nil {
		l.Errorf("Provide.examineeUserCase.GetExamineeByIds Failed, req:%v, err:%v", req, err.Error())
		return
	}
	//组装数据
	examineeSalesPaperAssociations := make([]*entity.ExamineeSalesPaperAssociation, 0, len(req.ExamineeIds))
	examineeEmailRecords := make([]*entity.ExamineeEmailRecord, 0, len(req.ExamineeIds))
	for _, examineeId := range req.ExamineeIds {
		examinee, ok := examineeMap[examineeId]
		if !ok {
			continue
		}
		id, e := isnowflake.SnowFlake.NextID(_const.SalesPaperDimensionQuestionPrefix)
		if e != nil {
			err = e
			return
		}
		examineeSalesPaperAssociations = append(examineeSalesPaperAssociations, &entity.ExamineeSalesPaperAssociation{
			ID:             id,
			SalesPaperID:   req.SalesPaperId,
			SalesPaperName: salesPaper.SalesPaper.SalesPaperName,
			ExamineeID:     examineeId,
			EmailStatus:    int32(v1.EmailStatus_NotSend),
			StageNumber:    int32(v1.StageNumber_NoStart),
			CreatedBy:      userId,
			UpdatedBy:      userId,
		})
		id, e = isnowflake.SnowFlake.NextID(_const.ExamineeEmailRecordPrefix)
		if e != nil {
			err = e
			return
		}
		emailContent, e := iemail.RenderEmail(iemail.EmailData{
			Name:         examinee.UserName,
			CompanyName:  "团智",
			ExamName:     salesPaper.SalesPaper.SalesPaperName,
			ExamURL:      "",
			Username:     examinee.Email,
			Password:     "邮箱用户名后六位，如不满六位前面补0",
			Duration:     fmt.Sprintf("%d", salesPaper.SalesPaper.RecommendTimeLim),
			ContactName:  "团智HR",
			ContactEmail: "hr@tuanzhi.com",
			ContactPhone: "18612894185",
			SendDate:     time.Now().Format(time.DateOnly),
		})
		if e != nil {
			err = e
			return
		}
		examineeEmailRecords = append(examineeEmailRecords, &entity.ExamineeEmailRecord{
			ID:                              id,
			SalesPaperID:                    req.SalesPaperId,
			ExamineeID:                      examineeId,
			ExamineeSalesPaperAssociationID: id,
			Title:                           "【重要】请尽快完成[团智]招聘测评 – 考试账号已生成",
			Content:                         emailContent,
			ReceiverEmail:                   examinee.Email,
			EmailStatus:                     int32(v1.EmailStatus_EmailStatusNoKnow),
			SenderEmail:                     "970259505@qq.com",
			CopyReceiverEmail:               "",
			Attachment:                      "",
			IsFalseAddress:                  false,
			CreatedBy:                       userId,
			UpdatedBy:                       userId,
		})
	}

	err = uc.repo.Provide(ctx, examineeSalesPaperAssociations, examineeEmailRecords)
	if err != nil {
		l.Errorf("Provide.repo.Provide Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	//发放成功之后将试卷置为已使用
	_, err = uc.salesPaperCase.SetSalesPaperUseStatus(ctx, req.SalesPaperId)
	if err != nil {
		l.Errorf("Provide.salesPaperCase.SetSalesPaperUseStatus Failed, req:%v, err:%v", req, err.Error())
		return resp, err
	}
	return resp, nil
}

func (uc *ExamineeSalesPaperAssociationUseCase) GetProvidePageList(ctx context.Context, req *v1.GetProvidePageListRequest) (resp *v1.GetProvidePageListResponse, err error) {
	if req.PageIndex == 0 {
		req.PageIndex = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp = &v1.GetProvidePageListResponse{ProvideData: make([]*v1.ProvideData, 0, req.PageSize)}
	l := uc.log.WithContext(ctx)
	userId, _ := icontext.UserIdFrom(ctx)
	res, total, err := uc.repo.GetProvidePageList(ctx, req, userId)
	if err != nil {
		l.Errorf("GetProvidePageList.repo.GetProvidePageList Failed, req:%v, err:%v", req, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	userMap := map[string]*entity.Administrator{}
	updatedIds := iutils.GetDistinctFields[*model.ExamineeSalesPaperAssociation, string](res, func(association *model.ExamineeSalesPaperAssociation) string {
		return association.UpdatedBy
	})
	//examineeMap := map[string]*entity.Examinee{}
	//examineeIds := iutils.GetDistinctFields[*entity.ExamineeSalesPaperAssociation, string](res, func(association *entity.ExamineeSalesPaperAssociation) string {
	//	return association.ExamineeID
	//})
	if len(updatedIds) > 0 {
		userMap = make(map[string]*entity.Administrator, len(updatedIds))
		userList, e := uc.userUseCase.GetUserListByIds(ctx, updatedIds)
		if e != nil {
			l.Errorf("GetProvidePageList.userUseCase.GetUserListByIds Failed, updatedIds:%v, err:%v", updatedIds, err.Error())
			err = innErr.ErrInternalServer
			return
		}
		for _, administrator := range userList {
			userMap[administrator.ID] = administrator
		}
	}
	//if len(examineeIds) > 0 {
	//	examineeMap = make(map[string]*entity.Examinee, len(examineeIds))
	//	examineeList, e := uc.examineeUserCase.GetExamineeByIds(ctx, examineeIds)
	//	if e != nil {
	//		l.Errorf("GetProvidePageList.examineeUserCase.GetExamineeByIds Failed, updatedIds:%v, err:%v", examineeIds, err.Error())
	//		err = innErr.ErrInternalServer
	//		return
	//	}
	//	for _, examinee := range examineeList {
	//		examineeMap[examinee.ID] = examinee
	//	}
	//}
	resp.Total = total
	for _, re := range res {
		updatedBy := ""
		if _, ok := userMap[re.UpdatedBy]; ok {
			updatedBy = userMap[re.UpdatedBy].UserName
		}
		//examineeName := ""
		//if _, ok := examineeMap[re.ExamineeID]; ok {
		//	examineeName = examineeMap[re.ExamineeID].UserName
		//}
		cur := &v1.ProvideData{
			ProvideId:      re.ID,
			ExamineeId:     re.ExamineeID,
			ExamineeName:   re.ExamineeName,
			ExamineeEmail:  re.ExamineeEmail,
			ExamineePhone:  re.ExamineePhone,
			SalesPaperId:   re.SalesPaperID,
			SalesPaperName: re.SalesPaperName,
			EmailStatus:    v1.EmailStatus(re.EmailStatus),
			StageNumber:    v1.StageNumber(re.StageNumber),
			ReportPath:     re.ReportPath,
			UpdatedAt:      re.UpdatedAt.Format(time.DateTime),
			UpdatedBy:      updatedBy,
		}
		resp.ProvideData = append(resp.ProvideData, cur)
	}
	return
}

func (uc *ExamineeSalesPaperAssociationUseCase) UpdateEmailStatus(ctx context.Context, examineeSalesPaperAssociationIds []string) (err error) {
	l := uc.log.WithContext(ctx)
	if len(examineeSalesPaperAssociationIds) == 0 {
		err = errors.New("参数无效")
		return
	}
	err = uc.repo.UpdateEmailStatus(ctx, examineeSalesPaperAssociationIds)
	if err != nil {
		l.Errorf("UpdateEmailStatus.repo.UpdateEmailStatus Failed, examineeSalesPaperAssociationIds:%v, err:%v", examineeSalesPaperAssociationIds, err.Error())
		err = innErr.ErrInternalServer
		return
	}
	return
}

func (uc *ExamineeSalesPaperAssociationUseCase) UpdateReport(tx *gorm.DB, reportPath, examineeSalesPaperAssociationId string) (err error) {
	if examineeSalesPaperAssociationId == "" || reportPath == "" {
		err = errors.New("参数无效")
		return
	}
	err = uc.repo.UpdateReport(tx, reportPath, examineeSalesPaperAssociationId)
	if err != nil {
		err = fmt.Errorf("UpdateReport.repo.UpdateReport Failed, examineeSalesPaperAssociationId:%v, reportPath:%v, err:%v", examineeSalesPaperAssociationId, reportPath, err.Error())
		return
	}
	return
}
