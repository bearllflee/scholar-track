package logic

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeleteLogic) FileDelete(in *storage.FileDeleteRequest) (*storage.FileDeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &storage.FileDeleteResponse{}, nil
}
