import {
  DraggablePermitParams,
  GetPermitListItem,
  GetRolePermitParams,
} from '/@/api/system/model/permitModel';
import { defHttp } from '/@/utils/http/axios';
enum Api {
  PermitList = '/admin/api/v1/getPermitList',
  DraggablePermit = '/admin/api/v1/dragglePermit',
  GetRolePermit = '/admin/api/v1/getRolePermit',
}

export const getPermitListApi = () => defHttp.get<GetPermitListItem>({ url: Api.PermitList });

export const draggablePermitApi = (params: DraggablePermitParams) =>
  defHttp.post({ url: Api.DraggablePermit, params });

export function getRolePermitApi(params: GetRolePermitParams) {
  return defHttp.get<{ code: Array<string> }>({
    url: Api.GetRolePermit,
    params,
  });
}
