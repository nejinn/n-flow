package keyWord

const GetRoleListSql = `
	select role_name as Name, role_code as Value from n_flow_role where 1=1
`

const GetDeptListSql = `
	select dept_name as Name, dept_code as Value, dept_parent as PValue
	from n_flow_dept
	where 1 = 1
`
