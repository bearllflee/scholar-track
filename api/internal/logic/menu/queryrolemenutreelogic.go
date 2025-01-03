package menu

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuTreeLogic {
	return &QueryRoleMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRoleMenuTreeLogic) QueryRoleMenuTree(req *types.QueryRoleMenuTreeReq) (resp *types.QueryRoleMenuTreeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
