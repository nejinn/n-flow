package admin

import (
	"net/http"

	"nFlow/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nFlow/flow/internal/logic/admin"
	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"
)

func KeyWordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestKeyWord
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewKeyWordLogic(r.Context(), svcCtx)
		resp, err := l.KeyWord(&req)
		result.HttpResult(r, w, resp, err)
	}
}
