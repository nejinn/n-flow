package dept

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/flow/internal/types"
)

// NFlowDeptList 获取部门列表接口
type NFlowDeptList interface {
	GetDeptList(context.Context, string, string) ([]types.ResponseGetDeptList, error)
}

// GetDeptList 获取部门列表
func (t NFlowGetDeptListParams) GetDeptList(c context.Context, topDeptCode, topDeptName string) (resp []types.ResponseGetDeptList, err error) {

	logx.Infof("t: %+v", t)

	finder := zorm.NewFinder()
	finderSql := GetDeptListSql
	finder.Append(finderSql)
	if t.Name != "" {
		finder.Append(" and dept_name like ?", "%"+t.Name+"%")
	}

	if t.Status != 0 {
		finder.Append(" and nfd.is_del = ?", t.Status)
	}

	var res []types.ResponseGetDeptList
	err = zorm.Query(c, finder, &res, nil)
	if err != nil {
		logx.Infof("err: %v", err)
		return []types.ResponseGetDeptList{}, err
	}

	if t.Top == 1 {
		resp = append(resp, types.ResponseGetDeptList{
			Id:         topDeptCode,
			DeptCode:   topDeptCode,
			DeptName:   topDeptName,
			DeptDesc:   topDeptName,
			DeptOrder:  0,
			DeptParent: "",
			Status:     0,
			CUser:      "",
			UUser:      "",
			CTime:      "",
			UTime:      "",
		})
		resp = append(resp, res...)
	} else {
		resp = res
	}

	logx.Infof("resp: %+v", resp)
	return resp, err
}
