import { defHttp } from '/@/utils/http/axios';
import { KeyWordParams, KeyWordResultModel } from '/@/api/sys/model/keyWordModel';
enum Api {
  KeyWord = '/admin/api/v1/keyWord',
}

export function keyWordApi(params: KeyWordParams) {
  return defHttp.get<KeyWordResultModel>({
    url: Api.KeyWord,
    params,
  });
}
