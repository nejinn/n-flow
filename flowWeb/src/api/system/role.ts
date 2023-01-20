import { GetRoleListParams, GetRoleListResultModel } from '/@/api/system/model/roleModel';
import { defHttp } from '/@/utils/http/axios';
enum Api {
  RoleList = '/admin/api/v1/getRoleList',
  SwitchRoleStatus = '/admin/api/v1/switchRoleStatus',
}

export const getRoleListApi = (params?: GetRoleListParams) =>
  defHttp.get<GetRoleListResultModel>({ url: Api.RoleList, params });

export const SwitchRoleStatusApi = (code: string, status: number) =>
  defHttp.post({ url: Api.SwitchRoleStatus, params: { code, status } });
