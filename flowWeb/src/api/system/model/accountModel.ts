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
  userRoleCodes: string;
  userDept: string;
  status: number;
  cTime: string;
  uTime: string;
};

export type GetAccountListResultModel = BasicFetchResult<GetAccountListItem>;

export type CheckUserAccountParams = {
  account: string;
};

export type AddAccountParams = {
  userAccount: string;
  userName: string;
  password: string;
  status: number;
  userRoles: string[];
  userDept: string;
  userMail: string;
};

export type UpdateAccountParams = {
  userAccount: string;
  userName: string;
  status: number;
  userRoles: string[];
  userDept: string;
  userMail: string;
  userCode: string;
};
