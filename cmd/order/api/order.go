package main

import (
	"cshop/pkg/middleware"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cshop/cmd/order/api/internal/config"
	"cshop/cmd/order/api/internal/handler"
	"cshop/cmd/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

func main() {
	flag.Parse()

	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/order-api.yaml")
	flag.Set("f", fullPath)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 获取用户身份中间件，从http header中获取AccountName设置到ctx中
	server.Use(middleware.NewIdentityMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
