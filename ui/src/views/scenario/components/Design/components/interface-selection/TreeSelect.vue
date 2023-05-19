<template>
  <div class="interface-tree-main dp-tree">
    <a-select
        v-model:value="serveId"
        :placeholder="'请选择服务'"
        :bordered="true"
        style="width: 150px;margin-right: 16px;"
        @change="selectServe">
      <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
    </a-select>
    <a-tree-select
        v-model:value="categoryId"
        style="width: 300px;"
        allow-clear
        :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
        :tree-data="treeDataCategory"
        placeholder="请选择分类"
        tree-default-expand-all
    />
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
  // console.log('loadCategory', serveId.value)
  const response = await loadCategory('endpoint');
  if (response.code === 0) {
    treeDataCategory.value = [response.data]
    // selectNode([response.data.id], null)
    // treeDataMapCategory = {}
    // getNodeMap(treeDataCategory.value, treeDataMapCategory)
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
  //width: 600px;
  display: flex;
  align-items: center;
  margin-right: 16px;
  margin-bottom: 16px;


}
</style>
