package user

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

// NFlowCheckUserAccount 检查账号是否存在接口
type NFlowCheckUserAccount interface {
	Check(context.Context) error
}

// Check 检查账号是否存在
func (t NFlowCheckUserAccountParams) Check(c context.Context) error {

	logx.Infof("t: %v", t)

	finder := zorm.NewFinder()
	finderSql := CheckUserAccountSql

	var count int64

	finder.Append(finderSql, t.Account)
	has, err := zorm.QueryRow(c, finder, &count)

	if err != nil {
		return errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR), "数据库查询错误, err: %v", err)
	}

	if has == false || count == 0 {
		return nil
	}

	return errors.Wrapf(xerr.NewErrMsg("账号已存在"), "账号已存在")
}
