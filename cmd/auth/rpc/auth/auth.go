// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package auth

import (
	"context"

	"cshop/cmd/auth/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = pb.GenerateTokenReq
	GenerateTokenResp = pb.GenerateTokenResp
	TokenValidateReq  = pb.TokenValidateReq
	TokenValidateResp = pb.TokenValidateResp

	Auth interface {
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		ValidateToken(ctx context.Context, in *TokenValidateReq, opts ...grpc.CallOption) (*TokenValidateResp, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

func (m *defaultAuth) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewAuthClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultAuth) ValidateToken(ctx context.Context, in *TokenValidateReq, opts ...grpc.CallOption) (*TokenValidateResp, error) {
	client := pb.NewAuthClient(m.cli.Conn())
	return client.ValidateToken(ctx, in, opts...)
}
