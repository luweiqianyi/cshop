package logic

import (
	"context"
	"cshop/cmd/auth/rpc/pb"
	"database/sql"
	"fmt"

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
	record, err := l.svcCtx.TbUserAccountModel.FindOneByAccountName(l.ctx, sql.NullString{
		String: req.AccountName,
		Valid:  true,
	})
	if err != nil {
		return &types.UnRegisterResp{
			CommonResp: types.CommonResp{
				Success: false,
			},
		}, fmt.Errorf("unregister failed, err: %v", err)
	}

	err = l.svcCtx.TbUserAccountModel.Delete(l.ctx, record.Id)
	if err != nil {
		return &types.UnRegisterResp{
			CommonResp: types.CommonResp{
				Success: false,
			},
		}, fmt.Errorf("unregister failed, err: %v", err)
	}

	_, err = l.svcCtx.AuthRpcClient.DeleteToken(l.ctx, &pb.DeleteTokenReq{
		AccountName: req.AccountName,
	})
	if err != nil {
		logx.Errorf("unregister: account[%v] delete token failed, err=%v", req.AccountName, err)
	}

	return &types.UnRegisterResp{
		CommonResp: types.CommonResp{
			Success: true,
		},
	}, nil
}
