package rolelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *system.UpdateRoleReq) (*system.UpdateRoleResp, error) {
	// 查询角色是否存在
	var c int64
	var roleModel model.Role
	err := global.DB.Where("role_name = ? AND id != ?", in.RoleName, in.Id).First(&roleModel).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrRoleHasExists
	}
	roleModel = model.Role{
		RoleName: in.RoleName,
		ParentId: in.ParentId,
	}
	err = global.DB.Model(&roleModel).Where("id = ?", in.Id).Updates(roleModel).Error
	if err != nil {
		return nil, err
	}

	return &system.UpdateRoleResp{
		Role: &system.RoleResp{
			Id:       roleModel.Id,
			RoleName: roleModel.RoleName,
			ParentId: roleModel.ParentId,
		},
	}, nil
}
