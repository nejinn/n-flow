package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/dept"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeptLogic) AddDept(req *types.RequestAddDept) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	userCode := l.ctx.Value("userCode").(string)

	var a dept.NFlowAddDept

	a = dept.NFlowAddDeptParams{
		DeptName:   req.DeptName,
		DeptDesc:   req.DeptDesc,
		DeptOrder:  req.DeptOrder,
		DeptParent: req.DeptParent,
		Status:     req.Status,
	}

	err := a.InsertData(l.ctx, userCode, l.svcCtx.Config.InitData.TopDeptCode)

	logx.Infof("resp: %v", err)
	
	return err
}
