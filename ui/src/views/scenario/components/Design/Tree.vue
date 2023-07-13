<template>
  <div class="scenario-tree-main">
    <div class="tree-container">
      <div class="tree-filter">
<!--        测试场景编排-->
        <a-input-search placeholder="输入关键字过滤"
                        class="search-input"
                        v-model:value="keywords" />
        <PlusOutlined class="plus-icon"/>
      </div>

      <div style="margin: 0 8px;">
        <a-tree
            class="deeptest-tree"
            draggable
            blockNode
            showIcon
            :expandAction="false"
            :expandedKeys="expandedKeys"
            :auto-expand-parent="autoExpandParent"
            v-model:selectedKeys="selectedKeys"
            @drop="onDrop"
            @expand="onExpand"
            @select="selectNode"
            :tree-data="treeData"
            :replace-fields="replaceFields">
          <template #switcherIcon>
            <CaretDownOutlined/>
          </template>
          <template #title="nodeProps">
            <div class="tree-title" :draggable="nodeProps.dataRef.id === -1">
              <span class="tree-title-text" v-if="nodeProps.dataRef.name.indexOf(keywords) > -1">
                <span>{{nodeProps.dataRef.name.substr(0, nodeProps.dataRef.name.indexOf(keywords))}}</span>
                <span style="color: #f50">{{keywords}}</span>
                <span>{{nodeProps.dataRef.name.substr(nodeProps.dataRef.name.indexOf(keywords) + keywords.length)}}</span>
              </span>
              <span class="tree-title-text" v-else>{{ nodeProps.dataRef.name }}</span>
              <span class="more-icon" v-if="nodeProps.dataRef.id > 0">
                  <a-dropdown>
                       <MoreOutlined/>
                      <template #overlay>
                        <TreeContextMenu :treeNode="nodeProps.dataRef" :onMenuClick="menuClick"/>
                      </template>
                    </a-dropdown>
                </span>
            </div>
          </template>
        </a-tree>
        <div v-if="!treeData" class="nodata-tip">请点击上方按钮添加分类 ~</div>
      </div>
    </div>

    <!--  编辑接口弹窗  -->
    <EditModal
        v-if="currentNode"
        :nodeInfo="currentNode"
        @ok="handleEditModalOk"
        @cancel="handleEditModalCancel" />

    <InterfaceSelectionFromDefine
        v-if="interfaceSelectionVisible && interfaceSelectionSrc==='fromDefine'"
        :onFinish="endpointInterfaceIdsSelectionFinish"
        :onCancel="interfaceSelectionCancel" />

    <InterfaceSelectionFromTest
        v-if="interfaceSelectionVisible && interfaceSelectionSrc==='fromTest'"
        :onFinish="diagnoseInterfaceNodesSelectionFinish"
        :onCancel="interfaceSelectionCancel" />

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch, getCurrentInstance} from "vue";

import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {PlusOutlined, CaretDownOutlined, MoreOutlined,} from '@ant-design/icons-vue';

import {expandAllKeys, expandOneKey} from "@/services/tree";

import {getExpandedKeys, getSelectedKey, setExpandedKeys} from "@/utils/cache";
import {StateType as ScenarioStateType} from "../../store";
import {isRoot, updateNodeName, isInterface} from "../../service";
import TreeContextMenu from "./components/TreeContextMenu.vue";
import EditModal from "./components/edit.vue";
import InterfaceSelectionFromDefine from "@/views/component/InterfaceSelectionFromDefine/main.vue";
import InterfaceSelectionFromTest from "@/views/component/InterfaceSelectionFromTest/main.vue";

const props = defineProps<{}>()

const useForm = Form.useForm;

const {t} = useI18n();
import {Scenario} from "@/views/scenario/data";
import {confirmToDelete} from "@/utils/confirm";
import {filterTree} from "@/utils/tree";
const store = useStore<{ Scenario: ScenarioStateType; }>();
const treeData = computed<any>(() => store.state.Scenario.treeData);
const treeDataMap = computed<any>(() => store.state.Scenario.treeDataMap);
const selectedNode = computed<any>(()=> store.state.Scenario.nodeData);
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);

watch(treeData, () => {
  console.log('watch', treeData)

  if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
    tips.value = '右键树状节点操作'
  }

  getExpandedKeysCall()
})

watch(() => {
  return detailResult.value.id
},async (newVal) => {
  if(newVal){
    // loadTree();
    await store.dispatch('Scenario/loadScenario', newVal);
  }
},{
  immediate:true
})

