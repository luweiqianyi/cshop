package logic

import (
	"context"
	"cshop/pkg/cryptx"
	"database/sql"
	"fmt"

	"cshop/cmd/account/api/internal/svc"
	"cshop/cmd/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	record, err := l.svcCtx.TbUserAccountModel.FindOneByAccountName(l.ctx, sql.NullString{
		String: req.AccountName,
		Valid:  true,
	})

	if err != nil {
		return &types.LoginResp{
			CommonResp: types.CommonResp{
				Success: false,
				Detail:  fmt.Sprintf("login failed, account not exist"),
			},
		}, fmt.Errorf("login failed, account not exist")
	}

	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if password != record.Password.String {
		return &types.LoginResp{
			CommonResp: types.CommonResp{
				Success: false,
				Detail:  fmt.Sprintf("login failed, password error"),
			},
		}, fmt.Errorf("login failed, password error")
	}

	return &types.LoginResp{
		CommonResp: types.CommonResp{
			Success: true,
		},
	}, nil
}
