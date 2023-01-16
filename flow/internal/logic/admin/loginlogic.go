package admin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/flow/internal/common/admin/login"
	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.RequestLogin) (resp *types.ResponseLogin, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %v", req)

	var a login.NFlowLogin
	a = login.NFlowLoginParams{
		Account:  req.Account,
		Password: req.Password,
	}

	now := time.Now().Unix()
	exp := now + l.svcCtx.Config.Auth.AccessExpire
	secretKey := l.svcCtx.Config.Auth.AccessSecret
	res, err := a.GenerateToken(l.ctx, exp, now, secretKey)

	if err != nil {
		return nil, err
	}

	resp = &types.ResponseLogin{
		Token: res,
	}
	logx.Infof("resp: %+v", resp)
	return &types.ResponseLogin{
		Token: res,
	}, nil
}
