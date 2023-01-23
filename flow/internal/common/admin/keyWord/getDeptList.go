package keyWord

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

func GetDeptList(c context.Context, key string) (resp interface{}, err error) {

	finder := zorm.NewFinder()
	finderSql := GetDeptListSql
	finder.Append(finderSql)

	var res []NFlowKeyWordRes
	err = zorm.Query(c, finder, &res, nil)

	if err != nil {
		return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
			"数据库查询错误, err: %v", err)
	}

	logx.Infof("resp: %+v", res)

	return res, nil
}
