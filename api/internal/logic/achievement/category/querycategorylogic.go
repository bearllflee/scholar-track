package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryLogic {
	return &QueryCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCategoryLogic) QueryCategory(req *types.QueryCategoryReq) (resp *types.QueryCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
