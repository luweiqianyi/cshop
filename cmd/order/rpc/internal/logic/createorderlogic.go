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

const orderCloseTime = time.Second * 60

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
		OrderStatus:       entity.OrderWaitForPay,
	})
	if err != nil {
		return &pb.CreateOrderResp{}, fmt.Errorf("create order failed, err: %v", err)
	}

	// 模拟订单到期，处理订单的业务过程
	go func(ctx context.Context, orderID string) {
		ticker := time.NewTicker(orderCloseTime)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				query, err := store.MemStoreInstance().Query(orderID)
				if err != nil {
					logx.Errorf("order[%v] not exist, err:%v", orderID, err)
					return
				}

				switch query.OrderStatus {
				case entity.OrderWaitForPay: // 如果订单到期还没有支付，则更新订单状态为“已取消”
					err := store.MemStoreInstance().UpdateOrderState(orderID, entity.OrderCanceled)
					if err != nil {
						logx.Errorf("update order[%v] status failed, err:%v", orderID, err)
						return
					}
					//err = store.MemStoreInstance().Delete(orderID)
					//if err != nil {
					//	logx.Errorf("delete order[%v] failed, err:%v", orderID, err)
					//	return
					//}
				}
				return // 结束业务
			}
		}
	}(l.ctx, orderID)

	return &pb.CreateOrderResp{
		OrderID: orderID,
		OrderAdditionalInfo: &pb.OrderAdditionalInfo{
			CreateTimestamp: ts,
		},
	}, nil
}
