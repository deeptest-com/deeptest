<template>

  <a-modal
    title="请选择接口"
    :destroy-on-close="true"
    :mask-closable="false"
    :visible="true"
    :onCancel="onCancel"
    :footer="null"
    wrapClassName="modal-tree-selection"
  >

  <div class="tree-main">
    <div class="toolbar">
      <div class="tips">
        <a-button type="primary" size="small" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
        <a-button @click="() => onCancel()" size="small" class="dp-btn-gap">取消</a-button>
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
          :tree-data="treeData"
          :replace-fields="replaceFields"
          show-icon
          @expand="expandNode"
          @check="checkNode"
          checkable

          v-model:checkedKeys="checkedKeys"
          v-model:expandedKeys="expandedKeys"

          class="interface-selection-tree"
      >
        <template #title="slotProps">
          <span>{{ slotProps.name }}</span>
        </template>

        <template #icon="slotProps">
          <FolderOutlined v-if="slotProps.isDir && !slotProps.expanded"/>
          <FolderOpenOutlined v-if="slotProps.isDir && slotProps.expanded"/>
          <FileOutlined v-if="!slotProps.isDir"/>
        </template>
      </a-tree>
    </div>

  </div>

  </a-modal>

</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, PropType, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Interface} from "@/views/interface/data";
import throttle from "lodash.debounce";
import {expandAllKeys, getNodeMap} from "@/services/tree";
import {useStore} from "vuex";
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";
import {StateType} from "@/views/interface/store";
import {StateType as ProjectStateType} from "@/store/project";
import {getExpandedKeys, setExpandedKeys} from "@/utils/cache";
import {isInArray} from "@/utils/array";

const {t} = useI18n();

const props = defineProps<{
  onCancel: Function,
  onFinish: Function,
}>()

const store = useStore<{ Interface: StateType, ProjectData: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectData.currProject);
const treeData = computed<any>(() => store.state.Interface.treeData);
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const queryTree = async () => {
  await store.dispatch('Interface/loadInterface');
}
queryTree();

const replaceFields = {key: 'id', title: 'name'};
let expandedKeys = ref<number[]>([]);
let checkedKeys = ref<number[]>([])
let isExpand = ref(false);

let tree = ref(null)
const expandNode = (keys: string[], e: any) => {
  console.log('expandNode', keys[0], e)

  setExpandedKeys(currProject.value.id, expandedKeys.value)
}

const checkNode = (keys, e) => {
  console.log('checkNode', checkedKeys)
}

const treeDataMap = {}
const getNodeMapCall = throttle(async () => {
  getNodeMap(treeData.value[0], treeDataMap)
}, 60)

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

watch(treeData, () => {
  console.log('watch', treeData)
  getNodeMapCall()

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

const onSubmit = async () => {
  console.log('onSubmit', checkedKeys)

  const selectedNodes = [] as any[]
  Object.keys(treeDataMap).forEach((id, index) => {
    if (!treeDataMap[id].isDir && isInArray(+id, checkedKeys.value)) {
      selectedNodes.push(treeDataMap[id])
    }
  })

  props.onFinish(selectedNodes);
}

onMounted(() => {
  console.log('onMounted')
})
onUnmounted(() => {
  console.log('onUnmounted')
})

</script>

<style lang="less">
.modal-tree-selection {
  .ant-modal-body {
    padding-top: 5px;
  }
}
</style>

<style lang="less" scoped>
.tree-main {
  .toolbar {
    display: flex;
    height: 32px;

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