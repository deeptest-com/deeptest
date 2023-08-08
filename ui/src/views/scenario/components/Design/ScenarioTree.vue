<template>
  <div class="scenario-tree-main">
    <div class="dp-tree-container">
      <div class="tree-filter">
        <a-input-search
            placeholder="输入关键字过滤"
            class="search-input"
            allowClear
            v-model:value="keywords"/>
        <TreeMenu @selectMenu="selectMenu" :treeNode="treeData?.[0]">
          <template #button>
            <PlusOutlined class="plus-icon" @click.prevent.stop/>
          </template>
        </TreeMenu>
      </div>
      <div class="tree-content">
        <a-tree
            class="deeptest-tree"
            draggable
            blockNode
            showIcon
            :expandAction="false"
            :expandedKeys="expandedKeys"
            :auto-expand-parent="autoExpandParent"
            v-model:selectedKeys="selectedKeys"
            @drop="onDrop"
            @expand="onExpand"
            @select="selectNode"
            :tree-data="treeDataNeedRender"
            :replace-fields="replaceFields">
          <template #switcherIcon>
            <CaretDownOutlined/>
          </template>
          <template #title="{dataRef}">
            <div class="tree-title" :draggable="dataRef.id === -1">
              <div class="title" :class="[dataRef.disable ? 'dp-disabled' : '']">
                <!-- 标题前缀 -->
                <span class="prefix-icon"
                      v-if="dataRef?.entityType !== 'processor_interface_default'">
                  <IconSvg v-if="DESIGN_TYPE_ICON_MAP[dataRef.entityType]"
                           :type="DESIGN_TYPE_ICON_MAP[dataRef.entityType]" class="prefix-icon-svg"/>
                </span>

                <!-- 请求：请求方法 -->
                <span class="prefix-req-method" v-if="dataRef.entityType === 'processor_interface_default'">
                  <a-tag class="method-tag" :color="getMethodColor(dataRef.method || 'GET', dataRef.disable)">{{
                      dataRef.method || "GET"
                    }}</a-tag>
                </span>
                <span class="title-text" :title="dataRef.name">
                {{ dataRef.name }}
              </span>
              </div>
              <div class="icon" v-if="dataRef.id > 0">
                <TreeMenu @selectMenu="selectMenu" :treeNode="dataRef">
                  <template #button>
                    <MoreOutlined/>
                  </template>
                </TreeMenu>
              </div>
            </div>
          </template>
        </a-tree>
        <div v-if="showAddTip" class="nodata-tip">请点击上方按钮添加编排场景 ~</div>
        <div v-if="showChangeKeywordTip" class="nodata-tip">搜索结果为空,请更换搜索关键词 ~</div>
      </div>
    </div>

    <!--  编辑接口弹窗  -->
    <EditModal
        v-if="currentNode"
        :nodeInfo="currentNode"
        @ok="handleEditModalOk"
        @cancel="handleEditModalCancel"/>

    <InterfaceSelectionFromDefine
        v-if="interfaceSelectionVisible && interfaceSelectionSrc.includes(ProcessorInterfaceSrc.Define)"
        :onFinish="endpointInterfaceIdsSelectionFinish"
        :onCancel="interfaceSelectionCancel"/>

    <InterfaceSelectionFromDiagnose
        v-if="interfaceSelectionVisible && interfaceSelectionSrc.includes(ProcessorInterfaceSrc.Diagnose)"
        :onFinish="diagnoseInterfaceNodesSelectionFinish"
        :onCancel="interfaceSelectionCancel"/>

    <!--  Curl导入弹窗  -->
    <InterfaceImportFromCurl
        v-if="interfaceSelectionVisible && interfaceSelectionSrc.includes(ProcessorInterfaceSrc.Curl)"
        @onFinish="interfaceImportFromCurlFinish"
        @onCancel="interfaceImportFromCurlCancel"/>

    <InterfaceSelectionFromDefineCase
        v-if="interfaceSelectionVisible && interfaceSelectionSrc.includes(ProcessorInterfaceSrc.Case)"
        :onFinish="endpointCaseSelectionFinish"
        :onCancel="endpointCaseSelectionCancel"/>

  </div>
