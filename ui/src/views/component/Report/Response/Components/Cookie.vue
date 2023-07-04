<template>
    <div class="cookies">
        <a-table :dataSource="data" :columns="columns"
                 :rowKey="(record,index) => {return index}"
                 :rowClassName="(record,index) => {return record.name==='' ? 'hidden' : ''}" />
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    data: Array<any>
}>();

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Value',
        dataIndex: 'value',
    },
    {
        title: 'Domain',
        dataIndex: 'domain',
    },
    {
        title: 'Path',
        dataIndex: 'path',
    },
    {
        title: 'Expires',
        dataIndex: 'expires',
    },
    {
        title: 'MaxAge',
        dataIndex: 'MaxAge',
    },
    {
        title: 'HttpOnly',
        dataIndex: 'httpOnly',
    },
    {
        title: 'Secure',
        dataIndex: 'secure',
    }
];
</script>
<style scoped lang="less">
.cookies {
    padding-top: 20px;
}
</style>
