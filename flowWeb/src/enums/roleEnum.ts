export enum RoleEnum {
  ADMINLOGIN = 'admin:login', // 登录
  ADMINGETUSERINFO = 'admin:getUserInfo', // 获取用户信息
  ADMINSYSTEMDEPTGETDEPTLIST = 'admin:system:dept:getDeptList', // 获取部门信息
  ADMINSYSTEMDEPTADDDEPT = 'admin:system:dept:addDept', // 新增部门
  ADMINSYSTEM = 'admin:system', // 系统管理菜单
  ADMINSYSTEMDEPT = 'admin:system:dept', // 系统管理-部门管理菜单
  ADMINSYSTEMDEPTUPDATEDEPT = 'admin:system:dept:updateDept', // 修改部门
}
