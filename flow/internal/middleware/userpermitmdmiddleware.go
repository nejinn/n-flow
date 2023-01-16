package middleware

import (
	"gitee.com/chunanyong/zorm"
	"nFlow/common/permit"
	"nFlow/common/result"
	"nFlow/common/xerr"
	"net/http"
)

type UserPermitMdMiddleware struct {
}

func NewUserPermitMdMiddleware() *UserPermitMdMiddleware {
	return &UserPermitMdMiddleware{}
}

func (m *UserPermitMdMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userCode := r.Context().Value("userCode")
		if userCode == nil {
			result.MapErrorResult(r, w, xerr.TOKEN_EXPIRE_ERROR, "登录过期")
			return
		}

		permitObj, err := permit.MapPermitUrlGetObj(r.URL.Path)

		if err != nil {
			result.MapErrorResult(r, w, xerr.SERVER_COMMON_ERROR, "服务器错误")
			return
		}

		var count int64

		finder := zorm.NewFinder()
		finderSql := `
			select count(nfp.id)
			from n_flow_permission nfp
					 left join n_flow_permission_role nfpr on nfp.permit_code = nfpr.permit_code
					 left join n_flow_role nfr on nfpr.role_code = nfr.role_code
					 left join n_flow_user_role nfur on nfr.role_code = nfur.role_code
			where nfp.permit_code = ?
			  and nfp.permit_method = ?
			  and nfp.is_del = 1
			  and nfr.is_del = 1
			  and nfur.user_code = ?
		`
		finder.Append(finderSql, permitObj.Code, permitObj.ApiMethod, userCode)

		has, err := zorm.QueryRow(r.Context(), finder, &count)

		if has == false || count == 0 {
			result.MapErrorResult(r, w, xerr.PERMIT_NO_ACCESS, "您无此权限")
			return
		}

		if err != nil {
			result.MapErrorResult(r, w, xerr.SERVER_COMMON_ERROR, "服务器错误")
			return
		}
		next(w, r)
	}
}
