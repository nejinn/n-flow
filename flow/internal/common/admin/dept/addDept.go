package dept

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/generateCode"
	"nFlow/common/xerr"
)

// NFlowAddDept 新增部门数据接口
type NFlowAddDept interface {
	ValidParams() error
	InsertData(context.Context, string, string) error
}

// ValidParams 校验新增部门数据的参数
func (t NFlowAddDeptParams) ValidParams() (err error) {

	logx.Infof("t: %+v", t)

	if t.Status < 1 || t.Status > 2 {
		return errors.Wrapf(xerr.NewErrMsg("状态参数非法"), "状态参数 status 传值范围不对, status: %d", t.Status)
	}

	if t.DeptOrder < 1 {
		return errors.Wrapf(xerr.NewErrMsg("排序参数非法"), "排序参数 deptOrder 传值范围不对, deptOrder: %d", t.DeptOrder)
	}
	return nil
}

// InsertData 插入新增的部门数据
func (t NFlowAddDeptParams) InsertData(c context.Context, userCode string, topDeptCode string) (err error) {

	logx.Infof("t: %+v", t)

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		deptCode := generateCode.GenerateCode(1)
		deptName := t.DeptName
		deptDesc := t.DeptDesc
		deptOrder := t.DeptOrder
		deptParent := topDeptCode
		if t.DeptParent != "" {
			deptParent = t.DeptParent
		}
		isDel := t.Status
		cUser := userCode
		uUser := userCode

		finder := zorm.NewFinder()
		finderSql := InsertDeptSql

		finder.Append(finderSql, deptCode, deptName,
			deptDesc, deptOrder, deptParent, isDel,
			cUser, uUser)
		_, err = zorm.UpdateFinder(ctx, finder)

		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试"), "插入部门数据失败， err： %v", err)
	}
	return nil
}