const keywords = ref('');
watch(keywords, (newVal) => {
  console.log('watch keywords', keywords)
  expandedKeys.value = filterTree(treeData.value, newVal)
  autoExpandParent.value = true;
});

const replaceFields = {key: 'id', title: 'name'};

let expandedKeys = ref<number[]>([]);
const autoExpandParent = ref<boolean>(false);

let selectedKeys = ref<number[]>([]);
let isExpand = ref(false);

let tree = ref(null)
const expandNode = (keys: string[], e: any) => {
  console.log('expandNode', keys[0], e)

  setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
}

const selectNode = (keys, e) => {
  console.log('selectNode', keys, e?.node.dataRef.id)

  if (keys.length === 0 && e) {
    selectedKeys.value = [e.node.dataRef.id] // cancel un-select
    return
  } else {
    selectedKeys.value = keys
  }

  const selectedData = treeDataMap.value[selectedKeys.value[0]]
  if (selectedData && isRoot(selectedData.entityCategory)) {
    store.dispatch('Scenario/getNode', null)
    return
  }

  store.dispatch('Scenario/getNode', selectedData).then((ok) => {
    if (ok && selectedNode.value.processorType === 'processor_interface_default') {
      // will cause watch event to load debug data in components/interface/interface.vue
      store.dispatch('Scenario/setScenarioProcessorIdForDebug', selectedNode.value.processorID)
      // store.dispatch('Scenario/setEndpointInterfaceIdForDebug', selectedNode.value.endpointInterfaceId)
    }
  })
}

