// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package server

import (
	"context"

	"cshop/cmd/auth/rpc/internal/logic"
	"cshop/cmd/auth/rpc/internal/svc"
	"cshop/cmd/auth/rpc/pb"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAuthServer
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *AuthServer) ValidateToken(ctx context.Context, in *pb.TokenValidateReq) (*pb.TokenValidateResp, error) {
	l := logic.NewValidateTokenLogic(ctx, s.svcCtx)
	return l.ValidateToken(in)
}

func (s *AuthServer) DeleteToken(ctx context.Context, in *pb.DeleteTokenReq) (*pb.DeleteTokenResp, error) {
	l := logic.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}
