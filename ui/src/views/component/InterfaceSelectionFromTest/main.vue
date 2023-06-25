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
        <Tree :selectInterfaces="onSelectInterfaces" />
    </div>

    <template #footer>
      <a-button @click="onCancel">取消</a-button>
      <a-button @click="onSubmit" type="primary">确定</a-button>
    </template>
  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, ref} from "vue";

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

const onSelectInterfaces = (data) => {
  console.log('onSelectInterfaces', data)
}

const onSubmit = (nodes) => {
  console.log('onSubmit', nodes)
  // props.onFinish(nodes)
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
}
</style>
