package svc

import (
	"cshop/cmd/auth/api/internal/config"
	"cshop/cmd/auth/rpc/auth"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AuthRpcCli auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AuthRpcCli: auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConf)),
	}
}
