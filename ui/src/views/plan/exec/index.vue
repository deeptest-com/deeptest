<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ execResult.name || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo :items="execResult.basicInfoList || {}" :show-operation="false" :scene="scene" />
            <StatisticTable :exec-status="execResult.progressStatus" :scene="scene" :data="execResult.statisticData" />
            <Progress :exec-status="execResult.progressStatus" v-if="scene !== ReportDetailType.QueryDetail" :percent="execResult.progressValue || 10"
                @exec-cancel="execCancel" />
            <template v-for="scenarioReportItem in execResult.scenarioReports" :key="scenarioReportItem.id">
                <ScenarioCollapsePanel :show-scenario-info="showScenarioInfo" :expand-active="expandActive"
                    :record="scenarioReportItem">
                    <template #endpointData>
                        <EndpointCollapsePanel v-if="scenarioReportItem.requestLogs.length > 0"
                            :recordList="scenarioReportItem.requestLogs" />
                    </template>
                </ScenarioCollapsePanel>
            </template>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, computed, watch } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel, Progress } from '@/views/component/Report/Components';

import { StateType as PlanStateType } from "../store";
import { ReportDetailType } from '@/utils/enum';
import settings from "@/config/settings";
import bus from "@/utils/eventBus";

import { useExec } from '../hooks/exec';

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scenarioExpandActive: boolean
    showScenarioInfo: boolean
    scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
    reportId?: number
}>();

const emits = defineEmits(['onClose']);

const store = useStore<{ Plan: PlanStateType }>();
const execResult = computed<any>(() => store.state.Plan.execResult);
const expandActive = ref(props.scenarioExpandActive || false);
const { execCancel, execStart, OnWebSocketMsg, onWebSocketConnStatusMsg } = useExec();

function onClose() {
    emits('onClose');
}

watch(props, (val) => {
    if (val.drawerVisible) {
        execStart();
        bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
        bus.on(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
    } else {
        execCancel();
        bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
        bus.off(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
    }
}, {
    immediate: true,
    deep: true
});
</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}
</style>