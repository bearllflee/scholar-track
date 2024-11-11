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
	// // 1. 查询角色树
	// total, roleTrees, err := l.svcCtx.RoleModel.FindRoleTree(l.ctx, in)
	// if err != nil {
	// 	return nil, err
	// }
	// // 2. 返回角色树
	// var roleTreeListResp rpcRole.RoleTreeListResp
	// roleTreeListResp.Total = total
	// roleTreeListResp.Page = in.Page
	// roleTreeListResp.PageSize = in.PageSize
	// roleTreeListResp.Roles = buildRoleTreeResp(roleTrees)
	// return &roleTreeListResp, nil
	return nil, nil
}

// func buildRoleTreeResp(roleTrees []*modelRole.RoleTree) []*rpcRole.RoleTreeResp {
// 	var roleTreeListResp []*rpcRole.RoleTreeResp
// 	for _, roleTree := range roleTrees {
// 		roleTreeListResp = append(roleTreeListResp, &rpcRole.RoleTreeResp{
// 			Role:     &rpcRole.RoleResp{
// 				Id:       roleTree.Role.Id,
// 				RoleName: roleTree.Role.RoleName,
// 				ParentId: roleTree.Role.ParentId,
// 				CreatedAt: roleTree.Role.CreatedAt.Unix(),
// 				UpdatedAt: roleTree.Role.UpdatedAt.Unix(),
// 			},
// 			Children: buildRoleTreeResp(roleTree.Children),
// 		})
// 	}
// 	return roleTreeListResp
// }
