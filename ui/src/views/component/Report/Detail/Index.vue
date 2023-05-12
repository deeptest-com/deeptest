<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ title || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo :basic-info="reportData" />
            <StatisticTable />
            <Progress :percent="60" @exec-cancel="execCancel" />
            <template v-for="logItem in reportData.logList" :key="logItem.id">
                <ScenarioCollapsePanel :show-scenario-info="showScenarioInfo" :expand-active="expandActive"
                    :record="logItem">
                    <template #endpointData>
                        <EndpointCollapsePanel :recordList="logItem.reponseList" />
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
import { ExecStatus } from "@/store/exec";
import settings from "@/config/settings";
import { WebSocket } from "@/services/websocket";
import { WsMsg } from "@/types/data";
import bus from "@/utils/eventBus";
import { getToken } from "@/utils/localToken";
import { reportData } from './testData';

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scenarioExpandActive: boolean
    showScenarioInfo: boolean
}>();
const emits = defineEmits(['onClose']);
const store = useStore<{ Plan: PlanStateType, Global: GlobalStateType, Exec: ExecStatus }>();
const currPlan = computed<any>(() => store.state.Plan.currPlan);
const execResult = computed<any>(() => store.state.Plan.execResult);
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
    }

    WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({ act: 'execPlan', planExecReq: data }));
};

const execCancel = () => {
    console.log('execCancel');
    const msg = { act: 'stop', execReq: { scenarioId: currPlan.value && currPlan.value.id } };
    WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
};

const OnWebSocketMsg = (data: any) => {
    console.log('--- WebsocketMsgEvent in exec info', data)

    if (!data.msg) return

    const wsMsg = JSON.parse(data.msg) as WsMsg
    if (wsMsg.category == 'result') { // update result
        result.value = wsMsg.data
        console.log('=====', result.value)
        return
    } else if (wsMsg.category !== '') { // update status
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
    } else {
        if (!logMap.value[log.parentId]) logMap.value[log.parentId] = {};
        if (!logMap.value[log.parentId].logs) logMap.value[log.parentId].logs = [];
        logMap.value[log.parentId].logs.push(log);
    }

    console.log('logMap Data--------', logMap.value);
    console.log('logTree Data--------', logTreeData.value);
    console.log('execResult Data -----', execResult);
};

function onClose() {
    emits('onClose');
}
console.log(props.drawerVisible);
watch(() => {
    return props;
}, val => {
    if (val.drawerVisible) {
        bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
        execStart();
    } else {
        execCancel();
        bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
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