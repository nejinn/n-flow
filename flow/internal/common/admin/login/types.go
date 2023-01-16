package login

import "nFlow/flow/internal/types"

// NFlowLoginParams 登录接口
type NFlowLoginParams types.RequestLogin

// ResLoginUserInfo 登录用户的基本信息
type ResLoginUserInfo struct {
	UserCode string `json:"userCode"`
	IsDel    int64  `json:"isDel"`
}
