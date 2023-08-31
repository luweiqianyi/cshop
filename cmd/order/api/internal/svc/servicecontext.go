package svc

import (
	"cshop/cmd/order/api/internal/config"
	"cshop/cmd/order/rpc/orderrpcservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpcService orderrpcservice.OrderRPCService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		OrderRpcService: orderrpcservice.NewOrderRPCService(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
