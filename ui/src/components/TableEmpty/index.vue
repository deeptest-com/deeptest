<template>
    <!-- 表格无数据时的 展示 使用ConfigProvider定义组件空状态 -->
    <a-config-provider>
        <template #renderEmpty>
            <div v-if="!loading" class="empty-container">
                <img src="@/assets/images/empty.png" alt="">
                <span>{{ text || '暂无数据' }}</span>
            </div>
            <div v-else class="empty-loading">
                <a-spin></a-spin>
            </div>
        </template>
        <slot name="content"></slot>
    </a-config-provider>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    text?: String;
    loading?: Boolean;
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

.empty-loading {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 40px;
}
</style>