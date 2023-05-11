<template>
    <TableFilter :show-operation="false" @get-list="refreshList" />
    <a-table :columns="columns" :data-source="data" />
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import TableFilter from "@/views/component/Report/List/TableFilter.vue";

const columns = [
    {
        title: '编号',
        dataIndex: 'serialNumber',
        slots: { customRender: 'serialNumber' },
        width: 120
    },
    {
        title: '测试通过率',
        dataIndex: 'interfacePassRate',
        width: 120,
    },
    {
        title: '场景通过率',
        dataIndex: 'interfacePassRate',
        width: 120,
    },
    {
        title: '执行耗时',
        dataIndex: 'duration',
        width: 80,
        slots: { customRender: 'duration' },
    },
    {
        title: '执行人',
        dataIndex: 'createUserName',
        width: 80,
    },
    {
        title: '执行时间',
        dataIndex: 'executionTime',
        width: 200,
        slots: { customRender: 'executionTime' },
    },
];

const data = [];

function refreshList(params: any) {
    console.log('get test-plan reportList ---- [tableFilter] paramsData', params);
}
</script>
  
  