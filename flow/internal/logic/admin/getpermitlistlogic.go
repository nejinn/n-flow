package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/permit"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermitListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermitListLogic {
	return &GetPermitListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermitListLogic) GetPermitList() (resp []*types.ResponseGetPermitList, err error) {
	// todo: add your logic here and delete this line

	resp, _ = permit.GetPermitList(l.ctx)

	return resp, nil
}
