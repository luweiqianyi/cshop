package logic

import (
	"context"

	"cshop/cmd/account/api/internal/svc"
	"cshop/cmd/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnRegisterLogic {
	return &UnRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnRegisterLogic) UnRegister(req *types.UnRegisterReq) (resp *types.UnRegisterResp, err error) {
	return &types.UnRegisterResp{
		CommonResp: types.CommonResp{
			Success: true,
		},
	}, nil
}
