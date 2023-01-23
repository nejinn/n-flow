export enum RoleEnum {
  ADMINLOGIN = 'admin:login', // 登录
  ADMINGETUSERINFO = 'admin:getUserInfo', // 获取用户信息
  ADMINSYSTEMDEPTGETDEPTLIST = 'admin:system:dept:getDeptList', // 获取部门信息
  ADMINSYSTEMDEPTADDDEPT = 'admin:system:dept:addDept', // 新增部门
  ADMINSYSTEM = 'admin:system', // 系统管理菜单
  ADMINSYSTEMDEPT = 'admin:system:dept', // 系统管理-部门管理菜单
  ADMINSYSTEMDEPTUPDATEDEPT = 'admin:system:dept:updateDept', // 修改部门
  ADMINSYSTEMACCOUNT = 'admin:system:account', // 系统管理-账号管理菜单
  ADMINSYSTEMACCOUNTGETUSERLIST = 'admin:system:account:getUserList', // 获取账户菜单
  ADMINSYSTEMROLE = 'admin:system:role', // 系统管理-角色管理菜单
  ADMINSYSTEMROLEGETROLELIST = 'admin:system:role:getRoleList', // 获取角色列表
  ADMINSYSTEMYROLESWITCHROLESTATUS = 'admin:system:role:switchRoleStatus', // 修改角色状态
  ADMINSYSTEMPERMIT = 'admin:system:permit', // 系统管理-权限管理菜单
  ADMINSYSTEMPERMITGETPERMITLIST = 'admin:system:permit:getPermitList', // 获取权限列表
  ADMINSYSTEMPERMITDRAGGLEPERMIT = 'admin:system:permit:dragglePermit', //拖拽权限树
  ADMINSYSTEMROLEGETROLEPERMIT = 'admin:system:role:getRolePermit', // 获取角色所有权限
  ADMINSYSTEMROLEADDROLE = 'admin:system:role:addRole', // 添加角色
  ADMINSYSTEMROLEUPDATEROLE = 'admin:system:role:updateRole', // 修改角色
}
