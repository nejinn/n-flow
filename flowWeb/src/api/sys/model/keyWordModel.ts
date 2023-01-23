export type KeyWordParams = {
  category: string;
  word?: string;
};

export type KeyWordResult = {
  name: string;
  value: string;
  pValue?: string;
};

export type KeyWordResultModel = {
  [key: string]: KeyWordResult[];
};
