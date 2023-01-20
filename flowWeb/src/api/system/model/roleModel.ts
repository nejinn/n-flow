import { BasicFetchResult } from '/@/api/model/baseModel';

export type GetRoleListParams = {
  roleName: string;
  status: string;
};

export type GetRoleListItem = {
  roleCode: string;
  roleDesc: string;
  roleName: string;
  status: number;
  roleOrder: number;
  cTime: string;
  uTime: string;
  cUser: string;
  uUser: string;
};

export type GetRoleListResultModel = BasicFetchResult<GetRoleListItem>;
