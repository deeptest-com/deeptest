<template>
    <div class="endpoint-collapse-panel">
        <a-collapse v-model:activeKey="activeKey" @change="changeActivekey">
            <template v-for="item in recordList" :key="item.requestId">
                <EndpointCollapse :collapseKey="item.requestId" :endpointData="item"></EndpointCollapse>
            </template>
        </a-collapse>
    </div>
</template>
<script lang="ts" setup>
import { ref, defineProps } from 'vue';
import EndpointCollapse from './EndpointCollapse.vue';

defineProps(['recordList'])

const activeKey = ref([]);

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