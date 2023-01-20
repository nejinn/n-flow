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
		InitUserCode      string // 初始化用户code
		InitUserName      string // 初始化用户name
		InitSuperRoleCode string // 初始化超级管理员角色code
		InitSuperRoleName string // 初始化超级管理员角色name
		InitTopPermitCode string // 初始化第一级权限code
		InitTopPermitName string // 初始化第一级权限name
		InitTopDeptCode   string // 初始化第一级部门code
		InitTopDeptName   string // 初始化第一级部门name
		InitAvatar        string //初始化头像url
	}
}
