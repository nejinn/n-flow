package account

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
       group_concat(nfr.role_name)                  as UserRoles
from n_flow_user nfu
         left join n_flow_user_dept nfud on nfu.user_code = nfud.user_code
         left join n_flow_dept nfd on nfd.dept_code = nfud.dept_code
         left outer join n_flow_user_role nfur on nfu.user_code = nfur.user_code
         left join n_flow_role nfr on nfur.role_code = nfr.role_code
where 1 = 1
`

const GetUserBaseInfoListCountSql = `
	select count(nfu.id)
	from n_flow_user nfu
			 left join n_flow_user_dept nfud on nfu.user_code = nfud.user_code
			 left join n_flow_dept nfd on nfd.dept_code = nfud.dept_code where 1=1
`
