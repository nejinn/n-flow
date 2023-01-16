package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/getUserInfo"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.ResponseGetUserInfo, err error) {

	avatar := l.svcCtx.Config.InitData.Avatar
	var a getUserInfo.NFlowGetUserInfo
	a = getUserInfo.NFlowGetUserInfoParams{UserCode: l.ctx.Value("userCode").(string)}
	res, err := a.GetUserInfo(l.ctx, avatar)

	logx.Infof("resp: %+v", res)

	return &res, err
}
