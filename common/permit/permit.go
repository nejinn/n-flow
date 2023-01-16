package permit

import "github.com/pkg/errors"

type BasePermission struct {
	Code            string
	ApiUrl          string
	ApiMethod       string
	UserLogType     string
	UserLogTypeName string
}

var AdminLoginPermit = BasePermission{
	Code:            "admin:login",
	ApiUrl:          "/nflow/admin/api/v1/login",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminGetUserInfoPermit = BasePermission{
	Code:            "admin:getUserInfo",
	ApiUrl:          "/nflow/admin/api/v1/getUserInfo",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminSystemMenuPermit = BasePermission{
	Code:            "admin:system",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminSystemDeptMenuPermit = BasePermission{
	Code:            "admin:system:dept",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminGetDeptListPermit = BasePermission{
	Code:            "admin:system:dept:getDeptList",
	ApiUrl:          "/nflow/admin/api/v1/getDeptList",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminAddDeptPermit = BasePermission{
	Code:            "admin:system:dept:addDept",
	ApiUrl:          "/nflow/admin/api/v1/addDept",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
}

var AdminUpdateDeptPermit = BasePermission{
	Code:            "admin:system:dept:updateDept",
	ApiUrl:          "/nflow/admin/api/v1/updateDept",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
}

var permitMap map[string]BasePermission

func init() {
	permitMap = make(map[string]BasePermission)
	permitMap[AdminLoginPermit.ApiUrl] = AdminLoginPermit
	permitMap[AdminGetUserInfoPermit.ApiUrl] = AdminGetUserInfoPermit
	permitMap[AdminGetDeptListPermit.ApiUrl] = AdminGetDeptListPermit
	permitMap[AdminAddDeptPermit.ApiUrl] = AdminAddDeptPermit
	permitMap[AdminUpdateDeptPermit.ApiUrl] = AdminUpdateDeptPermit
}

// MapPermitUrlGetObj 根据Url获取对应的权限对象
func MapPermitUrlGetObj(apiUrl string) (resp BasePermission, err error) {
	if permit, ok := permitMap[apiUrl]; ok {
		return permit, nil
	} else {
		return BasePermission{
			Code:            "",
			ApiUrl:          "",
			ApiMethod:       "",
			UserLogType:     "",
			UserLogTypeName: "",
		}, errors.New("非法请求")
	}
}
