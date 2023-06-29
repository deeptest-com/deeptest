<template>
  <div class="interface-tree-main dp-tree">
    <div class="toolbar">
      <div class="tips">
        <a-select
            v-model:value="serveId"
            :placeholder="'请选择服务'"
            :bordered="true"
            @change="selectServe"
            size="small"
            class="dp-no-border">
          <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
        </a-select>
      </div>

      <div class="buttons">
        <a-button @click="expandAll" type="link" class="dp-color-primary">
          <span v-if="!isExpand">展开</span>
          <span v-if="isExpand">收缩</span>
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

          v-model:selectedKeys="selectedKeys"
          v-model:expandedKeys="expandedKeys"
      >
        <template #title="slotProps">
          <span class="name-editor">
            {{ slotProps.name }}
          </span>
        </template>

        <template #icon="slotProps">
          <FolderOutlined v-if="!slotProps.isLeaf && !slotProps.expanded"/>
          <FolderOpenOutlined v-if="!slotProps.isLeaf && slotProps.expanded"/>
          <FileOutlined v-if="slotProps.isLeaf"/>
        </template>
      </a-tree>
    </div>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch} from "vue";

import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {CloseOutlined, FileOutlined, FolderOutlined, FolderOpenOutlined, CheckOutlined} from "@ant-design/icons-vue";

import {expandAllKeys, expandOneKey, getNodeMap} from "@/services/tree";
import {listServe} from "@/services/serve";
import {loadCategory} from "@/services/category";

const useForm = Form.useForm;

const {t} = useI18n();

const props = defineProps({
  selectCategory: {
    type: Function,
    required: true,
  }
})

watch(props, () => {
  console.log('watch props for reload', props)

}, {deep: true})


const serves = ref([] as any[]);
const serveId = ref(0)
const categoryId = ref(0)

const loadServe = async () => {
  listServe().then((json) => {
    serves.value = json.data.serves

    if (serves.value.length > 0) {
      serveId.value = serves.value[0].id
      loadCategoryByServe()
    }
  })
}
loadServe()

const selectServe = () => {
  console.log('selectServe', serveId.value)
  props.selectCategory('未分类') // TODO:

  loadCategoryByServe()
}

const treeDataCategory = ref([] as any[])
let treeDataMapCategory = {}
const loadCategoryByServe = async () => {
  console.log('loadCategory', serveId.value)

  const response = await loadCategory('endpoint');
  if (response.code === 0) {
    treeDataCategory.value = [response.data]

    selectNode([response.data.id], null)
    treeDataMapCategory = {}
    getNodeMap(treeDataCategory.value, treeDataMapCategory)
  }
}

const selectCategory = async (id) => {
  console.log('selectCategory', id)
  categoryId.value = id
  props.selectCategory(id)
}

const replaceFields = {key: 'id', title: 'name'};
let expandedKeys = ref<number[]>([]);
let selectedKeys = ref<number[]>([]);
let isExpand = ref(false);

let tree = ref(null)
const expandNode = (keys: string[], e: any) => {
  console.log('expandNode', keys[0], e)
}

const selectNode = (keys, e) => {
  console.log('selectNode', keys)
  if (keys.length === 0 && e) {
    selectedKeys.value = [e.node.dataRef.id] // cancel un-select
    return
  } else {
    selectedKeys.value = keys
  }
  props.selectCategory(selectedKeys.value[0])
}

const expandAll = () => {
  console.log('expandAll')
  isExpand.value = !isExpand.value
  expandedKeys.value = expandAllKeys(treeDataMapCategory, isExpand.value)
}

onMounted(() => {
  console.log('onMounted')
})
onUnmounted(() => {
  console.log('onUnmounted')
})

</script>

<style lang="less" scoped>
.interface-tree-main {
  .toolbar {
    display: flex;
    .tips {
      flex: 1;
      .ant-select {
        width: 100%;
      }
    }
    .buttons {
      width: 50px;
    }
  }

  .tree-panel {
    width: 260px;
    border-right: 1px solid #f0f0f0;
    height: calc(100vh - 140px);
    overflow-y:scroll;
  }

}
</style>
