package user

import "nFlow/flow/internal/types"

// NFlowGetUserInfoParams 获取用户信息查询参数
type NFlowGetUserInfoParams struct {
	UserCode string `json:"userCode"`
}

// NFlowCheckUserAccountParams 检查用账号是否存在
type NFlowCheckUserAccountParams types.RequestCheckUserAccount