</template>
<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch, getCurrentInstance} from "vue";

import {useI18n} from "vue-i18n";
import {Form, message, Modal} from 'ant-design-vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {confirmToDelete} from "@/utils/confirm";
import {filterTree, filterByKeyword} from "@/utils/tree";
import {ProcessorInterface, ProcessorInterfaceSrc} from "@/utils/enum";
import {DESIGN_TYPE_ICON_MAP, menuKeyMapToProcessorCategory} from "./config";
import {getMethodColor} from "@/utils/dom";
import {DropEvent, TreeDragEvent} from "ant-design-vue/es/tree/Tree";
import {PlusOutlined, CaretDownOutlined, MoreOutlined, FolderOpenOutlined, FolderOutlined} from '@ant-design/icons-vue';
import {expandAllKeys, expandOneKey} from "@/services/tree";
import TreeMenu from "./components/TreeMenu/index.vue";
import IconSvg from "@/components/IconSvg";
import {getExpandedKeys, getSelectedKey, setExpandedKeys} from "@/utils/cache";
import {StateType as ScenarioStateType} from "../../store";
import {isRoot, updateNodeName, isInterface} from "../../service";
import {Scenario} from "@/views/scenario/data";
import EditModal from "./components/edit.vue";
import InterfaceSelectionFromDefine from "@/views/component/InterfaceSelectionFromDefine/main.vue";
import InterfaceSelectionFromDiagnose from "@/views/component/InterfaceSelectionFromDiagnose/main.vue";
import cloneDeep from "lodash/cloneDeep";
import InterfaceImportFromCurl from "@/views/component/InterfaceImportFromCurl/index.tsx";
import InterfaceSelectionFromDefineCase from "@/views/component/InterfaceSelectionFromDefineCase/index.vue";

const props = defineProps<{}>()
const {t} = useI18n();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const treeData = computed<any>(() => store.state.Scenario.treeData);
const treeDataNeedRender = computed<any>(() => {
  const children = cloneDeep(treeData.value?.[0]?.children);
  if (children?.length > 0) {
    return [...filterByKeyword(children, keywords.value, 'name')];
  }
  return []
});

const showChangeKeywordTip = computed(() => {
  if (showAddTip.value) {
    return false;
  }
  return keywords.value && treeDataNeedRender.value?.length === 0;
})

const showAddTip = computed(() => {
  const children = treeData.value?.[0]?.children;
  return !children?.length;
})

const treeDataMap = computed<any>(() => store.state.Scenario.treeDataMap);
const selectedNode = computed<any>(() => store.state.Scenario.nodeData);
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);

watch(treeData, () => {
  console.log('832 watch treeData1', treeData.value)
  console.log('832 watch treeData2', treeDataNeedRender.value)
  if (!treeData.value[0].children || treeData.value[0].children.length === 0) {
    tips.value = '右键树状节点操作'
  }
  getExpandedKeysCall()
})

watch(() => {
  return detailResult.value.id
}, async (newVal) => {
  if (newVal) {
    // loadTree();
    await store.dispatch('Scenario/loadScenario', newVal);
  }
}, {
  immediate: true
})

const keywords = ref('');
watch(keywords, (newVal) => {
  console.log('watch keywords', keywords);
  expandedKeys.value = filterTree(treeData.value, newVal)
  autoExpandParent.value = true;
});

const replaceFields = {key: 'id', title: 'name'};

let expandedKeys = ref<number[]>([]);
const autoExpandParent = ref<boolean>(false);

let selectedKeys = ref<number[]>([]);
let isExpand = ref(false);

let tree = ref(null)
const expandNode = (keys: string[], e: any) => {
  console.log('expandNode', keys[0], e)

  setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
}

