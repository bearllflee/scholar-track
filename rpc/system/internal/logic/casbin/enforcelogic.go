package casbinlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/initialize"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnforceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnforceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnforceLogic {
	return &EnforceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnforceLogic) Enforce(in *system.EnforceReq) (*system.EnforceResp, error) {
	e := initialize.Casbin(global.DB)
	success, err := e.Enforce(in.Role, in.Path, in.Method)
	if err != nil {
		return nil, err
	}
	return &system.EnforceResp{Result: success}, nil
}
