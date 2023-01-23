package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/user"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserAccountLogic {
	return &CheckUserAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckUserAccountLogic) CheckUserAccount(req *types.RequestCheckUserAccount) (resp *types.ResponseCheckUserAccount, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a user.NFlowCheckUserAccount
	a = user.NFlowCheckUserAccountParams{
		Account: req.Account,
	}

	err = a.Check(l.ctx)

	logx.Infof("err: %v", err)

	if err != nil {
		return &types.ResponseCheckUserAccount{Code: 1}, nil
	}
	return &types.ResponseCheckUserAccount{Code: 2}, nil
}
