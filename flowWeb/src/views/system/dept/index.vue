<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate" v-if="hasPermission([addPermit])">
          新增部门
        </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                onClick: handleEdit.bind(null, record),
              },
              // {
              //   icon: 'ant-design:delete-outlined',
              //   color: 'error',
              //   popConfirm: {
              //     title: '是否确认删除',
              //     placement: 'left',
              //     confirm: handleDelete.bind(null, record),
              //   },
              // },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <DeptModal @register="registerModal" @success="handleSuccess" />
  </div>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getDeptListApi } from '/@/api/system/dept';
  import { useModal } from '/@/components/Modal';
  import DeptModal from './DeptModal.vue';
  import { columns, searchFormSchema } from './dept.data';
  import { DeptListItem } from '/@/api/demo/model/systemModel';
  import { listToTree } from '/@/utils/helper/treeHelper';
  import { TreeItem } from '/@/components/Tree/src/types/tree';
  import { RoleEnum } from '/@/enums/roleEnum';
  import { usePermission } from '/@/hooks/web/usePermission';
  export default defineComponent({
    name: 'DeptManagement',
    components: { BasicTable, DeptModal, TableAction },
    setup() {
      const listPermit = RoleEnum.ADMINSYSTEMDEPTGETDEPTLIST;
      const addPermit = RoleEnum.ADMINSYSTEMDEPTADDDEPT;
      const updatePermit = RoleEnum.ADMINSYSTEMDEPTUPDATEDEPT;
      const { hasPermission } = usePermission();

      const [registerModal, { openModal }] = useModal();
      const [registerTable, { reload }] = useTable({
        title: '部门列表',
        api: getDeptListApi,
        columns,
        formConfig: {
          labelWidth: 120,
          schemas: searchFormSchema,
          showResetButton: hasPermission([listPermit]),
          showSubmitButton: hasPermission([listPermit]),
        },
        tableSetting: {
          redo: hasPermission([listPermit]),
          size: true,
          setting: true,
          fullScreen: true,
        },
        pagination: false,
        striped: false,
        useSearchForm: true,
        showTableSetting: true,
        bordered: true,
        showIndexColumn: false,
        canResize: false,
        actionColumn: {
          width: 80,
          title: '操作',
          dataIndex: 'action',
          // slots: { customRender: 'action' },
          fixed: 'right',
          ifShow: () => {
            return hasPermission([updatePermit]);
          },
        },
        afterFetch(res: DeptListItem[]) {
          const TreeConfig = res.map((e: DeptListItem) => {
            e.cId = e.id;
            e.id = e.deptCode;
            e.pid = e.deptParent;
            e.key = e.deptCode;
            return e;
          });
          return listToTree(TreeConfig) as any as TreeItem[];
        },
      });
      function handleCreate() {
        openModal(true, {
          isUpdate: false,
        });
      }
      function handleEdit(record: Recordable) {
        openModal(true, {
          record,
          isUpdate: true,
        });
      }
      function handleDelete(record: Recordable) {
        console.log(record);
      }
      function handleSuccess() {
        reload();
      }
      return {
        registerTable,
        registerModal,
        handleCreate,
        handleEdit,
        handleDelete,
        handleSuccess,
        hasPermission,
        addPermit,
      };
    },
  });
</script>
