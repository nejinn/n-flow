package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/role"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchRoleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchRoleStatusLogic {
	return &SwitchRoleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchRoleStatusLogic) SwitchRoleStatus(req *types.RequestSwitchStatus) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a role.NFlowSwitchRoleStatus

	a = role.NFlowSwitchRoleStatusParams{
		Code:   req.Code,
		Status: req.Status,
	}

	err := a.SwitchStatus(l.ctx, l.svcCtx.Config.InitData.InitSuperRoleCode)

	return err
}
