package role

import (
	"context"
	"github.com/bearllflee/scholar-track/rpc/system/client/role"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (*types.AddRoleResp, error) {
	// 1. 检查角色是否存在
	// exist, err := l.svcCtx.Role.FindOne(l.ctx, req.ParentId)
	// if err != nil {
	// 	return nil, err
	// }
	// 2. 添加角色
	resp, err := l.svcCtx.Role.AddRole(l.ctx, &role.AddRoleReq{
		RoleName: req.RoleName,
		ParentId: uint64(req.ParentId),
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleResp{
		Role: &types.Role{
			RoleName: resp.Role.RoleName,
			ParentId: int64(resp.Role.ParentId),
			ID:       int64(resp.Role.Id),
		},
	}, nil
}
