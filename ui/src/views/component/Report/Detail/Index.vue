<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ detailResult.name || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo :basic-info="detailResult.basicInfo || {}" :scene="scene" />
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
import { defineProps, defineEmits, ref, computed, watch } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel, Progress } from './Components';

import { StateType as PlanStateType } from '@/views/plan/store';
import { StateType as GlobalStateType } from "@/store/global";
import { StateType as ReportStateType } from "../store";
import { StateType as ProjectSettingStateType } from '@/views/projectSetting/store';
import { ExecStatus } from "@/store/exec";
import settings from "@/config/settings";
import { WebSocket } from "@/services/websocket";
import { WsMsg } from "@/types/data";
import bus from "@/utils/eventBus";
import { getToken } from "@/utils/localToken";
import { ReportDetailType, WsMsgCategory } from '@/utils/enum';

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scenarioExpandActive: boolean
    showScenarioInfo: boolean
    scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
    reportId?: number
}>();

const emits = defineEmits(['onClose']);

const store = useStore<{ 
    Plan: PlanStateType,
    Global: GlobalStateType, 
    Exec: ExecStatus, 
    Report: ReportStateType,
    ProjectSetting: ProjectSettingStateType
 }>();

const currPlan = computed<any>(() => store.state.Plan.currPlan);
const execResult = computed<any>(() => store.state.Plan.execResult);
const detailResult = computed<any>(() => store.state.Report.detailResult);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);

const expandActive = ref(props.scenarioExpandActive || false);
const logTreeData = ref<any>([]);
const result = ref({} as any);
const logMap = ref({} as any);

const execStart = async () => {
    console.log('execStart');

    logTreeData.value = [];

    const data = {
        serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
        token: await getToken(),
        planId: currPlan.value && currPlan.value.id,
        envId: currEnvId.value
    }

    WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({ act: 'execPlan', planExecReq: data }));
};

const execCancel = () => {
    console.log('execCancel');
    const msg = { act: 'stop', execReq: { scenarioId: currPlan.value && currPlan.value.id } };
    WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
};

const OnWebSocketMsg = (data: any) => {
    if (!data.msg) return
    const wsMsg = JSON.parse(data.msg) as WsMsg
    console.log('--- WebsocketMsgEvent in exec info', wsMsg);
    if (wsMsg.category == 'result') { // update result
        result.value = wsMsg.data
        console.log('=====', result.value)
        return
    } else if (wsMsg.category === WsMsgCategory.InProgress ||  wsMsg.category === WsMsgCategory.End) { // update status
        execResult.value.progressStatus = wsMsg.category
        if (wsMsg.category === 'in_progress') {
            result.value = {};
        }
        return
    }

    const log: any = wsMsg.data;
    logMap.value[log.id] = log;

    if (log.parentId === 0) { // root
        logTreeData.value.push(log);
        if (logTreeData.value.find(e => e.id === log.id)) {
            const findIndex = logTreeData.value.find(e => e.id === log.id);
            logTreeData.value.splice(findIndex, 1, log);
        }
    } else {
        if (!logMap.value[log.parentId]) logMap.value[log.parentId] = {};
        if (!logMap.value[log.parentId].logs) logMap.value[log.parentId].logs = [];
        logMap.value[log.parentId].logs.push(log);
    }

    console.log('logMap Data--------', logMap.value);
    console.log('logTree Data--------', logTreeData.value);
    console.log('execResult Data -----', execResult);
};

// websocket 连接状态 查询
const onWebSocketConnStatusMsg = (data: any) => {
    console.log('join websocket room', data);
    if (!data.msg) {
        return;
    }
    const { conn }: any = JSON.parse(data.msg);
    execResult.value.progressStatus = conn === 'success' ? WsMsgCategory.InProgress : 'failed';
}

function onClose() {
    emits('onClose');
}
watch(() => {
    return props;
}, val => {
    if (val.scene === ReportDetailType.QueryDetail && val.drawerVisible) {
        store.dispatch('Report/get', val.reportId || 1);
    } else if (val.scene !== ReportDetailType.QueryDetail) {
        if (val.drawerVisible) {
            store.dispatch('Report/initExecResult');
            bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
            bus.on(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
            execStart();
        } else {
            execCancel();
            bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
            bus.off(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
        }
    } 
}, { immediate: true, deep: true });

</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}
</style>