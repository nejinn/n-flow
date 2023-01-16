package getUserInfo

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
	"nFlow/flow/internal/types"
)

// NFlowGetUserInfo 获取用户数据信息接口
type NFlowGetUserInfo interface {
	QueryBaseInfo(c context.Context) (types.ResponseGetUserInfo, error)              // 获取用户基本信息
	GetRoles(c context.Context) ([]types.ResponseGetUserInfoRoleItem, error)         // 获取用户角色信息
	GetUserInfo(c context.Context, avatar string) (types.ResponseGetUserInfo, error) // 获取用户信息
}

// QueryBaseInfo 获取用户基本信息
func (t NFlowGetUserInfoParams) QueryBaseInfo(c context.Context) (resp types.ResponseGetUserInfo, err error) {

	logx.Infof("t: %+v", t)

	var user types.ResponseGetUserInfo
	finder := zorm.NewFinder()
	finderSql := QueryBaseInfoSql
	finder.Append(finderSql, t.UserCode)
	_, err = zorm.QueryRow(c, finder, &user)

	if err != nil {
		return resp, errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试"), "查询用户信息出错, err: %v", err)
	}

	logx.Infof("resp: %+v", user)

	return user, err
}

// GetRoles 获取用户角色信息
func (t NFlowGetUserInfoParams) GetRoles(c context.Context) (resp []types.ResponseGetUserInfoRoleItem, err error) {

	logx.Infof("t: %+v", t)

	var roles []types.ResponseGetUserInfoRoleItem
	finder := zorm.NewFinder()
	finderSql := QueryRoleSql
	finder.Append(finderSql, t.UserCode)
	err = zorm.Query(c, finder, &roles, nil)

	if err != nil {
		return resp, errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试"), "查询用户信角色权限出错, err: %v", err)
	}

	logx.Infof("resp: %+v", roles)

	return roles, err
}

// GetUserInfo 获取用户信息
func (t NFlowGetUserInfoParams) GetUserInfo(c context.Context, avatar string) (resp types.ResponseGetUserInfo, err error) {

	logx.Infof("t: %+v", t)

	res, err := t.QueryBaseInfo(c)

	if err != nil {
		return resp, err
	}

	roles, err := t.GetRoles(c)

	if err != nil {
		return resp, err
	}

	if res.Avatar == "" {
		resp = types.ResponseGetUserInfo{
			Roles:    roles,
			RealName: res.RealName,
			Avatar:   avatar,
			UserId:   res.UserId,
			Username: res.Username,
			Desc:     res.Desc,
		}
	} else {
		resp = types.ResponseGetUserInfo{
			Roles:    roles,
			RealName: res.RealName,
			Avatar:   res.Avatar,
			UserId:   res.UserId,
			Username: res.Username,
			Desc:     res.Desc,
		}
	}

	logx.Infof("resp: %+v", resp)

	return resp, err
}
