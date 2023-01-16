import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';

export const columns: BasicColumn[] = [
  {
    title: '部门名称',
    dataIndex: 'deptName',
    width: 160,
    align: 'left',
  },
  {
    title: '排序',
    dataIndex: 'deptOrder',
    width: 50,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 80,
    customRender: ({ record }) => {
      const status = record.status;
      const enable = ~~status === 1;
      const color = enable ? 'green' : 'red';
      const text = enable ? '启用' : '停用';
      return h(Tag, { color: color }, () => text);
    },
  },
  {
    title: '描述',
    dataIndex: 'deptDesc',
  },
  {
    title: '创建人',
    dataIndex: 'cUser',
    width: 180,
  },
  {
    title: '更新人',
    dataIndex: 'uUser',
    width: 180,
  },
  {
    title: '创建时间',
    dataIndex: 'cTime',
    width: 180,
  },
  {
    title: '更新时间',
    dataIndex: 'uTime',
    width: 180,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: '部门名称',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'status',
    label: '状态',
    component: 'Select',
    componentProps: {
      options: [
        { label: '启用', value: 1 },
        { label: '停用', value: 2 },
      ],
    },
    colProps: { span: 8 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'deptCode',
    label: '部门编码',
    component: 'Input',
    required: false,
    show() {
      return false;
    },
  },
  {
    field: 'deptName',
    label: '部门名称',
    component: 'Input',
    required: true,
  },
  {
    field: 'deptParent',
    label: '上级部门',
    component: 'TreeSelect',
    componentProps: {
      fieldNames: {
        label: 'deptName',
        key: 'deptCode',
        value: 'deptCode',
      },
      getPopupContainer: () => document.body,
    },
    required: true,
  },
  {
    field: 'deptOrder',
    label: '排序',
    defaultValue: 1,
    component: 'InputNumber',
    required: true,
    dynamicRules: () => {
      return [
        {
          required: true,
          validator: (_, value) => {
            if (value < 1) {
              return Promise.reject('排序必须大于0');
            }
            if (!value) {
              return Promise.reject('请输入排序数字');
            }
            return Promise.resolve();
          },
        },
      ];
    },
  },
  {
    field: 'status',
    label: '状态',
    component: 'RadioButtonGroup',
    defaultValue: 1,
    componentProps: {
      options: [
        { label: '启用', value: 1 },
        { label: '停用', value: 2 },
      ],
    },
    required: true,
  },
  {
    label: '描述',
    field: 'deptDesc',
    component: 'InputTextArea',
  },
];
