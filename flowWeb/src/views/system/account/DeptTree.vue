<template>
  <div class="bg-white m-4 mr-0 overflow-hidden">
    <div class="p-2">
      <BasicTree
        title="部门列表"
        toolbar
        search
        checkStrictly
        showLine
        clickRowToExpand
        :treeData="treeData"
        @select="handleSelect"
      />
    </div>
  </div>
</template>
<script lang="ts">
  import { defineComponent, onMounted, ref } from 'vue';

  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { getDeptListApi } from '/@/api/system/dept';
  import { DeptListItem } from '/@/api/system/model/deptModel';
  import { eachTree, listToTree } from '/@/utils/helper/treeHelper';

  export default defineComponent({
    name: 'DeptTree',
    components: { BasicTree },

    emits: ['select'],
    setup(_, { emit }) {
      const treeData = ref<TreeItem[]>([]);

      async function fetch() {
        treeData.value;
        const res = (await getDeptListApi()) as unknown as TreeItem[];
        const TreeConfig = res.map((e: DeptListItem) => {
          e.cId = e.id;
          e.id = e.deptCode;
          e.pid = e.deptParent;
          e.key = e.deptCode;
          e.title = e.deptName;
          return e;
        });
        treeData.value = listToTree(TreeConfig) as any as TreeItem[];
      }

      function handleSelect(_, f) {
        const keys: Array<string> = [];
        const getKey = (element: any) => {
          keys.push(element.key);
        };
        eachTree(f.selectedNodes, getKey);
        emit('select', keys.join(','));
      }

      onMounted(() => {
        fetch();
      });
      return { treeData, handleSelect };
    },
  });
</script>
