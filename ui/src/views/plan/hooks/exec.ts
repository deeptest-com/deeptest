import { ref, computed } from 'vue';
import { useStore } from 'vuex';

import { StateType as PlanStateType } from '@/views/plan/store';
import { StateType as GlobalStateType } from "@/store/global";
import { StateType as ReportStateType } from "../store";
import { ExecStatus } from "@/store/exec";
import { StateType as ProjectSettingStateType } from '@/views/projectSetting/store';

import settings from "@/config/settings";
import { WebSocket } from "@/services/websocket";
import { WsMsg } from "@/types/data";
import { getToken } from "@/utils/localToken";
import { WsMsgCategory } from '@/utils/enum';
import { momentUtc } from '@/utils/datetime';

export function useExec() {
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
    const logTreeData = ref<any>([]);
    const logDetailData = ref<any>({});
    const result = ref({} as any);
    const logMap = ref({} as any);

    const execStart = async () => {
        console.log('execStart');

        logTreeData.value = [];
        logDetailData.value = { basicInfo: {}, scenarioReports: [] };

        const data = {
            serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
            token: await getToken(),
            planId: currPlan.value && currPlan.value.id,
            environmentId: currEnvId.value
        }

        WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({ act: 'execPlan', planExecReq: data }));
    };

    const execCancel = () => {
        console.log('execCancel');
        const msg = { act: 'stop', execReq: { planId: currPlan.value && currPlan.value.id } };
        WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
    };

    const OnWebSocketMsg = (data: any) => {
        if (!data.msg) return
        const wsMsg = JSON.parse(data.msg) as WsMsg;
        const log = JSON.parse(JSON.stringify(wsMsg.data));
        console.log('--- WebsocketMsgEvent in exec info', wsMsg);
        if (wsMsg.category == 'result') { // update result
            console.log('=====', result.value)
            // scenarioId === 0  测试计划的数据
            if (wsMsg.data.scenarioId === 0) {
                logDetailData.value.basicInfo = {
                    name: log.name || '',
                    startTime: (log.startTime && momentUtc(log.startTime)) || '',
                    execEnv: log.execEnv || '',
                    createUserName: log.createUserName || ''
                };
                logDetailData.value.statisticData = {
                    "duration": log.duration, //执行耗时（单位：s)
                    "totalScenarioNum": log.totalScenarioNum, //场景总数
                    "passScenarioNum": log.passScenarioNum, //通过场景数
                    "failScenarioNum": log.failScenarioNum, //失败场景数
                    "totalInterfaceNum": log.totalInterfaceNum, //接口总数
                    "passInterfaceNum": log.passInterfaceNum,
                    "failInterfaceNum": log.failInterfaceNum,
                    "totalRequestNum": log.totalRequestNum,
                    "passRequestNum": log.passRequestNum,
                    "failRequestNum": log.failRequestNum,
                    "totalAssertionNum": log.totalAssertionNum, //检查点总数
                    "passAssertionNum": log.passAssertionNum, //通过检查点数
                    "failAssertionNum": log.failAssertionNum, //失败检查点数
                };
            } else { // scenarioId !== 0 场景的执行结果汇总
                logDetailData.value.scenarioReports
                let scenarioReports = logDetailData.value.scenarioReports || [];
                scenarioReports = scenarioReports.map(e => {
                    if (e.name === data.name) {
                        return { ...data, ...e };
                    }
                });
            }

            return
        } else if (wsMsg.category === WsMsgCategory.InProgress || wsMsg.category === WsMsgCategory.End) { // update status
            execResult.value.progressStatus = wsMsg.category
            if (wsMsg.category === 'in_progress') {
                result.value = {};
            }
            return
        }

        if (log.id) {
            logMap.value[log.id] = log;
        }

        if (log.parentId === 0) { // 场景数据
            console.log('~~~~ log ~~~', log);
            log.requestLogs = [];
            const findScenarioReport = logDetailData.value.scenarioReports.find(e => e.id === log.id);
            console.log(findScenarioReport);
            if (!findScenarioReport) {
                logDetailData.value.scenarioReports.push(log);
            }
        } else {
            logDetailData.value.scenarioReports = logDetailData.value.scenarioReports.map(e => {
                if (log.parentId === e.id) {
                    const findRequestLog = e.requestLogs.find(request => request.id === log.id);
                    if (!findRequestLog) {
                        e.requestLogs.push(log);
                    }
                    return e;
                }
                return e;
            })
        }

        console.log('~~~~~~ get logDetailData ~~~~~', logDetailData.value);
        store.dispatch('Report/setExecResult', logDetailData.value);
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

    return {
        logDetailData,
        execStart,
        execCancel,
        onWebSocketConnStatusMsg,
        OnWebSocketMsg
    }
}