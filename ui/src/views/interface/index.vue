<template>
  <div id="main">
    <div id="left">
      <div class="toolbar">
        <a-button @click="expandAll" type="link">
          <span v-if="!isExpand">展开全部</span>
          <span v-if="isExpand">收缩全部</span>
        </a-button>
      </div>
      <div>
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
        >
          <template #icon="slotProps">
            <icon-svg v-if="slotProps.isDir" type="folder-outlined"></icon-svg>
            <icon-svg v-if="!slotProps.isDir" type="file-outlined"></icon-svg>
          </template>
        </a-tree>

        <div v-if="treeNode.id >= 0" :style="menuStyle" class="tree-context-menu">
          <a-menu @click="menuClick" mode="inline">
            <a-menu-item key="add_brother_node" class="menu-item" v-if="treeNode.id > 0">
              <PlusOutlined/>
              <span>创建兄弟节点</span>
            </a-menu-item>
            <a-menu-item key="add_child_node" class="menu-item" v-if="treeNode.isDir">
              <PlusOutlined/>
              <span>创建子节点</span>
            </a-menu-item>

            <a-menu-item key="add_brother_dir" class="menu-item" v-if="treeNode.id > 0">
              <PlusOutlined/>
              <span>创建兄弟目录</span>
            </a-menu-item>
            <a-menu-item key="add_child_dir" class="menu-item" v-if="treeNode.isDir">
              <PlusOutlined/>
              <span>创建子目录</span>
            </a-menu-item>


            <a-menu-item key="remove" class="menu-item" v-if="treeNode.id > 0">
              <CloseOutlined/>
              <span v-if="treeNode.isDir">删除目录</span>
              <span v-if="!treeNode.isDir">删除节点</span>
            </a-menu-item>
          </a-menu>
        </div>
      </div>
    </div>
    <div id="resize"></div>
    <div id="content">

    </div>

    <create-form
        :visible="createFormVisible"
        :onCancel="() => setCreateFormVisible(false)"
        :onSubmitLoading="createSubmitLoading"
        :onSubmit="createSubmit"
    />

  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {Props} from 'ant-design-vue/lib/form/useForm';
