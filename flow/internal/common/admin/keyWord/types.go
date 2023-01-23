package keyWord

import "nFlow/flow/internal/types"

// NFlowKeyWordParams 搜索关键词传入参数
type NFlowKeyWordParams types.RequestKeyWord

type NFlowKeyWordRes struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	PValue string `json:"pValue"`
}
