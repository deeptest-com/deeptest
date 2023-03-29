<template>
  <div class="interface-tree-main dp-tree">
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
import {computed, defineProps, onMounted, onUnmounted, ref, watch, getCurrentInstance, defineEmits} from "vue";

import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";

import {expandAllKeys, expandOneKey} from "@/services/tree";

import {getExpandedKeys, getSelectedKey, setExpandedKeys, setSelectedKey} from "@/utils/cache";
import {getContextMenuStyle} from "@/utils/dom";
import {StateType as InterfaceStateType} from "../store";
import {StateType as ProjectStateType} from "@/store/project";
import {updateCategoryName} from "@/services/category";
import TreeContextMenu from "./tree-context-menu.vue";

const useForm = Form.useForm;

const {t} = useI18n();

const store = useStore<{ Interface: InterfaceStateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Interface.treeDataCategory);
const treeDataMapCategory = computed<any>(() => store.state.Interface.treeDataMapCategory);
const nodeDataCategory = computed<any>(()=> store.state.Interface.nodeDataCategory);

const emit = defineEmits(['select']);


watch(treeDataCategory, () => {
  console.log('watch treeDataCategory', treeDataCategory)
  selectStoredKeyCall()
  getExpandedKeysCall()
  if (!treeDataCategory.value[0].children || treeDataCategory.value[0].children.length === 0) {
    tips.value = '右键树状节点操作'
  }
})

const loadTree = debounce(async () => {
  await store.dispatch('Interface/loadCategory');
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

  setExpandedKeys('category-interface', currProject.value.id, expandedKeys.value)
}

const selectNode = (keys, e) => {
  console.log('selectNode', keys)

  if (keys.length === 0 && e) {
    selectedKeys.value = [e.node.dataRef.id] // cancel un-select
    return
  } else {
    selectedKeys.value = keys
  }

  setSelectedKey('category-interface', currProject.value.id, selectedKeys.value[0])

  emit('select', selectedKeys.value[0]);

  const selectedData = treeDataMapCategory.value[selectedKeys.value[0]]
  store.dispatch('Interface/getCategoryNode', selectedData)
}

const updateName = (id) => {
  const name = editedData.value[id]
  console.log('updateName', id, name)

  updateCategoryName(id, name).then((json) => {
    if (json.code === 0) {
      store.dispatch('Interface/saveTreeMapItemPropCategory', {id: id, prop: 'name', value: name})
      store.dispatch('Interface/saveTreeMapItemPropCategory', {id: id, prop: 'isEdit', value: false})

      if (id === nodeDataCategory.value.processorId) {
        store.dispatch('Interface/getCategory', {id: id})
      }
    }
  })
}
const cancelUpdate = (id) => {
  console.log('cancelUpdate', id)
  store.dispatch('Interface/saveTreeMapItemPropCategory', {id: id, prop: 'isEdit', value: false})
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
  getExpandedKeys('category-interface', currProject.value.id).then(async keys => {
    console.log('keys', keys)
    if (keys)
      expandedKeys.value = keys

    if (!expandedKeys.value || expandedKeys.value.length === 0) {
      getOpenKeys(treeDataCategory.value[0], false) // expend first level folder
      console.log('expandedKeys.value', expandedKeys.value)
      await setExpandedKeys('category-interface', currProject.value.id, expandedKeys.value)
    }
  })
}, 300)

const selectStoredKeyCall = debounce(async () => {
  console.log('selectStoredKeyCall')
  getSelectedKey('category-interface', currProject.value.id).then(async key => {
    console.log('key', key)
    key = key ? key : treeDataCategory.value[0].id
    if (key)
      selectNode([key], null)
  })
}, 300)

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

  setExpandedKeys('category-interface', currProject.value.id, expandedKeys.value)
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
    store.dispatch('Interface/saveTreeMapItemPropCategory', {id: key, prop: 'isEdit', value: false})
  })
  store.dispatch('Interface/saveTreeMapItemPropCategory', {id: targetModelId, prop: 'isEdit', value: true})
  setTimeout(() => {
    console.log('==', currentInstance.ctx.$refs[`name-editor-${targetModelId}`])
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.focus()
    currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.select();
  }, 50)
}

const addNode = (mode, targetId) => {
  console.log('addNode', mode, targetId)

    store.dispatch('Interface/createCategoryNode',
        {mode, targetId, name: '新分类', type: 'interface'}).then((newNode) => {
      console.log('createCategoryNode successfully', newNode)
      selectNode([newNode.id], null)
      expandOneKey(treeDataMapCategory.value, mode === 'parent' ? newNode.id : newNode.parentId, expandedKeys.value) // expend new node
      setExpandedKeys('category-interface', currProject.value.id, expandedKeys.value)
    })
}

const removeNode = () => {
  console.log('removeNode')
  store.dispatch('Interface/removeCategoryNode', targetModelId);
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
  if (treeDataMapCategory.value[dropKey].isLeaf && dropPosition === 0) {
    dropPosition = 1
  }
  console.log(dragKey, dropKey, dropPosition);

  store.dispatch('Interface/moveCategoryNode',
      {dragKey: dragKey, dropKey: dropKey, dropPos: dropPosition, type: 'interface'}).then(
      (result) => {
        if (result) {
          expandOneKey(treeDataMapCategory.value, dropKey, expandedKeys.value) // expend parent node
          setExpandedKeys('category-interface', currProject.value.id, expandedKeys.value)
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
.interface-tree-main {
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
.interface-tree-main {
}
</style>
