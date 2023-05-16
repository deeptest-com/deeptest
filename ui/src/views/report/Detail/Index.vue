<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ detailResult.name || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo  :items="detailResult.basicInfoList" :scene="scene" />
            <StatisticTable :scene="scene" :data="detailResult.statisticData" />
            <Progress :exec-status="execResult.progressStatus" v-if="scene !== ReportDetailType.QueryDetail" :percent="60" @exec-cancel="execCancel" />
            <template v-for="scenarioReportItem in detailResult.scenarioReports" :key="scenarioReportItem.id">
                <ScenarioCollapsePanel :show-scenario-info="showScenarioInfo" :expand-active="expandActive"
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
import { defineProps, defineEmits, ref, computed } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel, Progress } from '@/views/component/Report/Components';

import { StateType as ReportStateType } from "../store";
import { StateType as PlanStateType } from '@/views/plan/store';
import { ReportDetailType } from '@/utils/enum';

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scenarioExpandActive: boolean
    showScenarioInfo: boolean
    scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
    reportId?: number
}>();

const emits = defineEmits(['onClose', 'execCancel']);

const store = useStore<{ Report: ReportStateType, Plan: PlanStateType }>();
const execResult = computed<any>(() => store.state.Plan.execResult);
const detailResult = computed<any>(() => store.state.Report.detailResult);
const expandActive = ref(props.scenarioExpandActive || false);

function onClose() {
    emits('onClose');
}

function execCancel() {
    emits('execCancel');
}

</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}
</style>
