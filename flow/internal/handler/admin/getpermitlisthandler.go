package admin

import (
	"net/http"

	"nFlow/common/result"
	"nFlow/flow/internal/logic/admin"
	"nFlow/flow/internal/svc"
)

func GetPermitListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewGetPermitListLogic(r.Context(), svcCtx)
		resp, err := l.GetPermitList()
		result.HttpResult(r, w, resp, err)
	}
}
