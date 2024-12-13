package rolelogic

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleTreeLogic {
	return &RoleTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleTreeLogic) RoleTree(in *system.RoleTreeReq) (*system.RoleTreeListResp, error) {
	return nil, nil
}
