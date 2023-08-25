package svc

import (
	"cshop/cmd/account/api/internal/config"
	"cshop/cmd/account/api/model"
	"cshop/cmd/auth/rpc/auth"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	TbUserAccountModel model.TbUserAccountModel
	AuthRpcClient      auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:             c,
		TbUserAccountModel: model.NewTbUserAccountModel(conn),
		AuthRpcClient:      auth.NewAuth(zrpc.MustNewClient(c.AuthRpcClientConf)),
	}
}
