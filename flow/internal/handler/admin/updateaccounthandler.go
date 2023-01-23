package admin

import (
	"net/http"

	"nFlow/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nFlow/flow/internal/logic/admin"
	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"
)

func UpdateAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestUpdateAccount
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewUpdateAccountLogic(r.Context(), svcCtx)
		err := l.UpdateAccount(&req)
		result.HttpResult(r, w, nil, err)
	}
}
