package account

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/flow/internal/types"
	"strings"
)

// NFlowGetUserList 获取用户列表接口
type NFlowGetUserList interface {
	GetUserBaseInfoList(ctx context.Context) (*types.ResponseGetUserListRes, error)
}

// GetUserBaseInfoList 获取用户基本信息数据
func (t NFlowGetAccountListParams) GetUserBaseInfoList(c context.Context) (resp *types.ResponseGetUserListRes, err error) {

	logx.Infof("t: %+v", t)

	finder := zorm.NewFinder()
	countFinder := zorm.NewFinder()
	finderSql := GetUserBaseInfoListSql
	countSql := GetUserBaseInfoListCountSql

	finder.Append(finderSql)
	countFinder.Append(countSql)

	if t.Dept != "" {
		finder.Append(" and nfd.dept_code in (?)", strings.Split(t.Dept, ","))
		countFinder.Append(" and nfd.dept_code in (?)", strings.Split(t.Dept, ","))
	}

	if t.UserAccount != "" {
		finder.Append(" and nfu.user_account like ?", "%"+t.UserAccount+"%")
		countFinder.Append(" and nfu.user_account like ?", "%"+t.UserAccount+"%")
	}

	if t.UserName != "" {
		finder.Append(" and nfu.user_name like ?", "%"+t.UserName+"%")
		countFinder.Append(" and nfu.user_name like ?", "%"+t.UserName+"%")
	}

	if t.Status != 0 {
		finder.Append(" and nfu.is_del = ?", t.Status)
		countFinder.Append(" and nfu.is_del = ?", t.Status)
	}

	finder.Append(" group by nfu.user_code order by nfu.id desc")

	page := zorm.Page{
		PageNo:   int(t.Page),
		PageSize: int(t.PageSize),
	}
	var count int64
	_, _ = zorm.QueryRow(c, countFinder, &count)

	var res []types.ResponseGetUserList
	err = zorm.Query(c, finder, &res, &page)

	logx.Infof("err: %v", err)

	resp = &types.ResponseGetUserListRes{
		Items: res,
		Total: count,
	}
	logx.Infof("resp: %v", resp)

	return resp, nil
}
