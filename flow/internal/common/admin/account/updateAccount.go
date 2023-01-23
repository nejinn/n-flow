package account

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
	"regexp"
)

type NFlowUpdateAccount interface {
	ValidParams(string) error
	UpdateData(context.Context, string) error
}

func (t NFlowUpdateAccountParams) ValidParams(initUserCode string) error {

	logx.Infof("t: %+v", t)

	if t.UserCode == initUserCode {
		return errors.Wrapf(xerr.NewErrMsg("初始超级管理员不允许修改"),
			"初始超级管理员不允许修改 userCode: %s", t.UserCode)
	}

	if len(t.UserRoles) == 0 {
		return errors.Wrapf(xerr.NewErrMsg("角色参数非法"),
			"角色参数 userRoles 不对, userRoles: %v", t.UserRoles)
	}

	if t.Status < 1 || t.Status > 2 {
		return errors.Wrapf(xerr.NewErrMsg("状态参数非法"),
			"状态参数 status 传值范围不对, status: %d", t.Status)
	}

	emailRegexp := regexp.MustCompile("^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$")

	if !emailRegexp.MatchString(t.UserMail) {
		return errors.Wrapf(xerr.NewErrMsg("邮箱参数非法"),
			"邮箱格式不对: %s", t.UserMail)
	}

	return nil
}

func (t NFlowUpdateAccountParams) UpdateData(c context.Context, initUserCode string) error {

	err := t.ValidParams(initUserCode)

	if err != nil {
		return err
	}

	logx.Infof("t: %+v", t)

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		userCode := t.UserCode
		setUser := c.Value("userCode").(string)

		finder := zorm.NewFinder()
		finderSql := UpdateAccountSql
		finder.Append(finderSql, t.UserAccount, t.UserName, t.UserMail, t.Status, userCode)
		_, err := zorm.UpdateFinder(ctx, finder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库修改用户数据错误: %v", err)
		}

		delUserRoleFinder := zorm.NewFinder()
		delUserRoleFinderSql := DelUserRoleSql
		delUserRoleFinder.Append(delUserRoleFinderSql, userCode)
		_, err = zorm.UpdateFinder(ctx, delUserRoleFinder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库删除用户角色数据错误: %v", err)
		}

		insertUserRoleFinder := zorm.NewFinder()
		insertUserRoleFinderSql := InsertUserRoleSql
		insertUserRoleFinder.Append(insertUserRoleFinderSql)

		for idx, item := range t.UserRoles {
			if idx == 0 {
				insertUserRoleFinder.Append(" (?, ?, ?, ?)", userCode, item, setUser, setUser)
			} else {
				insertUserRoleFinder.Append(" ,(?, ?, ?, ?)", userCode, item, setUser, setUser)
			}
		}

		_, err = zorm.UpdateFinder(ctx, insertUserRoleFinder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库插入用户角色数据错误: %v", err)
		}

		delUserDeptFinder := zorm.NewFinder()
		delUserDeptFinderSql := DelUserDeptSql
		delUserDeptFinder.Append(delUserDeptFinderSql, userCode)
		_, err = zorm.UpdateFinder(ctx, delUserDeptFinder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库删除用户部门数据错误: %v", err)
		}

		insertUserDeptFinder := zorm.NewFinder()
		insertUserDeptFinderSql := InsertUserDeptSql
		insertUserDeptFinder.Append(insertUserDeptFinderSql, userCode, t.UserDept, setUser, setUser)
		_, err = zorm.UpdateFinder(ctx, insertUserDeptFinder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库插入用户部门数据错误: %v", err)
		}

		return nil, nil
	})

	logx.Infof("err: %v", err)

	return err
}
