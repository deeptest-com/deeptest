<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ execResult.name || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo :basic-info="execResult.basicInfo || {}" :scene="scene" />
            <StatisticTable :scene="scene" :data="execResult.statisticData" />
            <template v-for="scenarioReportItem in execResult.scenarioReports" :key="scenarioReportItem.id">
                <ScenarioCollapsePanel :show-scenario-info="true" :expand-active="false"
                    :record="scenarioReportItem">
                    <template #endpointData>
                        <EndpointCollapsePanel v-if="scenarioReportItem.requestLogs.length > 0" :recordList="scenarioReportItem.requestLogs" />
                    </template>
                </ScenarioCollapsePanel>
            </template>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, computed, watch } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel } from '@/views/component/Report/Components';

import { StateType as PlanStateType } from "../../store";

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
    reportId: number
}>();

const emits = defineEmits(['onClose']);

const store = useStore<{ Plan: PlanStateType }>();
const execResult = computed<any>(() => store.state.Plan.execResult);

function onClose() {
    emits('onClose');
}


watch(() => {
    return props;
}, val => {
    if (val.drawerVisible) {
        store.dispatch('Plan/getExecDetail', val.reportId);
    }
}, {
    immediate: true,
    deep: true
})
</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}
</style>