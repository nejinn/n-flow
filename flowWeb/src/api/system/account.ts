import { GetAccountListParams, GetAccountListResultModel } from '/@/api/system/model/accountModel';
import { defHttp } from '/@/utils/http/axios';
enum Api {
  AccountList = '/admin/api/v1/getUserList',
}

export const getAccountListApi = (params?: GetAccountListParams) =>
  defHttp.get<GetAccountListResultModel>({ url: Api.AccountList, params });
