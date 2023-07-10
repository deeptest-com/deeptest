<template>
  <div class="case-design-main">
    <a-modal title="设计用例" width="1000px" height="90%"
             :visible="visible"
             @ok="onFinish"
             @cancel="onCancel">

      <EndpointDebug />

    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import {defineProps} from 'vue';
import {Form} from "ant-design-vue";
import EndpointDebug from '../Debug/index.vue';

const useForm = Form.useForm;

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  model: {
    required: true,
    type: Object,
  },
  onFinish: {
    type: Function,
    required: true,
  },
  onCancel: {
    type: Function,
    required: true,
  },
})

const finish = () => {
  console.log('finish')
}

const cancel = () => {
  console.log('cancel')
}

</script>

<style lang="less" scoped>
.case-design-main {

}
</style>
