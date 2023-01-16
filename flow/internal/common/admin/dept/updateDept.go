package dept

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

// NFlowUpdateDept 修改部门数据接口
type NFlowUpdateDept interface {
	ValidParams() error
	UpdateData(context.Context, string, string) error
}

// ValidParams 校验修改部门数据传入参数
func (t NFlowUpdateDeptParams) ValidParams() error {

	logx.Infof("t: %+v", t)

	if t.DeptCode == "" {
		return errors.Wrapf(xerr.NewErrMsg("部门编码参数非法"), "部门编码参数不能为空, deptCode: %s", t.DeptCode)
	}

	if t.Status < 1 || t.Status > 2 {
		return errors.Wrapf(xerr.NewErrMsg("状态参数非法"), "状态参数 status 传值范围不对, status: %d", t.Status)
	}

	if t.DeptOrder < 1 {
		return errors.Wrapf(xerr.NewErrMsg("排序参数非法"), "排序参数 deptOrder 传值范围不对, deptOrder: %d", t.DeptOrder)
	}
	return nil
}

// UpdateData 修改部门数据
func (t NFlowUpdateDeptParams) UpdateData(c context.Context, userCode string, topDeptCode string) (err error) {

	logx.Infof("t: %+v", t)

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		deptCode := t.DeptCode
		deptName := t.DeptName
		deptDesc := t.DeptDesc
		deptOrder := t.DeptOrder
		deptParent := topDeptCode
		if t.DeptParent != "" {
			deptParent = t.DeptParent
		}
		isDel := t.Status
		uUser := userCode

		finder := zorm.NewFinder()
		finderSql := UpdateDeptSql

		finder.Append(finderSql, deptName,
			deptDesc, deptOrder, deptParent, isDel,
			uUser, deptCode)
		_, err = zorm.UpdateFinder(ctx, finder)

		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试"), "修改部门数据失败， err： %v", err)
	}
	return nil
}
