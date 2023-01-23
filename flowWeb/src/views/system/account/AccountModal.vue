<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm">
      <template #role="{ model, field }">
        <Select
          v-model:value="model[field]"
          show-search
          mode="multiple"
          placeholder="请输入角色"
          :default-active-first-option="false"
          :filter-option="false"
          :not-found-content="null"
          :options="roleSelectData"
          @search="handleRoleSearch"
        />
      </template>
    </BasicForm>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { accountFormSchema } from './account.data';
  import { listToTree } from '/@/utils/helper/treeHelper';
  import { Select } from 'ant-design-vue';
  import { keyWordApi } from '/@/api/sys/keyWord';
  import { KeyWordParams, KeyWordResult } from '/@/api/sys/model/keyWordModel';
  import { TreeItem } from '/@/components/Tree';
  import { addAccountApi, updateAccountApi } from '/@/api/system/account';
  import { AddAccountParams, UpdateAccountParams } from '/@/api/system/model/accountModel';

  export default defineComponent({
    name: 'AccountModal',
    components: { BasicModal, BasicForm, Select },
    emits: ['success', 'register'],
    setup(_, { emit }) {
      const isUpdate = ref(true);
      const roleSelectData = ref<{ value: string; label: string }[]>([]);
      const depthSelectData = ref<{ value: string; title: string; key: string }[]>([]);
      let timeout: any;

      const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
        labelWidth: 100,
        schemas: accountFormSchema,
        showActionButtonGroup: false,
        actionColOptions: {
          span: 23,
        },
      });

      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
        resetFields();
        setModalProps({ confirmLoading: false });
        isUpdate.value = !!data?.isUpdate;

        if (unref(isUpdate)) {
          console.log(data.record);
          setFieldsValue({
            ...data.record,
            initUserAccount: data.record.userAccount,
            userRoles: data.record.userRoleCodes ? data.record.userRoleCodes.split(',') : [],
          });
        }
        await getRole();
        await getDept();
        updateSchema({
          field: 'userDept',
          componentProps: { treeData: depthSelectData.value as any as TreeItem[] },
        });

        updateSchema({
          field: 'password',
          ifShow: !unref(isUpdate),
        });
      });

      const getTitle = computed(() => (!unref(isUpdate) ? '新增账号' : '编辑账号'));

      async function getRole(word = '') {
        const params: KeyWordParams = {
          category: 'role',
        };

        if (word) {
          params.word = word;
        }
        const roleSelectDataRes = await keyWordApi(params);
        roleSelectData.value = roleSelectDataRes.role.map((e: KeyWordResult) => {
          return {
            value: e.value,
            label: e.name,
          };
        });
      }

      async function getDept() {
        const params: KeyWordParams = {
          category: 'dept',
        };
        const deptTreeData = await keyWordApi(params);
        const deptTreeConfig = deptTreeData.dept.map((e: KeyWordResult) => {
          return {
            cId: e.value,
            id: e.value,
            pid: e.pValue,
            key: e.value,
            title: e.name,
            value: e.value,
          };
        });
        depthSelectData.value = listToTree(deptTreeConfig);
      }

      function handleRoleSearch(val: string) {
        if (timeout) {
          clearTimeout(timeout);
          timeout = null;
        }

        timeout = setTimeout(async () => {
          await getRole(val);
        }, 1000);
      }

      async function handleSubmit() {
        try {
          const values = await validate();
          setModalProps({ confirmLoading: true });

          if (unref(isUpdate)) {
            const params: UpdateAccountParams = {
              userAccount: values.userAccount,
              userName: values.userName,
              userCode: values.userCode,
              status: values.status,
              userRoles: values.userRoles,
              userDept: values.userDept,
              userMail: values.userMail,
            };
            await updateAccountApi(params);
          } else {
            const params: AddAccountParams = {
              userAccount: values.userAccount,
              userName: values.userName,
              password: values.password,
              status: values.status,
              userRoles: values.userRoles,
              userDept: values.userDept,
              userMail: values.userMail,
            };
            await addAccountApi(params);
          }
          // TODO custom api
          closeModal();
          emit('success');
        } finally {
          setModalProps({ confirmLoading: false });
        }
      }

      return {
        registerModal,
        registerForm,
        getTitle,
        handleSubmit,
        roleSelectData,
        handleRoleSearch,
      };
    },
  });
</script>
