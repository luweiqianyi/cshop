// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"cshop/cmd/order/rpc/internal/logic"
	"cshop/cmd/order/rpc/internal/svc"
	"cshop/cmd/order/rpc/pb"
)

type OrderRPCServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderRPCServiceServer
}

func NewOrderRPCServiceServer(svcCtx *svc.ServiceContext) *OrderRPCServiceServer {
	return &OrderRPCServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderRPCServiceServer) CreateOrder(ctx context.Context, in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderRPCServiceServer) QueryOrder(ctx context.Context, in *pb.QueryOrderReq) (*pb.QueryOrderResp, error) {
	l := logic.NewQueryOrderLogic(ctx, s.svcCtx)
	return l.QueryOrder(in)
}
