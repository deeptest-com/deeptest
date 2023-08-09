<template>
  <div class="diagnose-tree-main">
    <div class="dp-tree-container">
      <div class="tree-filter">
        <a-input-search
            class="search-input"
            v-model:value="keywords"
            placeholder="输入关键字过滤"/>
        <div class="add-btn" @click="create(treeData?.[0]?.id, 'dir')">
          <PlusOutlined style="font-size: 16px;"/>
        </div>
      </div>

      <div class="tree-content">
        <a-tree
            class="deeptest-tree"
            draggable
            blockNode
            showIcon
            :expandedKeys="expandedKeys"
            :auto-expand-parent="autoExpandParent"
            :selectedKeys="selectedKeys"
            @drop="onDrop"
            @expand="onExpand"
            @select="selectNode"
            :tree-data="treeData?.[0]?.children"
            :replace-fields="replaceFields">

          <template #switcherIcon>
            <CaretDownOutlined/>
          </template>

          <template #title="nodeProps">
            <div class="tree-title" :title="nodeProps.dataRef.title" :draggable="nodeProps.dataRef.id === -1">
              <span class="tree-title-text" v-if="nodeProps.dataRef.title.indexOf(keywords) > -1">
                <span>{{nodeProps.dataRef.title.substr(0, nodeProps.dataRef.title.indexOf(keywords))}}</span>
                <span style="color: #f50">{{keywords}}</span>
                <span>{{nodeProps.dataRef.title.substr(nodeProps.dataRef.title.indexOf(keywords) + keywords.length)}}</span>
              </span>
              <span class="tree-title-text" v-else>{{ nodeProps.dataRef.title }}</span>

              <span class="more-icon" v-if="nodeProps.dataRef.id > 0">
                  <a-dropdown>
                       <MoreOutlined/>
                      <template #overlay>
                        <a-menu>
                          <a-menu-item v-if="nodeProps.dataRef.type === 'dir'" key="0" @click="create(nodeProps.dataRef.id, 'dir')">
                             新建目录
                          </a-menu-item>
                          <a-menu-item v-if="nodeProps.dataRef.type === 'dir'" key="1" @click="create(nodeProps.dataRef.id, 'interface')">
                             新建接口
                          </a-menu-item>
                          <a-menu-item v-if="nodeProps.dataRef.id !== -1" key="2" @click="edit(nodeProps)">
                           {{'编辑' + (nodeProps.dataRef.type === 'interface' ? '接口' : '目录')}}
                          </a-menu-item>
                          <a-menu-item v-if="nodeProps.dataRef.id !== -1" key="3" @click="deleteNode(nodeProps.dataRef)">
                            {{'删除' + (nodeProps.dataRef.type === 'interface' ? '接口' : '目录')}}
                          </a-menu-item>
                          <a-menu-item v-if="nodeProps.dataRef.type === 'dir'" key="4" @click="importInterfaces(nodeProps.dataRef)">
                             导入接口
                          </a-menu-item>
                          <a-menu-item v-if="nodeProps.dataRef.type === 'dir'" key="5" @click="importCurl(nodeProps.dataRef)">
                             导入cURL
                          </a-menu-item>
                        </a-menu>
                      </template>
                    </a-dropdown>
                </span>
            </div>
          </template>
        </a-tree>
        <div v-if="!treeData?.[0]?.children?.length" class="nodata-tip">请点击上方按钮添加目录 ~</div>
      </div>
    </div>

    <!--  编辑接口弹窗  -->
    <EditModal
        v-if="currentNode"
        :nodeInfo="currentNode"
        @ok="handleModalOk"
        @cancel="handleModalCancel" />

    <!--  导入接口弹窗  -->
    <InterfaceSelectionFromDefine
        v-if="interfaceSelectionVisible"
        :onFinish="interfaceSelectionFinish"
        :onCancel="interfaceSelectionCancel" />

    <!--  导入cURL弹窗  -->
    <CurlImportModal
        v-if="curlImportVisible"
        :visible="curlImportVisible"
        :onFinish="importCurlFinish"
        :onCancel="importCurlCancel" />

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
import {getSelectedKey, setExpandedKeys, setSelectedKey} from "@/utils/cache";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as DiagnoseInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";

import {expandOneKey} from "@/services/tree";
import EditModal from './edit.vue'
import {filterTree} from "@/utils/tree";
import {confirmToDelete} from "@/utils/confirm";
import debounce from "lodash.debounce";
import InterfaceSelectionFromDefine from "@/views/component/InterfaceSelectionFromDefine/main.vue";
import CurlImportModal from "./curl.vue";

