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
            :showIcon="false"
            :expandAction="false"
            :auto-expand-parent="autoExpandParent"
            :expandedKeys="expandedKeys"
            v-model:selectedKeys="selectedKeys"
            @drop="onDrop"
            @dragstart="onDragstart"
            @dragend="onDragEnd"
            @expand="onExpand"
            @select="selectNode"
            :tree-data="treeDataNeedRender"
            :replace-fields="replaceFields">
          <template #switcherIcon>
            <CaretDownOutlined/>
          </template>
          <template #title="{dataRef}">
            <div class="tree-title"
                 :class="{
                  'dp-tree-border':showLineScenarioType.includes(dataRef?.entityType),
                  'dp-tree-if':dataRef?.entityType === 'processor_logic_if',
                  'dp-tree-else':dataRef?.entityType === 'processor_logic_else',
                  'dp-tree-if-has-else':dataRef?.entityType === 'processor_logic_if' && checkIfHasElse(dataRef),
                 }"
                 :draggable="dataRef.id === -1">
              <div class="title" :class="[dataRef.disable ? 'dp-disabled' : '']">
                <!-- 标题前缀 -->
                <span class="prefix-icon"
                      v-if="!['processor_interface_default'].includes(dataRef?.entityType)">
                  <IconSvg v-if="DESIGN_TYPE_ICON_MAP[dataRef.entityType]"
                           :type="DESIGN_TYPE_ICON_MAP[dataRef.entityType]" class="prefix-icon-svg"/>
                </span>
                <!-- 逻辑：if/else -->
                <span class="prefix-logic"
                      v-if="['processor_logic_if','processor_logic_else'].includes(dataRef?.entityType)">
                      <a-typography-text
                          strong
                          style=" display: inline-block; text-align: left;margin-right: 4px;"
                          :type="dataRef.entityType === 'processor_logic_if' ? 'success' : 'danger'">{{
                          dataRef.entityType === 'processor_logic_if' ? 'IF' : 'ELSE'
                        }}</a-typography-text>
                </span>
                <!-- 请求：请求方法 -->
                <span class="prefix-req-method" v-if="dataRef.entityType === 'processor_interface_default'">
                  <a-tag class="method-tag" :color="getMethodColor(dataRef.method || 'GET', dataRef.disable)">{{
                      dataRef.method || "GET"
                    }}</a-tag>
                </span>
                <!-- 节点名称 -->
                <span class="title-text" :title="dataRef.name"
                      v-if="dataRef.name && needHandleShowName.includes(dataRef.entityType)">
                  {{ dataRef.name }}
                </span>
                <span class="title-text" :title="dataRef.name" v-else-if= "!needHandleShowName.includes(dataRef.entityType)">
                  {{
                    dataRef.name ? `${scenarioTypeMapToText[dataRef.entityType]} - ${dataRef.name}` : `${scenarioTypeMapToText[dataRef.entityType]}`
                  }}
                </span>
              </div>
              <div class="icon" v-if="dataRef.id > 0"
                   :style="dataRef.entityType === 'processor_logic_if'? {width:'60px'} : null">
                <a href="javascript:void (0)" type="link" @click.stop="() => {
                  addElse(dataRef)
                }" class="add-else-tag" v-if="showElse(dataRef)">+ else</a>
                <TreeMenu @selectMenu="selectMenu" :treeNode="dataRef">
                  <template #button>
                    <MoreOutlined style="min-width: 14px"/>
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
import {
  DESIGN_TYPE_ICON_MAP,
  menuKeyMapToProcessorCategory,
  scenarioTypeMapToBindText,
  scenarioTypeMapToText,
  needHandleShowName
} from "./config";
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
import {showLineScenarioType, onlyShowDisableAndDeleteTypes, loopIteratorTypes} from "./config";
import EditAndShow from "@/components/EditAndShow/index.vue";

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

const treeDataMap = computed<any>(() => store.state.Scenario.treeDataMap)
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult)
const scenarioProcessorIdForDebug = computed<number>(() => store.state.Scenario.scenarioProcessorIdForDebug)
const scenarioCount = computed<any>(() => store.state.Scenario.scenarioCount)

watch(scenarioCount, () => {
  console.log('watch scenarioCount', scenarioCount.value)
  selectedKeys.value = []
})

