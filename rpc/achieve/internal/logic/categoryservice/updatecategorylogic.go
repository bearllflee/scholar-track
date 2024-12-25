package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *achieve.UpdateCategoryReq) (*achieve.UpdateCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.UpdateCategoryResp{}, nil
}
