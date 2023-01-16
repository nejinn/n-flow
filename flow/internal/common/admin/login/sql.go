package login

// QueryUserInfoSql 获取用户信息sql
const QueryUserInfoSql = `
	select user_code as UserCode, is_del as IsDel
	from n_flow_user
	where user_account = ?
	  and user_pwd = ?
`

// CheckLoginPermitSql 检测登录权限
const CheckLoginPermitSql = `
	select count(nfur.id)
	from n_flow_user_role nfur
			 left join n_flow_role nfr on nfur.role_code = nfr.role_code
			 left join n_flow_permission_role nfpr on nfr.role_code = nfpr.role_code
			 left join n_flow_permission nfp on nfpr.permit_code = nfp.permit_code
	where nfur.user_code = ?
	  and nfr.is_del = 1
	  and nfp.is_del = 1
	  and nfp.permit_code = ?
	  and nfp.permit_method = ?
`
