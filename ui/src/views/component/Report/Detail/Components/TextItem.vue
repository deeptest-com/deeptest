<template>
    <div class="text-wrapper">
        <span :class="['text-label', labelClassName]" :style="labelStyle" v-if="label">{{ label }}</span>
        <span class="text-value">
            <template v-if="value">{{ value }}</template>
            <template v-if="$slots.value">
                <slot name="value"></slot>    
            </template>
        </span>    
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    value?: String
    label?: String
    labelStyle?: any
    labelClassName?: String
}>()
</script>
<style scoped lang="less">
.text-wrapper {
    display: flex;
    align-items: center;
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
        display: inline-flex;
        align-items: center;
        

        &:before {
            content: '';
            display: block;
            margin-right: 7px;
            border-radius: 50%;
        } 
        
        &.success:before {
            width: 6px;
            height: 6px;
            background-color: #04C495;
        }

        &.failed:before {
            width: 6px;
            height: 6px;
            background-color: #F63838;
        }

        &.notest:before {
            width: 6px;
            height: 6px;
            background-color: rgba(0, 0, 0, 0.28);
        }
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