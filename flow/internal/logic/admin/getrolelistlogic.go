package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/role"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.RequestGetRoleList) (resp *types.ResponseGetRoleListRes, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a role.NFlowGetRoleList

	a = role.NFlowGetRoleListParams{
		RoleName: req.RoleName,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	resp = a.GetRoleList(l.ctx)

	logx.Infof("resp: %+v", resp)

	return resp, nil
}
