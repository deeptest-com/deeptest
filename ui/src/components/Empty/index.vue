<template>
     <div class="empty-container">
        <img src="@/assets/images/empty.png" alt="">
        <span>{{ text || '暂无数据' }}</span>
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    text?: String
}>();   
</script>

<style scoped lang="less">
.empty-container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px;

    img {
        width: 90px;
        height: 90px;
        margin-bottom: 8px;
    }

    span {
        font-size: 12px;
        line-height: 20px;
        text-align: center;
        color: rgba(0, 0, 0, 0.46);
    }
}
</style>