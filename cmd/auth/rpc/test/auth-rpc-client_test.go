package test

import (
	"context"
	"cshop/cmd/auth/rpc/auth"
	"cshop/cmd/auth/rpc/pb"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"

	"testing"
)

func TestRpcClientSendValidateTokenRequest(t *testing.T) {
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:9000", // auth-rpc服务端地址
	}
	client := auth.NewAuth(zrpc.MustNewClient(c))
	resp, err := client.ValidateToken(context.Background(), &pb.TokenValidateReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiYWNjb3VudE5hbWVcIjpcImxlZWJhaVwifSIsImV4cCI6MTY5NDA4MDA0Nn0.uzPqBT-de3nb1OhKDWqO5XUzGYdvwgKl4qCj4SCLKJQ",
	})
	fmt.Printf("resp: %#v\nerr: %v\n", resp, err)
}

func TestRpcClientSendGenerateTokenRequest(t *testing.T) {
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:9000", // auth-rpc服务端地址
	}
	client := auth.NewAuth(zrpc.MustNewClient(c))

	resp, err := client.GenerateToken(
		context.Background(),
		&pb.GenerateTokenReq{
			AccountName: "liubei",
		})
	fmt.Printf("resp: %#v\nerr: %v\n", resp, err)
}

func TestRpcClientSendDeleteTokenRequest(t *testing.T) {
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:9000", // auth-rpc服务端地址
	}
	client := auth.NewAuth(zrpc.MustNewClient(c))

	resp, err := client.DeleteToken(
		context.Background(),
		&pb.DeleteTokenReq{
			AccountName: "liubei",
		})
	fmt.Printf("resp: %#v\nerr: %v\n", resp, err)
}
