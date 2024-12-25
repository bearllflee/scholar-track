package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryDetailLogic {
	return &QueryCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryCategoryDetailLogic) QueryCategoryDetail(in *achieve.QueryCategoryDetailReq) (*achieve.QueryCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.QueryCategoryResp{}, nil
}
