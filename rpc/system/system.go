package main

import (
	"flag"
	"fmt"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/initialize"
	casbinServer "github.com/bearllflee/scholar-track/rpc/system/internal/server/casbin"
	roleServer "github.com/bearllflee/scholar-track/rpc/system/internal/server/role"
	userServer "github.com/bearllflee/scholar-track/rpc/system/internal/server/user"

	"github.com/bearllflee/scholar-track/rpc/system/internal/config"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	initialize.MustNewGrom(c.DataSource)
	initialize.Casbin(global.DB)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		system.RegisterCasbinServer(grpcServer, casbinServer.NewCasbinServer(ctx))
		system.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		system.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting system server at %s...\n", c.ListenOn)
	s.Start()
}
