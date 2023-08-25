package svc

import (
	"cshop/cmd/account/api/internal/config"
	"cshop/cmd/account/api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	TbUserAccountModel model.TbUserAccountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:             c,
		TbUserAccountModel: model.NewTbUserAccountModel(conn),
	}
}
