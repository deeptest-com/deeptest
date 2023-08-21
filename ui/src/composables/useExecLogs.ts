/**
 * 根据【场景执行】、【测试计划】执行产出的 [动态log]， 生成按场景编排的执行记录树
 * 1. 生成场景执行记录树
 * 2. 生成场景执行动态结果数据
 */
import {computed, ref} from 'vue';
import {getDivision, getPercentStr} from "@/utils/number";
import {ProcessorCategory} from "@/utils/enum";

// 打平的执行记录
const execLogs: any = ref([]);
// 场景的执行结果列表
const execResults: any = ref([]);


const reportsMap:any = ref({} as any)
// 组装后的场景的执行报告树
const scenarioReports = ref([] as any[]);

// 更新场景的执行结果
function updateExecResult(res) {
    // 1. 更新执行结果
    if (execResults.value.some((item: any) => item.scenarioId === res.scenarioId)) {
        for (let item of execResults.value) {
            if (item.scenarioId === res.scenarioId) {
                item = {...item, ...res};
                break;
            }
        }
        // 2. 新增执行结果
    } else {
        execResults.value.push(res);
    }
}




// 更新场景的执行日志，不包括场景的执行结果。
function updateExecLogs(processor) {
    console.log('updateExecLogs', 'logId=' + processor.logId, 'parentLogId=' + processor.parentLogId);

    /**
     * 更新执行日志至打平的执行记录，用于动态 更 新执行进度
     * */
    function hasSameId(log, item) {
        return item?.logId === log?.logId && item?.scenarioId === log?.scenarioId;
    }
    const isExist = execLogs.value.some((item: any) => {
        return hasSameId(processor, item);
    });
    // 1. 更新执行记录
    if (isExist) {
        for (let item of execLogs.value) {
            if (hasSameId(processor, item)) {
                item = {...item, ...processor};
                break;
            }
        }
        // 2. 新增执行记录
    } else {
        execLogs.value.push(processor);
    }

    /**
     * 组装执行记录树，动态执行树最多展示50条记录
     * */
    if (processor.processorCategory === ProcessorCategory.ProcessorRoot) { // reset
        reportsMap.value = {}
        // scenarioReports.value = []
    }

    reportsMap.value[processor.logId] = processor;
    if (processor.processorCategory === ProcessorCategory.ProcessorRoot) {
        scenarioReports.value.push(processor);
        return
    }
    if (reportsMap.value[processor.parentLogId]) {
        if (!reportsMap.value[processor.parentLogId].logs) {
            reportsMap.value[processor.parentLogId].logs = []
        }
        if (reportsMap.value[processor.parentLogId].logs.length >= 50) {
            reportsMap.value[processor.parentLogId].logs.shift()
        }
        reportsMap.value[processor.parentLogId].logs.push(processor);
    }

    const elems = document.getElementsByClassName('scenario-exec-log-tree')
    if (elems && elems.length > 0) {
        elems[0].scrollTop = elems[0].scrollHeight + 1000
    }

    console.log('scenarioReports832222', scenarioReports.value);
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
    console.log(execLogs.value,'***',interfaceNum);
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

// init data
function initData(res: any) {
    updateStatFromLog(res);
}

function resetData() {
    execLogs.value = [];
    execResults.value = [];
    scenarioReports.value = [];
    reportsMap.value = {};
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
    statisticData,
    progressStatus, progressValue,
    statInfo,
    execLogs,
    execResults,
    updateExecLogs,
    updateExecResult,
    resetData,
    initData,
    updatePlanRes,
};
