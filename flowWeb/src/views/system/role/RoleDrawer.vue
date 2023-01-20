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
          v-if="treeData.length > 0"
          :treeData="treeData"
          :fieldNames="{ title: 'title', key: 'key' }"
          :checked-keys="checkedKeys"
          checkable
          title="权限分配"
          @check="handleCheck"
          defaultExpandAll
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
  import { getPermitListApi, getRolePermitApi } from '/@/api/system/permit';
  import { GetPermitListItem, GetRolePermitParams } from '/@/api/system/model/permitModel';
  import { listToTree } from '/@/utils/helper/treeHelper';
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
        // if (unref(treeData).length === 0) {
        //   treeData.value = (await getMenuList()) as any as TreeItem[];
        // }
        // isUpdate.value = !!data?.isUpdate;
        // if (unref(isUpdate)) {
        //   setFieldsValue({
        //     ...data.record,
        //   });
        // }
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
            setFieldsValue({
              ...data.record,
              isUpdate: isUpdate.value,
            });
            currentCode.value = data.record.roleCode;
            const params: GetRolePermitParams = {
              code: data.record.roleCode,
            };
            const res = await getRolePermitApi(params);
            console.log(res);
            treeCheckedKey.value = res.code || [];
            getLastKey(treeData.value);
            setFieldsValue({ permit: checkedKeys.value });
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
          // TODO custom api
          console.log(values);
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
