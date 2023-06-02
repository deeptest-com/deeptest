/**
 * 根据【场景执行】、【测试计划】执行产出的 [动态log]， 生成按场景编排的执行记录树
 * 1. 生成场景执行记录树
 * 2. 生成场景执行动态结果数据
 */
import {computed, ComputedRef, onMounted, ref, Ref, watch} from 'vue';
import {getDivision, getPercentStr} from "@/utils/number";

// 打平的执行记录
const execLogs: any = ref([]);
// 场景的执行结果列表
const execRes: any = ref([]);

// 更新场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
function updateExecRes(res) {
    // 1. 更新执行结果
    if (execRes.value.some((item: any) => item.scenarioId === res.scenarioId)) {
        for (let item of execRes.value) {
            if (item.scenarioId === res.scenarioId) {
                item = {...item, ...res};
                break;
            }
        }
        // 2. 新增执行结果
    } else {
        execRes.value.push(res);
    }
}

/**
 * @description 生成执行记录树
 * @param execLogs 场景的执行日志 数组
 * @param execRes  场景的执行结果
 * */
function genLogTreeView(execLogs, execRes) {
    // 用于存储根节点，即场景节点,即 processorCategory为 processor_root的节点
    const scenarioReports: any = [];
    execLogs.forEach((item: any) => {
        if (item.processorCategory === "processor_root") {
            const res = execRes.find((log) => log.scenarioId === item.scenarioId) || {};
            item.logs = [{
                ...item,
                ...res,
                logs: [],
            }];
            scenarioReports.push(item);
        }
    });
    scenarioReports.forEach((scenario) => {
        function fn(array, rootId) {
            const res: any = [];
            // 用于存储 树节点的 map
            const map = {};
            array.forEach((item) => {
                map[item.logId] = {...item}
            });
            array.forEach((item) => {
                const {logId, parentLogId} = item;
                const mapItem = map[logId];
                if (!mapItem) return;
                if (!parentLogId) return;
                if (parentLogId === rootId) {
                    res.push(mapItem);
                } else if(map[parentLogId]){
                    if (map[parentLogId]?.logs) {
                        const hasSameId = map[parentLogId].logs.some((log) => log.logId === logId);
                        !hasSameId && map[parentLogId].logs.push(mapItem);
                    } else {
                        map[parentLogId].logs = [mapItem];
                    }
                }
            })
            return res;
        }
        scenario.logs[0].logs = fn(execLogs, scenario.logId);
    });

    console.log("场景里每条编排的执行记录 2222", scenarioReports)
    return scenarioReports;
}

// // 处理动态执行的展开收起标识
// function handleActiveKey(log) {
//     log.activeKey = [log.logId];
//     if (log.logs) {
//         log.logs.forEach((item) => {
//             handleActiveKey(item);
//         })
//     }
// }

// 更新场景的执行记录，不包括场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
function updateExecLogs(log) {
    function hasSameId(log, item) {
        return item?.logId === log?.logId && item?.scenarioId === log?.scenarioId;
    }

    const isExist = execLogs.value.some((item: any) => {
        return hasSameId(log, item);
    });
    // 1. 更新执行记录
    if (isExist) {
        for (let item of execLogs.value) {
            if (hasSameId(log, item)) {
                item = {...item, ...log};
                break;
            }
        }
        // 2. 新增执行记录
    } else {
        execLogs.value.push(log);
    }
}


const scenarioReports = computed(() => {
    return [...genLogTreeView(execLogs.value, execRes.value)];
})

const expandKeys = computed(() => {
    return scenarioReports.value.map((item: any) => item.key);
})

function clearLog() {
    execLogs.value = [];
    execRes.value = [];
}

