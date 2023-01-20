package role

// GetRoleListSql 获取角色列表sql
const GetRoleListSql = `
	select nfr.role_code                                as RoleCode,
		   nfr.role_name                                as RoleName,
		   nfr.role_desc                                as RoleDesc,
		   nfr.role_order                               as RoleOrder,
		   nfr.is_del                                   as Status,
		   nfu.user_name                                as CUser,
		   n.user_name                                  as UUser,
		   date_format(nfr.c_time, "%Y-%m-%d %H:%i:%s") as CTime,
		   date_format(nfr.u_time, "%Y-%m-%d %H:%i:%s") as UTime
	from n_flow_role nfr
			 left join n_flow_user nfu on nfr.c_user = nfu.user_code
			 left join n_flow_user n on nfr.u_user = n.user_code
	where 1 = 1
`

// GetRoleListCountSql 获取角色数量Sql
const GetRoleListCountSql = `
	select count(id)
	from n_flow_role nfr
	where 1 = 1
`

// SwitchRoleStatusSql 修改角色状态sql
const SwitchRoleStatusSql = `
	update n_flow_role set is_del = ? where role_code = ?
`

// GetRolePermitSql 获取角色权限Sql
const GetRolePermitSql = `
	select permit_code as Code
	from n_flow_permission_role
	where role_code = ?
`
