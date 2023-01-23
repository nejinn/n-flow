import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Switch } from 'ant-design-vue';
import { SwitchRoleStatusApi } from '/@/api/system/role';
import { useMessage } from '/@/hooks/web/useMessage';
import { usePermission } from '/@/hooks/web/usePermission';
import { RoleEnum } from '/@/enums/roleEnum';
const { hasPermission } = usePermission();
export const columns: BasicColumn[] = [
  {
    title: '角色名称',
    dataIndex: 'roleName',
    width: 200,
    fixed: 'left',
  },
  {
    title: '角色值',
    dataIndex: 'roleCode',
  },
  {
    title: '排序',
    dataIndex: 'roleOrder',
    width: 50,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 120,
    customRender: ({ record }) => {
      if (!Reflect.has(record, 'pendingStatus')) {
        record.pendingStatus = false;
      }
      return h(Switch, {
        checked: record.status === 1,
        checkedChildren: '已启用',
        unCheckedChildren: '已禁用',
        loading: record.pendingStatus,
        disabled: !hasPermission([RoleEnum.ADMINSYSTEMYROLESWITCHROLESTATUS]),
        onChange(checked: boolean) {
          record.pendingStatus = true;
          const newStatus = checked ? 1 : 2;
          const { createMessage } = useMessage();
          SwitchRoleStatusApi(record.roleCode, newStatus)
            .then(() => {
              record.status = newStatus;
              createMessage.success(`已成功修改角色状态`);
            })
            .catch(() => {
              createMessage.error('修改角色状态失败');
            })
            .finally(() => {
              record.pendingStatus = false;
            });
        },
      });
    },
  },
  {
    title: '描述',
    dataIndex: 'roleDesc',
    width: 300,
  },
  {
    title: '创建人',
    dataIndex: 'cUser',
  },
  {
    title: '更新人',
    dataIndex: 'uUser',
  },
  {
    title: '创建时间',
    dataIndex: 'cTime',
    width: 180,
  },
  {
    title: '更新时间',
    dataIndex: 'uTime',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'roleName',
    label: '角色名称',
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
    field: 'roleName',
    label: '角色名称',
    required: true,
    component: 'Input',
  },
  {
    field: 'roleCode',
    label: '角色值',
    component: 'Input',
    show() {
      return false;
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
  },
  {
    label: '排序',
    field: 'roleOrder',
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
    label: '描述',
    field: 'roleDesc',
    required: true,
    component: 'InputTextArea',
  },
  {
    label: ' ',
    field: 'permit',
    slot: 'permit',
    required: true,
    component: 'Input',
    dynamicRules: () => {
      return [
        {
          required: true,
          validator: (_, value) => {
            if (!value || value.length === 0) {
              return Promise.reject('权限不能为空');
            }
            return Promise.resolve();
          },
        },
      ];
    },
  },
  {
    label: '',
    field: 'isUpdate',
    component: 'Input',
    ifShow: () => {
      return false;
    },
  },
];
