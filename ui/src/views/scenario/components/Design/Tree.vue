<template>
  <div class="scenario-tree-main dp-tree">
    <div class="toolbar">
      <div class="tips">
        <span>{{ tips }}</span>
      </div>
      <div class="buttons">
        <a-button @click="expandAll" type="link" class="dp-color-primary">
          <span v-if="!isExpand">展开全部</span>
          <span v-if="isExpand">收缩全部</span>
        </a-button>
      </div>
    </div>

    <div class="tree-panel">
      <a-tree
          ref="tree"
          :tree-data="treeData"
          :replace-fields="replaceFields"
          show-icon
          @expand="expandNode"
          @select="selectNode"
          @check="checkNode"
          @rightClick="onRightClick"

          v-model:selectedKeys="selectedKeys"
          v-model:checkedKeys="checkedKeys"
          v-model:expandedKeys="expandedKeys"

          draggable
          @dragenter="onDragEnter"
          @drop="onDrop"

          class="interf-tree"
      >
        <template #title="slotProps">
          <span v-if="treeDataMap[slotProps.id] && treeDataMap[slotProps.id].isEdit" class="name-editor">
            <a-input v-model:value="editedData[slotProps.id]"
                     @keyup.enter=updateName(slotProps.id)
                     @click.stop
                     :ref="'name-editor-' + slotProps.id" />

            <span class="btns">
              <CheckOutlined @click.stop="updateName(slotProps.id)"/>
              <CloseOutlined @click.stop="cancelUpdate(slotProps.id)"/>
            </span>
          </span>

          <span v-else>
            {{ slotProps.name }}
          </span>
        </template>

        <template #icon="slotProps">
          <template v-if="!slotProps.isLeaf">
            <FolderOutlined v-if="!slotProps.expanded"/>
            <FolderOpenOutlined v-if="slotProps.expanded"/>
          </template>

          <FileOutlined v-else />
        </template>
      </a-tree>

      <div v-if="contextNode.id >= 0" :style="menuStyle">
        <TreeContextMenu :treeNode="contextNode" :onMenuClick="menuClick"/>
      </div>
    </div>

    <InterfaceSelectionFromDefine
        v-if="interfaceSelectionVisible && interfaceSelectionSrc==='fromDefine'"
        :onFinish="endpointInterfaceIdsSelectionFinish"
        :onCancel="interfaceSelectionCancel" />

    <InterfaceSelectionFromTest
        v-if="interfaceSelectionVisible && interfaceSelectionSrc==='fromTest'"
        :onFinish="testInterfaceNodesSelectionFinish"
        :onCancel="interfaceSelectionCancel" />

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch, getCurrentInstance} from "vue";

import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";

import {expandAllKeys, expandOneKey} from "@/services/tree";

import {getExpandedKeys, setExpandedKeys} from "@/utils/cache";
import {getContextMenuStyle} from "@/utils/dom";
import {StateType as ScenarioStateType} from "../../store";
import {isRoot, updateNodeName, isInterface} from "../../service";
import TreeContextMenu from "./components/TreeContextMenu.vue";
import InterfaceSelectionFromDefine from "@/views/component/InterfaceSelectionFromDefine/main.vue";
import InterfaceSelectionFromTest from "@/views/component/InterfaceSelectionFromTest/main.vue";

const props = defineProps<{}>()

const useForm = Form.useForm;

const {t} = useI18n();
import {Scenario} from "@/views/scenario/data";
import {confirmToDelete} from "@/utils/confirm";
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

const replaceFields = {key: 'id', title: 'name'};
let expandedKeys = ref<number[]>([]);
let selectedKeys = ref<number[]>([]);
let checkedKeys = ref<number[]>([])
let isExpand = ref(false);

const editedData = ref<any>({})

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

const checkNode = (keys, e) => {
  console.log('checkNode', checkedKeys)
}

const updateName = (id) => {
  const name = editedData.value[id]
  console.log('updateName', id, name)

  updateNodeName(id, name).then((json) => {
    if (json.code === 0) {
      store.dispatch('Scenario/saveTreeMapItemProp', {id: id, prop: 'name', value: name})
      store.dispatch('Scenario/saveTreeMapItemProp', {id: id, prop: 'isEdit', value: false})

      if (id === selectedNode.value.processorId) {
        store.dispatch('Scenario/getNode', {id: id})
      }
    }
  })
}
const cancelUpdate = (id) => {
  console.log('cancelUpdate', id)
  store.dispatch('Scenario/saveTreeMapItemProp', {id: id, prop: 'isEdit', value: false})
}

let contextNode = ref({} as any)
let menuStyle = ref({} as any)
let tips = ref('')
let rightVisible = false

const onRightClick = (e) => {
  console.log('onRightClick', e)
  const {event, node} = e

  const y = event.currentTarget.getBoundingClientRect().top
  const x = event.currentTarget.getBoundingClientRect().right

  contextNode.value = {
    pageX: x,
    pageY: y,
    id: node.eventKey,
    title: node.title,
    entityCategory: node.dataRef.entityCategory,
    isLeaf: node.dataRef.isLeaf,
    entityType: node.dataRef.entityType,
    entityId: node.dataRef.entityId,
    interfaceId: node.dataRef.interfaceId,
    parentId: node.dataRef.parentId,
  }

  menuStyle.value = getContextMenuStyle(
      event.currentTarget.getBoundingClientRect().right, event.currentTarget.getBoundingClientRect().top, 120)
}

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

  if (menuKey === 'rename') {
    renameNode()
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

const renameNode = () => {
  selectedKeys.value = [targetModelId]
  selectNode(selectedKeys.value, null)
  editedData.value[targetModelId] = treeDataMap.value[targetModelId].name

  Object.keys(treeDataMap.value).forEach((key) => {
    store.dispatch('Scenario/saveTreeMapItemProp', {id: key, prop: 'isEdit', value: false})
  })
  store.dispatch('Scenario/saveTreeMapItemProp', {id: targetModelId, prop: 'isEdit', value: true})

  setTimeout(() => {
    console.log('==', currentInstance.ctx.$refs[`name-editor-${targetModelId}`])
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.focus()
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.select();
  }, 50)
}

const addNode = (mode, processorCategory, processorType,
                 targetProcessorCategory, targetProcessorType, targetProcessorId) => {
  console.log('addNode', mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  if (processorCategory === 'interface') { // select a interface
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

const testInterfaceNodesSelectionFinish = (interfaceNodes) => {
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
  console.log(node)

  const title = '确定删除该' + (node.isLeaf?'接口':'目录') + '吗？'
  const context = !node.isLeaf?'删除后所有所有子目录都会被删除。' : ''

  confirmToDelete(title, context, () => {
    store.dispatch('Scenario/removeNode', targetModelId);
    selectNode([], null)
  })
}

const clearMenu = () => {
  console.log('clearMenu')
  contextNode.value = ref({})
}

const onDragEnter = (info: TreeDragEvent) => {
  console.log(info);
  // expandedKeys.value = info.expandedKeys
};
const onDrop = (info: DropEvent) => {
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const dropPos = info.node.pos.split('-');
  let dropPosition = info.dropPosition - Number(dropPos[dropPos.length - 1]);
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
  height: 100%;
}
</style>
