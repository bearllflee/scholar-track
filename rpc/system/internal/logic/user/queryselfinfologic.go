package userlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySelfInfoLogic {
	return &QuerySelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySelfInfoLogic) QuerySelfInfo(in *system.QuerySelfInfoReq) (*system.QuerySelfInfoResp, error) {
	var userModel model.User
	err := global.DB.Where(in.Id).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &system.QuerySelfInfoResp{
		User: &system.QueryUserDetailResp{
			Id:        int64(userModel.ID),
			CreatedAt: userModel.CreatedAt.Unix(),
			UpdatedAt: userModel.UpdatedAt.Unix(),
			Username:  userModel.Username,
			Email:     userModel.Email,
			Avatar:    userModel.Avatar,
			Role:      int64(userModel.Role),
			Status:    int32(userModel.Status),
			Nickname:  userModel.Nickname,
			Phone:     userModel.Phone,
			Gender:    int32(userModel.Gender),
			Major:     userModel.Major,
			College:   userModel.College,
			Grade:     userModel.Grade,
			Class:     userModel.Class,
			Realname:  userModel.Realname,
		},
	}, nil

}
