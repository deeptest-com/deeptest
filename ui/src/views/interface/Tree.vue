<template>
  <div class="tree-main dp-tree">
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

          class="interface-tree"
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

    <ImportModal
        v-if="showImport"
        :isVisible="showImport"
        :submit="importSubmit"
        :cancel="importClose"
    />

  </div>
</template>

<script lang="ts">
import {computed, defineComponent, getCurrentInstance, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";
import {Interface} from "@/views/interface/data";
import throttle from "lodash.debounce";
import {importSpec, updateNodeName} from "@/views/interface/service";
import {expandAllKeys, expandOneKey, getNodeMap} from "@/services/tree";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {StateType as ProjectStateType} from "@/store/project";

import {getExpandedKeys, setExpandedKeys} from "@/utils/cache";
import {getContextMenuStyle} from "@/utils/dom";
import {NotificationKeyCommon} from "@/utils/const";

import TreeContextMenu from "./components/TreeContextMenu.vue";
import ImportModal from "./components/ImportModal.vue";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceTree',
  props: {},
  components: {
    TreeContextMenu, ImportModal,
    CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined,
  },
  setup(props) {
    const {t} = useI18n();

    const store = useStore<{ Interface: StateType, ProjectGlobal: ProjectStateType }>();
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

    const treeData = computed<any>(() => store.state.Interface.treeData);
    const treeDataMap = computed<any>(() => store.state.Interface.treeDataMap);
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    watch(currProject, () => {
      console.log('watch currProject', currProject.value.id)
      queryTree();
      selectNode([], null)
    }, {deep: false})

    watch(treeData, () => {
      console.log('watch', treeData)

      if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
        tips.value = '右键树状节点操作'
      }

      getExpandedKeysCall()
    })

    const queryTree = throttle(async () => {
      await store.dispatch('Interface/loadInterface');
    }, 100)
    queryTree();

    const replaceFields = {key: 'id', title: 'name'};
    let expandedKeys = ref<number[]>([]);
    let selectedKeys = ref<number[]>([]);
    let checkedKeys = ref<number[]>([])
    let isExpand = ref(false);

    const editedData = ref<any>({})

    const isLeaf = computed<boolean>(() => {
      return contextNode.value && contextNode.value.isLeaf;
    })

    let tree = ref(null)

    const selectNode = (keys, e) => {
      console.log('selectNode', keys, e?.node.dataRef.id)

      if (keys.length === 0 && e) {
        selectedKeys.value = [e.node.dataRef.id] // cancel un-select
        return
      } else {
        selectedKeys.value = keys
      }

      if (!selectedKeys.value || selectedKeys.value.length === 0) { // not to display the design page
        store.dispatch('Interface/getInterface', {isLeaf: false})
        return
      }

      const selectedData = treeDataMap.value[selectedKeys.value[0]]
      if (!selectedData) return

      store.dispatch('Interface/getInterface', {id: selectedData.id, isLeaf: selectedData.isLeaf})
      if (selectedData.isLeaf) {
        store.dispatch('Interface/getLastInvocationResp', selectedData.id)
        store.dispatch('Interface/listInvocation', selectedData.id)
      }
    }

    const checkNode = (keys, e) => {
      console.log('checkNode', checkedKeys)
    }

    const updateName = (id) => {
      const name = editedData.value[id]
      console.log('updateName', id, name)
      updateNodeName(id, name).then((json) => {
        if (json.code === 0) {
          store.dispatch('Interface/saveTreeMapItemProp', {id: id, prop: 'name', value: name})
          store.dispatch('Interface/saveTreeMapItemProp', {id: id, prop: 'isEdit', value: false})
        }
      })
    }
    const cancelUpdate = (id) => {
      console.log('cancelUpdate', id)
      store.dispatch('Interface/saveTreeMapItemProp', {id: id, prop: 'isEdit', value: false})
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

      const contextNodeData = treeDataMap.value[node.eventKey]
      contextNode.value = {
        pageX: x,
        pageY: y,
        id: node.eventKey,
        title: node.title,
        isLeaf: contextNodeData.isLeaf,
        parentId: node.dataRef.parentId
      }

      menuStyle.value = getContextMenuStyle(
          event.currentTarget.getBoundingClientRect().right, event.currentTarget.getBoundingClientRect().top, 185)
    }

    const getExpandedKeysCall = throttle(async () => {
      getExpandedKeys('interface', currProject.value.id).then(async keys => {
        console.log('keys', keys)
        if (keys)
          expandedKeys.value = keys

        if (!expandedKeys.value || expandedKeys.value.length === 0) { // init, expend first level folder
          getOpenKeys(treeData.value[0], false)
          console.log('expandedKeys.value', expandedKeys.value)
          await setExpandedKeys('interface', currProject.value.id, expandedKeys.value)
        }
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
    getOpenKeys(treeData.value[0], false)
    console.log('expandedKeys.value', expandedKeys.value)

    const expandNode = (keys: string[], e: any) => {
      console.log('expandNode', keys[0], e)

      setExpandedKeys('interface', currProject.value.id, expandedKeys.value)
    }
    const expandAll = () => {
      console.log('expandAll')
      isExpand.value = !isExpand.value
      expandedKeys.value = expandAllKeys(treeDataMap.value, isExpand.value)

      setExpandedKeys('interface', currProject.value.id, expandedKeys.value)
    }

    let targetModelId = 0
    const menuClick = (menuKey: string, targetId: number) => {
      console.log('menuClick', menuKey, targetId)
      targetModelId = targetId

      if (menuKey === 'export-spec') {
        showImport.value = true
        return
      }

      if (menuKey === 'rename') {
        renameNode()
        return
      }

      if (menuKey === 'remove') {
        removeNode()
        return
      }

      // add
      const arr = menuKey.split('-')
      const mode = arr[1]
      const type = arr[2]

      addNode(mode, type)

      clearMenu()
    }

    const renameNode = () => {
      selectedKeys.value = [targetModelId]
      selectNode(selectedKeys.value, null)
      editedData.value[targetModelId] = treeDataMap.value[targetModelId].name

      Object.keys(treeDataMap.value).forEach((key) => {
        store.dispatch('Interface/saveTreeMapItemProp', {id: key, prop: 'isEdit', value: false})
      })
      // treeDataMap.value[targetModelId].isEdit = true
      store.dispatch('Interface/saveTreeMapItemProp', {id: targetModelId, prop: 'isEdit', value: true})

      setTimeout(() => {
        console.log('==', currentInstance.ctx.$refs[`name-editor-${targetModelId}`])
        currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.focus()
        currentInstance.ctx.$refs[`name-editor-${targetModelId}`]?.select();
      }, 50)
    }

    const addNode = (mode, type) => {
      console.log('addNode', targetModelId)
      store.dispatch('Interface/createInterface',
          {target: targetModelId, name: type === 'dir' ? '新目录' : '新接口', mode: mode, type: type})
          .then((newNode) => {
            console.log('newNode', newNode)

            targetModelId = newNode.id
            renameNode()
            // selectNode([newNode.id])

            expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
            setExpandedKeys('interface', currProject.value.id, expandedKeys.value)
          })
    }
    const removeNode = () => {
      console.log('removeNode')
      store.dispatch('Interface/deleteInterface', targetModelId);
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
      if (treeDataMap.value[dropKey].isLeaf && dropPosition === 0) dropPosition = 1
      console.log(dragKey, dropKey, dropPosition);

      store.dispatch('Interface/moveInterface', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPosition});
    }

    const showImport = ref(false)
    const importSubmit = (data) => {
      console.log('importSpec', data)
      importSpec(data, targetModelId).then((json) => {
        if (json.code === 0) {
          showImport.value = false
          store.dispatch('Interface/loadInterface').then((result) => {
              console.log(result)
              expandOneKey(treeDataMap.value, targetModelId, expandedKeys.value) // expend parent node
              setExpandedKeys('interface', currProject.value.id, expandedKeys.value)
            })
        } else {
          notification.error({
            key: NotificationKeyCommon,
            message: '导入失败',
          });
        }
      })
    }

    const importClose = () => {
      console.log('showImportClose')
      showImport.value = false
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

    return {
      treeData,
      treeDataMap,
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
      isLeaf,
      onDragEnter,
      onDrop,

      updateName,
      cancelUpdate,

      showImport,
      importSubmit,
      importClose,

      tips,
    }
  }
})
</script>

<style lang="less">

.interface-tree {
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

}
</style>