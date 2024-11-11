package role

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleTreeLogic {
	return &RoleTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleTreeLogic) RoleTree(req *types.RoleTreeReq) (resp *types.RoleTreeResp, err error) {
	// 获取角色树
	// roleTreeListResp, err := l.svcCtx.Role.RoleTree(l.ctx, &rpcRole.RoleTreeReq{
	// 	RoleName: req.RoleName,
	// 	Page:     uint64(req.Page),
	// 	PageSize: uint64(req.PageSize),
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// resp = &types.RoleTreeResp{}
	// resp.Roles = buildRoleTree(roleTreeListResp.Roles)
	// resp.Total = int64(roleTreeListResp.Total)
	// resp.Page = int32(roleTreeListResp.Page)
	// resp.PageSize = int32(roleTreeListResp.PageSize)
	return
}

// func buildRoleTree(roles []*rpcRole.RoleTreeResp) []*types.RoleTree {
// 	resRoles := make([]*types.RoleTree, 0)
// 	for _, role := range roles {
// 		resRoles = append(resRoles, &types.RoleTree{
// 			Role: &types.Role{
// 				ID:        int64(role.Role.Id),
// 				RoleName:  role.Role.RoleName,
// 				ParentId:  int64(role.Role.ParentId),
// 				CreatedAt: role.Role.CreatedAt,
// 				UpdatedAt: role.Role.UpdatedAt,
// 			},
// 			Children: buildRoleTree(role.Children),
// 		})
// 	}
// 	return resRoles
// }
