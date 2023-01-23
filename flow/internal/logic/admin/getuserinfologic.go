package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/user"
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

	avatar := l.svcCtx.Config.InitData.InitAvatar
	var a user.NFlowGetUserInfo
	a = user.NFlowGetUserInfoParams{UserCode: l.ctx.Value("userCode").(string)}
	res, err := a.GetUserInfo(l.ctx, avatar)

	logx.Infof("resp: %+v", res)

	return &res, err
}
