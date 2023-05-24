<template>

  <a-modal
    title="请选择接口"
    :destroy-on-close="true"
    :mask-closable="false"
    :visible="true"
    :onCancel="onCancel"
    wrapClassName="modal-tree-selection"
    width="1000px">
    <div class="interface-selection-main">
      <div class="left tree">
        <Tree :selectCategory="selectCategory"/>
      </div>
      <div class="right">
        <List :selectInterface="onSelectInterface"></List>
      </div>
    </div>
    <template #footer>
      <a-button @click="onCancel">取消</a-button>
      <a-button @click="onSubmit" type="primary">确定</a-button>
    </template>
  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, ref, watch} from "vue";
import {listServe} from "@/services/serve";

import Tree from "./tree.vue"
import List from "./list.vue"

const props = defineProps({
  onFinish: {
    type: Function,
    required: true,
  },
  onCancel: {
    type: Function,
    required: true,
  },
})

const categoryId = ref(0)
const interfaceIds = ref([])

const selectCategory = async (id) => {
  console.log('selectCategory', id)
  categoryId.value = id;
}

const onSelectInterface = async (ids) => {
  console.log('onSelectInterface', ids)
  interfaceIds.value = ids
}

const onSubmit = () => {
  console.log('onSubmit')
  props.onFinish(interfaceIds.value)
}

const onCancel = () => {
  console.log('onCancel')
  props.onCancel()
}

</script>

<style scoped lang="less">
.modal-tree-selection {
  .ant-modal-body {
    padding-top: 5px;
  }
}
</style>

<style lang="less" scoped>
.interface-selection-main {
  display: flex;
  .left {
    width: 260px;
  }
  .right {
    flex: 1;
    margin-left: 16px;
  }
}
</style>
