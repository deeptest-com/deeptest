<template>
  <div class="scenario-tree-main">
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
          <span v-if="!slotProps.isEdit">{{ slotProps.name }}</span>

          <span v-else class="name-editor">
              <a-input v-model:value="editedData[slotProps.id]"
                       @keyup.enter=updateName(slotProps.id)
                       @click.stop
                       :ref="'name-editor-' + slotProps.id" />

              <span class="btns">
                <CheckOutlined @click.stop="updateName(slotProps.id)"/>
                <CloseOutlined @click.stop="cancelUpdate(slotProps.id)"/>
              </span>
            </span>
        </template>

        <template #icon="slotProps">
          <FolderOutlined v-if="slotProps.entityCategory !==  'processor_interface' && !slotProps.expanded"/>
          <FolderOpenOutlined v-if="slotProps.entityCategory !==  'processor_interface' && slotProps.expanded"/>
          <FileOutlined v-if="slotProps.entityCategory ===  'processor_interface'"/>
        </template>
      </a-tree>

      <div v-if="contextNode.id >= 0" :style="menuStyle">
        <TreeContextMenu :treeNode="contextNode" :onMenuClick="menuClick"/>
      </div>
    </div>

    <InterfaceSelection
        v-if="interfaceSelectionVisible"
        :onFinish="interfaceSelectionFinish"
        :onCancel="interfaceSelectionCancel"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch, getCurrentInstance} from "vue";

import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {isRoot, isProcessor, isInterface} from '../../service'
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";

import debounce from "lodash.debounce";
import {expandAllKeys, expandOneKey, getNodeMap} from "@/services/tree";
import {getProcessorTypeNames, updateNodeName} from "../../service";
import {useStore} from "vuex";
import {Scenario} from "../../data";

import {StateType as ScenarioStateType} from "../../store";
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as ProjectStateType} from "@/store/project";

import {getExpandedKeys, setExpandedKeys} from "@/utils/cache";

import TreeContextMenu from "./TreeContextMenu.vue";
import InterfaceSelection from "./InterfaceSelection.vue";
import {StateType} from "@/views/interface/store";
import {getContextMenuStyle} from "@/utils/dom";

const props = defineProps<{ scenarioId: number }>()

const useForm = Form.useForm;

const {t} = useI18n();

const store = useStore<{ Scenario: ScenarioStateType; Interface: InterfaceStateType, Project: ProjectStateType; }>();
const treeData = computed<any>(() => store.state.Scenario.treeData);
const selectedNode = computed<any>(()=> store.state.Scenario.nodeData);

const loadTree = debounce(async () => {
  await store.dispatch('Scenario/loadScenario', props.scenarioId);
}, 60)
loadTree();

const replaceFields = {key: 'id', title: 'name'};
let expandedKeys = ref<number[]>([]);
let selectedKeys = ref<number[]>([]);
let checkedKeys = ref<number[]>([])
let isExpand = ref(false);

const editedData = ref<any>({})

let tree = ref(null)
const expandNode = (keys: string[], e: any) => {
  console.log('expandNode', keys[0], e)

  setExpandedKeys(treeData.value[0].scenarioId, expandedKeys.value)
}

const selectNode = (keys) => {
  console.log('selectNode', keys)
  if (selectedKeys.value.length === 0) return

  const selectedData = treeDataMap[selectedKeys.value[0]]
  if (isRoot(selectedData.entityCategory)) return

  store.dispatch('Scenario/getNode', selectedData).then((ok) => {
    console.log('===', selectedNode.value)
    if (ok && selectedNode.value.processorType === 'processor_interface_default') {
      const interfaceId = selectedNode.value.interfaceId
      store.dispatch('Interface/getInterface', {
        id: interfaceId
      })

      store.dispatch('Interface/listInvocation', interfaceId)
      store.dispatch('Interface/listEnvironment')
      store.dispatch('Interface/getEnvironment', {id: 0, interfaceId: interfaceId})
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
      treeDataMap[id].name = name
      treeDataMap[id].isEdit = false

      if (id === selectedNode.value.processorId) {
        store.dispatch('Scenario/getNode', {id: id})
      }
    }
  })
}
const cancelUpdate = (id) => {
  console.log('cancelUpdate', id)
  treeDataMap[id].isEdit = false
}

let contextNode = ref({} as any)
let menuStyle = ref({} as any)
const treeDataMap = {}
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
    entityType: node.dataRef.entityType,
    entityId: node.dataRef.entityId,
    interfaceId: node.dataRef.interfaceId,
    parentId: node.dataRef.parentId,
  }

  menuStyle.value = getContextMenuStyle(
      event.currentTarget.getBoundingClientRect().right, event.currentTarget.getBoundingClientRect().top, 120)
}

