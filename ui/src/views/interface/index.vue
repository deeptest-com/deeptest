<template>
  <div id="main">
    <div id="left">
      <div class="toolbar">
        <div class="tips">
          <span>{{tips}}</span>
        </div>
        <div class="buttons">
          <a-button @click="expandAll" type="link">
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
        >
          <template #icon="slotProps">
            <FolderOutlined v-if="slotProps.isDir && !slotProps.expanded" />
            <FolderOpenOutlined v-if="slotProps.isDir && slotProps.expanded" />
            <FileOutlined v-if="!slotProps.isDir" />
          </template>
        </a-tree>

        <div v-if="treeNode.id >= 0" :style="menuStyle" class="tree-context-menu">
          <TreeContextMenu :onSubmit="menuClick" :treeNode="treeNode"/>
        </div>
      </div>
    </div>
    <div id="resize"></div>
    <div id="content">

    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {Form, message, Modal} from "ant-design-vue";
import {CloseOutlined, PlusOutlined, FolderOutlined, FileOutlined, FolderOpenOutlined} from "@ant-design/icons-vue";
import { TreeDataItem, TreeDragEvent, DropEvent } from 'ant-design-vue/es/tree/Tree';

import {getNodeMap, expandAllKeys, expandOneKey} from "./service";
import {StateType as ListStateType} from "./store";
import throttle from "lodash.debounce";

import TreeContextMenu from './components/TreeContextMenu.vue';
import {resizeWidth} from "@/utils/dom";

const useForm = Form.useForm;

interface InterfaceIndexPageSetupData {
  replaceFields: any,
  expandNode: (expandedKeys: string[], e: any) => void;
  selectNode: (keys, e) => void;
  checkNode: (keys, e) => void;
  isExpand: Ref<boolean>;
  expandAll: (e) => void;
  selectedKeys: Ref<string[]>
  checkedKeys: Ref<string[]>
  expandedKeys: Ref<number[]>
  tips: Ref<string>
  tree: Ref;
  treeNode: Ref;
  menuStyle: Ref;
  rightVisible: boolean
  onRightClick: (event, node) => void;
  menuClick: (selectedKey: string, targetId: number) => void;
  isDir: ComputedRef<boolean>;

  onDragEnter: (info) => void;
  onDrop: (info) => void;

