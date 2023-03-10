package account

// GetUserBaseInfoListSql 获取账户列表基础信息
const GetUserBaseInfoListSql = `
select nfu.user_code                                as UserCode,
       nfu.user_account                             as UserAccount,
       nfu.user_name                                as UserName,
       nfu.user_avatar                              as UserAvatar,
       nfu.user_mail                                as UserMail,
       nfu.is_del                                   as Status,
       date_format(nfu.c_time, "%Y-%m-%d %H:%i:%s") as CTime,
       date_format(nfu.u_time, "%Y-%m-%d %H:%i:%s") as Utime,
       nfd.dept_name                                as UserDept,
       group_concat(nfr.role_name)                  as UserRoles,
       group_concat(nfr.role_code)                  as UserRoleCodes
from n_flow_user nfu
         left join n_flow_user_dept nfud on nfu.user_code = nfud.user_code
         left join n_flow_dept nfd on nfd.dept_code = nfud.dept_code
         left outer join n_flow_user_role nfur on nfu.user_code = nfur.user_code
         left join n_flow_role nfr on nfur.role_code = nfr.role_code
where 1 = 1
`

// GetUserBaseInfoListCountSql 获取账户总数量
const GetUserBaseInfoListCountSql = `
	select count(nfu.id)
	from n_flow_user nfu
			 left join n_flow_user_dept nfud on nfu.user_code = nfud.user_code
			 left join n_flow_dept nfd on nfd.dept_code = nfud.dept_code where 1=1
`

// InsertUserSql 添加用户sql
const InsertUserSql = `
insert into n_flow_user(user_code,
                        user_account,
                        user_pwd,
                        user_name,
                        user_mail,
                        is_del)
values (?, ?, ?, ?, ?, ?)
`

// InsertUserRoleSql 插入用户角色sql
const InsertUserRoleSql = `
	insert into n_flow_user_role(user_code,
								 role_code,
								 c_user,
								 u_user)
	values
`

// InsertUserDeptSql 插入用户部门Sql
const InsertUserDeptSql = `
insert into n_flow_user_dept(user_code,
                             dept_code,
                             c_user,
                             u_user)
VALUES (?, ?, ?, ?)
`

// UpdateAccountSql 修改用户信息
const UpdateAccountSql = `
	update n_flow_user
	set user_account = ?,
		user_name    = ?,
		user_mail    = ?,
		is_del       = ?
	where user_code = ?
`

// DelUserRoleSql 删除用户角色数据
const DelUserRoleSql = `
	delete from n_flow_user_role where user_code = ?
`

// DelUserDeptSql 删除用户部门数据
const DelUserDeptSql = `
	delete from n_flow_user_dept where user_code = ?
`
