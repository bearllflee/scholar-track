package user

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySelfInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySelfInfoLogic {
	return &QuerySelfInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySelfInfoLogic) QuerySelfInfo(req *types.QuerySelfInfoReq) (resp *types.QuerySelfInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