// 统计聚合数据
const statInfo = ref({
    failAssertionNum: 0,
    failInterfaceNum: 0,
    failRequestNum: 0,
    failScenarioNum: 0,
    passAssertionNum: 0,
    passInterfaceNum: 0,
    passRequestNum: 0,
    passScenarioNum: 0,
    totalAssertionNum: 0,
    totalInterfaceNum: 0,
    totalRequestNum: 0,
    totalScenarioNum: 0,
    totalProcessorNum: 0,
    notTestNum: 0,
    finishProcessorNum: 0,
    duration: 0,
    cost: 0,
})
const statisticData = computed(() => {
    const {
        failAssertionNum,
        failInterfaceNum,
        failRequestNum,
        failScenarioNum,
        passAssertionNum,
        passInterfaceNum,
        passRequestNum,
        passScenarioNum,
        totalAssertionNum,
        totalInterfaceNum,
        totalRequestNum,
        totalScenarioNum,
        totalProcessorNum,
        duration,
        notTestNum,
        cost,
    } = statInfo.value;
    // 计算平均接口耗时
    let interfaceDuration = 0;
    let interfaceNum = 0;
    execLogs.value.forEach((item: any) => {
        if (item.processorCategory === "processor_interface") {
            interfaceDuration += (item.cost || 0);
            interfaceNum++;
        }
    });
    const passRate = getPercentStr(passAssertionNum, totalAssertionNum);
    const notPassRate = getPercentStr(failAssertionNum, totalAssertionNum);
    const notTestNumRate = getPercentStr(notTestNum, totalAssertionNum);
    // 平均接口耗时
    const avgInterfaceDuration = getDivision(interfaceDuration, interfaceNum);
    return [
        {
            label: '通过',
            value: `${passAssertionNum} 个`,
            rate: passRate,
            class: 'success',
        },
        {
            label: '接口总耗时',
            value: `${interfaceDuration} 毫秒`
        },
        {
            label: '失败',
            rate: notPassRate,
            value: `${failAssertionNum} 个`,
            class: 'fail',
        },
        {
            label: '平均接口耗时',
            value: `${avgInterfaceDuration} 毫秒`,
        },
        {
            label: '未测',
            value: `${notTestNum}个`,
            rate: notTestNumRate,
            class: 'notest',
        },
        {
            label: '检查点 (成功/失败)',
            value: `${totalAssertionNum} (${passAssertionNum}/${failAssertionNum})`,
        },
    ]
})

/**
 * 从每次返回的执行日志中更新统计数据
 * */
function updateStatFromLog(res: any) {
    const {
        failAssertionNum = 0,
        failInterfaceNum = 0,
        failRequestNum = 0,
        failScenarioNum = 0,
        passAssertionNum = 0,
        passInterfaceNum = 0,
        passRequestNum = 0,
        passScenarioNum = 0,
        totalAssertionNum = 0,
        totalInterfaceNum = 0,
        totalRequestNum = 0,
        totalScenarioNum = 0,
        totalProcessorNum = 0,
        finishProcessorNum = 0,
        duration = 0,
        cost = 0,
    }: any = res;
    console.log('updateStatFromLog', res);
    const notTestNum = totalAssertionNum - passAssertionNum - failAssertionNum;
    statInfo.value = {
        failAssertionNum,
        failInterfaceNum,
        failRequestNum,
        failScenarioNum,
        passAssertionNum,
        passInterfaceNum,
        passRequestNum,
        passScenarioNum,
        totalAssertionNum,
        totalInterfaceNum,
        totalRequestNum,
        totalScenarioNum,
        totalProcessorNum,
        finishProcessorNum,
        notTestNum: notTestNum >= 0 ? notTestNum : 0,
        duration,
        cost,
    }
}

// 更新统计聚合数据
function initData(res: any) {
    updateStatFromLog(res);
}

// 【计划】的执行最终结果 用于更新最终的执行结果
function updatePlanRes(res) {
    updateStatFromLog(res);
}

const progressStatus = ref('in_progress');
const progressValue = computed(() => {
    const {
        totalProcessorNum = 1,
        finishProcessorNum,
    } = statInfo.value;
    return Math.round(((finishProcessorNum || execLogs.value.length) / totalProcessorNum) * 100);
});

export {
    scenarioReports,
    expandKeys,
    statInfo,
    execLogs,
    execRes,
    updateExecLogs,
    updateExecRes,
    clearLog
    , statisticData,
    initData, progressStatus, progressValue, updatePlanRes,
};