const getNodeMapCall = debounce(async () => {
  getNodeMap(treeData.value[0], treeDataMap)
}, 500)
const getExpandedKeysCall = debounce(async () => {
  getExpandedKeys(treeData.value[0].scenarioId).then(async keys => {
    console.log('keys', keys)
    if (keys)
      expandedKeys.value = keys

    if (!expandedKeys.value || expandedKeys.value.length === 0) {
      getOpenKeys(treeData.value[0], false) // expend first level folder
      console.log('expandedKeys.value', expandedKeys.value)
      await setExpandedKeys(treeData.value[0].scenarioId, expandedKeys.value)
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

watch(treeData, () => {
  console.log('watch', treeData)
  getNodeMapCall()
  // console.log('treeMap', Object.keys(treeDataMap), treeDataMap)
  if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
    tips.value = '右键树状节点操作'
  }

  getExpandedKeysCall()
})

const expandAll = () => {
  console.log('expandAll')
  isExpand.value = !isExpand.value
  expandedKeys.value = expandAllKeys(treeDataMap, isExpand.value)

  setExpandedKeys(treeData.value[0].scenarioId, expandedKeys.value)
}

let targetModelId = 0
const menuClick = (menuKey: string, targetId: number) => {
  console.log('menuClick', menuKey, targetId)

  targetModelId = targetId

  if (menuKey === 'rename') {
    selectedKeys.value = [targetModelId]
    // selectNode(selectedKeys.value)
    editedData.value[targetModelId] = treeDataMap[targetModelId].name

    Object.keys(treeDataMap).forEach((key) => {
      treeDataMap[key].isEdit = false
    })
    treeDataMap[targetModelId].isEdit = true

    setTimeout(() => {
      console.log('==', currentInstance.ctx.$refs[`name-editor-${targetModelId}`])
      currentInstance.ctx.$refs[`name-editor-${targetModelId}`].focus()
    }, 50)

    return
  }

  if (menuKey === 'remove') {
    removeNode()
    return
  }

  // add-child-interface
  // add-child-processor_logic-processor_logic_if
  const arr = menuKey.split('-')
  const mode = arr[1]
  const processorCategory = arr[2]
  const processorType = arr[3]

  const targetProcessorCategory = treeDataMap[targetModelId].entityCategory
  const targetProcessorType = treeDataMap[targetModelId].entityType

  const targetProcessorId = targetModelId

  addNode(mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  clearMenu()
}

const addNode = (mode, processorCategory, processorType,
                 targetProcessorCategory, targetProcessorType, targetProcessorId) => {
  console.log('addNode', mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  if (processorCategory === 'interface') { // select a interface
    interfaceSelectionVisible.value = true
    return
  } else {
    store.dispatch('Scenario/addProcessor',
        {mode, processorCategory, processorType,
          targetProcessorCategory, targetProcessorType, targetProcessorId,
          name: t(processorType)
        }).then((newNode) => {
          console.log('addProcessor successfully', newNode)
          selectNode([newNode.id])
          expandOneKey(treeDataMap, mode === 'parent' ? newNode.id : newNode.parentId, expandedKeys.value) // expend new node
          setExpandedKeys(treeData.value[0].scenarioId, expandedKeys.value)
        })
  }
}

const interfaceSelectionVisible = ref(false)
const interfaceSelectionFinish = (selectedNodes) => {
  const targetNode = treeDataMap[targetModelId]
  console.log('interfaceSelectionFinish', selectedNodes, targetNode)

  store.dispatch('Scenario/addInterfaces',
      {
        selectedNodes: selectedNodes,
        targetId: targetNode.id,
      }).then((newNode) => {
        console.log('addInterfaces successfully', newNode)

        interfaceSelectionVisible.value = false
        selectNode([newNode.id])
        expandOneKey(treeDataMap, newNode.parentId, expandedKeys.value) // expend new node
        setExpandedKeys(treeData.value[0].scenarioId, expandedKeys.value)
      })
}

const interfaceSelectionCancel = () => {
  console.log('interfaceSelectionCancel')
  interfaceSelectionVisible.value = false
}

const removeNode = () => {
  console.log('removeNode')
  store.dispatch('Scenario/removeNode', targetModelId);
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
  if (isInterface(treeDataMap[dropKey].processorCategory) && dropPosition === 0) dropPosition = 1
  console.log(dragKey, dropKey, dropPosition);

  store.dispatch('Scenario/moveNode', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPosition});
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

  .toolbar {
    display: flex;
    height: 32px;
    border-bottom: 1px solid #D0D7DE;

    .tips {
      flex: 1;
      padding: 0px 3px 0 6px;
      line-height: 31px;
      color: #5a5e66;
    }

    .buttons {
      padding: 0px;
      width: 100px;
      text-align: right;
    }
  }

  .tree-panel {
    height: calc(100% - 32px);
    overflow: auto;
  }
}
</style>