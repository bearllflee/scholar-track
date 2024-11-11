package userlogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySelfInfoLogic {
	return &QuerySelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySelfInfoLogic) QuerySelfInfo(in *system.QuerySelfInfoReq) (*system.QuerySelfInfoResp, error) {
	// todo: add your logic here and delete this line

	return &system.QuerySelfInfoResp{}, nil
}
