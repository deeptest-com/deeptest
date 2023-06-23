<template>
  <div class="tree-main">
    <div class="tree-filters">
      <a-select :placeholder="'请选择服务'" :bordered="true"
          v-model:value="serveId"
          @change="selectServe">
        <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
      </a-select>
    </div>

    <div class="tree-container">
      <a-tree
          class="deeptest-tree"
          showIcon
          :checkable="true"
          :expandedKeys="expandedKeys"
          :auto-expand-parent="autoExpandParent"
          :tree-data="treeData"
          :replace-fields="replaceFields">

        <template #switcherIcon>
          <CaretDownOutlined/>
        </template>

        <template #title="nodeProps">
          <div class="tree-title" :draggable="nodeProps.dataRef.id === -1">
              <span class="tree-title-text" v-if="nodeProps.dataRef.title.indexOf(searchValue) > -1">
                <span>{{nodeProps.dataRef.title.substr(0, nodeProps.dataRef.title.indexOf(searchValue))}}</span>
                <span style="color: #f50">{{searchValue}}</span>
                <span>{{nodeProps.dataRef.title.substr(nodeProps.dataRef.title.indexOf(searchValue) + searchValue.length)}}</span>
              </span>
            <span class="tree-title-text" v-else>{{ nodeProps.dataRef.title }}</span>
          </div>
        </template>
      </a-tree>

      <div v-if="!treeData" class="nodata-tip">空</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  computed, ref, onMounted,
  watch, defineEmits, defineProps
} from 'vue';
import {
  PlusOutlined,
  CaretDownOutlined,
  MoreOutlined
} from '@ant-design/icons-vue';
import {message, Modal} from 'ant-design-vue';
import {DropEvent} from 'ant-design-vue/es/tree/Tree';
import {useStore} from "vuex";
import {setExpandedKeys, setSelectedKey} from "@/utils/cache";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as TestInterfaceStateType} from '@/views/debugger/store';
import {StateType as ServeStateType} from "@/store/serve";

import {expandOneKey} from "@/services/tree";
import {listServe} from "@/services/serve";

const store = useStore<{ TestInterface: TestInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const treeData = computed<any>(() => store.state.TestInterface.treeData);
const treeDataMap = computed<any>(() => store.state.TestInterface.treeDataMap);

const props = defineProps({
})

const serves = ref([] as any[]);
const serveId = ref(0)

const loadServe = async () => {
  listServe().then((json) => {
    serves.value = json.data.serves

    if (serves.value.length > 0) {
      serveId.value = serves.value[0].id
    }
  })
}
loadServe()

const replaceFields = {key: 'id'};
const searchValue = ref('');
const expandedKeys = ref<number[]>([]);
const autoExpandParent = ref<boolean>(false);

async function loadTreeData() {
  if (currProject?.value?.id > 0 && currServe?.value?.id > 0) {
    await store.dispatch('TestInterface/loadTree', {projectId: currProject.value.id, serveId: currServe.value.id});
    expandAll();
  }
}

async function getServeServers() {
  await store.dispatch('TestInterface/getServeServers', {
    id: currServe.value.id,
  })
}

const selectServe = () => {
  console.log('selectServe', serveId.value)
}

watch((currProject), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData();
  await getServeServers()
}, {
  immediate: true
})
watch((currServe), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData();
  await getServeServers()
}, {
  immediate: true
})

watch(searchValue, (newVal) => {
  if (!treeData.value) return
  // 打平树形结构
  function flattenTree(tree) {
    const nodes: Array<any> = [];

    function traverse(node) {
      nodes.push(node);
      if (node.children) {
        node.children.forEach(traverse);
      }
    }

    traverse(tree);
    return nodes;
  }

  const flattenTreeList = flattenTree(treeData.value[0]);

  function findParentIds(nodeId, tree) {
    let current: any = tree.find(node => node.id === nodeId);
    let parentIds: Array<string> = [];
    while (current && current.parentId) {
      parentIds.unshift(current.parentId); // unshift 方法可以将新元素添加到数组的开头
      current = tree.find(node => node.id === current.parentId);
    }
    return parentIds;
  }

  let parentKeys: any = [];
  for (let i = 0; i < flattenTreeList.length; i++) {
    let node = flattenTreeList[i];
    if (node.title.includes(newVal)) {
      parentKeys.push(node.parentId);
      parentKeys = parentKeys.concat(findParentIds(node.parentId, flattenTreeList));
    }
  }
  parentKeys = [...new Set(parentKeys)];
  expandedKeys.value = parentKeys;
  autoExpandParent.value = true;
});

// 展开所有
function expandAll() {
  const keys: any = [];
  const data = treeData.value;

  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    arr.forEach((item, index) => {
      keys.push(item.id);
      if (Array.isArray(item.children)) {
        fn(item.children)
      }
    });
  }
  fn(data);
  expandedKeys.value = keys;
}

let selectedKeys = ref<number[]>([]);
const emit = defineEmits(['select']);

const currentNode = ref(null as any);

function create(parentId, type) {
  console.log('create', parentId, type)
  currentNode.value = {parentId, type};
}
function edit(node) {
  currentNode.value = node;
}
async function deleteNode(node) {
  Modal.confirm({
    title: () => '确定删除该' + (node.type === 'interface'?'接口':'目录') + '吗？',
    content: () => node.type === 'dir'?'删除后所有所有子目录都会被删除':'',
    okText: () => '确定',
    okType: 'danger',
    cancelText: () => '取消',
    onOk: async () => {
      const res = await store.dispatch('TestInterface/removeInterface', {id: node.id, type: node.type});
      if (res) {
        message.success('删除成功');
      } else {
        message.error('删除失败');
      }
    },
    onCancel() {
      console.log('Cancel');
    },
  });
}

onMounted(async () => {
  console.log('onMounted')
})

</script>

<style scoped lang="less">
.tree-main {
  .tree-filters {
    margin-bottom: 16px;
  }
  .tree-container {
    background: #ffffff;
    .tree-title {
      position: relative;

      .tree-title-text {
        display: inline-block;
        white-space: nowrap;
      }

      .more-icon {
        position: absolute;
        right: -8px;
        width: 20px;
      }
    }

    .nodata-tip {
      margin-top: 8px;
      text-align: center;
    }
  }
}

</style>
