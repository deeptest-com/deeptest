<template>
  <div class="scenario-exec-info-main">
    <ReportBasicInfo :items="baseInfoList || []"
                     :showBtn="true"
                      :btnText="'生成报告'"
                     @handleBtnClick="genReport"/>
    <StatisticTable :data="statisticData" :value="statInfo"/>
    <Progress :exec-status="progressStatus"
              :percent="progressValue"
              @exec-cancel="execCancel"/>
    <LogTreeView :treeData="scenarioReports" :expandKeys="expandKeys"/>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {
  ReportBasicInfo,
  StatisticTable,
  LogTreeView,
  Progress
} from '@/views/component/Report/components';
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import bus from "@/utils/eventBus";
import {useI18n} from "vue-i18n";
import {getToken} from "@/utils/localToken";
import {Scenario} from "@/views/scenario/data";
import {message} from "ant-design-vue";
import {getDivision, getPercentStr} from "@/utils/number";

const {t} = useI18n();

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType, Global: GlobalStateType, Exec: ExecStatus, ProjectSetting, Environment }>();
const collapsed = computed<boolean>(() => store.state.Global.collapsed);
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
const envList = computed(() => store.state.ProjectSetting.envList);
const scenarioId = computed(() => {
  return detailResult.value.id
});

const reportId = ref('');
const baseInfoList = computed(() => {
  if (!detailResult.value) return [];
  console.log(envList.value)
  const curEnv = envList.value.find((item: any) => item.id === currEnvId.value)
  return [
    {value: detailResult?.value?.name || '暂无', label: '场景名称'},
    {value: curEnv?.name || '暂无', label: '执行环境'},
  ]
})
const statisticData = computed(() => {
  const {
    failAssertionNum,
    passAssertionNum,
    totalAssertionNum,
    notTestNum,
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
      value: `${passRate} ${passAssertionNum} 个`,
    },
    {
      label: '接口总耗时',
      value: `${interfaceDuration} 毫秒`
    },
    {
      label: '失败',
      value: `${notPassRate}  ${failAssertionNum} 个`,
      class: 'fail',
    },
    {
      label: '平均接口耗时',
      value: `${avgInterfaceDuration} 毫秒`,
    },
    {
      label: '未测',
      value: `${notTestNumRate} ${notTestNum}个`,
      class: 'fail',
    },
  ]
})
const progressValue = computed(() => {
  const {
    totalProcessorNum = 1,
    finishProcessorNum,
  } = statInfo.value;
  return Math.round(((finishProcessorNum || execLogs.value.length) / totalProcessorNum) * 100);
});
const progressStatus = ref('in_progress');
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
  totalScenarioNum: 0,
  totalProcessorNum: 0,
  notTestNum: 0,
  finishProcessorNum: 0,
  cost:0,
});

const scenarioReports = computed(() => {
  return [...genLogTreeView(execLogs.value, execRes.value)];
})

const expandKeys = computed(() => {
  return scenarioReports.value.map((item: any) => item.key);
})

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
        logs: []
      }];
      scenarioReports.push(item);
    }
  });
  scenarioReports.forEach((scenario) => {
    function fn(array, rootId) {
      const res: any = [];
      // 用于存储 树节点的map
      const map = {};
      array.forEach((item) => {
        map[item.id] = {...item}
      });
      array.forEach((item) => {
        const {id, parentId} = item;
        if (!map[id]) return;
        if (!parentId) return;
        if (parentId === rootId) {
          res.push(map[id]);
        } else {
          if (map[parentId]?.logs) {
            map[parentId].logs.push(map[id]);
          } else {
            map[parentId].logs = [map[id]];
          }
        }
      })
      return res;
    }
    scenario.logs[0].logs = fn(execLogs, scenario.id);
  });
  console.log('832 scenarioReports', scenarioReports)
  return scenarioReports;
}

const execStart = async () => {
  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    scenarioId: scenarioId.value,
    environmentId: currEnvId.value,
  }
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', scenarioExecReq: data}))
}

const execCancel = () => {
  progressStatus.value = 'cancel';
  const msg = {act: 'stop', execReq: {scenarioId: scenarioId.value}}
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
}

onMounted(async () => {
  await execStart();
  bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
})

onUnmounted(() => {
  bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
})

// 打平的执行记录
const execLogs: any = ref([]);
// 场景的执行结果列表
const execRes: any = ref([]);

// 更新场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
function updateExecRes(res) {
  // 1. 更新执行结果
  if (execRes.value.some((item: any) => item.scenarioId === res.scenarioId)) {
    execRes.value.forEach((item: any) => {
      if (item.scenarioId === res.scenarioId) {
        item = {...item, ...res};
      }
    });
    // 2. 新增执行结果
  } else {
    execRes.value.push(res);
  }
}

// 更新场景的执行记录，不包括场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
function updateExecLogs(log) {
  const isExist = execLogs.value.some((item: any) => {
    return item.logId === log.logId && item.scenarioId === log.scenarioId;
  });
  // 1. 更新执行记录
  if (isExist) {
    execLogs.value.forEach((item: any) => {
      if (item.logId === log.logId && item.scenarioId === log.scenarioId) {
        item = {...item, ...log};
      }
    });
    // 2. 新增执行记录
  } else {
    execLogs.value.push(log);
  }
}

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
    totalScenarioNum,
    totalProcessorNum,
    finishProcessorNum,
    cost,
    notTestNum: notTestNum >= 0 ? notTestNum : 0,
  }
}

// 更新统计聚合数据
function initData(res: any) {
  updateStatFromLog(res);
}

const OnWebSocketMsg = (data: any) => {
  if (!data.msg) return;
  if(progressStatus.value === 'cancel') return;
  const wsMsg = JSON.parse(data.msg);
  const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
  // 开始执行，初始化数据
  if (wsMsg.category == 'initialize') {
    initData(log);
    progressStatus.value = 'in_progress';
  }
  // 执行中
  else if (wsMsg.category == 'in_progress') {
    progressStatus.value = 'in_progress';
  }
  //  更新【场景】的执行结果
  else if (wsMsg.category == 'result' && log.scenarioId) {
    updateExecRes(log);
  }
  // 更新【场景中每条编排】的执行记录
  else if (wsMsg.category === "processor" && log.scenarioId) {
    console.log('场景里每条编排的执行记录', log)
    updateExecLogs(log);
  }
  // 执行完毕
  else if (wsMsg.category == 'end') {
    progressStatus.value = 'end';
  } else {
    console.log('wsMsg', wsMsg);
  }
}


async function genReport() {
  const res = await store.dispatch('Scenario/genReport', {
    id: reportId.value,
  });
  if(res){
    message.success('生成报告成功');
  }else {
    message.error('生成报告失败');
  }
}
</script>

<style lang="less" scoped>

.scenario-exec-info-main {
  height: 100%;
  padding: 6px;

  .scenario {
    .header {
      display: flex;
      padding: 0px 12px;
      background-color: #fafafa;
      border: 1px solid #f0f0f0;
      line-height: 32px;

      .title {
        flex: 1;
      }

      .progress {
        width: 100px;
      }

      .status {
        width: 100px;
      }

      .opt {
        width: 260px;
        text-align: right;
      }
    }

    .logs {
      padding: 0px 12px;
    }

    .result {
      padding: 5px 23px 6px 23px;

      .ant-row {
        margin: 6px 0;
      }
    }
  }
}

</style>
