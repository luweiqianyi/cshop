package logic

import (
	"context"
	"cshop/cmd/auth/rpc/pb"
	"fmt"

	"cshop/cmd/auth/api/internal/svc"
	"cshop/cmd/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req *types.AuthReq) (resp *types.AuthResp, err error) {
	resp = new(types.AuthResp)

	_, err = l.svcCtx.AuthRpcCli.ValidateToken(context.Background(), &pb.TokenValidateReq{
		Token: req.AccessToken,
	})
	if err != nil {
		resp.CommonResp.Success = false
		resp.CommonResp.Detail = fmt.Sprintf("auth failed: %v", err)
		return
	}

	resp.CommonResp.Success = true
	return
}
