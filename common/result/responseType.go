package result

import "nFlow/common/xerr"

type ResponseSuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

// Success 请求成功返回数据
func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{xerr.OK, xerr.MapErrMsg(xerr.OK), data}
}

type ResponseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

// Error 请求失败返回数据
func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
