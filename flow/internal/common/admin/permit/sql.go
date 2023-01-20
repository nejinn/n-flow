package permit

const GetPermitListSql = `
	select permit_code   as PermitCode,
		   permit_name   as PermitName,
		   permit_desc   as PermitDesc,
		   permit_order  as PermitOrder,
		   permit_type   as PermitType,
		   permit_method as PermitMethod,
		   permit_url    as PermitUrl,
		   permit_parent as PermitParent,
		   is_del        as status
	from n_flow_permission
`

const ValidPCodeInDbSql = `
select count(id) from n_flow_permission where permit_code = ?
`

const UpdatePermitPCodeSql = `
	update n_flow_permission set permit_parent = ?, u_user = ? WHERE permit_code = ?
`
