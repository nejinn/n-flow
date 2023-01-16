import { defHttp } from '/@/utils/http/axios';
import {
  DeptParams,
  DeptListGetResultModel,
  addDeptParams,
  updateDeptParams,
} from '/@/api/system/model/deptModel';
enum Api {
  DeptList = '/admin/api/v1/getDeptList',
  addDept = '/admin/api/v1/addDept',
  updateDept = '/admin/api/v1/updateDept',
}

export const getDeptListApi = (params?: DeptParams) =>
  defHttp.get<DeptListGetResultModel>({ url: Api.DeptList, params });

export const addDeptApi = (params?: addDeptParams) => defHttp.post({ url: Api.addDept, params });

export const updateDeptApi = (params?: updateDeptParams) =>
  defHttp.post({ url: Api.updateDept, params });
