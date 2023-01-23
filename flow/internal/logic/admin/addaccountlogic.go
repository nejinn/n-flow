package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/account"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAccountLogic {
	return &AddAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAccountLogic) AddAccount(req *types.RequestAddAccount) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a account.NFlowAddAccount
	a = account.NFlowAddAccountParams{
		UserAccount: req.UserAccount,
		UserName:    req.UserName,
		Password:    req.Password,
		Status:      req.Status,
		UserRoles:   req.UserRoles,
		UserDept:    req.UserDept,
		UserMail:    req.UserMail,
	}
	err := a.InsertData(l.ctx)

	logx.Infof("err: %v", err)

	return err
}
