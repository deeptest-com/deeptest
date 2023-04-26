<template>
    <div class="text-wrapper">
        <span class="text-label" v-if="label">{{ label }}</span>
        <span class="text-value">{{ value }}</span>
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    value: string
    label?: string
}>()
</script>
<style scoped lang="less">
.text-wrapper {
    .text-label {
        display: inline-block;
        width: 68px;
        font-family: 'PingFang SC';
        font-style: normal;
        font-weight: 400;
        font-size: 14px;
        line-height: 22px;
        /* identical to box height, or 157% */
        /* 中性色/文字/强调-标题 */

        color: rgba(0, 0, 0, 0.85);
        margin-right: 30px;
    }

    .text-value {
        font-family: 'PingFang SC';
        font-style: normal;
        font-weight: 400;
        font-size: 14px;
        line-height: 22px;
        /* identical to box height, or 157% */
        /* 中性色/文字/次强调-正文 */
        color: rgba(0, 0, 0, 0.65);
    }
}
</style>