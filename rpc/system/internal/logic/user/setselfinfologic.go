package userlogic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &system.SetSelfInfoResp{}, nil
}
