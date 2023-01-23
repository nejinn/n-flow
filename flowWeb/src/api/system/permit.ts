import { DraggablePermitParams, GetPermitListItem } from '/@/api/system/model/permitModel';
import { defHttp } from '/@/utils/http/axios';
enum Api {
  PermitList = '/admin/api/v1/getPermitList',
  DraggablePermit = '/admin/api/v1/dragglePermit',
}

export const getPermitListApi = () => defHttp.get<GetPermitListItem>({ url: Api.PermitList });

export const draggablePermitApi = (params: DraggablePermitParams) =>
  defHttp.post({ url: Api.DraggablePermit, params });
