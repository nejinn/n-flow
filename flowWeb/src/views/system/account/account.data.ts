import { Tag } from 'ant-design-vue';
import { h } from 'vue';
// import { keyWordApi } from '/@/api/sys/keyWord';
import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { checkUserAccountApi } from '/@/api/system/account';

export const columns: BasicColumn[] = [
  {
    title: '用户名',
    dataIndex: 'userAccount',
    width: 120,
    fixed: 'left',
  },
  {
    title: '昵称',
    dataIndex: 'userName',
    width: 120,
  },
  {
    title: '头像',
    dataIndex: 'userAvatar',
    width: 120,
  },
  {
    title: '邮箱',
    dataIndex: 'userMail',
    width: 120,
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
    title: '角色',
    dataIndex: 'userRoles',
  },
  {
    title: '部门',
    dataIndex: 'userDept',
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
    field: 'userAccount',
    label: '用户名',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'userName',
    label: '昵称',
    component: 'Input',
    colProps: { span: 8 },
  },
];

export const accountFormSchema: FormSchema[] = [
  {
    field: 'initUserAccount',
    label: '用户名',
    component: 'Input',
    show: false,
  },
  {
    field: 'userCode',
    label: '用户编码',
    component: 'Input',
    show: false,
  },
  {
    field: 'userAccount',
    label: '用户名',
    component: 'Input',
    colProps: { span: 24 },
    required: true,
    helpMessage: ['用户名可以用来作为登录账号'],
    dynamicRules: ({ values }) => {
      return [
        {
          required: true,
          message: '请输入用户名',
        },
        {
          validator(_, value) {
            return new Promise((resolve, reject) => {
              if (!values.initUserAccount || values.initUserAccount === value) {
                resolve();
              } else {
                checkUserAccountApi({ account: value })
                  .then((res) => {
                    if (res.code === 1) {
                      reject('账号已存在');
                    }
                    resolve();
                  })
                  .catch((err) => {
                    reject(err.message || '验证失败');
                  });
              }
            });
          },
        },
      ];
    },
  },
  {
    field: 'userName',
    label: '昵称',
    component: 'Input',
    required: true,
    colProps: { span: 24 },
  },
  {
    field: 'password',
    label: '密码',
    component: 'InputPassword',
    required: true,
    ifShow: false,
    colProps: { span: 24 },
  },
  {
    field: 'status',
    label: '状态',
    component: 'RadioButtonGroup',
    defaultValue: 1,
    required: true,
    componentProps: {
      options: [
        { label: '启用', value: 1 },
        { label: '停用', value: 2 },
      ],
    },
  },
  {
    label: '角色',
    field: 'userRoles',
    component: 'Input',
    slot: 'role',
    colProps: { span: 24 },
    required: true,
  },
  {
    field: 'userDept',
    label: '所属部门',
    // component: 'Input',
    component: 'TreeSelect',
    componentProps: {
      fieldNames: {
        label: 'title',
        key: 'key',
        value: 'value',
      },
      getPopupContainer: () => document.body,
    },
    required: true,
    colProps: { span: 24 },
    // slot: 'dept',
  },

  {
    label: '邮箱',
    field: 'userMail',
    component: 'Input',
    required: true,
    colProps: { span: 24 },
    dynamicRules: () => {
      return [
        {
          required: true,
          message: '请输入邮箱',
        },
        {
          validator(_, value) {
            if (!value) {
              return Promise.reject('邮箱不能为空');
            }
            const pwdRegExp = new RegExp(
              '^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$',
            );
            if (pwdRegExp.test(value) === false) {
              return Promise.reject('邮箱格式不对');
            }
            return Promise.resolve();
          },
        },
      ];
    },
  },
];
