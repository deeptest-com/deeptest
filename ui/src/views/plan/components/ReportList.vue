<template>
    <TableFilter :show-operation="false" @handle-filter="handleFilter" />
    <a-table
        row-key="id"
        :loading="loading"
        :columns="columns"
        :data-source="list"
        :pagination="{
            ...pagination,
            onChange: handlePageChanged
        }"
    >
        <template #duration="{ record }">
            <span v-html="formatWithSeconds(record.duration)"></span>
        </template>
        <template #startTime="{ record }">
            {{ momentUtc(record.startTime) }} ~  {{ momentUtc(record.endTime) }}
        </template>
        <template #operation="{ record }">
            <a-button type="primary" @click="queryDetail(record.id)">查看报告</a-button>
        </template>
    </a-table>
    <ExecDetail
        title="测试报告详情"
        :show-scenario-info="true"
        :scenario-expand-active="false"
        :drawer-visible="detailDrawerVisible"
        :scene="ReportDetailType.QueryDetail"
        @on-close="detailDrawerVisible = false" />
</template>
<script lang="ts" setup>
import { reactive, computed, defineProps, watch, defineEmits, ref } from 'vue';
import { useStore } from 'vuex';

import { TableFilter } from "@/views/component/Report/components";
import ExecDetail from '@/views/report/Detail/index.vue';

import { StateType as ReportStateType } from '@/views/report/store';
import { StateType as PlanStateType } from '../store';
import { momentUtc, formatWithSeconds } from '@/utils/datetime';
import { ReportDetailType } from '@/utils/enum';

const props = defineProps<{
    showReportList: Boolean
}>();

const columns = [
    {
        title: '编号',
        dataIndex: 'serialNumber',
        slots: { customRender: 'serialNumber' },
    },
    {
        title: '测试通过率',
        dataIndex: 'interfacePassRate',
    },
    {
        title: '场景通过率',
        dataIndex: 'interfacePassRate',
    },
    {
        title: '执行耗时',
        dataIndex: 'duration',
        slots: { customRender: 'duration' },
    },
    {
        title: '执行人',
        dataIndex: 'createUserName',
    },
    {
        title: '执行时间',
        width: 340,
        dataIndex: 'startTime',
        slots: { customRender: 'startTime' },
    },
    {
        title: '查看报告',
        dataIndex: 'operation',
        slots: { customRender: 'operation' },
    }
];

const store = useStore<{ Plan: PlanStateType, Report: ReportStateType }>();
const list = computed<any[]>(() => store.state.Report.listResult.list);
const currPlan = computed<any>(() => store.state.Plan.currPlan);
const detailDrawerVisible = ref(false);
let formState = reactive({});
const loading = ref(false);
let pagination = computed(() => store.state.Report.listResult.pagination);

function handleFilter(params) {
    formState = params;
    refreshList({});
}

function handlePageChanged(page) {
    pagination.value.current = page;
    refreshList({ page, pageSize: pagination.value.pageSize });
}

async function refreshList(params: any) {
    loading.value = true;
    await store.dispatch('Report/list', {
        page: pagination.value.current,
        pageSize: pagination.value.pageSize,
        planId: currPlan.value.id,
        ...formState,
        ...params
    })
    loading.value = false;
    console.log('get test-plan reportList ---- [tableFilter] paramsData', { ...formState, ...params });
}

async function queryDetail(id) {
    await store.dispatch('Report/initReportDetail');
    store.dispatch('Report/get', id);
    detailDrawerVisible.value = true;
}

watch(() => props.showReportList, val => {
    if (val) {
        refreshList({});
    }
}, {
    immediate: true
})
</script>

