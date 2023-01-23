package keyWord

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/mapFunc"
	"nFlow/common/xerr"
	"strings"
)

// NFlowKeyWord 搜索接口
type NFlowKeyWord interface {
	ConvertParams() []string
	GetResult(ctx context.Context) (interface{}, error)
}

var FcMap = map[string]interface{}{
	"role": GetRoleList,
	"dept": GetDeptList,
}

func (t NFlowKeyWordParams) ConvertParams() []string {

	logx.Infof("t: %+v", t)

	category := strings.Split(t.Category, ",")

	logx.Infof("category: %+v", category)

	return category
}

func (t NFlowKeyWordParams) GetResult(c context.Context) (resp interface{}, err error) {

	logx.Infof("t: %+v", t)

	category := t.ConvertParams()

	var bukErr error
	result := make(map[string]interface{})

	for _, item := range category {

		res, err := mapFunc.CallFcs(FcMap[item], c, t.Word)

		if err != nil {
			bukErr = err
		}
		result[item] = res
	}

	if bukErr != nil {
		return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
			"查询数据库错误, err: %v", bukErr)
	}

	logx.Infof("resp: %+v", result)

	return result, nil
}
