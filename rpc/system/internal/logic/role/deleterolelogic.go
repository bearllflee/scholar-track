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

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *system.DeleteRoleReq) (*system.DeleteRoleResp, error) {
	// 查询角色是否存在
	var c int64
	var roleModel model.Role
	err := global.DB.Where("id = ?", in.Id).First(&roleModel).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c == 0 {
		return nil, cerror.ErrRoleNotExists
	}
	err = global.DB.Where("id = ?", in.Id).Delete(&roleModel).Error
	if err != nil {
		return nil, err
	}

	return &system.DeleteRoleResp{
		Success: true,
	}, nil
}
