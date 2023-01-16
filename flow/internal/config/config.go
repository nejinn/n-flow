package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	LogConf  logx.LogConf
	MaxBytes int64 `json:",default=1073741824"`
	Mysql    struct {
		DataSource string
	}
	InitData struct {
		InitUser      string
		SuperRoleCode string
		SuperRoleName string
		TopPermitCode string
		TopPermitName string
		TopDeptCode   string
		TopDeptName   string
		Avatar        string
	}
}
