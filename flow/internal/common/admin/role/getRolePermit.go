package role

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
	"nFlow/flow/internal/types"
)

// NFlowGetRolePermit 获取角色权限接口
type NFlowGetRolePermit interface {
	GetRolePermit(ctx context.Context) (*types.ResponseGetRolePermit, error)
}

// GetRolePermit 获取角色权限
func (t NFlowGetRolePermitParams) GetRolePermit(c context.Context) (resp *types.ResponseGetRolePermit, err error) {

	logx.Infof("t: %+v", t)
	finder := zorm.NewFinder()
	finderSql := GetRolePermitSql
	finder.Append(finderSql, t.Code)

	var permits []string
	err = zorm.Query(c, finder, &permits, nil)

	if err != nil {
		return resp, errors.Wrapf(xerr.NewErrMsg("服务器错误,请重试!"),
			"查询角色权限出错, err: %v", err)
	}

	logx.Infof("resp: %+v", resp)

	return &types.ResponseGetRolePermit{Code: permits}, nil
}