const store = useStore<{ DiagnoseInterface: DiagnoseInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const treeData = computed<any>(() => store.state.DiagnoseInterface.treeData);
const treeDataMap = computed<any>(() => store.state.DiagnoseInterface.treeDataMap);

const props = defineProps({
  serveId: {
    required: false,
    type: Number || String,
  },
})

const keywords = ref('');
const replaceFields = {key: 'id'};
const expandedKeys = ref<number[]>([]);
const autoExpandParent = ref<boolean>(false);

async function loadTreeData() {
  if (currProject?.value?.id > 0 && currServe?.value?.id > 0) {
    await store.dispatch('DiagnoseInterface/loadTree', {projectId: currProject.value.id, serveId: currServe.value.id});
    expandAll();
  }
}

async function getServeServers() {
  await store.dispatch('DiagnoseInterface/getServeServers', {
    id: currServe.value.id,
  })
}

watch((currProject), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData();
  await getServeServers()
}, {immediate: true})
watch((currServe), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData();
  await getServeServers()
  selectStoredKeyCall()
}, {immediate: true})

watch(keywords, (newVal) => {
  expandedKeys.value = filterTree(treeData.value, newVal)
  autoExpandParent.value = true;
});

const getSelectedKeyName = () => {
  return `diagnose-interface-` + currServe.value.id
}
const selectStoredKeyCall = debounce(async () => {
  console.log('selectStoredKeyCall')
  let key = await getSelectedKey(getSelectedKeyName(), currProject.value.id)
  console.log('key', key)
  selectNode([key], null)
}, 300)

const onExpand = (keys: number[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

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

function selectNode(keys, e) {
  console.log('selectNode', keys, treeDataMap.value)

  if (keys.length === 0 && e) {
    selectedKeys.value = [e.node.dataRef.id] // un-select
    return
  } else {
    selectedKeys.value = keys
  }
  setSelectedKey(getSelectedKeyName(), currProject.value.id, selectedKeys.value[0])

  // if (e) {
    const selectedItem = treeDataMap.value[selectedKeys.value[0]]
    store.dispatch('DiagnoseInterface/openInterfaceTab', selectedItem);
  // }
}

const currentNode = ref(null as any);

function create(parentId, type) {
  console.log('create', parentId, type)
  currentNode.value = {parentId, type};
}
function edit(node) {
  currentNode.value = node;
}
async function deleteNode(node) {
  const title = node.type === 'dir' ?'将级联删除目录下的所有子目录、快捷调试' : ''
  const context = '删除后无法恢复，请确认是否删除？'

  confirmToDelete(title, context, () => {
    store.dispatch('DiagnoseInterface/removeInterface', {id: node.id, type: node.type});
  })
}

async function handleModalOk(model) {
  console.log('handleModalOk')
  Object.assign(model, {
    projectId: currProject.value.id,
    serveId: currServe.value.id,
  })

  const res = await store.dispatch('DiagnoseInterface/saveInterface', model)
  if (res) {
    currentNode.value = null
    expandOneKey(treeDataMap.value, model.parentId, expandedKeys.value)
  }
}

function handleModalCancel() {
  console.log('handleModalCancel')
  currentNode.value = null
}

// import interfaces
const importTarget = ref(null as any)
const interfaceSelectionVisible = ref(false)
const importInterfaces = (target) => {
  console.log('importInterfaces', target)

  importTarget.value = target
  interfaceSelectionVisible.value = true
}
const interfaceSelectionFinish = (interfaceIds) => {
  console.log('interfaceSelectionFinish', interfaceIds, importTarget.value)

  store.dispatch('DiagnoseInterface/importInterfaces', {
    interfaceIds: interfaceIds,
    targetId: importTarget.value.id,
  }).then((newNode) => {
    console.log('importInterfaces successfully', newNode)

    interfaceSelectionVisible.value = false
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })
}
function interfaceSelectionCancel() {
  console.log('handleModalCancel')
  interfaceSelectionVisible.value = false
}

// import cURL
const curlImportVisible = ref(false)
const importCurl = (target) => {
  console.log('importCurl', target)

  importTarget.value = target
  curlImportVisible.value = true
}
const importCurlFinish = (model) => {
  curlImportVisible.value = false
  console.log('importCurlFinish', model)

  store.dispatch('DiagnoseInterface/importCurl', {
    content: model.content,
    targetId: importTarget.value.id,
  }).then((newNode) => {
    console.log('importCurl successfully', newNode)
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })
}
const importCurlCancel = () => {
  console.log('importCurlCancel')
  curlImportVisible.value = false
}

async function onDrop(info: DropEvent) {
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const pos = info.node.pos.split('-');
  const dropPosition = info.dropPosition - Number(pos[pos.length - 1]);

  const res = await store.dispatch('DiagnoseInterface/moveInterface', {
    "dragKey": dragKey, // 移动谁
    "dropKey": dropKey,  // 移动那儿
    "dropPos": dropPosition // 0 表示移动到目标节点的子节点，-1 表示移动到目标节点的前面， 1表示移动到目标节点的后面
  });
  if (res) {
    // 移动到目标节点的子节点，则需要展开目标节点
    if (dropPosition === 0) {
      expandedKeys.value = [...new Set([...expandedKeys.value, dropKey])];
    }
    message.success('移动成功');
  } else {
    message.error('移动失败');
  }
}

onMounted(async () => {
  console.log('onMounted')
})

</script>

<style scoped lang="less">
.diagnose-tree-main {
  height: 100%;
  background: #ffffff;
}
</style>
