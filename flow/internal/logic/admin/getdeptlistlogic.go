package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/dept"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptListLogic {
	return &GetDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeptListLogic) GetDeptList(req *types.RequestGetDeptList) (resp []types.ResponseGetDeptList, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("req: %+v", req)

	var a dept.NFlowDeptList
	a = dept.NFlowGetDeptListParams{
		Name:   req.Name,
		Status: req.Status,
		Top:    req.Top,
	}

	resp, err = a.GetDeptList(l.ctx, l.svcCtx.Config.InitData.TopDeptCode, l.svcCtx.Config.InitData.TopDeptName)

	logx.Infof("resp: %+v, err: %v", resp, err)
	return resp, err
}
