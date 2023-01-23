package admin

import (
	"context"
	"nFlow/flow/internal/common/admin/keyWord"

	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeyWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyWordLogic {
	return &KeyWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeyWordLogic) KeyWord(req *types.RequestKeyWord) (resp interface{}, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("req: %+v", req)

	var a keyWord.NFlowKeyWord

	a = keyWord.NFlowKeyWordParams{
		Category: req.Category,
		Word:     req.Word,
	}

	resp, err = a.GetResult(l.ctx)

	logx.Infof("resp: %+v, err: %v", resp, err)

	return resp, err
}
