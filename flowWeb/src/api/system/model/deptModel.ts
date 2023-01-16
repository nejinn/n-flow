import { BasicFetchResult } from '/@/api/model/baseModel';

export type DeptParams = {
  deptName?: string;
  status?: string;
  top?: number;
};

export type addDeptParams = {
  deptName: string;
  deptDesc: string;
  deptOrder: number;
  deptParent: string;
  status: number;
};

export type updateDeptParams = {
  deptCode: string;
  deptName: string;
  deptDesc: string;
  deptOrder: number;
  deptParent: string;
  status: number;
};

export interface DeptListItem {
  id: string | number;
  deptCode: string;
  deptName: string;
  deptDesc: string;
  deptOrder: string;
  deptParent: string;
  status: number;
  cUser: string;
  uUser: string;
  cTime: string;
  uTime: string;
  cId: number | string;
  pid: string;
  key: string;
}

export type DeptListGetResultModel = BasicFetchResult<DeptListItem>;