const onExpand = (keys: number[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

const currentNode = ref(null as any);
function create(parentId, type) {
  console.log('create', parentId, type)
  currentNode.value = {parentId, type};
}
function edit(node) {
  console.log('edit', node)
  currentNode.value = node;
}

let contextNode = ref({} as any)
let tips = ref('')

const getExpandedKeysCall = debounce(async () => {
  getExpandedKeys('scenario', treeData.value[0].scenarioId).then(async keys => {
    console.log('keys', keys)
    if (keys)
      expandedKeys.value = keys

    if (!expandedKeys.value || expandedKeys.value.length === 0) {
      getOpenKeys(treeData.value[0], false) // expend first level folder
      console.log('expandedKeys.value', expandedKeys.value)
      await setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
    }
  })
}, 500)

const getOpenKeys = (treeNode, isAll) => {
  if (!treeNode) return

  expandedKeys.value.push(treeNode.id)
  if (treeNode.children && isAll) {
    treeNode.children.forEach((item, index) => {
      getOpenKeys(item, isAll)
    })
  }
}
getOpenKeys(treeData.value[0], false)
console.log('expandedKeys.value', expandedKeys.value)

const expandAll = () => {
  console.log('expandAll')
  isExpand.value = !isExpand.value
  expandedKeys.value = expandAllKeys(treeDataMap.value, isExpand.value)

  setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
}

let targetModelId = 0
const menuClick = (menuKey: string, targetId: number) => {
  console.log('menuClick', menuKey, targetId)
  targetModelId = targetId

  if (menuKey === 'edit') {
    edit(treeDataMap.value[targetModelId])
    return
  }

  if (menuKey === 'remove') {
    removeNode()
    return
  }

  // add-child-interface-define
  // add-child-interface-debug
  // add-child-processor_logic-processor_logic_if
  const arr = menuKey.split('-')
  const mode = arr[1]
  const processorCategory = arr[2]
  const processorType = arr[3]

  const targetProcessorId = targetModelId
  const targetProcessorCategory = treeDataMap.value[targetModelId].entityCategory
  const targetProcessorType = treeDataMap.value[targetModelId].entityType

  addNode(mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  clearMenu()
}

async function handleEditModalOk(model) {
  console.log('handleEditModalOk')
  Object.assign(model, {
    // projectId: currProject.value.id,
    // serveId: currServe.value.id,
  })

  const res = await store.dispatch('Scenario/saveProcessorInfo', model)
  if (res) {
    currentNode.value = null
    expandOneKey(treeDataMap.value, model.parentId, expandedKeys.value)
  }
}

function handleEditModalCancel() {
  console.log('handleEditModalCancel')
  currentNode.value = null
}

const addNode = (mode, processorCategory, processorType,
                 targetProcessorCategory, targetProcessorType, targetProcessorId) => {
  console.log('addNode', mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  if (processorCategory === 'interface') { // show popup to select a interface
    interfaceSelectionSrc.value = processorType
    interfaceSelectionVisible.value = true
    return

  } else {
    store.dispatch('Scenario/addProcessor',
        {mode, processorCategory, processorType,
          targetProcessorCategory, targetProcessorType, targetProcessorId,
          name: t(processorType)
        }).then((newNode) => {
          console.log('addProcessor successfully', newNode)
          selectNode([newNode.id], null)
          expandOneKey(treeDataMap.value, mode === 'parent' ? newNode.id : newNode.parentId, expandedKeys.value) // expend new node
          setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
        })
  }
}

const interfaceSelectionVisible = ref(false)
const interfaceSelectionSrc = ref('')

const endpointInterfaceIdsSelectionFinish = (interfaceIds) => {
  const targetNode = treeDataMap.value[targetModelId]
  console.log('endpointInterfaceIdsSelectionFinish', interfaceIds, targetNode)

  store.dispatch('Scenario/addInterfacesFromDefine', {
    interfaceIds: interfaceIds,
    targetId: targetNode.id,
  }).then((newNode) => {
    console.log('addInterfaces successfully', newNode)

    interfaceSelectionVisible.value = false
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })
}

const diagnoseInterfaceNodesSelectionFinish = (interfaceNodes) => {
  const targetNode = treeDataMap.value[targetModelId]
  console.log('endpointInterfaceIdsSelectionFinish', interfaceNodes, targetNode)

  store.dispatch('Scenario/addInterfacesFromTest', {
    selectedNodes: interfaceNodes,
    targetId: targetNode.id,
  }).then((newNode) => {
    console.log('addInterfaces successfully', newNode)

    interfaceSelectionVisible.value = false
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })
}

const interfaceSelectionCancel = () => {
  console.log('interfaceSelectionCancel')
  interfaceSelectionVisible.value = false
}

const removeNode = () => {
  console.log('removeNode')

  const node = treeDataMap.value[targetModelId]

  const title = `确定删除名为${node.name}的节点吗？`
  const context = '该节点的所有子节点都将被删除！'

  confirmToDelete(title, context, () => {
    store.dispatch('Scenario/removeNode', targetModelId);
    selectNode([], null)
  })
}

const clearMenu = () => {
  console.log('clearMenu')
  contextNode.value = ref({})
}

async function onDrop(info: DropEvent) {
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const pos = info.node.pos.split('-');
  let dropPosition = info.dropPosition - Number(pos[pos.length - 1]);

  if (isInterface(treeDataMap.value[dropKey].processorCategory) && dropPosition === 0) dropPosition = 1
  console.log(dragKey, dropKey, dropPosition);

  store.dispatch('Scenario/moveNode', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPosition}).then(
      (result) => {
        if (result) {
          expandOneKey(treeDataMap.value, dropKey, expandedKeys.value) // expend parent node
          setExpandedKeys('category', treeData.value[0].scenarioId, expandedKeys.value)
        }
      }
  )
}

let currentInstance
onMounted(() => {
  console.log('onMounted')
  currentInstance = getCurrentInstance()
  document.addEventListener("click", clearMenu)
})
onUnmounted(() => {
  document.removeEventListener("click", clearMenu)
})

</script>

<style lang="less">
.interf-tree {
  .ant-tree-iconEle {
    height: 20px !important;
    line-height: 20px !important;
  }

  .name-editor {
    vertical-align: 5px;

    input {
      margin-top: -2px;
      height: 24px;
      width: 120px;
      background-color: transparent;
    }

    .btns {
      display: inline-block;
      padding-left: 4px;

      .anticon {
        padding: 0 2px;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.scenario-tree-main {
  background: #ffffff;
  .tree-container {
    .tree-filter {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 50px;
      margin-top: 8px;
      .search-input {
        margin-left: 16px;
        margin-right: 8px;
      }

      .add-btn {
        margin-left: 2px;
        margin-right: 16px;
        cursor: pointer;
      }
    }

    .deeptest-tree {
      .tree-title {
        position: relative;
        display: inline-block;
        width: 100%;
        .tree-title-text {
          display: inline-block;
          width: calc(100% - 24px);
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        }

        //&:hover{
        //  .more-icon {
        //    background-color: #f5f5f5;
        //  }
        //}
        .more-icon {
          position: absolute;
          right: -8px;
          width: 20px;
        }
      }
    }

    .nodata-tip {
      margin-top: 8px;
      text-align: center;
    }
  }

  .plus-icon{
    margin-right: 8px;
    margin-left: 4px;
    font-size: 18px;
    cursor: pointer;
  }
}


</style>
