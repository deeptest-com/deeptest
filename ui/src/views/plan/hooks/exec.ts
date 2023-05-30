import {ref, computed, watch} from 'vue';
import {useStore} from 'vuex';

import {StateType as PlanStateType} from '@/views/plan/store';
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as ReportStateType} from "../store";
import {StateType as UserStateType} from '@/store/user';
import {ExecStatus} from "@/store/exec";
import {StateType as ProjectSettingStateType} from '@/views/projectSetting/store';

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {WsMsg} from "@/types/data";
import {getToken} from "@/utils/localToken";
import {WsMsgCategory} from '@/utils/enum';
import {momentUtc} from '@/utils/datetime';
import {formatData} from '@/utils/formatData';

export function useExec() {
    const store = useStore<{
        Plan: PlanStateType,
        Global: GlobalStateType,
        Exec: ExecStatus,
        Report: ReportStateType,
        ProjectSetting: ProjectSettingStateType,
        User: UserStateType
    }>();
    const currPlan = computed<any>(() => store.state.Plan.currPlan);
    const execResult = computed<any>(() => store.state.Plan.execResult);
    const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
    const currUser = computed(() => store.state.User.currentUser);
    const logDetailData = ref<any>({});
    const processNum = ref(0);

    const transformWithUndefined = (num: number | undefined) => {
        return num || 0;
    }

    const calcNum = (currNum, lastNum) => {
        return currNum + transformWithUndefined(lastNum);
    }

    const execStart = async () => {
        console.log('execStart');
        logDetailData.value = {basicInfoList: [], scenarioReports: [], statisticData: {}};
        processNum.value = 0;
        flattenExecLogs.value = [];
        const data = {
            serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
            token: await getToken(),
            planId: currPlan.value && currPlan.value.id,
            environmentId: currEnvId.value
        }
        WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execPlan', planExecReq: data}));
    };

    const execCancel = () => {
        console.log('execCancel');
        flattenExecLogs.value = [];
        const msg = {act: 'stop', execReq: {planId: currPlan.value && currPlan.value.id}};
        WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
    };

    /**
     * @description 生成执行记录树
     * @param data 打平的执行记录
     * */
    function genLogTreeView(data: any) {
        // function fn(array, rootId) {
        //     const map = {};
        //     array.forEach((item) => {
        //         map[item.id] = {...item} // 浅拷贝
        //     });
        //     const res = [];
        //     array.forEach((item:any) => {
        //         let {id, parentId} = item;
        //         if (parentId === rootId) {
        //             res.push(map[id]);
        //         } else {
        //             if (map[parentId].logs) {
        //                 map[parentId].logs.push(map[id]);
        //             } else {
        //                 map[parentId].logs = [map[id]];
        //             }
        //         }
        //     })
        //     return res;
        // }
        // return fn(data, parent);
        return null;

    }

    // 打平的执行记录
    const flattenExecLogs: any = ref([]);

    // 更新执行记录
    function updateExecLogs(log) {
        // 1. 更新执行记录
        if(flattenExecLogs.value.some((item: any) => item.id === log.id)) {
            flattenExecLogs.value.forEach((item: any) => {
                if(item.id === log.id) {
                    item = {...log};
                }
            });
        // 2. 新增执行记录
        }else {
            flattenExecLogs.value.push(log);
        }
    }

    // 监听执行记录, 生成执行记录树
    watch(flattenExecLogs.value,
        (newValue, oldValue) => {
            console.log(832,newValue);
            logDetailData.value.scenarioReports = genLogTreeView(newValue);
        }, {
            deep: true
        });

    const OnWebSocketMsg = (data: any) => {
        if (!data.msg) return;
        const wsMsg = JSON.parse(data.msg);
        console.log('832 msg', wsMsg);
        const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
        const statisticData = JSON.parse(JSON.stringify(logDetailData.value.statisticData));
        // :::: 构造 测试计划的基本信息 + 执行的基本统计信息 ::::
        if (wsMsg.category == 'result' && !wsMsg.data.scenarioId && wsMsg.data.planId) {
            // 测试计划的基本信息
            logDetailData.value.basicInfoList = [
                {
                    label: '测试计划',
                    value: log.planName || '-'
                },
                {
                    label: '开始时间',
                    value: momentUtc(new Date())
                },
                {
                    label: '执行环境',
                    value: log.execEnv || '--'
                },
                {
                    label: '创建人',
                    value: currUser.value.username || '--'
                },
            ];
            //  测试计划的统计信息
            logDetailData.value.statisticData = {
                ...statisticData,
                totalScenarioNum: log.totalScenarioNum,
                totalInterfaceNum: log.totalInterfaceNum,
                totalAssertionNum: log.totalAssertionNum
            };
            store.dispatch('Plan/setExecResult', logDetailData.value);
        }

        if (wsMsg.category === "processor"  || wsMsg.category === "result") { // update progress
            if('scenarioId' in log){
               updateExecLogs(log);
            }
        }


        // if (wsMsg.category == 'result') { // update result
        //
        //
        //     if (!wsMsg.data.scenarioId && wsMsg.data.planId) {
        //         // 测试计划的基本信息
        //         logDetailData.value.basicInfoList = [
        //             {
        //                 label: '测试计划',
        //                 value: log.planName || '-'
        //             },
        //             {
        //                 label: '开始时间',
        //                 value: momentUtc(new Date())
        //             },
        //             {
        //                 label: '执行环境',
        //                 value:log.execEnv || '--'
        //             },
        //             {
        //                 label: '创建人',
        //                 value: currUser.value.username || '--'
        //             },
        //         ],
        //         //  测试计划的统计信息
        //         logDetailData.value.statisticData = {
        //             ...statisticData,
        //             totalScenarioNum: log.totalScenarioNum,
        //             totalInterfaceNum: log.totalInterfaceNum,
        //             totalAssertionNum: log.totalAssertionNum
        //         };
        //         store.dispatch('Plan/setExecResult', logDetailData.value);
        //     }
        //
        //     else if (log.scenarioId !== 0) { // scenarioId !== 0 则是单条场景的执行结果
        //         const isExsitData = logDetailData.value.scenarioReports.find(e => e.id === log.id);
        //         if (isExsitData) {
        //             return;
        //         }
        //         logDetailData.value.statisticData = {
        //             ...statisticData,
        //             "duration": log.duration + transformWithUndefined(statisticData.duration), //执行耗时（单位：s)
        //             "passScenarioNum": calcNum(log.resultStatus === 'fail' ? 0 : 1, statisticData.passScenarioNum), //通过场景数
        //             "failScenarioNum": calcNum(log.resultStatus === 'fail' ? 1 : 0, statisticData.failScenarioNum), //失败场景数
        //             "passInterfaceNum": calcNum(log.passInterfaceNum, statisticData.passInterfaceNum),
        //             "failInterfaceNum": calcNum(log.failInterfaceNum, statisticData.failInterfaceNum),
        //             "totalRequestNum": calcNum(log.totalRequestNum, statisticData.totalRequestNum),
        //             "passRequestNum": calcNum(log.passRequestNum, statisticData.passRequestNum),
        //             "failRequestNum": calcNum(log.failRequestNum, statisticData.failRequestNum),
        //             "passAssertionNum": calcNum(log.passAssertionNum, statisticData.passAssertionNum), //通过检查点数
        //             "failAssertionNum": calcNum(log.failAssertionNum, statisticData.failAssertionNum), //失败检查点数
        //         };
        //         const requestLogs = (log.logs && log.logs[0].logs) ? formatData(log.logs[0].logs) : [];
        //         logDetailData.value.scenarioReports.push({
        //             ...log,
        //             requestLogs
        //         })
        //         processNum.value++;
        //         logDetailData.value.progressValue = statisticData.totalScenarioNum ?  processNum.value / statisticData.totalScenarioNum * 100 : 10;
        //         store.dispatch('Plan/setExecResult', logDetailData.value);
        //     } else { // scenarioId = 0 为整个计划的执行结果.
        //         logDetailData.value.statisticData = {
        //             ...statisticData,
        //             "duration": log.duration, //执行耗时（单位：ms)
        //             "passScenarioNum": log.passScenarioNum, //通过场景数
        //             "failScenarioNum": log.failScenarioNum, //失败场景数
        //             "passInterfaceNum": log.passInterfaceNum,
        //             "failInterfaceNum": log.failInterfaceNum,
        //             "totalRequestNum": log.totalRequestNum,
        //             "passRequestNum": log.passRequestNum,
        //             "failRequestNum": log.failRequestNum,
        //             "passAssertionNum": log.passAssertionNum, //通过检查点数
        //             "failAssertionNum": log.failAssertionNum, //失败检查点数
        //             "totalAssertionNum": log.totalAssertionNum
        //         };
        //         store.dispatch('Plan/setExecResult', logDetailData.value);
        //     }
        //     return;
        // }
        //
        // else if (wsMsg.category === WsMsgCategory.InProgress || wsMsg.category === WsMsgCategory.End) { // update status
        //     logDetailData.value.progressStatus = wsMsg.category;
        //     if (wsMsg.category === WsMsgCategory.End) {
        //         logDetailData.value.progressValue = 100;
        //     }
        //     store.dispatch('Plan/setExecResult', logDetailData.value);
        //     return;
        // }
    };

    // websocket 连接状态 查询
    const onWebSocketConnStatusMsg = (data: any) => {
        if (!data.msg) {
            return;
        }
        const {conn}: any = JSON.parse(data.msg);
        logDetailData.value.progressStatus = conn === 'success' ? WsMsgCategory.InProgress : 'failed';
        store.dispatch('Plan/setExecResult', logDetailData.value);
    }

    return {
        logDetailData,
        execStart,
        execCancel,
        onWebSocketConnStatusMsg,
        OnWebSocketMsg
    }
}
