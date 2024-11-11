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

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *system.AddRoleReq) (*system.AddRoleResp, error) {
	var c int64
	err := global.DB.Where("role_name = ?", in.RoleName).First(&model.Role{}).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrRoleHasExists
	}
	roleModel := &model.Role{
		RoleName: in.RoleName,
		ParentID: uint(in.ParentId),
	}
	err = global.DB.Create(roleModel).Error
	if err != nil {
		return nil, err
	}

	return &system.AddRoleResp{
		Role: &system.RoleResp{
			Id:       uint64(roleModel.ID),
			RoleName: roleModel.RoleName,
			ParentId: uint64(roleModel.ParentID),
		},
	}, nil
}
