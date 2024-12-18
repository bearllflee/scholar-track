package rolelogic

import (
	"context"
	"strconv"

	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRolePoliciesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRolePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRolePoliciesLogic {
	return &SetRolePoliciesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRolePoliciesLogic) SetRolePolicies(in *system.SetRolePoliciesReq) (*system.SetRolePoliciesResp, error) {
	// todo: 查看角色是否存在
	// todo: 如果api已经完成了，也要查看api是否存在
	authorityId := strconv.Itoa(int(in.RoleId))
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&gormadapter.CasbinRule{}).Where("v0 = ? AND ptype != ?", authorityId, "g").First(&gormadapter.CasbinRule{}).Error
		if err != nil {
			return err
		}
		err = l.svcCtx.CasbinService.UpdateCasbin(in.RoleId, in.Rules)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &system.SetRolePoliciesResp{}, nil
}
