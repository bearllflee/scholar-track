package dictionary_detail

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryDetailListLogic {
	return &QueryDictionaryDetailListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictionaryDetailListLogic) QueryDictionaryDetailList(req *types.QueryDictionaryDetailListReq) (resp *types.QueryDictionaryDetailListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
