package category

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) (resp *types.UpdateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
