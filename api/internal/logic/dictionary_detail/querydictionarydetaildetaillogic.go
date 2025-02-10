package dictionary_detail

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictionaryDetailDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailDetailLogic {
	return &QueryDictionaryDetailDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictionaryDetailDetailLogic) QueryDictionaryDetailDetail(req *types.QueryDictionaryDetailDetailReq) (resp *types.QueryDictionaryDetailDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
