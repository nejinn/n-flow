package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"nFlow/flow/internal/config"
	"nFlow/flow/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	UserLogMd rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserLogMd: middleware.NewUserLogMdMiddleware().Handle,
	}
}
