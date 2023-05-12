<template>
    <div class="progress-container">
        <a-progress :percent="percent || 30" :showInfo="false" />
        <a-button class="exec-cancel" type="default" @click="handleExecCancel">终止执行</a-button>
    </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

defineProps<{
    percent: number
}>();

const emits = defineEmits(['execCancel']);

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