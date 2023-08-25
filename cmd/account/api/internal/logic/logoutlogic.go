package logic

import (
	"context"
	"cshop/cmd/auth/rpc/pb"
	"fmt"

	"cshop/cmd/account/api/internal/svc"
	"cshop/cmd/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	_, err = l.svcCtx.AuthRpcClient.DeleteToken(l.ctx, &pb.DeleteTokenReq{
		AccountName: req.AccountName,
	})

	if err != nil {
		return &types.LogoutResp{
			CommonResp: types.CommonResp{
				Success: false,
				Detail:  fmt.Sprintf("logout failed, err: %v", err),
			},
		}, fmt.Errorf("logout failed, err: %v", err)
	}

	return &types.LogoutResp{
		CommonResp: types.CommonResp{
			Success: true,
		},
	}, nil
}
