package admin

import (
	"net/http"

	"nFlow/common/result"
	"nFlow/flow/internal/logic/admin"
	"nFlow/flow/internal/svc"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		result.HttpResult(r, w, resp, err)
	}
}
