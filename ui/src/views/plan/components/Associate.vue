<template>
    <a-modal 
        class="associate-scenario-modal"
        title="关联测试场景" 
        :visible="associateModalVisible" 
        :closable="true"
        @cancel="handleCancel" 
        @ok="onOk"
        width="1000px">
        <ScenarioList :show-scenario-operation="false" @select-row-keys="handleSelectRowKeys" />
    </a-modal>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref } from 'vue';
import ScenarioList from './ScenarioList.vue';

// import { StateType as ScenarioStateType } from '@/views/scenario/store';

defineProps<{
    associateModalVisible: Boolean
}>();

const emits = defineEmits(['onCancel', 'onOk']);
const selectedScenarioIds = ref<any[]>([]);

function handleSelectRowKeys(value: any[]) {
    selectedScenarioIds.value = value;
}

function handleCancel() {
    emits('onCancel');
}

function onOk() {
    emits('onOk', selectedScenarioIds);
}
</script>
<style scoped lang="less">
:deep(.ant-modal.associate-scenario-modal) {
    top: 50%;
    transform: translateY(-50%);
}

</style>