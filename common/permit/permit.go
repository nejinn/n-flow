package permit

import "github.com/pkg/errors"

type BasePermission struct {
	Code            string
	ApiUrl          string
	ApiMethod       string
	UserLogType     string
	UserLogTypeName string
	ParentCode      string
}

var AdminLoginPermit = BasePermission{
	Code:            "admin:login",
	ApiUrl:          "/nflow/admin/api/v1/login",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "",
}

var AdminKeyWordPermit = BasePermission{
	Code:            "admin:keyWord",
	ApiUrl:          "/nflow/admin/api/v1/keyWord",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "",
}

var AdminGetUserInfoPermit = BasePermission{
	Code:            "admin:getUserInfo",
	ApiUrl:          "/nflow/admin/api/v1/getUserInfo",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "",
}

var AdminSystemMenuPermit = BasePermission{
	Code:            "admin:system",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "",
}

var AdminSystemDeptMenuPermit = BasePermission{
	Code:            "admin:system:dept",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system",
}

var AdminGetDeptListPermit = BasePermission{
	Code:            "admin:system:dept:getDeptList",
	ApiUrl:          "/nflow/admin/api/v1/getDeptList",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:dept",
}

var AdminAddDeptPermit = BasePermission{
	Code:            "admin:system:dept:addDept",
	ApiUrl:          "/nflow/admin/api/v1/addDept",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:dept",
}

var AdminUpdateDeptPermit = BasePermission{
	Code:            "admin:system:dept:updateDept",
	ApiUrl:          "/nflow/admin/api/v1/updateDept",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:dept",
}

var AdminSystemAccountMenuPermit = BasePermission{
	Code:            "admin:system:account",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system",
}

var AdminGetAccountListPermit = BasePermission{
	Code:            "admin:system:account:getUserList",
	ApiUrl:          "/nflow/admin/api/v1/getUserList",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:account",
}

var AdminSystemRoleMenuPermit = BasePermission{
	Code:            "admin:system:role",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system",
}

var AdminGetRoleListPermit = BasePermission{
	Code:            "admin:system:role:getRoleList",
	ApiUrl:          "/nflow/admin/api/v1/getRoleList",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:role",
}

var AdminSwitchRoleStatusPermit = BasePermission{
	Code:            "admin:system:role:switchRoleStatus",
	ApiUrl:          "/nflow/admin/api/v1/switchRoleStatus",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:role",
}

var AdminSystemPermitMenuPermit = BasePermission{
	Code:            "admin:system:permit",
	ApiUrl:          "",
	ApiMethod:       "",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system",
}

var AdminGetPermitListPermit = BasePermission{
	Code:            "admin:system:permit:getPermitList",
	ApiUrl:          "/nflow/admin/api/v1/getPermitList",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:permit",
}

var AdminDragglePermitPermit = BasePermission{
	Code:            "admin:system:permit:dragglePermit",
	ApiUrl:          "/nflow/admin/api/v1/dragglePermit",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:permit",
}

var AdminGetRolePermitPermit = BasePermission{
	Code:            "admin:system:role:getRolePermit",
	ApiUrl:          "/nflow/admin/api/v1/getRolePermit",
	ApiMethod:       "GET",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:role",
}

var AdminAddRolePermit = BasePermission{
	Code:            "admin:system:role:addRole",
	ApiUrl:          "/nflow/admin/api/v1/addRole",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:role",
}

var AdminUpdateRolePermit = BasePermission{
	Code:            "admin:system:role:updateRole",
	ApiUrl:          "/nflow/admin/api/v1/updateRole",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:role",
}

var AdminCheckUserAccountPermit = BasePermission{
	Code:            "admin:system:account:checkUserAccount",
	ApiUrl:          "/nflow/admin/api/v1/checkUserAccount",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:account",
}

var AdminAddAccountPermit = BasePermission{
	Code:            "admin:system:account:addAccount",
	ApiUrl:          "/nflow/admin/api/v1/addAccount",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:account",
}

var AdminUpdateAccountPermit = BasePermission{
	Code:            "admin:system:account:updateAccount",
	ApiUrl:          "/nflow/admin/api/v1/updateAccount",
	ApiMethod:       "POST",
	UserLogType:     "",
	UserLogTypeName: "",
	ParentCode:      "admin:system:account",
}

var permitMap map[string]BasePermission

func init() {
	permitMap = make(map[string]BasePermission)
	permitMap[AdminLoginPermit.ApiUrl] = AdminLoginPermit
	permitMap[AdminGetUserInfoPermit.ApiUrl] = AdminGetUserInfoPermit
	permitMap[AdminGetDeptListPermit.ApiUrl] = AdminGetDeptListPermit
	permitMap[AdminAddDeptPermit.ApiUrl] = AdminAddDeptPermit
	permitMap[AdminUpdateDeptPermit.ApiUrl] = AdminUpdateDeptPermit
	permitMap[AdminGetAccountListPermit.ApiUrl] = AdminGetAccountListPermit
	permitMap[AdminGetRoleListPermit.ApiUrl] = AdminGetRoleListPermit
	permitMap[AdminSwitchRoleStatusPermit.ApiUrl] = AdminSwitchRoleStatusPermit
	permitMap[AdminGetPermitListPermit.ApiUrl] = AdminGetPermitListPermit
	permitMap[AdminDragglePermitPermit.ApiUrl] = AdminDragglePermitPermit
	permitMap[AdminGetRolePermitPermit.ApiUrl] = AdminGetRolePermitPermit
	permitMap[AdminAddRolePermit.ApiUrl] = AdminAddRolePermit
	permitMap[AdminUpdateRolePermit.ApiUrl] = AdminUpdateRolePermit
	permitMap[AdminCheckUserAccountPermit.ApiUrl] = AdminCheckUserAccountPermit
	permitMap[AdminKeyWordPermit.ApiUrl] = AdminKeyWordPermit
	permitMap[AdminAddAccountPermit.ApiUrl] = AdminAddAccountPermit
	permitMap[AdminUpdateAccountPermit.ApiUrl] = AdminUpdateAccountPermit

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
