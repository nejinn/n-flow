package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"nFlow/common/database"
	"nFlow/flow/internal/config"
	"nFlow/flow/internal/middleware"
)

type ServiceContext struct {
	Config       config.Config
	UserPermitMd rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	database.InitZorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		UserPermitMd: middleware.NewUserPermitMdMiddleware().Handle,
	}
}
