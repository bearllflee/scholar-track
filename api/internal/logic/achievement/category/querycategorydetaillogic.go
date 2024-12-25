package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryDetailLogic {
	return &QueryCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCategoryDetailLogic) QueryCategoryDetail(req *types.QueryCategoryDetailReq) (resp *types.QueryCategoryDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
