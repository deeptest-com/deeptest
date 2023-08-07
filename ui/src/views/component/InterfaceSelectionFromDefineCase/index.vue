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
  
  const selectInterfaces = ref([])
  const onSelectInterfaces = (data) => {
    console.log('onSelectInterfaces', data)
    selectInterfaces.value = data
  }
  
  const onSubmit = (e) => {
    console.log('onSubmit')
    props.onFinish(selectInterfaces.value)
  }
  
  const onCancel = () => {
    console.log('onCancel')
    props.onCancel()
  }
  
  </script>
  
  <style lang="less">
  .modal-tree-selection {
    .ant-modal {
      .ant-modal-content {
        .ant-modal-body {
        }
      }
    }
  }
  </style>
  
  <style lang="less" scoped>
  .interface-selection-main {
  }
  </style>
  