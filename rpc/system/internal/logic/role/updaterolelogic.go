package rolelogic

import (
	"context"
	"time"

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
	var roleModel model.Role
	var c int64
	// 先查看角色是否存在
	global.DB.Model(&roleModel).Where("id = ?", in.Id).Count(&c)
	if c == 0 {
		return nil, cerror.ErrRoleNotExists
	}
	// 看看修改后的角色名称是否存在
	global.DB.Model(&roleModel).Where("role_name = ? AND id != ?", in.RoleName, in.Id).Count(&c)
	if c > 0 {
		return nil, cerror.ErrRoleHasExists
	}
	// 看看父角色是否存在
	if in.ParentId != 0 {
		global.DB.Model(&roleModel).Where("id = ?", in.ParentId).Count(&c)
		if c == 0 {
			return nil, cerror.ErrParentRoleNotExists
		}
	}
	roleModel.RoleName = in.RoleName
	roleModel.ParentId = in.ParentId
	roleModel.Id = in.Id
	roleModel.UpdatedAt = time.Now()
	roleModel.CreatedAt = time.Now()
	// 更新角色信息
	err := global.DB.Save(&roleModel).Error
	if err != nil {
		return nil, err
	}

	return &system.UpdateRoleResp{
		Id:       roleModel.Id,
		RoleName: roleModel.RoleName,
		ParentId: roleModel.ParentId,
	}, nil
}
