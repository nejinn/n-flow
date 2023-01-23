package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/role"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RequestUpdateRole) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a role.NFlowUpdateRole
	a = role.NFlowUpdateRoleParams{
		RoleName:  req.RoleName,
		RoleCode:  req.RoleCode,
		RoleDesc:  req.RoleDesc,
		RoleOrder: req.RoleOrder,
		Status:    req.Status,
		Permit:    req.Permit,
	}

	err := a.UpdateData(l.ctx)

	logx.Infof("err: %v", err)

	return err
}