watch(treeData, () => {
  console.log('watch treeData1', treeData.value, treeDataNeedRender.value)

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

watch(scenarioProcessorIdForDebug, async (newVal) => {
  console.log('watch scenarioProcessorIdForDebug', scenarioProcessorIdForDebug.value)
  if (scenarioProcessorIdForDebug.value) {
    // just select node, not to fire getNode event again
    selectedKeys.value = [scenarioProcessorIdForDebug.value]
  }
}, {immediate: false})

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
  let mode = 'child';
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
  // 如果是 逻辑 else，则需要添加到父节点，即 if 节点下
  if (key === 'processor_logic_else') {
    // 如果已经存在 else 节点，则不允许添加
    // 另外目标节点已经有 else节点了，也不能再添加
    const repeat = checkIfHasElse(treeNode);
    if (repeat) {
      return;
    }
    mode = 'siblings';
  }

  const processorCategory = menuKeyMapToProcessorCategory[key];
  if (!processorCategory) return;
  const targetProcessorId = targetModelId
  const targetProcessorCategory = treeDataMap.value[targetModelId].entityCategory
  const targetProcessorType = treeDataMap.value[targetModelId].entityType
  addNode(mode, processorCategory, processorType,
      targetProcessorCategory, targetProcessorType, targetProcessorId);


}

/**
 * 判断当前的 if 节点是否有已经配对的 else 节点
 *
 * */

function showElse(dataRef) {
  if (dataRef?.entityType !== 'processor_logic_if') return false;
  let hasElse = true;
  const id = dataRef.id;
  const parentId = treeDataMap?.value[id]?.parentId;
  const children = treeDataMap?.value?.[parentId]?.children || [];
  const index = children?.findIndex(item => item.id === id);
  const nextNode = children?.[index + 1];
  if (nextNode?.entityType === 'processor_logic_else') {
    hasElse = false;
  }
  return hasElse;
}

/**
 * 检测 当前 if 节点 是否有匹配的 else 节点
 * */
function checkIfHasElse(node) {
  // 如果已经存在 else 节点，则不允许添加
  let exist = false;
  const parentId = node?.parentId;
  const children = treeDataMap?.value?.[parentId]?.children;
  const currentIndex = children?.findIndex(item => item.id === node.id);
  const nextNode = children?.[currentIndex + 1];
  if (nextNode?.entityType === 'processor_logic_else') {
    exist = true;
  }
  return exist;
}

/**
 * 检测 当前 if 节点 是否有匹配的 else 节点，用户控制 if 不能拖动到自己匹配的 else 节点下
 * */
function checkIfElseIsMatch(ifNode, elseNode) {
  let match = false;
  const parentId = ifNode?.parentId;
  const children = treeDataMap?.value?.[parentId]?.children;
  const ifNodeIndex = children?.findIndex(item => item.id === ifNode.id);
  const nextNode = children?.[ifNodeIndex + 1];
  if (nextNode?.entityType === 'processor_logic_else') {
    if (nextNode?.id === elseNode?.id) {
      match = true;
    }
  }
  return match;
}


const addElse = (treeNode) => {
  targetModelId = treeNode?.id;
  if (!targetModelId) return;
  // 另外目标节点已经有 else节点了，也不能再添加
  const repeat = checkIfHasElse(treeNode);
  if (repeat) {
    message.warning('已经存在 else 节点，不允许再添加');
    return;
  }

  const targetProcessorId = targetModelId
  const targetProcessorCategory = treeDataMap.value[targetModelId].entityCategory
  const targetProcessorType = treeDataMap.value[targetModelId].entityType
  const mode = 'siblings';
  addNode(mode, 'processor_logic', 'processor_logic_else',
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
          name: ''
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
  const name = node.name? scenarioTypeMapToText[node.entityType]+"-"+node.name : scenarioTypeMapToText[node.entityType];
  const title = `确定删除名为${name}的节点吗？`
  let context = node.entityCategory === "processor_interface" ? '' : '该节点的所有子节点都将被删除！'

  // 如果是 if 节点，则需要判断是否有 else 节点，如果有，则需要提示
  if (node.entityType === 'processor_logic_if') {
    const parentNode = treeDataMap.value[node.parentId];
    const currentIndex = parentNode?.children?.findIndex(item => item.id === node.id);
    const nextNode = parentNode?.children?.[currentIndex + 1];
    if (nextNode && nextNode.entityType === 'processor_logic_else') {
      context = '该节点的所有子节点都将被删除, 且将同时删除关联的 Else处理器，Else处理器不可单独使用,删除后将无法恢复，是否确定删除？'
    }
  }
  confirmToDelete(title, context, () => {
    store.dispatch('Scenario/removeNode', targetModelId);
    selectNode([], null)
  })
}

const disableNodeOrNot = () => {
  const node = treeDataMap.value[targetModelId]
  const action = node.disable ? '启用' : '禁用';
  const content = node.disable ? '将同时启用该步骤下的所有子步骤，是否确定启用该步骤？' : '禁用后该步骤及所有子步骤在场景测试运行时不会被执行，是否确定禁用？';
  const name = node.name? scenarioTypeMapToText[node.entityType]+"-"+node.name : scenarioTypeMapToText[node.entityType];
  Modal.confirm({
    okType: 'danger',
    title: `确定${action}名为【${name}】的场景步骤吗？`,
    content: content,
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
  const event = info.event;
  // 源节点
  const dragNodeInfo = info.dragNode.dataRef;
  // 目标节点
  const dropNodeInfo = info.node.dataRef;

  const pos = info.node.pos.split('-');
  let dropPosition = info.dropPosition - Number(pos[pos.length - 1]);

  if (isInterface(treeDataMap.value[dropKey].processorCategory) && dropPosition === 0) dropPosition = 1
  console.log(dragKey, dropKey, dropPosition);

  if (elseNodeRef.value) {
    elseNodeRef.value.hidden = false;
  }


  // 且 else节点需要拖动到 if节点后
  // 0  表示移动到目标节点的子节点
  //-1 表示移动到目标节点的前面，
  // 1  表示移动到目标节点的后面
  // else 节点不能移动到非 if节点后
  if (dragNodeInfo?.entityType === 'processor_logic_else' && dropNodeInfo?.entityType !== 'processor_logic_if') {
    message.warning('else 节点只能拖动到 If节点后');
    event.preventDefault();
    return;
  } else if (dragNodeInfo?.entityType === 'processor_logic_else' && dropNodeInfo?.entityType === 'processor_logic_if') {
    // 另外目标节点已经有 else节点了，也不能再添加
    const repeat = checkIfHasElse(dropNodeInfo);
    if (repeat) {
      message.warning('已经存在 else 节点，不允许再添加');
      return;
    }
    dropPosition = 1;
  }
  // 非 else 节点不能移动到 if节点后
  else if (dragNodeInfo?.entityType !== 'processor_logic_else' && dropNodeInfo?.entityType === 'processor_logic_if' && dropPosition === 1) {
    message.warning('非 else 节点不能拖动到 If节点后');
    event.preventDefault();
    return;
  }
  //  if 节点不能移动到 自己匹配的 else 节点下
  else if (dragNodeInfo?.entityType === 'processor_logic_if' && dropNodeInfo?.entityType === 'processor_logic_else' && dropPosition === 0) {
    if (checkIfElseIsMatch(dragNodeInfo, dropNodeInfo)) {
      message.warning('if 节点不能移动到 自己匹配的 else 节点下');
      event.preventDefault();
      return;
    }
  }
  // 任何节点都不能移动到 else节点前
  else if (dropNodeInfo?.entityType === 'processor_logic_else' && dropPosition === -1) {
    message.warning('任何节点都不能拖动到 Else 节点前');
    event.preventDefault();
    return;
  }
  // 跳出迭代不能 移动到 非迭代场景节点 下
  else if (dragNodeInfo?.entityType === 'processor_loop_break') {
    if (!loopIteratorTypes.includes(dropNodeInfo?.entityType)) {
      event.preventDefault();
      message.warning('跳出循环只能移动到循环迭代器里');
      return;
    }
    if (dropPosition !== 0) {
      event.preventDefault();
      message.warning('跳出循环只能移动到循环迭代器里');
      return;
    }
  }
  // 以下是叶子节点，不能移动到叶子节点下
  else if (onlyShowDisableAndDeleteTypes.includes(dropNodeInfo?.entityType) && dropPosition === 0) {
    message.warning('不能移动该改场景下');
    event.preventDefault();
    return;
  } else {
    console.log('else 其他拖动场景');
  }


  store.dispatch('Scenario/moveNode', {dragKey: dragKey, dropKey: dropKey, dropPos: dropPosition}).then(
      (result) => {
        if (result) {
          selectNode([dragKey], null);
          expandOneKey(treeDataMap.value, dragKey, expandedKeys.value) // expend parent node
          setExpandedKeys('scenario', treeData.value[0].scenarioId, expandedKeys.value);
        }
      }
  )
}

watch(() => {
  return expandedKeys.value
}, (value, oldValue, onCleanup) => {
  console.log('expandedKeys', value, oldValue)
})


const elseNodeRef: any = ref(null);

/**
 * 开始拖拽，阻止某些节点不让多动
 * */
function onDragstart({event, node}) {
  const nodeInfo = node.dataRef;
  // else节点 只能由 if 节点拖动带过去
  if (nodeInfo?.entityType === 'processor_logic_if' && checkIfHasElse(nodeInfo)) {
    event.target.parentNode.nextSibling.hidden = false
    elseNodeRef.value = event?.target?.parentNode?.nextSibling;
    if (elseNodeRef.value) {
      elseNodeRef.value.hidden = true;
    }
  }
}

function onDragEnd() {
  if (elseNodeRef.value) {
    elseNodeRef.value.hidden = false;
  }
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
  console.log('scenario tree onMounted')
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

  :deep(.ant-tree-child-tree-open:has(.tree-title.dp-tree-border)) {
    position: relative;

    &:before {
      content: '';
      position: absolute;
      top: 0;
      left: 11px;
      height: 100%;
      width: 1px;
      background: #f0f0f0;
    }
  }

  :deep(.ant-tree-treenode-switcher-close .tree-title.dp-tree-else  .prefix-icon) {
    //visibility: hidden;
  }
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
  height: 24px;

  .title {
    flex: 1;
    width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    //display: inline-block;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .icon {
    width: 16px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }

  .prefix-req-method {

  }

  .add-else-tag {
    margin-right: 6px;
    color: #1890ff;
    cursor: pointer;
    //transform: scale(0.8);
  }


  .draggable {
    width: 100px;
    height: 100px;
    background-color: lightblue;
    margin: 10px;
    cursor: grab;
  }

}

</style>
