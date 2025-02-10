package dictionaryservicelogic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *system.UpdateDictionaryDetailReq) (*system.UpdateDictionaryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &system.UpdateDictionaryDetailResp{}, nil
}
