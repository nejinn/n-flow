package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/role"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.RequestAddRole) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a role.NFlowAddRole

	a = role.NFlowAddRoleParams{
		RoleOrder: req.RoleOrder,
		RoleName:  req.RoleName,
		RoleDesc:  req.RoleDesc,
		Status:    req.Status,
		Permit:    req.Permit,
	}

	err := a.InsertData(l.ctx)

	logx.Infof("err: %v", err)

	return err
}
