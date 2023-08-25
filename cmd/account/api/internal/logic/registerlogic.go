package logic

import (
	"context"
	"cshop/cmd/account/api/model"
	"database/sql"
	"fmt"

	"cshop/cmd/account/api/internal/svc"
	"cshop/cmd/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	record := &model.TbUserAccount{
		AccountName: sql.NullString{
			String: req.AccountName,
			Valid:  true,
		},
		Password: sql.NullString{
			String: req.Password,
			Valid:  true,
		},
	}

	_, err = l.svcCtx.TbUserAccountModel.Insert(l.ctx, record)
	if err != nil {
		return &types.RegisterResp{
			CommonResp: types.CommonResp{
				Success: false,
			},
		}, fmt.Errorf("register failed, err: %v", err)
	}

	return &types.RegisterResp{
		CommonResp: types.CommonResp{
			Success: true,
		},
	}, nil
}
