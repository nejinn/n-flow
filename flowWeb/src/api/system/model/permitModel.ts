export type GetPermitListItem = {
  permitCode: string;
  permitName: string;
  permitDesc: string;
  permitOrder: number;
  permitMethod: string;
  permitType: number;
  permitUrl: string;
  permitParent: string;
  id: number | string;
  cId: number | string;
  pid: string;
  key: string;
  title: string;
};

export type DraggablePermitParams = {
  code: string;
  pCode: string;
};

export type GetRolePermitParams = {
  code: string;
};