import {Form, message, Modal} from "ant-design-vue";
import {CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";

import {getNodeMap, getOpenKeys} from "./service";
import {StateType as ListStateType} from "./store";
import throttle from "lodash.debounce";
import IconSvg from "@/components/IconSvg";

import CreateForm from './components/CreateForm.vue';
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
  tree: Ref;
  treeNode: Ref;
  menuStyle: Ref;
  rightVisible: boolean
  onRightClick: (event, node) => void;
  menuClick: (event) => void;
  isDir: ComputedRef<boolean>;

  treeData: ComputedRef<any[]>;
  treeLoading: Ref<boolean>;
  getTree: (current: number) => Promise<void>;
  createFormVisible: Ref<boolean>;
  setCreateFormVisible: (val: boolean) => void;
  createSubmitLoading: Ref<boolean>;
  createSubmit: (values: any, resetFields: (newValues?: Props | undefined) => void) => Promise<void>;

  modelData: ComputedRef;
  getLoading: Ref<number[]>;
  editInterface: (id: number) => Promise<void>;
  updateFormVisible: Ref<boolean>;
  updateFormCancel: () => void;
  updateSubmitLoading: Ref<boolean>;
  updateSubmit: (values: any, resetFields: (newValues?: Props | undefined) => void) => Promise<void>;

  deleteLoading: Ref<number[]>;
  deleteInterface: (id: number) => void;
}

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    PlusOutlined,
    CloseOutlined,
    IconSvg,
    CreateForm,
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
      console.log('selectNode', selectedKeys)
    }
    const checkNode = (keys, e) => {
      console.log('checkNode', checkedKeys)
    }

    let treeNode = ref({} as any)
    let menuStyle = ref({} as any)
    const treeMap = {}
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
        parentID: node.dataRef.parentID || null
      }
      console.log('---', treeNode)

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

    const getNodeMapCall = throttle(async () => {
      getNodeMap(treeData.value[0], treeMap)}, 300)
    watch(treeData, () => {
      getNodeMapCall()
    })

    const treeLoading = ref<boolean>(true);
    const getTree = async (): Promise<void> => {
      treeLoading.value = true;

      treeLoading.value = false;
    }

    const expandAll = () => {
      console.log('expandAll')
      isExpand.value = !isExpand.value
      expandedKeys.value = getOpenKeys(treeMap, isExpand.value)
    }

    let targetModelId = 0
    const menuClick = (e) => {
      console.log('menuClick', e, treeNode)

      targetModelId = treeNode.value.id
      if (e.key === 'remove') {
        removeNode()
        return
      }

      const arr = e.key.split('_')
      addNode(arr[1], arr[2])

      clearMenu()
    }

    const addNode = (mode, type) => {
      console.log('addNode', targetModelId)
      store.dispatch('Interface/createInterface',
          {mode: mode, type: type, target: targetModelId, name: type === 'dir' ? '新目录' : '新接口'});
    }
    const removeNode = () => {
      store.dispatch('Interface/deleteInterface', targetModelId);
    }
    const clearMenu = () => {
      console.log('clearMenu')
      treeNode.value = ref({})
    }

    // 创建
    const createFormVisible = ref<boolean>(false);
    const setCreateFormVisible = (val: boolean) => {
      createFormVisible.value = val;
    };
    const createSubmitLoading = ref<boolean>(false);
    const createSubmit = async (values: any, resetFields: (newValues?: Props | undefined) => void) => {
      createSubmitLoading.value = true;
      const res: boolean = await store.dispatch('Interface/createInterface', values);
      if (res === true) {
        resetFields();
        setCreateFormVisible(false);
        message.success('新增成功！');
        getTree();
      }
      createSubmitLoading.value = false;
    }

    // 更新
    const updateFormVisible = ref<boolean>(false);
    const setUpdateFormVisible = (val: boolean) => {
      updateFormVisible.value = val;
    }
    const updateFormCancel = () => {
      setUpdateFormVisible(false);
      store.commit('ListInterface/setItem', {});
    }
    const updateSubmitLoading = ref<boolean>(false);
    const updateSubmit = async (values: any, resetFields: (newValues?: Props | undefined) => void) => {
      updateSubmitLoading.value = true;
      const res: boolean = await store.dispatch('ListInterface/updateInterface', values);
      if (res === true) {
        updateFormCancel();
        message.success('编辑成功！');
        getTree();
      }
      updateSubmitLoading.value = false;
    }

    // 编辑
    const getLoading = ref<number[]>([]);
    const editInterface = async (id: number) => {
      getLoading.value = [id];
      const res: boolean = await store.dispatch('Interface/getInterface', id);
      if (res === true) {
        setUpdateFormVisible(true);
      }
      getLoading.value = [];
    }

    // 删除
    const deleteLoading = ref<number[]>([]);
    const deleteInterface = (id: number) => {
      Modal.confirm({
        title: '删除脚本',
        content: '确定删除吗？',
        okText: '确认',
        cancelText: '取消',
        onOk: async () => {
          deleteLoading.value = [id];
          const res: boolean = await store.dispatch('ListInterface/deleteInterface', id);
          if (res === true) {
            message.success('删除成功！');
            await getTree();
          }
          deleteLoading.value = [];
        }
      });
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

      treeLoading,
      getTree,
      createFormVisible,
      setCreateFormVisible,
      createSubmitLoading,
      createSubmit,
      getLoading,
      editInterface,

      modelData,
      updateFormVisible,
      updateFormCancel,
      updateSubmitLoading,
      updateSubmit,
      deleteLoading,
      deleteInterface,
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
      height: 32px;
      border-bottom: 1px solid #D0D7DE;
      text-align: right;
    }
  }

  #right {
    flex: 1;
    height: 100%;
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