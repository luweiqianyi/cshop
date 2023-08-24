package main

import (
	"cshop/pkg/store"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"os"
	"path/filepath"
	"time"

	"cshop/cmd/auth/rpc/internal/config"
	"cshop/cmd/auth/rpc/internal/server"
	"cshop/cmd/auth/rpc/internal/svc"
	"cshop/cmd/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auth-rpc.yaml", "the config file")
var redisConfigFile = flag.String("redis", "etc/redis.yaml", "redis config file")

func main() {
	flag.Parse()

	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/auth-rpc.yaml")
	flag.Set("f", fullPath)

	fullPath = filepath.Join(dir, "./etc/redis.yaml")
	flag.Set("redis", fullPath)
	var redisConf redis.RedisConf
	conf.MustLoad(*redisConfigFile, &redisConf)
	redisConf.PingTimeout = time.Second * 60
	logx.Infof("redis config: %#v", redisConf)
	store.MustUseRedisStore(redisConf)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAuthServer(grpcServer, server.NewAuthServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
