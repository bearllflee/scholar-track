package basic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAchievementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAchievementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAchievementLogic {
	return &UploadAchievementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAchievementLogic) UploadAchievement(req *types.UploadAchievementReq) (resp *types.UploadAchievementResp, err error) {
	files := make([]*achieve.File, 0)
	for _, file := range req.Files {
		files = append(files, &achieve.File{
			Name:    file.Name,
			Size:    file.Size,
			Content: file.Content,
		})
	}
	otherInfo, err := structpb.NewStruct(req.OtherInfo)
	if err != nil {
		return nil, err
	}
	uploadAchieveResp, err := l.svcCtx.Achieve.UploadAchieve(l.ctx, &achieve.UploadAchieveReq{
		Code:        req.Code,
		Name:        req.Name,
		Status:      req.Status,
		Level:       req.Level,
		Rank:        req.Rank,
		AwardTime:   req.AwardTime,
		Share:       req.Share,
		Star:        req.Star,
		Uploader:    req.Uploader,
		Team:        req.Team,
		TeamMembers: req.TeamMembers,
		Description: req.Description,
		CategoryId:  req.CategoryId,
		OtherInfo:   otherInfo,
		File:        files,
	})
	if err != nil {
		return nil, err
	}
	resp.Id = uploadAchieveResp.Id
	return
}
