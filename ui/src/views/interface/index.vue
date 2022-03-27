<template>
  <div id="main">
    <div id="left-panel">
      <div class="toolbar">
        <div class="tips">
          <span>{{tips}}</span>
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
          <template #icon="slotProps">
            <FolderOutlined v-if="slotProps.isDir && !slotProps.expanded" />
            <FolderOpenOutlined v-if="slotProps.isDir && slotProps.expanded" />
            <FileOutlined v-if="!slotProps.isDir" />
          </template>
        </a-tree>

        <div v-if="contextNode.id >= 0" :style="menuStyle">
          <TreeContextMenu :treeNode="contextNode" :onSubmit="menuClick"/>
        </div>
      </div>
    </div>
    <div id="splitter-h"></div>
    <div id="right-panel">
      <InterfaceDesigner v-if="requestData.id" :onSubmit="saveInterface"></InterfaceDesigner>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {Form} from "ant-design-vue";
import {FolderOutlined, FileOutlined, FolderOpenOutlined} from "@ant-design/icons-vue";
import { TreeDragEvent, DropEvent } from 'ant-design-vue/es/tree/Tree';

import {getNodeMap, expandAllKeys, expandOneKey} from "./service";
import {StateType} from "./store";
import throttle from "lodash.debounce";

import TreeContextMenu from './components/TreeContextMenu.vue';
import InterfaceDesigner from './components/Designer.vue';
import {resizeWidth} from "@/utils/dom";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    FolderOutlined, FolderOpenOutlined, FileOutlined,
    TreeContextMenu, InterfaceDesigner,
  },
  setup() {
    const router = useRouter();
    const store = useStore<{ Interface: StateType }>();

    const treeData = computed<any>(() => store.state.Interface.treeData);
    const requestData = computed<any>(() => store.state.Interface.requestData);

    const queryTree = throttle(async () => {
      await store.dispatch('Interface/loadInterface');
    }, 600)
    queryTree();

    const replaceFields = {key: 'id', title: 'name'};
    let expandedKeys = ref<number[]>([]);
    let selectedKeys = ref<string[]>([]);
    let checkedKeys = ref<string[]>([])
    let isExpand = ref(false);

    const isDir = computed<boolean>(() => {
      return contextNode.value && contextNode.value.isDir;
    })

    let tree = ref(null)
    const expandNode = (keys: string[], e: any) => {
      console.log('expandNode', keys[0], e)
    }

    const selectNode = (keys, e) => {
      console.log('selectNode')
      if (selectedKeys.value.length === 0) return

      const selectedData = treeDataMap[selectedKeys.value[0]]
      store.dispatch('Interface/getInterface', {id: selectedData.id, isDir: selectedData.isDir})
    }
    const saveInterface= (data) => {
      console.log('saveInterface', data)
    }

    const checkNode = (keys, e) => {
      console.log('checkNode', checkedKeys)
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

      const contextNodeData = treeDataMap[node.eventKey]
      contextNode.value = {
        pageX: x,
        pageY: y,
        id: node.eventKey,
        title: node.title,
        isDir: contextNodeData.isDir,
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
      resizeWidth('main', 'left-panel', 'splitter-h', 'right-panel', 260, 800)
      document.addEventListener("click", clearMenu)
    })
    onUnmounted(() => {
      console.log('onUnmounted')
      document.removeEventListener("click", clearMenu)
    })

    const getNodeMapCall = throttle(async () => {getNodeMap(treeData.value[0], treeDataMap)}, 300)
    watch(treeData, () => {
      console.log('watch', treeData)
      getNodeMapCall()
      console.log('treeMap', Object.keys(treeDataMap), treeDataMap)
      if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
        tips.value = '右键树状节点操作'
      }
    })

    const expandAll = () => {
      console.log('expandAll')
      isExpand.value = !isExpand.value
      expandedKeys.value = expandAllKeys(treeDataMap, isExpand.value)
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
            expandOneKey(treeDataMap, newNode.parentId, expandedKeys.value) // expend new node
          }
      )
    }
    const removeNode = () => {
      store.dispatch('Interface/deleteInterface', targetModelId);
    }
    const clearMenu = () => {
      // console.log('clearMenu')
      contextNode.value = ref({})
    }

    const onDragEnter = (info: TreeDragEvent) => {
      console.log(info);
      // expandedKeys.value = info.expandedKeys
    };
    const onDrop = (info: DropEvent) => {
      const dragKey = info.dragNode.eventKey;
      const dropKey = info.node.eventKey;
      let dropPos = info.dropPosition > 1? 1: info.dropPosition;
      if (!treeDataMap[dropKey].isDir && dropPos === 0) dropPos = 1
      console.log(dragKey, dropKey, dropPos);

      store.dispatch('Interface/moveInterface', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPos});
    }

    return {
      treeData,
      requestData,

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
      contextNode,
      menuStyle,
      rightVisible,
      onRightClick,
      menuClick,
      saveInterface,
      isDir,
      onDragEnter,
      onDrop,

      tips,
    }
  }

})
</script>

<style lang="less">
.interf-tree {
  .ant-tree-iconEle {
    height: 20px !important;
    line-height: 20px !important;
  }
}
</style>

<style lang="less" scoped>
#main {
  display: flex;
  height: 100%;

  #left-panel {
    width: 260px;
    height: 100%;

    .toolbar {
      display: flex;
      height: 34px;
      border-bottom: 1px solid #D0D7DE;
      .tips {
        flex: 1;
        padding: 2px 3px 0 6px;
        line-height: 31px;
        color: #5a5e66;
      }
      .buttons {
        padding: 2px;
        width: 100px;
        text-align: right;
      }
    }
    .tree-panel {
      height: calc(100% - 32px);
      overflow: auto;
    }
  }

  #right-panel {
    flex: 1;
    height: 100%;
    overflow: auto;
  }

  #splitter-h {
    width: 3px;
    height: 100%;
    background-color: #e6e9ec;
    cursor: ew-resize;

    &:hover {
      width: 3px;
      background-color: #D0D7DE;
    }

    &.active {
      width: 3px;
      background-color: #a9aeb4;
    }
  }
}

</style>