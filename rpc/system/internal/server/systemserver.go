// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: system.proto

package server

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/system/internal/logic"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"
)

type SystemServer struct {
	svcCtx *svc.ServiceContext
	system.UnimplementedSystemServer
}

func NewSystemServer(svcCtx *svc.ServiceContext) *SystemServer {
	return &SystemServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemServer) Ping(ctx context.Context, in *system.Request) (*system.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}