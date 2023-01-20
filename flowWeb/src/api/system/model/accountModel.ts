import { BasicFetchResult } from '/@/api/model/baseModel';

export type GetAccountListParams = {
  userAccount: string;
  userName: string;
  status: string;
};

export type GetAccountListItem = {
  userCode: string;
  userName: string;
  userAccount: string;
  userMail: string;
  userAvatar: string;
  userRoles: string;
  userDept: string;
  status: number;
  cTime: string;
  uTime: string;
};

export type GetAccountListResultModel = BasicFetchResult<GetAccountListItem>;
