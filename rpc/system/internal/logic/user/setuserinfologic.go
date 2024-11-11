package userlogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserInfoLogic) SetUserInfo(in *system.SetUserInfoReq) (*system.SetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &system.SetUserInfoResp{}, nil
}
