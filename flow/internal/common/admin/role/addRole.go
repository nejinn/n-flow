package role

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/generateCode"
	"nFlow/common/xerr"
)

type NFlowAddRole interface {
	ValidParams() error
	InsertData(ctx context.Context) error
}

// ValidParams 校验参数
func (t NFlowAddRoleParams) ValidParams() (err error) {

	logx.Infof("t: %+v", t)

	if t.Status < 1 || t.Status > 2 {
		return errors.Wrapf(xerr.NewErrMsg("状态参数非法"), "状态参数 status 传值范围不对, status: %d", t.Status)
	}

	if t.RoleOrder < 1 {
		return errors.Wrapf(xerr.NewErrMsg("排序参数非法"), "排序参数 roleOrder 传值范围不对, roleOrder: %d", t.RoleOrder)
	}

	return nil

}

func (t NFlowAddRoleParams) InsertData(c context.Context) (err error) {

	logx.Infof("t: %+v", t)

	err = t.ValidParams()

	if err != nil {
		logx.Infof("err: %v", err)
		return err
	}

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		roleCode := generateCode.GenerateCode(2)
		userCode := c.Value("userCode").(string)
		finder := zorm.NewFinder()
		finderSql := InsertRoleSql
		finder.Append(finderSql, roleCode, t.RoleName, t.RoleDesc, t.RoleOrder, t.Status, userCode, userCode)
		_, err := zorm.UpdateFinder(ctx, finder)
		if err != nil {
			return nil, err
		}

		permitFinder := zorm.NewFinder()
		permitFinderSql := InsertRolePermitSql
		permitFinder.Append(permitFinderSql)
		for idx, item := range t.Permit {
			if idx == 0 {
				permitFinder.Append(" (?, ?, ?, ?)", item, roleCode, userCode, userCode)
			} else {
				permitFinder.Append(" ,(?, ?, ?, ?)", item, roleCode, userCode, userCode)
			}
		}
		_, err = zorm.UpdateFinder(ctx, permitFinder)

		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误,请重试!"),
			"添加角色错误, err: %v", err)
	}

	return nil
}
