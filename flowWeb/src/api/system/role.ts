import {
  AddRoleParams,
  GetRoleListParams,
  GetRoleListResultModel,
  GetRolePermitParams,
  UpdateRoleParams,
} from '/@/api/system/model/roleModel';
import { defHttp } from '/@/utils/http/axios';

enum Api {
  RoleList = '/admin/api/v1/getRoleList',
  SwitchRoleStatus = '/admin/api/v1/switchRoleStatus',
  AddRole = '/admin/api/v1/addRole',
  GetRolePermit = '/admin/api/v1/getRolePermit',
  UpdateRole = '/admin/api/v1/updateRole',
}

export const getRoleListApi = (params?: GetRoleListParams) =>
  defHttp.get<GetRoleListResultModel>({ url: Api.RoleList, params });

export const SwitchRoleStatusApi = (code: string, status: number) =>
  defHttp.post({ url: Api.SwitchRoleStatus, params: { code, status } });

export const addRoleApi = (params: AddRoleParams) => defHttp.post({ url: Api.AddRole, params });

export function getRolePermitApi(params: GetRolePermitParams) {
  return defHttp.get<{ code: Array<string> }>({
    url: Api.GetRolePermit,
    params,
  });
}

export const updateRoleApi = (params: UpdateRoleParams) =>
  defHttp.post({ url: Api.UpdateRole, params });
