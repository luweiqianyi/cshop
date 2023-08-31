package logic

import (
	"context"
	"cshop/cmd/order/rpc/pb"
	"cshop/pkg/ctxdata"
	"fmt"

	"cshop/cmd/order/api/internal/svc"
	"cshop/cmd/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommitOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommitOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommitOrderLogic {
	return &CommitOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommitOrderLogic) CommitOrder(req *types.CommitOrderReq) (resp *types.CommitOrderResp, err error) {
	accountName := ctxdata.GetAccountNameFromCtx(l.ctx)

	rpcResp, err := l.svcCtx.OrderRpcService.CreateOrder(
		l.ctx,
		&pb.CreateOrderReq{
			OrderCreatorID: accountName,
			OrderInfo: &pb.OrderInfo{
				ProductName:          req.ProductName,
				ProductNumber:        req.ProductNumber,
				PayMethod:            req.PayMethod,
				DeliveryMethod:       req.DeliveryMethod,
				ExpectedDeliveryTime: req.ExpectedDeliveryTime,
			},
		},
	)

	resp = new(types.CommitOrderResp)
	if err != nil {
		resp.CommonResp.Success = false
		resp.Detail = fmt.Sprintf("commitOrder failed, err:%v", err)
		return
	}

	resp.CommonResp.Success = true
	resp.OrderID = rpcResp.OrderID
	return
}
