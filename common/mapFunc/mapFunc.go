package mapFunc

import (
	"github.com/pkg/errors"
	"nFlow/common/xerr"
	"reflect"
)

// CallFc 动态调用函数
func CallFc(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.Wrapf(xerr.NewErrMsg("the number of params is not adapted"),
			"参数数量不对, number: %d", len(params))
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

// CallFcs 动态调用函数
// 依次传入参数就可以调用对应参数
func CallFcs(fc interface{}, params ...interface{}) (resp interface{}, err error) {
	f := reflect.ValueOf(fc)
	if f.Kind() != reflect.Func {
		return nil, errors.Wrapf(xerr.NewErrMsg("param fc is not available, must be func"),
			"fc 参数必须是函数类型, %v", f.Kind())
	}
	if len(params) != f.Type().NumIn() {
		return nil, errors.Wrapf(xerr.NewErrMsg("the number of params is not adapted"),
			"参数数量不对, number: %d", len(params))
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	res := f.Call(in)

	if res[0].Interface() == nil {
		resp = nil
	} else {
		resp = reflect.ValueOf(res[0].Interface()).Interface()
	}

	if res[1].Interface() == nil {
		err = nil
	} else {
		err = res[1].Interface().(error)
	}

	return resp, err
}
