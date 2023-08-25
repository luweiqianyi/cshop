package logic

import (
	"context"
	"cshop/cmd/auth/rpc/store"
	"fmt"

	"cshop/cmd/auth/rpc/internal/svc"
	"cshop/cmd/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *pb.DeleteTokenReq) (*pb.DeleteTokenResp, error) {
	err := store.DeleteTokenByAccountName(in.AccountName)
	if err != nil {
		return &pb.DeleteTokenResp{
			Success: false,
		}, fmt.Errorf("delete token failed, err: %v", err)
	}

	return &pb.DeleteTokenResp{
		Success: true,
	}, nil
}
