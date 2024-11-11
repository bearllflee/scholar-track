package userlogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserRoleLogic) SetUserRole(in *system.SetUserRoleReq) (*system.SetUserRoleResp, error) {
	// todo: add your logic here and delete this line

	return &system.SetUserRoleResp{}, nil
}
