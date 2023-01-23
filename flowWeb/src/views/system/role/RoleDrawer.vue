<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="getTitle"
    width="500px"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #permit>
        <BasicTree
          toolbar
          :treeData="treeData"
          :fieldNames="{ title: 'title', key: 'key' }"
          :checked-keys="checkedKeys"
          checkable
          title="权限分配"
          @check="handleCheck"
        />
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './role.data';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { getPermitListApi } from '/@/api/system/permit';
  import { GetPermitListItem } from '/@/api/system/model/permitModel';
  import { listToTree } from '/@/utils/helper/treeHelper';
  import {
    AddRoleParams,
    GetRolePermitParams,
    UpdateRoleParams,
  } from '/@/api/system/model/roleModel';
  import { addRoleApi, getRolePermitApi, updateRoleApi } from '/@/api/system/role';
  export default defineComponent({
    name: 'RoleDrawer',
    components: { BasicDrawer, BasicForm, BasicTree },
    emits: ['success', 'register'],
    setup(_, { emit }) {
      const isUpdate = ref(true);
      const treeData = ref<TreeItem[]>([]);
      const checkedKeys = ref<Array<string>>([]);
      const treeCheckedKey = ref<Array<string>>([]);
      const currentCode = ref('');

      const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
        labelWidth: 90,
        baseColProps: { span: 24 },
        schemas: formSchema,
        showActionButtonGroup: false,
      });
      const [registerDrawer, { setDrawerProps, closeDrawer }] = useDrawerInner(async (data) => {
        checkedKeys.value = [];
        treeCheckedKey.value = [];
        resetFields();
        setDrawerProps({ confirmLoading: false, loading: true });
        // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
        try {
          let res = (await getPermitListApi()) as unknown as GetPermitListItem[];
          const permitTreeConfig = res.map((e: GetPermitListItem) => {
            e.id = e.permitCode;
            e.pid = e.permitParent;
            e.key = e.permitCode;
            e.title = e.permitName;
            return e;
          });
          treeData.value = listToTree(permitTreeConfig) as any as TreeItem[];
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            currentCode.value = data.record.roleCode;
            const params: GetRolePermitParams = {
              code: data.record.roleCode,
            };
            const res = await getRolePermitApi(params);
            treeCheckedKey.value = res.code || [];
            getLastKey(treeData.value);
            setFieldsValue({
              ...data.record,
              isUpdate: isUpdate.value,
              permit: checkedKeys.value,
            });
          }
        } finally {
          setDrawerProps({ loading: false });
        }
      });

      const getTitle = computed(() => (!unref(isUpdate) ? '新增角色' : '编辑角色'));

      function getLastKey(data: TreeItem[]) {
        data.forEach((element) => {
          if (element.children && element.children.length > 0) {
            getLastKey(element.children);
          } else {
            if (treeCheckedKey.value.indexOf(element.key as string) !== -1) {
              checkedKeys.value.push(element.key as string);
            }
          }
        });
      }

      function handleCheck(e: Array<string>, f: any) {
        const permits = [...e, ...f.halfCheckedKeys];
        setFieldsValue({ permit: permits });
      }

      async function handleSubmit() {
        try {
          const values = await validate();
          setDrawerProps({ confirmLoading: true });
          if (unref(isUpdate)) {
            const params: UpdateRoleParams = {
              roleCode: values.roleCode,
              roleName: values.roleName,
              roleDesc: values.roleDesc,
              roleOrder: values.roleOrder,
              status: values.status,
              permit: values.permit,
            };
            await updateRoleApi(params);
          } else {
            const params: AddRoleParams = {
              roleName: values.roleName,
              roleDesc: values.roleDesc,
              roleOrder: values.roleOrder,
              status: values.status,
              permit: values.permit,
            };
            await addRoleApi(params);
          }
          treeData.value = [];
          checkedKeys.value = [];
          closeDrawer();
          emit('success');
        } finally {
          setDrawerProps({ confirmLoading: false });
        }
      }

      return {
        registerDrawer,
        registerForm,
        getTitle,
        handleSubmit,
        treeData,
        checkedKeys,
        handleCheck,
      };
    },
  });
</script>
