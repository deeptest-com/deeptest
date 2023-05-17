<template>
  <a-modal
      class="associate-scenario-modal"
      title="关联测试场景"
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

const props = defineProps<{
  visible: Boolean,
}>();

const emits = defineEmits(['cancal', 'ok']);

function handleCancel() {
  emits('cancal');
}

const selectedKeys = ref<any[]>([]); // Check here to configure the default column
function handleChange(keys: any[]) {
  selectedKeys.value= keys;
}

async function ok() {
  emits('ok',selectedKeys.value);
}

</script>
<style scoped lang="less">
:deep(.ant-modal.associate-scenario-modal) {
  top: 50%;
  transform: translateY(-50%);
}

</style>
