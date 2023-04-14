<template>
    <div class="con-box-title" :style="backgroundStyle">{{ title }}</div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';
defineProps<{
    title: string
    backgroundStyle?: string
}>()
</script>
<style scoped lang="less">
.con-box-title {
    padding-left: 8px;
    width: 100%;
    height: 32px;
    display: flex;
    align-items: center;
    text-align: left;
    font-size: 14px;
    color: rgba(0, 0, 0, 0.85);
    position: relative;
    font-weight: bold;

    &:before {
        content: '';
        position: absolute;
        width: 3px;
        height: 14px;
        background: #447DFD;
        top: 9px;
        left: 0;
    }
}
</style>