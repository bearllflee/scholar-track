package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryReq) (resp *types.DeleteCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
