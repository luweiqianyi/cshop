package logic

import (
	"context"
	"cshop/cmd/order/rpc/internal/svc"
	"cshop/cmd/order/rpc/pb"
	"cshop/cmd/order/rpc/store"
	"cshop/cmd/order/rpc/store/entity"
	"cshop/cmd/order/rpc/util"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	orderID := util.OrderIdGenerator().OrderID("Common")
	ts := time.Now().UnixMilli()

	err := store.MemStoreInstance().Add(orderID, entity.Order{
		OrderID:           orderID,
		OrderCreatorID:    in.OrderCreatorID,
		CreateTimestampMs: ts,
	})
	if err != nil {
		return &pb.CreateOrderResp{}, fmt.Errorf("create order failed, err: %v", err)
	}

	return &pb.CreateOrderResp{
		OrderID: orderID,
		OrderAdditionalInfo: &pb.OrderAdditionalInfo{
			CreateTimestamp: ts,
		},
	}, nil
}
