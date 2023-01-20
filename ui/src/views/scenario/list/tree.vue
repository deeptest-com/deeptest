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
          :tree-data="treeDataCategory"
          :replace-fields="replaceFields"
          show-icon
          @expand="expandNode"
          @select="selectNode"
          @rightClick="onRightClick"

          v-model:selectedKeys="selectedKeys"
          v-model:checkedKeys="checkedKeys"
          v-model:expandedKeys="expandedKeys"

          draggable
          @dragenter="onDragEnter"
          @drop="onDrop"

          class="interface-tree"
      >
        <template #title="slotProps">
          <span v-if="treeDataMapCategory[slotProps.id] && treeDataMapCategory[slotProps.id].isEdit" class="name-editor">
            <a-input v-model:value="editedData[slotProps.id]"
                     @keyup.enter=updateName(slotProps.id)
                     @click.stop
                     :ref="'name-editor-' + slotProps.id" />

              <span class="btns">
                <CheckOutlined @click.stop="updateName(slotProps.id)"/>
                <CloseOutlined @click.stop="cancelUpdate(slotProps.id)"/>
              </span>
          </span>

          <span v-else class="name-editor">
            {{ slotProps.name }}
          </span>
        </template>

        <template #icon="slotProps">
          <FolderOutlined v-if="!slotProps.isLeaf && !slotProps.expanded"/>
          <FolderOpenOutlined v-if="!slotProps.isLeaf && slotProps.expanded"/>
          <FileOutlined v-if="slotProps.isLeaf"/>
        </template>
      </a-tree>

      <div v-if="contextNode.id >= 0" :style="menuStyle">
        <TreeContextMenu :treeNode="contextNode" :onMenuClick="menuClick"/>
      </div>
    </div>

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
import {StateType as ScenarioStateType} from "../store";
import {StateType as ProjectStateType} from "@/store/project";
import {isRoot, updateNodeName, isInterface} from "../service";
import TreeContextMenu from "./tree-context-menu.vue";

const useForm = Form.useForm;

const {t} = useI18n();

const store = useStore<{ Scenario: ScenarioStateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Scenario.treeDataCategory);
const treeDataMapCategory = computed<any>(() => store.state.Scenario.treeDataMapCategory);
const nodeDataCategory = computed<any>(()=> store.state.Scenario.nodeDataCategory);

watch(treeDataCategory, () => {
  console.log('watch', treeDataCategory)

  if (!treeDataCategory.value[0].children || treeDataCategory.value[0].children.length === 0) {
    tips.value = '右键树状节点操作'
  }

  getExpandedKeysCall()
})

const loadTree = debounce(async () => {
  await store.dispatch('Scenario/loadCategory');
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

  setExpandedKeys(treeDataCategory.value[0].scenarioId, expandedKeys.value)
}

const selectNode = (keys, e) => {
  console.log('selectNode', keys, e?.node.dataRef.id)

  if (keys.length === 0 && e) {
    selectedKeys.value = [e.node.dataRef.id] // cancel un-select
    return
  } else {
    selectedKeys.value = keys
  }

  if (!selectedKeys.value || selectedKeys.value.length === 0) {
    store.dispatch('Scenario/getCategoryNode', null)
    return
  }

  const selectedData = treeDataMapCategory.value[selectedKeys.value[0]]
  if (selectedData && isRoot(selectedData.entityCategory)) {
    store.dispatch('Scenario/getCategoryNode', null)
    return
  }

  store.dispatch('Scenario/getCategoryNode', selectedData)
}

const updateName = (id) => {
  const name = editedData.value[id]
  console.log('updateName', id, name)

  updateNodeName(id, name).then((json) => {
    if (json.code === 0) {
      store.dispatch('Scenario/saveTreeMapItemProp', {id: id, prop: 'name', value: name})
      store.dispatch('Scenario/saveTreeMapItemProp', {id: id, prop: 'isEdit', value: false})

      if (id === nodeDataCategory.value.processorId) {
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
  getExpandedKeys(treeDataCategory.value[0].scenarioId).then(async keys => {
    console.log('keys', keys)
    if (keys)
      expandedKeys.value = keys

    if (!expandedKeys.value || expandedKeys.value.length === 0) {
      getOpenKeys(treeDataCategory.value[0], false) // expend first level folder
      console.log('expandedKeys.value', expandedKeys.value)
      await setExpandedKeys(treeDataCategory.value[0].scenarioId, expandedKeys.value)
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
getOpenKeys(treeDataCategory.value[0], false)
console.log('expandedKeys.value', expandedKeys.value)

const expandAll = () => {
  console.log('expandAll')
  isExpand.value = !isExpand.value
  expandedKeys.value = expandAllKeys(treeDataMapCategory.value, isExpand.value)

  setExpandedKeys(treeDataCategory.value[0].scenarioId, expandedKeys.value)
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

  // add-child-interface
  // add-child-processor_logic-processor_logic_if
  const arr = menuKey.split('-')
  const mode = arr[1]

  addNode(mode, targetModelId)

  clearMenu()
}

const renameNode = () => {
  selectedKeys.value = [targetModelId]
  selectNode(selectedKeys.value, null)
  editedData.value[targetModelId] = treeDataMapCategory.value[targetModelId].name

  Object.keys(treeDataMapCategory.value).forEach((key) => {
    store.dispatch('Scenario/saveTreeMapItemProp', {id: key, prop: 'isEdit', value: false})
  })
  store.dispatch('Scenario/saveTreeMapItemProp', {id: targetModelId, prop: 'isEdit', value: true})

  setTimeout(() => {
    console.log('==', currentInstance.ctx.$refs[`name-editor-${targetModelId}`])
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.focus()
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.select();
  }, 50)
}

const addNode = (mode, targetId) => {
  console.log('addNode', mode, targetId)

    store.dispatch('Scenario/createCategoryNode', {mode, targetId, name: '新分类'}).then((newNode) => {
      console.log('createCategoryNode successfully', newNode)
      selectNode([newNode.id], null)
      expandOneKey(treeDataMapCategory.value, mode === 'parent' ? newNode.id : newNode.parentId, expandedKeys.value) // expend new node
      setExpandedKeys(currProject.value.id, expandedKeys.value)
    })
}

const interfaceSelectionVisible = ref(false)
const interfaceSelectionFinish = (selectedNodes) => {
  const targetNode = treeDataMapCategory.value[targetModelId]
  console.log('interfaceSelectionFinish', selectedNodes, targetNode)

  store.dispatch('Scenario/addInterfaces',
      {
        selectedNodes: selectedNodes,
        targetId: targetNode.id,
      }).then((newNode) => {
    console.log('addInterfaces successfully', newNode)

    interfaceSelectionVisible.value = false
    selectNode([newNode.id], null)
    expandOneKey(treeDataMapCategory.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys(treeDataCategory.value[0].scenarioId, expandedKeys.value)
  })
}

const interfaceSelectionCancel = () => {
  console.log('interfaceSelectionCancel')
  interfaceSelectionVisible.value = false
}

const removeNode = () => {
  console.log('removeNode')
  store.dispatch('Scenario/removeNode', targetModelId);
  selectNode([], null)
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
  if (isInterface(treeDataMapCategory.value[dropKey].processorCategory) && dropPosition === 0) dropPosition = 1
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
.scenario-tree-main {
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
}
</style>