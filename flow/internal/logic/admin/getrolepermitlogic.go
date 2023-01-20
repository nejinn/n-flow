package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/role"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermitLogic {
	return &GetRolePermitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermitLogic) GetRolePermit(req *types.RequestGetRolePermit) (resp *types.ResponseGetRolePermit, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a role.NFlowGetRolePermit

	a = role.NFlowGetRolePermitParams{
		Code: req.Code,
	}

	resp, err = a.GetRolePermit(l.ctx)

	logx.Infof("resp: %+v; err: %v", resp, err)

	return
}
