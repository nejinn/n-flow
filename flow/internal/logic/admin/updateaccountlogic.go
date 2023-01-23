package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/account"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountLogic {
	return &UpdateAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountLogic) UpdateAccount(req *types.RequestUpdateAccount) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a account.NFlowUpdateAccount
	a = account.NFlowUpdateAccountParams{
		UserCode:    req.UserCode,
		UserAccount: req.UserAccount,
		UserName:    req.UserName,
		Status:      req.Status,
		UserRoles:   req.UserRoles,
		UserDept:    req.UserDept,
		UserMail:    req.UserMail,
	}

	err := a.UpdateData(l.ctx, l.svcCtx.Config.InitData.InitUserCode)

	logx.Infof("err: %v", err)
	return err
}
