<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './dept.data';
  import { addDeptApi, getDeptListApi, updateDeptApi } from '/@/api/system/dept';
  import {
    DeptParams,
    DeptListItem,
    addDeptParams,
    updateDeptParams,
  } from '/@/api/system/model/deptModel';
  import { listToTreeRemove, listToTree } from '/@/utils/helper/treeHelper';
  import { TreeItem } from '/@/components/Tree/src/types/tree';
  export default defineComponent({
    name: 'DeptModal',
    components: { BasicModal, BasicForm },
    emits: ['success', 'register'],
    setup(_, { emit }) {
      const isUpdate = ref(true);
      const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
        labelWidth: 100,
        baseColProps: { span: 24 },
        schemas: formSchema,
        showActionButtonGroup: false,
      });
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
        resetFields();
        setModalProps({ confirmLoading: false });
        isUpdate.value = !!data?.isUpdate;
        if (unref(isUpdate)) {
          setFieldsValue({
            ...data.record,
          });
        }
        const params: DeptParams = {
          top: 1,
        };
        const treeData = await getDeptListApi(params);
        // console.log(treeData);
        const TreeConfig = (treeData as unknown as DeptListItem[]).map((e: DeptListItem) => {
          e.cId = e.id;
          e.id = e.deptCode;
          e.pid = e.deptParent;
          e.key = e.deptCode;
          return e;
        });

        let treeDataConvert = [];
        if (unref(isUpdate)) {
          treeDataConvert = listToTreeRemove(TreeConfig, {}, data.record.key);
        } else {
          treeDataConvert = listToTree(TreeConfig);
        }

        updateSchema({
          field: 'deptParent',
          componentProps: { treeData: treeDataConvert as any as TreeItem[] },
        });
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '????????????' : '????????????'));
      async function handleSubmit() {
        try {
          const values = await validate();
          setModalProps({ confirmLoading: true });
          // TODO custom api
          if (unref(isUpdate)) {
            const params: updateDeptParams = {
              ...values,
            };
            await updateDeptApi(params);
          } else {
            const params: addDeptParams = {
              ...values,
            };
            await addDeptApi(params);
          }
          closeModal();
          emit('success');
        } finally {
          setModalProps({ confirmLoading: false });
        }
      }
      return { registerModal, registerForm, getTitle, handleSubmit };
    },
  });
</script>