const selectNode = (keys, e) => {
  console.log('selectNode', keys, e?.node?.dataRef, e?.node?.dataRef?.id)
  if (keys.length === 0 && e && e?.node?.dataRef?.id) {
    selectedKeys.value = [e.node.dataRef.id] // cancel un-select
    return
  } else {
    selectedKeys.value = keys
  }

  const selectedData = treeDataMap.value[selectedKeys.value[0]]
  console.log('selectedData', selectedData)

  if (selectedData && isRoot(selectedData.entityCategory)) {
    store.dispatch('Scenario/getNode', null)
    return
  }

  store.dispatch('Scenario/getNode', selectedData).then((ok) => {
    if (ok && selectedData?.entityType === ProcessorInterface.Interface) {
      // will cause watch event to load debug data in components/interface/interface.vue
      store.dispatch('Scenario/setScenarioProcessorIdForDebug', selectedData.id)
    }
  })
}

const onExpand = (keys: number[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

const currentNode = ref(null as any);

function create(parentId, type) {
  console.log('create', parentId, type)
  currentNode.value = {parentId, type};
}

function edit(node) {
  console.log('edit', node)
  currentNode.value = node;
}

let contextNode = ref({} as any)
let tips = ref('')

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
  console.log('menuClick', menuKey)
  targetModelId = targetId
  if (menuKey === 'edit') {
    edit(treeDataMap.value[targetModelId])
    return
  }
  if (menuKey === 'remove') {
    removeNode()
    return
  }
  // add-child-interface-interface
  // add-child-interface-diagnose
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

/**
 * 选中的菜单key，对应的处理器类型
 * */
function selectMenu(menuInfo, treeNode) {
  targetModelId = treeNode?.id;
  const key = menuInfo.key;
  const mode = 'child';
  const processorType = key;
  // 检验必要字段
  if (!targetModelId) return;
  if (!key) return;

  if (key === 'edit') {
    edit(treeDataMap.value[targetModelId])
    return
  }
  if (key === 'remove') {
    removeNode()
    return
  }
  if (key === 'disable' || key === 'enable') {
    disableNodeOrNot()
    return
  }

  const processorCategory = menuKeyMapToProcessorCategory[key];
  if (!processorCategory) return;
  const targetProcessorId = targetModelId
  const targetProcessorCategory = treeDataMap.value[targetModelId].entityCategory
  const targetProcessorType = treeDataMap.value[targetModelId].entityType
  addNode(mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId);
}

async function handleEditModalOk(model) {
  console.log('handleEditModalOk')

  // convert data
  model.processorCategory = model.entityCategory
  model.processorType = model.entityType

  if (!model.id && model.entityType === ProcessorInterface.Interface) { // create interface
    store.dispatch('Scenario/addProcessor', model).then((newNode) => {
      console.log('addProcessor successfully', newNode)
      currentNode.value = null

      selectNode([newNode.id], null)
      expandOneKey(treeDataMap.value, model.mode === 'parent' ? newNode.id : newNode.parentId, expandedKeys.value) // expend new node
      setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
    })

    return
  }

  const res = await store.dispatch('Scenario/saveProcessorInfo', model)
  if (res) {
    currentNode.value = null
    expandOneKey(treeDataMap.value, model.parentId, expandedKeys.value)
  }
}

function handleEditModalCancel() {
  console.log('handleEditModalCancel')
  currentNode.value = null
}

const addNode = (mode, processorCategory, processorType,
                 targetProcessorCategory, targetProcessorType, targetProcessorId) => {
  console.log('addNode', mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId)

  // 如果是添加接口，会弹框选择接口类型
  if (processorCategory === 'interface' || processorCategory === 'processor_interface') { // show popup to select a interface
    interfaceSelectionSrc.value = processorType.substr(processorType.lastIndexOf('-') + 1)
    if (interfaceSelectionSrc.value === '' + ProcessorInterfaceSrc.Custom) { // show interface create popup
      currentNode.value = {
        name: '',
        entityCategory: processorCategory,
        entityType: ProcessorInterface.Interface,
        processorInterfaceSrc: interfaceSelectionSrc.value,
        targetProcessorCategory,
        targetProcessorType,
        targetProcessorId,
        mode,
      }
    } else { // show selection popup
      interfaceSelectionVisible.value = true
    }

    return

  } else {
    store.dispatch('Scenario/addProcessor',
        {
          mode, processorCategory, processorType,
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

const diagnoseInterfaceNodesSelectionFinish = (interfaceNodes) => {
  const targetNode = treeDataMap.value[targetModelId]
  console.log('endpointInterfaceIdsSelectionFinish', interfaceNodes, targetNode)

  store.dispatch('Scenario/addInterfacesFromDiagnose', {
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
  const title = `确定删除名为${node.name}的节点吗？`
  const context = '该节点的所有子节点都将被删除！'
  confirmToDelete(title, context, () => {
    store.dispatch('Scenario/removeNode', targetModelId);
    selectNode([], null)
  })
}

const disableNodeOrNot = () => {
  const node = treeDataMap.value[targetModelId]
  const action = node.disable ? '启用' : '禁用'

  Modal.confirm({
    okType: 'danger',
    title: `确定${action}名为${node.name}的节点吗？`,
    content: '将同时禁用步骤下的所有子步骤，禁用后该步骤及所有子步骤在场景测试运行时不会被执行，是否确定禁用？',
    okText: () => '确定',
    cancelText: () => '取消',
    onOk: async () => {
      await store.dispatch('Scenario/disableNodeOrNot', targetModelId);
      selectNode([], null)
    },
    onCancel() {
      console.log('Cancel');
    },
  });
};

const clearMenu = () => {
  console.log('clearMenu')
  contextNode.value = ref({})
}

async function onDrop(info: DropEvent) {
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const pos = info.node.pos.split('-');
  let dropPosition = info.dropPosition - Number(pos[pos.length - 1]);

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

const interfaceImportFromCurlFinish = (content: string) => {
  const targetNode = treeDataMap.value[targetModelId]
  store.dispatch('Scenario/importCurl', {
    content: content, targetId: targetNode.id
  }).then((newNode) => {
    console.log('importUrl successfully', newNode)
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.id, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })

  interfaceSelectionVisible.value = false
}

const interfaceImportFromCurlCancel = () => {
  interfaceSelectionVisible.value = false
}

const endpointCaseSelectionFinish = (interfaceNodes: any) => {
  const targetNode = treeDataMap.value[targetModelId]
  console.log('endpointCaseSelectionFinish', interfaceNodes, targetNode)

  store.dispatch('Scenario/addInterfacesFromCase', {
    selectedNodes: interfaceNodes,
    targetId: targetNode.id,
  }).then((newNode) => {
    console.log('addInterfaces successfully', newNode)

    interfaceSelectionVisible.value = false
    selectNode([newNode.id], null)
    expandOneKey(treeDataMap.value, newNode.parentId, expandedKeys.value) // expend new node
    setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value)
  })

  interfaceSelectionVisible.value = false
}

const endpointCaseSelectionCancel = () => {
  interfaceSelectionVisible.value = false
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
<style lang="less" scoped>
.scenario-tree-main {
  background: #ffffff;
}

.tree-filter {
  margin-top: 8px;
}

.plus-icon {
  margin: 0 12px 0 6px;
  cursor: pointer;
}

.prefix-icon {
  margin-right: 6px;
}

.method-tag {
  transform: scale(0.7);
  margin-left: -6px; // move to left to offset the space caused by transform
  margin-right: 3px;
}

// 给下面样式添加省略号
.tree-title {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .title {
    flex: 1;
    width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    display: inline-block;
  }

  .icon {
    width: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}

</style>
