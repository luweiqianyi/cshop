package logic

import (
	"context"
	"cshop/cmd/order/rpc/store"
	"fmt"

	"cshop/cmd/order/rpc/internal/svc"
	"cshop/cmd/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderLogic {
	return &QueryOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryOrderLogic) QueryOrder(in *pb.QueryOrderReq) (*pb.QueryOrderResp, error) {
	order, err := store.MemStoreInstance().Query(in.OrderID)
	if err != nil {
		return &pb.QueryOrderResp{}, fmt.Errorf("query order failed, err: %v", err)
	}

	return &pb.QueryOrderResp{
		OrderID:   order.OrderID,
		OrderInfo: &pb.OrderInfo{},
		OrderAdditionalInfo: &pb.OrderAdditionalInfo{
			CreateTimestamp: order.CreateTimestampMs,
		},
	}, nil
}
