package casbinlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPolicyLogic) AddPolicy(in *system.AddPolicyReq) (*system.AddPolicyResp, error) {
	var casbinRules []gormadapter.CasbinRule
	for _, rule := range in.Rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rule.Rules[0],
			V1:    rule.Rules[1],
			V2:    rule.Rules[2],
		})
	}
	err := global.DB.Create(casbinRules).Error
	if err != nil {
		return nil, err
	}
	return &system.AddPolicyResp{}, nil
}
