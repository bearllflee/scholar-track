package user

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserRoleLogic) SetUserRole(req *types.SetUserRoleReq) (resp *types.SetUserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
