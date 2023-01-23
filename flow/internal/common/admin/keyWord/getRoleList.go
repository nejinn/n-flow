package keyWord

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

func GetRoleList(c context.Context, key string) (resp interface{}, err error) {

	logx.Infof("key: %s", key)

	finder := zorm.NewFinder()
	finderSql := GetRoleListSql
	finder.Append(finderSql)

	if key != "" {
		finder.Append(" and role_name like ?", "%"+key+"%")
	}

	pageObj := zorm.Page{
		PageNo:   1,
		PageSize: 10,
	}

	var res []NFlowKeyWordRes
	err = zorm.Query(c, finder, &res, &pageObj)

	if err != nil {
		return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
			"数据库查询错误, err: %v", err)
	}

	logx.Infof("resp: %+v", res)

	return res, nil
}
