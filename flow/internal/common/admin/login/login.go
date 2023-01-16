package login

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/md5"
	"nFlow/common/permit"
	"nFlow/common/xerr"
)

type NFlowLogin interface {
	QueryUserInfo(c context.Context) (ResLoginUserInfo, error)                               // 查询登录用户信息
	ValidUserStatus(c context.Context) (string, error)                                       //校验用户状态
	CheckLoginPermit(c context.Context) (string, error)                                      // 校验用户是否有的登录权限
	GenerateToken(c context.Context, exp int64, iat int64, secretKey string) (string, error) //生成token
}

// QueryUserInfo 查询登录用户信息
func (t NFlowLoginParams) QueryUserInfo(c context.Context) (resp ResLoginUserInfo, err error) {

	logx.Infof("t: %+v", t)

	md5Pwd, _ := md5.MdFiveString(t.Password)

	finder := zorm.NewFinder()
	finderSql := QueryUserInfoSql
	finder.Append(finderSql, t.Account, md5Pwd)

	has, err := zorm.QueryRow(c, finder, &resp)
	if has == false {
		return resp, errors.Wrapf(xerr.NewErrMsg("账户或者密码错误"), "未查询到此用户, has: %t", has)
	}

	if err != nil {
		return resp, errors.Wrapf(xerr.NewErrMsg("账户或者密码错误"), "数据库查询错误, err: %v", err)
	}

	logx.Infof("resp: %+v", resp)
	return resp, nil

}

// ValidUserStatus 校验用户状态
func (t NFlowLoginParams) ValidUserStatus(c context.Context) (resp string, err error) {

	logx.Infof("t: %+v", t)

	res, err := t.QueryUserInfo(c)

	if err != nil {
		return resp, err
	}

	if res.IsDel != 1 {
		return resp, errors.Wrapf(xerr.NewErrMsg("用户已停用"), "用户状态为2, 用户已停用, isDel: %d", res.IsDel)
	}

	logx.Infof("resp: %s", res.UserCode)

	return res.UserCode, nil
}

// CheckLoginPermit 校验用户是否有的登录权限
func (t NFlowLoginParams) CheckLoginPermit(c context.Context) (userCode string, err error) {

	logx.Infof("t: %+v", t)

	res, err := t.ValidUserStatus(c)

	if err != nil {
		return userCode, err
	}

	finder := zorm.NewFinder()
	finderSql := CheckLoginPermitSql
	finder.Append(finderSql, res, permit.AdminLoginPermit.Code, permit.AdminLoginPermit.ApiMethod)

	var count int64
	has, err := zorm.QueryRow(c, finder, &count)

	if has == false || count == 0 {
		return userCode, errors.Wrapf(xerr.NewErrMsg("非法登录"), "此用户无登录权限, has: %t, count: %d", has, count)

	}

	if err != nil {
		return userCode, errors.Wrapf(xerr.NewErrMsg("非法登录"), "数据库查询错误, err: %v", err)
	}

	logx.Infof("resp: %s", res)

	return res, nil
}

// GenerateToken 生成token
func (t NFlowLoginParams) GenerateToken(c context.Context, exp int64, iat int64, secretKey string) (token string, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("exp: %d, iat: %d", exp, iat)

	res, err := t.CheckLoginPermit(c)

	if err != nil {
		return token, err
	}

	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userCode"] = res
	tokenObj := jwt.New(jwt.SigningMethodHS256)
	tokenObj.Claims = claims
	token, err = tokenObj.SignedString([]byte(secretKey))

	logx.Infof("token: %s, err: %+v", token, err)
	return token, err
}
