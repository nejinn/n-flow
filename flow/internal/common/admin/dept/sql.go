package dept

// GetDeptListSql 获取部门泪飙数据查询sql
const GetDeptListSql = `
	select nfd.id                                       as Id,
		   nfd.dept_code                                as DeptCode,
		   nfd.dept_name                                as DeptName,
		   nfd.dept_desc                                as DeptDesc,
		   nfd.dept_order                               as DeptOrder,
		   nfd.dept_parent                              as DeptParent,
		   nfd.is_del                                   as Status,
		   nfu.user_name                            as CUser,
		   nfu1.user_name                           as UUser,
		   date_format(nfd.c_time, "%Y-%m-%d %H:%i:%s") as CTime,
		   date_format(nfd.u_time, "%Y-%m-%d %H:%i:%s") as UTime
	from n_flow_dept nfd
			 left join n_flow_user nfu on nfd.c_user = nfu.user_code
			 left join n_flow_user nfu1 on nfd.u_user = nfu1.user_code
	where 1 = 1
`

// InsertDeptSql 插入部门数据sql
const InsertDeptSql = `
	insert into n_flow_dept(dept_code, dept_name, dept_desc, dept_order, dept_parent, is_del, c_user, u_user)
	values (?, ?, ?, ?, ?, ?, ?, ?)
`

// UpdateDeptSql 修改部门数据sql
const UpdateDeptSql = `
	update n_flow_dept
	set dept_name=?,
		dept_desc=?,
		dept_order=?,
		dept_parent=?,
		is_del=?,
		u_user=?
	where dept_code = ?
`
