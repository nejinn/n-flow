package getUserInfo

// QueryBaseInfoSql 获取用户基本信息
const QueryBaseInfoSql = `
	select id          as UserId,
		   user_name   as UserName,
		   user_avatar as RealName
	from n_flow_user
	where user_code = ?
`

// QueryRoleSql 获取角色信息
const QueryRoleSql = `
	select nfp.permit_name as RoleName,
		   nfp.permit_code as Value
	from n_flow_permission nfp
			 left join n_flow_permission_role nfpr on nfp.permit_code = nfpr.permit_code
			 left join n_flow_role nfr on nfpr.role_code = nfr.role_code
			 left join n_flow_user_role nfur on nfr.role_code = nfur.role_code
	where nfp.is_del = 1
	  and nfr.is_del = 1
	  and nfur.user_code = ?
`
