package categoryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/achieve"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryListLogic {
	return &QueryCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryCategoryListLogic) QueryCategoryList(in *achieve.QueryCategoryListReq) (*achieve.QueryCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &achieve.QueryCategoryListResp{}, nil
}
