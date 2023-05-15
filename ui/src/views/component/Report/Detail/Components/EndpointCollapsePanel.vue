<template>
    <div class="endpoint-collapse-panel">
        <a-collapse v-model:activeKey="activeKey" @change="changeActivekey">
            <template v-for="item in recordList" :key="item.id">
                <EndpointCollapse :collapseKey="item.id" :endpointData="item" @query-detail="logResponseDetailVisible = true"></EndpointCollapse>
            </template>
        </a-collapse>
    </div>
    <ResponseDrawer 
        :response-drawer-visible="logResponseDetailVisible"
        @onClose="logResponseDetailVisible = false" />
</template>
<script lang="ts" setup>
import { ref, defineProps } from 'vue';

import ResponseDrawer from '@/views/component/Report/Response/index.vue';
import EndpointCollapse from './EndpointCollapse.vue';

defineProps(['recordList'])

const activeKey = ref([]);
// 接口响应详情
const logResponseDetailVisible = ref(false);

const changeActivekey = (key: string) => {
    console.log(key);
};

</script>
<style scoped lang="less">
.endpoint-collapse-panel {
    :deep(.ant-collapse) {
        border: 0;
        background-color: #fbfbfb;
    }
}
</style>