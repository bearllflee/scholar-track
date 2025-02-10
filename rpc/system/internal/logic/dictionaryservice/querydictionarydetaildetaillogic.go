package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictionaryDetailDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailDetailLogic {
	return &QueryDictionaryDetailDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryDictionaryDetailDetailLogic) QueryDictionaryDetailDetail(in *system.QueryDictionaryDetailDetailReq) (*system.QueryDictionaryDetailDetailResp, error) {
	// todo: add your logic here and delete this line

	return &system.QueryDictionaryDetailDetailResp{}, nil
}
