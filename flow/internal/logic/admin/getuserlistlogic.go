package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/account"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.RequestGetUserList) (resp *types.ResponseGetUserListRes, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a account.NFlowGetUserList

	a = account.NFlowGetAccountListParams{
		UserAccount: req.UserAccount,
		UserName:    req.UserName,
		Status:      req.Status,
		Dept:        req.Dept,
		Page:        req.Page,
		PageSize:    req.PageSize,
	}

	resp, _ = a.GetUserBaseInfoList(l.ctx)

	return resp, nil
}
