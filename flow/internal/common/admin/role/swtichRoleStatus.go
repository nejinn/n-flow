package role

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

// NFlowSwitchRoleStatus 修改角色状态接口
type NFlowSwitchRoleStatus interface {
	ValidParams(string) error
	SwitchStatus(context.Context, string) error
}

// ValidParams 校验参数
func (t NFlowSwitchRoleStatusParams) ValidParams(initSuperRoleCode string) (err error) {

	logx.Infof("t: %+v", t)

	if t.Status > 2 || t.Status < 1 {
		return errors.Wrapf(xerr.NewErrMsg("status参数非法"),
			"状态参数 status 传值范围不对, status: %d", t.Status)
	}

	if t.Code == initSuperRoleCode {
		return errors.Wrapf(xerr.NewErrMsg("禁止禁用超级管理员角色"),
			"code参数不能是超级管理员，超级管理员禁止停用, code: %s", t.Code)
	}
	return nil
}

// SwitchStatus 修改角色状态
func (t NFlowSwitchRoleStatusParams) SwitchStatus(c context.Context, initSuperRoleCode string) (err error) {

	logx.Infof("t: %+v", t)

	err = t.ValidParams(initSuperRoleCode)

	if err != nil {
		return err
	}

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		finder := zorm.NewFinder()
		finderSql := SwitchRoleStatusSql
		finder.Append(finderSql, t.Status, t.Code)

		_, err := zorm.UpdateFinder(ctx, finder)
		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试"), "修改部门数据失败， err： %v", err)
	}

	return nil
}
