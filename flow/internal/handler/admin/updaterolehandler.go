package admin

import (
	"net/http"

	"nFlow/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nFlow/flow/internal/logic/admin"
	"nFlow/flow/internal/svc"
	"nFlow/flow/internal/types"
)

func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestUpdateRole
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewUpdateRoleLogic(r.Context(), svcCtx)
		err := l.UpdateRole(&req)
		result.HttpResult(r, w, nil, err)
	}
}
