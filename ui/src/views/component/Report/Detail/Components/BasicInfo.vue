<template>
    <div class="report-basicinfo">
        <TextItem label="测试计划" :value="basicInfo.logName" />
        <TextItem label="执行时间" :value="basicInfo.logTime" />
        <TextItem label="执行环境" :value="basicInfo.logEnv" />
        <TextItem label="执行人" :value="basicInfo.logExecutor" />
        <a-button type="primary" class="report-export">导出报告</a-button>
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';
import TextItem from './TextItem.vue';

defineProps<{
    basicInfo: any
}>();
</script>
<style scoped lang="less">
.report-basicinfo {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    position: relative;
    margin-bottom: 24px;

    .text-wrapper {
        width: 40%;
        margin-bottom: 4px;
    }

    .report-export {
        position: absolute;
        right: 0;
        top: 0;
    }
}
</style>