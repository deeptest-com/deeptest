<template>
    <TableFilter :show-operation="false" @handle-filter="handleFilter" />
    <a-table 
        row-key="id"
        :columns="columns" 
        :data-source="list"
        :pagination="{
            ...pagination,
            onChange: handlePageChanged
        }"
    />
</template>
<script lang="ts" setup>
import { reactive, computed, defineProps, watch } from 'vue';
import { useStore } from 'vuex';

import TableFilter from "@/views/component/Report/List/TableFilter.vue";

import { StateType as ReportStateType } from '@/views/component/Report/store';
import { StateType as PlanStateType } from '../store';  

const props = defineProps<{
    showReportList: Boolean
}>();

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

const store = useStore<{ Plan: PlanStateType, Report: ReportStateType }>();
const list = computed<any[]>(() => store.state.Report.listResult.list);
const currPlan = computed<any>(() => store.state.Plan.currPlan);
let formState = reactive({});
let pagination = computed(() => store.state.Report.listResult.pagination);

function handleFilter(params) {
    formState = params;
    refreshList({});
}

function handlePageChanged(page) {
    pagination.value.current = page;
    refreshList({ page, pageSize: pagination.value.pageSize });
}

function refreshList(params: any) {
    store.dispatch('Report/list', {
        page: pagination.value.current,
        pageSize: pagination.value.pageSize,
        planId: currPlan.value.id,
        ...formState,
        ...params
    })
    console.log('get test-plan reportList ---- [tableFilter] paramsData', { ...formState, ...params });
}

watch(() => {
    return props.showReportList;
}, val => {
    console.log(val);
    if (val) {
        refreshList({});
    }
})
</script>
  
  