package achieveservicelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"time"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/storage/storage_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAchieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadAchieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAchieveLogic {
	return &UploadAchieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadAchieveLogic) UploadAchieve(in *achieve.UploadAchieveReq) (*achieve.UploadAchieveResp, error) {
	// 使用分布式事务 saga 给补偿方式
	// 1. 先上传成果，插入成果库
	// 2. 上传文件，插入文件库
	// 3. 如果出现错误，给出第一步的补偿
	var achieveBasic = &model.AchieveBasic{
		StModel:     global.StModel{},
		Code:        in.Code,
		CategoryId:  in.CategoryId,
		Level:       in.Level,
		Rank:        in.Rank,
		AwardTime:   time.UnixMilli(in.AwardTime),
		Status:      in.Status,
		Name:        in.Name,
		Share:       in.Share,
		Star:        0,
		Description: in.Description,
		Uploader:    in.Uploader,
		Team:        in.Team,
		TeamMembers: in.TeamMembers,
		OtherInfo:   in.OtherInfo.AsMap(),
	}
	uploadAchieve, err := l.svcCtx.AchieveService.UploadAchieve(l.ctx, achieveBasic)
	if err != nil {
		return nil, err
	}

	for _, file := range in.File {
		_, err = l.svcCtx.StorageClient.FileUpload(l.ctx, &storage_client.FileUploadRequest{
			FileData:      file.Content,
			FileName:      file.Name,
			FileType:      file.Type,
			BussinessId:   uploadAchieve.Id,
			BussinessName: "achieve",
		})
		if err != nil {
			err = l.svcCtx.AchieveService.DeleteAchieve(l.ctx, uploadAchieve.Id)
			return nil, err
		}
	}
	return &achieve.UploadAchieveResp{Id: uploadAchieve.Id}, nil
}
