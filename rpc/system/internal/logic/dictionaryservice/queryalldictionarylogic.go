package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllDictionaryLogic {
	return &QueryAllDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllDictionaryLogic) QueryAllDictionary(in *system.QueryAllDictionaryReq) (*system.QueryAllDictionaryResp, error) {
	// todo: add your logic here and delete this line

	return &system.QueryAllDictionaryResp{}, nil
}
