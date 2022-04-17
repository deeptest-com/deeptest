<template>
  <div class="tree-main">
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
                       @click.stop/>

              <span class="btns">
                <CheckOutlined @click.stop="updateName(slotProps.id)"/>
                <CloseOutlined @click.stop="cancelUpdate(slotProps.id)"/>
              </span>
            </span>
        </template>

        <template #icon="slotProps">
          <FolderOutlined v-if="slotProps.isDir && !slotProps.expanded"/>
          <FolderOpenOutlined v-if="slotProps.isDir && slotProps.expanded"/>
          <FileOutlined v-if="!slotProps.isDir"/>
        </template>
      </a-tree>

      <div v-if="contextNode.id >= 0" :style="menuStyle">
        <TreeContextMenu :treeNode="contextNode" :onSubmit="menuClick"/>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import {computed, defineComponent, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";
import {Interface} from "@/views/interface/data";
import throttle from "lodash.debounce";
import {expandAllKeys, expandOneKey, getNodeMap, updateNameReq} from "@/views/interface/service";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";

import TreeContextMenu from "./TreeContextMenu.vue";
import {StateType as ProjectStateType} from "@/store/project";
import {getExpandedKeys, setExpandedKeys} from "@/utils/cache";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceTree',
  props: {},
  components: {
    TreeContextMenu,
    CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined,
  },
  setup(props) {
    const {t} = useI18n();

    const projectStore = useStore<{ project: ProjectStateType }>();
    const currProject = computed<any>(() => projectStore.state.project.currProject);

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

      setExpandedKeys(currProject.value.id, expandedKeys.value)
    }

    const selectNode = (keys) => {
      console.log('selectNode', keys)
      if (selectedKeys.value.length === 0) return

      const selectedData = treeDataMap[selectedKeys.value[0]]
      store.dispatch('Interface/getInterface', {id: selectedData.id, isDir: selectedData.isDir})
      if (!selectedData.isDir) {
        store.dispatch('Interface/listRequest', selectedData.id)
      }
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

    const getNodeMapCall = throttle(async () => {
      getNodeMap(treeData.value[0], treeDataMap)
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
    getOpenKeys(treeData.value[0], false)
    console.log('expandedKeys.value', expandedKeys.value)

    watch(treeData, () => {
      console.log('watch', treeData)
      getNodeMapCall()
      console.log('treeMap', Object.keys(treeDataMap), treeDataMap)
      if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
        tips.value = '右键树状节点操作'
      }

      getExpandedKeys(currProject.value.id).then(async keys => {
        console.log('keys', keys)
        if (keys)
          expandedKeys.value = keys

        if (!expandedKeys.value || expandedKeys.value.length === 0) {
          getOpenKeys(treeData.value[0], false) // expend first level folder
          console.log('expandedKeys.value', expandedKeys.value)
          await setExpandedKeys(currProject.value.id, expandedKeys.value)
        }
      })
    })

    const expandAll = () => {
      console.log('expandAll')
      isExpand.value = !isExpand.value
      expandedKeys.value = expandAllKeys(treeDataMap, isExpand.value)

      setExpandedKeys(currProject.value.id, expandedKeys.value)
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
      let dropPos = info.dropPosition > 1 ? 1 : info.dropPosition;
      if (!treeDataMap[dropKey].isDir && dropPos === 0) dropPos = 1
      console.log(dragKey, dropKey, dropPos);

      store.dispatch('Interface/moveInterface', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPos});
    }

    onMounted(() => {
      console.log('onMounted')
      document.addEventListener("click", clearMenu)
    })
    onUnmounted(() => {
      document.removeEventListener("click", clearMenu)
    })

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
.tree-main {
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