  treeData: ComputedRef<any[]>;
  treeLoading: Ref<boolean>;
  getTree: (current: number) => Promise<void>;
  modelData: ComputedRef;
}

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    FolderOutlined, FolderOpenOutlined, FileOutlined,
    TreeContextMenu,
  },
  setup(): InterfaceIndexPageSetupData {
    const router = useRouter();
    const store = useStore<{ Interface: ListStateType }>();

    const treeData = computed<any>(() => store.state.Interface.treeResult);
    const modelData = computed<any>(() => store.state.Interface.detailResult);

    const queryTree = throttle(async () => {
      treeLoading.value = true;
      await store.dispatch('Interface/loadInterface');
      treeLoading.value = false;
    }, 600)
    queryTree();

    const replaceFields = {key: 'id', title: 'name'};
    let expandedKeys = ref<number[]>([]);
    let selectedKeys = ref<string[]>([]);
    let checkedKeys = ref<string[]>([])
    let isExpand = ref(false);

    const isDir = computed<boolean>(() => {
      return treeNode.value && treeNode.value.isDir;
    })

    let tree = ref(null)
    const expandNode = (keys: string[], e: any) => {
      console.log('expandNode', keys[0], e)
    }
    const selectNode = (keys, e) => {
      if (selectedKeys.value.length === 0) return

      const nodeData = treeMap[selectedKeys.value[0]]
      console.log('selectNode', nodeData.id)
    }
    const checkNode = (keys, e) => {
      console.log('checkNode', checkedKeys)
    }

    let treeNode = ref({} as any)
    let menuStyle = ref({} as any)
    const treeMap = {}
    let tips = ref('')
    let rightVisible = false
    const onRightClick = (e) => {
      console.log('onRightClick', e)
      const {event, node} = e

      const y = event.currentTarget.getBoundingClientRect().top
      const x = event.currentTarget.getBoundingClientRect().right

      const treeNodeData = treeMap[node.eventKey]
      treeNode.value = {
        pageX: x,
        pageY: y,
        id: node.eventKey,
        title: node.title,
        isDir: treeNodeData.isDir,
        parentId: node.dataRef.parentId
      }

      menuStyle.value = {
        position: 'fixed',
        maxHeight: 40,
        textAlign: 'center',
        left: `${x + 10}px`,
        top: `${y + 6}px`
        // display: 'flex',
        // flexDirection: 'row'
      }
    }

    onMounted(() => {
      console.log('onMounted')
      resizeWidth('main', 'left', 'resize', 'content', 280, 800)
      getTree();
      document.addEventListener("click", clearMenu)
    })
    onUnmounted(() => {
      console.log('onUnmounted')
      document.removeEventListener("click", clearMenu)
    })

    const getNodeMapCall = throttle(async () => {getNodeMap(treeData.value[0], treeMap)}, 300)
    watch(treeData, () => {
      console.log('watch', treeData)
      getNodeMapCall()
      console.log('treeMap', Object.keys(treeMap), treeMap)
      if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
        tips.value = '右键树状节点操作'
      }
    })

    const treeLoading = ref<boolean>(true);
    const getTree = async (): Promise<void> => {
      treeLoading.value = true;

      treeLoading.value = false;
    }

    const expandAll = () => {
      console.log('expandAll')
      isExpand.value = !isExpand.value
      expandedKeys.value = expandAllKeys(treeMap, isExpand.value)
    }

    let targetModelId = 0
    const menuClick = (selectedKey: string, targetId: number) => {
      console.log('menuClick', selectedKey, targetId)

      targetModelId = targetId
      if (selectedKey === 'remove') {
        removeNode()
        return
      }

      const arr = selectedKey.split('_')
      addNode(arr[1], arr[2])

      clearMenu()
    }

    const addNode = (mode, type) => {
      console.log('addNode', targetModelId)
      store.dispatch('Interface/createInterface',
          {mode: mode, type: type, target: targetModelId, name: type === 'dir' ? '新目录' : '新接口'})
          .then((newNode) => {
            console.log('newNode', newNode)
            selectedKeys.value = [newNode.id] // select new node
            expandOneKey(treeMap, newNode.parentId, expandedKeys.value) // expend new node
          }
      )
    }
    const removeNode = () => {
      store.dispatch('Interface/deleteInterface', targetModelId);
    }
    const clearMenu = () => {
      // console.log('clearMenu')
      treeNode.value = ref({})
    }

    const onDragEnter = (info: TreeDragEvent) => {
      console.log(info);
      // expandedKeys.value = info.expandedKeys
    };
    const onDrop = (info: DropEvent) => {
      const dragKey = info.dragNode.eventKey;
      const dropKey = info.node.eventKey;
      let dropPos = info.dropPosition > 1? 1: info.dropPosition;
      if (!treeMap[dropKey].isDir && dropPos === 0) dropPos = 1
      console.log(dragKey, dropKey, dropPos);

      store.dispatch('Interface/moveInterface', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPos});
    }

    return {
      treeData,
      replaceFields,
      expandedKeys,
      selectedKeys,
      checkedKeys,
      isExpand,
      expandAll,
      expandNode,
      selectNode,
      checkNode,
      tree,
      treeNode,
      menuStyle,
      rightVisible,
      onRightClick,
      menuClick,
      isDir,
      onDragEnter,
      onDrop,

      treeLoading,
      tips,
      getTree,
      modelData,
    }
  }

})
</script>

<style lang="less" scoped>
#main {
  display: flex;
  height: 100%;

  #left {
    width: 300px;
    height: 100%;

    .toolbar {
      display: flex;
      height: 32px;
      border-bottom: 1px solid #D0D7DE;
      .tips {
        flex: 1;
        padding: 0 3px 0 6px;
        line-height: 31px;
        color: #5a5e66;
      }
      .buttons {
        width: 100px;
        text-align: right;
      }
    }
    .tree-panel {
      height: calc(100% - 32px);
      overflow: auto;
    }
  }

  #right {
    flex: 1;
    height: 100%;
    overflow: auto;
  }

  #resize {
    width: 2px;
    height: 100%;
    background-color: #D0D7DE;
    cursor: ew-resize;

    &.active {
      background-color: #a9aeb4;
    }
  }
}

</style>