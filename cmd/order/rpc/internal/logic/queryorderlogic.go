package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.QueryOrderResp{}, nil
}
