package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/permit"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DragglePermitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDragglePermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DragglePermitLogic {
	return &DragglePermitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DragglePermitLogic) DragglePermit(req *types.RequestDragglePermit) error {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a permit.NFlowDragglePermit

	a = permit.NFlowDragglePermitParams{
		Code:  req.Code,
		PCode: req.PCode,
	}

	err := a.UpdateData(l.ctx, l.svcCtx.Config.InitData.InitTopPermitCode)

	logx.Infof("err: %v", err)

	return err
}
