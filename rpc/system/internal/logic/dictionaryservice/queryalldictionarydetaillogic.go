package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllDictionaryDetailLogic {
	return &QueryAllDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllDictionaryDetailLogic) QueryAllDictionaryDetail(in *system.QueryAllDictionaryDetailReq) (*system.QueryAllDictionaryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &system.QueryAllDictionaryDetailResp{}, nil
}
