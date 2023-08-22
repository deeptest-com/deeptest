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
const reportsMap: any = ref({} as any)
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

/**
 * 最终的计划的执行结果，需要从打平的执行结果中计算得出
 * */
const reports = computed(() => {
    const res = execResults.value;
    console.log('reports111', res, scenarioReports.value);
    return scenarioReports.value.map((item: any) => {
        const report = res.find((r: any) => r.scenarioId === item.scenarioId);
        return {
            ...item, ...{
                resultStatus: report?.resultStatus,
                totalProcessorNum: report?.totalProcessorNum
            }
        };
    });
})


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
    checkpointFail: 0,
    checkpointPass: 0,
    interfaceCount: 0,
    interfaceDurationAverage: 0,
    interfaceDurationTotal: 0,
    interfaceFail: 0,
    interfacePass: 0,
    interfaceSkip: 0,
});

const statisticData = computed(() => {
    const {
        checkpointFail,
        checkpointPass,
        interfaceCount,
        interfaceDurationAverage = 0,
        interfaceDurationTotal = 0,
        interfaceFail,
        interfacePass,
        interfaceSkip,
    } = statInfo.value;

    const passRate = getPercentStr(interfacePass, interfaceCount);
    const notPassRate = getPercentStr(interfaceFail, interfaceCount);
    const notTestNumRate = getPercentStr(interfaceSkip, interfaceCount);
    return [
        {
            label: '通过接口',
            value: `${interfacePass} 个`,
            rate: passRate,
            class: 'success',
        },
        {
            label: '接口总耗时',
            value: `${interfaceDurationTotal} 毫秒`
        },
        {
            label: '失败接口',
            rate: notPassRate,
            value: `${interfaceFail} 个`,
            class: 'fail',
        },
        {
            label: '平均接口耗时',
            value: `${interfaceDurationAverage} 毫秒`,
        },
        {
            label: '未测接口',
            value: `${interfaceSkip}个`,
            rate: notTestNumRate,
            class: 'notest',
        },
        {
            label: '检查点 (成功/失败)',
            value: `${checkpointPass + checkpointFail} (${checkpointPass}/${checkpointFail})`,
        },
    ]
})

/**
 * 从每次返回的执行日志中更新统计数据
 * */
function updateStatFromLog(res: any) {
    statInfo.value = {
        ...statInfo.value,
        ...res,
    }
}


// init data
function initData(res: any) {
    statInfo.value = {
        checkpointFail: 0,
        checkpointPass: 0,
        interfaceCount: 0,
        interfaceDurationAverage: 0,
        interfaceDurationTotal: 0,
        interfaceFail: 0,
        interfacePass: 0,
        interfaceSkip: 0,
    }
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
// todo 进度处理，目前先写死
const progressValue = computed(() => {
    return 10;
});

export {
    scenarioReports,
    statisticData,
    progressStatus,
    progressValue,
    statInfo,
    execLogs,
    execResults,
    updateExecLogs,
    updateExecResult,
    resetData,
    initData,
    updatePlanRes,
    reports,
    updateStatFromLog,
};
