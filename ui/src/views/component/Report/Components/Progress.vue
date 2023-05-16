<template>
    <div class="progress-container">
        <a-progress :strokeColor="execStatusColorMap.get(execStatus)" :percent="percent || 30" :showInfo="false" />
        <a-button class="exec-cancel" type="default" @click="handleExecCancel">{{ execStatusTextMap.get(execStatus) }}</a-button>
    </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

defineProps<{
    percent: number
    execStatus: string
}>();

const emits = defineEmits(['execCancel']);
const execStatusTextMap = new Map([['in_progress', '终止执行'], ['end', '执行成功'], ['failed', '执行失败']]);
const execStatusColorMap = new Map([['in_progress', '#1890ff'], ['end', '#04C495'], ['failed', '#F63838']]);
const handleExecCancel = () => {
    emits('execCancel');
};
</script>

<style scoped lang="less">
.progress-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 20px;

    :deep(.ant-progress-bg) {
        height: 22px !important;
    }

    .exec-cancel {
        margin-left: 40px;
    }
}
</style>