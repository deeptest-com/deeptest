<template>
  <a-modal
      class="associate-scenario-modal"
      title="关联测试计划"
      :visible="visible"
      :closable="true"
      @cancel="handleCancel"
      @ok="ok"
      width="1000px">
    <PlanList :linked="false" @select="handleChange"/>
  </a-modal>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref} from 'vue';
import PlanList from './PlanList.vue';
import {message} from "ant-design-vue";

const props = defineProps<{
  visible: Boolean,
}>();

const emits = defineEmits(['cancel', 'ok']);

function handleCancel() {
  emits('cancel');
}

const selectedKeys = ref<any[]>([]);
function handleChange(keys: any[]) {
  selectedKeys.value= keys;
}

async function ok() {
  if(selectedKeys.value.length === 0) {
    message.info('请选择测试场景');
    return;
  }
  emits('ok',selectedKeys.value);
}

</script>
<style scoped lang="less">
:deep(.ant-modal.associate-scenario-modal) {
  top: 50%;
  transform: translateY(-50%);
}

</style>
