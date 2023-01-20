package role

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/flow/internal/types"
)

// NFlowGetRoleList 获取角色列表接口
type NFlowGetRoleList interface {
	GetRoleList(ctx context.Context) *types.ResponseGetRoleListRes
}

// GetRoleList 获取角色列表
func (t NFlowGetRoleListParams) GetRoleList(c context.Context) (resp *types.ResponseGetRoleListRes) {

	logx.Infof("t: %+v", t)

	finder := zorm.NewFinder()
	finderSql := GetRoleListSql

	countFinder := zorm.NewFinder()
	countFinderSql := GetRoleListCountSql

	finder.Append(finderSql)
	countFinder.Append(countFinderSql)

	if t.Status != 0 {
		finder.Append(" and nfr.is_del = ?", t.Status)
		countFinder.Append(" and nfr.is_del = ?", t.Status)
	}

	if t.RoleName != "" {
		finder.Append(" and nfr.role_name like ?", "%"+t.RoleName+"%")
		countFinder.Append(" and nfr.role_name like ?", "%"+t.RoleName+"%")
	}

	finder.Append(" order by nfr.id desc")

	page := zorm.Page{
		PageNo:   int(t.Page),
		PageSize: int(t.PageSize),
	}
	var count int64
	_, _ = zorm.QueryRow(c, countFinder, &count)

	var res []types.ResponseGetRoleList
	err := zorm.Query(c, finder, &res, &page)

	logx.Infof("err: %v", err)

	resp = &types.ResponseGetRoleListRes{
		Items: res,
		Total: count,
	}
	logx.Infof("resp: %v", resp)

	return resp

}
