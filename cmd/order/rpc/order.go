package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cshop/cmd/order/rpc/internal/config"
	"cshop/cmd/order/rpc/internal/server"
	"cshop/cmd/order/rpc/internal/svc"
	"cshop/cmd/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order-rpc.yaml", "the config file")

func main() {
	flag.Parse()

	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/order-rpc.yaml")
	flag.Set("f", fullPath)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderRPCServiceServer(grpcServer, server.NewOrderRPCServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
