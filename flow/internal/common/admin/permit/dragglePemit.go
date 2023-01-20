package permit

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/xerr"
)

type NFlowDragglePermit interface {
	ValidParams(c context.Context, topPermitCode string) error
	UpdateData(c context.Context, topPermitCode string) error
}

// ValidParams 校验参数
func (t NFlowDragglePermitParams) ValidParams(c context.Context, topPermitCode string) (err error) {

	logx.Infof("t: %+v", t)

	if t.PCode == topPermitCode || t.PCode == "" {
		return nil
	}

	finder := zorm.NewFinder()
	finderSql := ValidPCodeInDbSql
	finder.Append(finderSql, t.PCode)
	var count int64
	has, err := zorm.QueryRow(c, finder, &count)

	if has == false || count == 0 || err != nil {
		logx.Infof("has: %t; err: %v", has, err)
		return errors.Wrapf(xerr.NewErrMsg("PCode参数非法"),
			"PCode参数非法，此值不存在与数据库, Pcode: %s", t.PCode)
	}
	return nil
}

// UpdateData 修改数据
func (t NFlowDragglePermitParams) UpdateData(c context.Context, topPermitCode string) (err error) {

	logx.Infof("t: %+v", t)

	err = t.ValidParams(c, topPermitCode)

	if err != nil {
		return err
	}

	_, err = zorm.Transaction(c, func(ctx context.Context) (interface{}, error) {
		finder := zorm.NewFinder()
		finderSql := UpdatePermitPCodeSql
		PCode := t.PCode
		if t.PCode == "" {
			PCode = topPermitCode
		}
		finder.Append(finderSql, PCode, c.Value("userCode").(string), t.Code)

		_, err = zorm.UpdateFinder(ctx, finder)
		return nil, err
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("服务器错误, 请重试！"),
			"修改权限出错, err: %v", err)
	}

	return nil
}
