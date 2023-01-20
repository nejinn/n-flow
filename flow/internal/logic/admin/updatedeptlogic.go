package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/dept"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeptLogic) UpdateDept(req *types.RequestUpdateDept) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	userCode := l.ctx.Value("userCode").(string)

	var a dept.NFlowUpdateDept
	a = dept.NFlowUpdateDeptParams{
		DeptCode:   req.DeptCode,
		DeptParent: req.DeptParent,
		DeptOrder:  req.DeptOrder,
		DeptDesc:   req.DeptDesc,
		DeptName:   req.DeptName,
		Status:     req.Status,
	}

	err := a.UpdateData(l.ctx, userCode, l.svcCtx.Config.InitData.InitTopDeptCode)

	logx.Infof("resp: %v", err)

	return err
}
