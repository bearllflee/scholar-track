package userlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSelfInfoLogic {
	return &SetSelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetSelfInfoLogic) SetSelfInfo(in *system.SetSelfInfoReq) (*system.SetSelfInfoResp, error) {
	modelUser := model.User{
		StModel: global.StModel{
			ID: uint(in.Id),
		},
		Username: in.Username,
		Email:    in.Email,
		Avatar:   in.Avatar,
		Role:     uint(in.Role),
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Gender:   int8(in.Gender),
		Major:    in.Major,
		College:  in.College,
		Grade:    in.Grade,
		Class:    in.Class,
		Realname: in.Realname,
	}
	err := IfUniqueHasExists(&modelUser)
	if err != nil {
		return nil, err
	}
	err = global.DB.Model(&model.User{}).Select(
		"username", "email", "avatar", "role", "nickname", "phone", "gender", "major", "college", "grade", "class", "realname",
	).Where("id = ?", modelUser.ID).Updates(&modelUser).Error
	if err != nil {
		return nil, err
	}
	return &system.SetSelfInfoResp{}, nil
}
