<template>
    <div class="request">
        <ConBoxTitle title="请求URL"></ConBoxTitle>
        <div class="request-url">
            {{ data?.url || '' }}
        </div>
        <ConBoxTitle title="Header"></ConBoxTitle>
        <div class="request-headers">
            <a-table
                :data="data?.headers || []"
                :columns="columns" />
        </div>
        <ConBoxTitle title="Body"></ConBoxTitle>
        <div class="request-body">
            <span>Body类型: {{ data?.bodyLang || '' }}</span>
        </div>
    </div>
</template>
<script setup lang="ts">
import { defineProps, watch } from 'vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';

const props = defineProps({
    data: {
        type: Object
    }
});

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Value',
        dataIndex: 'value',
    }
];

watch(() => {
    return props.data;
}, val => {
    console.log(val);
}, {
    immediate: true,
    deep: true
})
</script>
<style scoped lang="less">
.request {
    padding-top: 20px;

    .request-url,.request-headers,.request-body {
        padding: 10px;
    }
}
</style>