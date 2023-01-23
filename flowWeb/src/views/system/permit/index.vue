<template>
  <div class="bg-white m-4 overflow-hidden">
    <div class="p-2">
      <BasicTree
        toolbar
        search
        checkStrictly
        showLine
        clickRowToExpand
        :tree-data="treeData"
        draggable
        :loading="treeLoading"
        @drop="onDrop"
      >
        <template #title="item">
          <Tag color="pink"> {{ item.permitName }}</Tag>
          <Tag color="purple"> {{ item.permitCode }}</Tag>
          <Tag :color="item.permitType === 1 ? 'red' : 'blue'">
            {{ item.permitType === 1 ? '菜单' : '功能' }}</Tag
          >
          <template v-if="item.permitUrl">
            <Tag :color="item.permitMethod === 'POST' ? 'green' : 'cyan'">
              {{ item.permitMethod }}
            </Tag>
            <Tag :color="item.permitMethod === 'POST' ? 'green' : 'cyan'">
              {{ item.permitUrl }}</Tag
            >
          </template>
          <Tag> {{ item.permitDesc }}</Tag>
        </template>
      </BasicTree>
    </div>
  </div>
</template>
<script lang="ts">
  import { defineComponent, onMounted, ref } from 'vue';
  import { Tag } from 'ant-design-vue';

  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { listToTree } from '/@/utils/helper/treeHelper';
  import { getPermitListApi, draggablePermitApi } from '/@/api/system/permit';
  import { DraggablePermitParams, GetPermitListItem } from '/@/api/system/model/permitModel';
  import { useMessage } from '/@/hooks/web/useMessage';
  export default defineComponent({
    name: 'PermitManagement',
    components: { BasicTree, Tag },
    setup() {
      const treeData = ref<TreeItem[]>([]);
      const treeLoading = ref(true);
      const { createMessage } = useMessage();

      async function fetch() {
        treeData.value = [];
        try {
          const res = (await getPermitListApi()) as unknown as TreeItem[];
          const TreeConfig = res.map((e: GetPermitListItem) => {
            e.cId = e.permitCode;
            e.id = e.permitCode;
            e.pid = e.permitParent;
            e.key = e.permitCode;
            e.title = e.permitName;
            return e;
          });

          treeData.value = listToTree(TreeConfig) as any as TreeItem[];
        } finally {
          treeLoading.value = false;
        }
      }

      onMounted(() => {
        fetch();
      });

      async function onDrop(e: {
        dragNode: { eventKey: any; permitParent: any };
        node: { permitParent: any; eventKey: any };
        dropToGap: any;
      }) {
        if (
          e.dragNode.permitParent === e.node.eventKey ||
          (e.dropToGap && e.dragNode.permitParent === e.node.permitParent)
        ) {
          createMessage.error('同一级拖拽无效');
          return;
        }

        treeLoading.value = true;
        const params: DraggablePermitParams = {
          code: e.dragNode.eventKey,
          pCode: e.dropToGap ? '' : e.node.eventKey,
        };
        try {
          await draggablePermitApi(params);
          fetch();
        } finally {
          treeLoading.value = false;
        }
      }

      return { treeData, onDrop, treeLoading };
    },
  });
</script>
