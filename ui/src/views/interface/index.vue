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
          <template #title="slotProps">
            <span v-if="!slotProps.isEdit">{{ slotProps.name }}</span>

            <span v-else class="name-editor">
              <a-input v-model:value="editedData[slotProps.id]"
                       @keyup.enter=updateName(slotProps.id)
                       @click.stop />

              <span class="btns">
                <CheckOutlined @click.stop="updateName(slotProps.id)" />
                <CloseOutlined @click.stop="cancelUpdate(slotProps.id)" />
              </span>
            </span>
          </template>

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
      <InterfaceDesigner v-if="interfaceData.id" :onSubmit="saveInterface"></InterfaceDesigner>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {Form} from "ant-design-vue";
import {FolderOutlined, FileOutlined, FolderOpenOutlined,
  CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";
import { TreeDragEvent, DropEvent } from 'ant-design-vue/es/tree/Tree';

import {getNodeMap, expandAllKeys, expandOneKey, updateNameReq} from "./service";
import {StateType} from "./store";
import throttle from "lodash.debounce";

import TreeContextMenu from './components/TreeContextMenu.vue';
import InterfaceDesigner from './components/Designer.vue';
import {resizeWidth} from "@/utils/dom";
import {Interface} from "@/views/interface/data";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    FolderOutlined, FolderOpenOutlined, FileOutlined,
    CheckOutlined, CloseOutlined,
    TreeContextMenu, InterfaceDesigner,
  },
  setup() {
    const router = useRouter();
    const store = useStore<{ Interface: StateType }>();

    const treeData = computed<any>(() => store.state.Interface.treeData);
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const queryTree = throttle(async () => {
      await store.dispatch('Interface/loadInterface');
    }, 600)
    queryTree();

    const replaceFields = {key: 'id', title: 'name'};
    let expandedKeys = ref<number[]>([]);
    let selectedKeys = ref<number[]>([]);
    let checkedKeys = ref<number[]>([])
    let isExpand = ref(false);

    const editedData = ref<any>({})

    const isDir = computed<boolean>(() => {
      return contextNode.value && contextNode.value.isDir;
    })

    let tree = ref(null)
    const expandNode = (keys: string[], e: any) => {
      console.log('expandNode', keys[0], e)
    }

    const selectNode = (keys) => {
      console.log('selectNode', keys)
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

    const updateName = (id) => {
      const name = editedData.value[id]
      console.log('updateName', id, name)
      updateNameReq(id, name).then((json) => {
        if (json.code === 0) {
          treeDataMap[id].name = name
          treeDataMap[id].isEdit = false
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

      if (selectedKey === 'rename') {
        selectedKeys.value = [targetModelId]
        selectNode(selectedKeys.value)
        console.log('rename', treeDataMap[targetModelId])
        editedData.value[targetModelId] = treeDataMap[targetModelId].name

        treeDataMap[targetModelId].isEdit = true
        return
      }

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
      interfaceData,
      editedData,

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

      updateName,
      cancelUpdate,

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
#main {
  display: flex;
  height: 100%;

  #left-panel {
    width: 260px;
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

  #right-panel {
    flex: 1;
    height: 100%;
    overflow-y: auto;
    overflow-x: hidden;
  }

  #splitter-h {
    width: 1px;
    height: 100%;
    background-color: #e6e9ec;
    cursor: ew-resize;

    &:hover {
      width: 1px;
      background-color: #D0D7DE;
    }

    &.active {
      width: 1px;
      background-color: #a9aeb4;
    }
  }
}

</style>