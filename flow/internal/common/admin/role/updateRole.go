package role

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

// NFlowUpdateRole 修改角色接口
type NFlowUpdateRole interface {
	ValidParams() error
	UpdateData(c context.Context) error
}

// ValidParams 校验参数
func (t NFlowUpdateRoleParams) ValidParams() error {

	logx.Infof("t: %+v", t)

	if t.RoleCode == "" {
		return errors.Wrapf(xerr.NewErrMsg("角色编码非法"),
			"角色编码值为空, roleCode: %s", t.RoleCode)
	}

	if t.Status < 1 || t.Status > 2 {
		return errors.Wrapf(xerr.NewErrMsg("状态参数非法"), "状态参数 status 传值范围不对, status: %d", t.Status)
	}

	if t.RoleOrder < 1 {
		return errors.Wrapf(xerr.NewErrMsg("排序参数非法"), "排序参数 roleOrder 传值范围不对, roleOrder: %d", t.RoleOrder)
	}

	return nil
}

// UpdateData 校验参数
func (t NFlowUpdateRoleParams) UpdateData(c context.Context) error {

	logx.Infof("t: %+v", t)

	_, err := zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		userCode := c.Value("userCode").(string)
		finder := zorm.NewFinder()
		finderSql := UpdateRoleSql
		finder.Append(finderSql, t.RoleName, t.RoleDesc, t.RoleOrder, t.Status, userCode, t.RoleCode)
		_, err := zorm.UpdateFinder(ctx, finder)
		if err != nil {
			return nil, err
		}

		delPermitFinder := zorm.NewFinder()
		delPermitFinderSql := DelRolePermitSql
		delPermitFinder.Append(delPermitFinderSql, t.RoleCode)
		_, err = zorm.UpdateFinder(ctx, delPermitFinder)

		if err != nil {
			return nil, err
		}

		permitFinder := zorm.NewFinder()
		permitFinderSql := InsertRolePermitSql
		permitFinder.Append(permitFinderSql)
		for idx, item := range t.Permit {
			if idx == 0 {
				permitFinder.Append(" (?, ?, ?, ?)", item, t.RoleCode, userCode, userCode)
			} else {
				permitFinder.Append(" ,(?, ?, ?, ?)", item, t.RoleCode, userCode, userCode)
			}
		}
		_, err = zorm.UpdateFinder(ctx, permitFinder)

		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误,请重试!"),
			"修改角色错误, err: %v", err)
	}
	return nil
}
