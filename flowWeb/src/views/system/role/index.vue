<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate" v-if="hasPermission([addPermit])">
          新增角色
        </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                onClick: handleEdit.bind(null, record),
                auth: [updatePermit],
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
    <RoleDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getRoleListApi } from '/@/api/system/role';
  import { useDrawer } from '/@/components/Drawer';
  import RoleDrawer from './RoleDrawer.vue';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { columns, searchFormSchema } from './role.data';
  import { RoleEnum } from '/@/enums/roleEnum';
  export default defineComponent({
    name: 'RoleManagement',
    components: { BasicTable, RoleDrawer, TableAction },
    setup() {
      const listPermit = RoleEnum.ADMINSYSTEMROLEGETROLELIST;
      const addPermit = RoleEnum.ADMINSYSTEMROLEADDROLE;
      const updatePermit = RoleEnum.ADMINSYSTEMROLEUPDATEROLE;
      const { hasPermission } = usePermission();

      const [registerDrawer, { openDrawer }] = useDrawer();
      const [registerTable, { reload }] = useTable({
        title: '角色列表',
        api: getRoleListApi,
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
        useSearchForm: true,
        showTableSetting: true,
        bordered: true,
        showIndexColumn: false,
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
      });
      function handleCreate() {
        openDrawer(true, {
          isUpdate: false,
        });
      }
      function handleEdit(record: Recordable) {
        openDrawer(true, {
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
        registerDrawer,
        handleCreate,
        handleEdit,
        handleDelete,
        handleSuccess,
        addPermit,
        updatePermit,
        hasPermission,
      };
    },
  });
</script>
