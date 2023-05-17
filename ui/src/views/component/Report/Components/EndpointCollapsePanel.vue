<template>
    <div class="endpoint-collapse-panel">
        <a-collapse v-model:activeKey="activeKey" @change="changeActivekey">
            <template v-for="item in recordList" :key="item.id">
                <EndpointCollapse :collapseKey="item.id" :endpointData="item" @query-detail="handleQueryDetail"></EndpointCollapse>
            </template>
        </a-collapse>
    </div>
    <ResponseDrawer 
        :data="currRespDetail"
        :response-drawer-visible="logResponseDetailVisible"
        @onClose="logResponseDetailVisible = false" />
</template>
<script lang="ts" setup>
import { ref, defineProps } from 'vue';

import ResponseDrawer from '@/views/component/Report/Response/index.vue';
import EndpointCollapse from './EndpointCollapse.vue';

defineProps(['recordList'])

const activeKey = ref([]);
const currRespDetail = ref({});
// 接口响应详情
const logResponseDetailVisible = ref(false);

const changeActivekey = (key: string) => {
    console.log(key);
};

const handleQueryDetail = ({ resContent, reqContent }) => {
    console.log(resContent, reqContent);
    currRespDetail.value = { resContent, reqContent };
    logResponseDetailVisible.value = true;
}

</script>
<style scoped lang="less">
.endpoint-collapse-panel {
    :deep(.ant-collapse) {
        border: 0;
        background-color: #fbfbfb;
    }
}
</style>