package account

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/generateCode"
	"nFlow/common/md5"
	"nFlow/common/xerr"
	"regexp"
)

type NFlowAddAccount interface {
	ValidParams() error
	InsertData(context.Context) error
}

func (t NFlowAddAccountParams) ValidParams() error {

	logx.Infof("t: %+v", t)

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

func (t NFlowAddAccountParams) InsertData(c context.Context) error {

	err := t.ValidParams()

	if err != nil {
		return err
	}

	logx.Infof("t: %+v", t)

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {

		userCode := generateCode.GenerateCode(3)
		pwd, _ := md5.MdFiveString(t.Password)
		setUser := c.Value("userCode").(string)

		finder := zorm.NewFinder()
		finderSql := InsertUserSql
		finder.Append(finderSql, userCode, t.UserAccount, pwd, t.UserName, t.UserMail, t.Status)
		_, err := zorm.UpdateFinder(ctx, finder)

		if err != nil {
			return nil, errors.Wrapf(xerr.CodeErrMsg(xerr.SERVER_COMMON_ERROR),
				"数据库插入用户数据错误: %v", err)
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
