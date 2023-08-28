package test

import (
	"context"
	"cshop/cmd/order/rpc/orderrpcservice"
	"cshop/cmd/order/rpc/pb"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"

	"testing"
)

func TestRpcClientSendCreateOrderRequest(t *testing.T) {
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:9004", // order-rpc服务端地址
	}
	client := orderrpcservice.NewOrderRPCService(zrpc.MustNewClient(c))
	resp, err := client.CreateOrder(
		context.Background(),
		&pb.CreateOrderReq{
			OrderCreatorID: "1",
			OrderInfo: &orderrpcservice.OrderInfo{
				Info: "simulate information of order",
			},
		},
	)
	fmt.Printf("resp: %+v\nerr: %v\n", resp, err)
}

func TestRpcClientSendQueryOrderRequest(t *testing.T) {
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:9004", // order-rpc服务端地址
	}
	client := orderrpcservice.NewOrderRPCService(zrpc.MustNewClient(c))

	resp, err := client.CreateOrder(
		context.Background(),
		&pb.CreateOrderReq{
			OrderCreatorID: "1",
			OrderInfo: &orderrpcservice.OrderInfo{
				Info: "simulate information of order",
			},
		},
	)

	queryResp, err := client.QueryOrder(
		context.Background(),
		&pb.QueryOrderReq{
			OrderID: resp.OrderID,
		})
	fmt.Printf("queryResp: %+v\nerr: %v\n", queryResp, err)
}
