import {
  GetAccountListParams,
  GetAccountListResultModel,
  CheckUserAccountParams,
  AddAccountParams,
  UpdateAccountParams,
} from '/@/api/system/model/accountModel';
import { defHttp } from '/@/utils/http/axios';
enum Api {
  AccountList = '/admin/api/v1/getUserList',
  CheckUserAccount = '/admin/api/v1/checkUserAccount',
  AddAccount = '/admin/api/v1/addAccount',
  UpdateAccount = '/admin/api/v1/updateAccount',
}

export const getAccountListApi = (params?: GetAccountListParams) =>
  defHttp.get<GetAccountListResultModel>({ url: Api.AccountList, params });

export const checkUserAccountApi = (params: CheckUserAccountParams) =>
  defHttp.post<{ code: number }>({ url: Api.CheckUserAccount, params });

export const addAccountApi = (params: AddAccountParams) =>
  defHttp.post({ url: Api.AddAccount, params });

export const updateAccountApi = (params: UpdateAccountParams) =>
  defHttp.post({ url: Api.UpdateAccount, params });
