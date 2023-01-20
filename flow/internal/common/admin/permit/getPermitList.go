package permit

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/flow/internal/types"
)

func GetPermitList(c context.Context) (resp []*types.ResponseGetPermitList, err error) {

	finder := zorm.NewFinder()
	finderSql := GetPermitListSql
	finder.Append(finderSql)
	err = zorm.Query(c, finder, &resp, nil)

	if err != nil {
		logx.Infof("err: %v", err)
	}

	logx.Infof("resp: %+v", resp)

	return resp, nil
}
