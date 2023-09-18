<template>
  <div class="scenario-exec-info-main">
    <ReportBasicInfo :items="baseInfoList || []"
                     :showBtn="show"
                     :btnText="'另存为报告'"
                     @handleBtnClick="genReport"/>
    <StatisticTable :data="statisticData" :value="statInfo"/>
    <Progress :exec-status="progressStatus"
              :percent="progressValue"
              :key="progressKey"
              @exec-cancel="execCancel" />
    <LogTreeView class="scenario-exec-log-tree"
        :treeData="scenarioReports"
                 :isSingleScenario="true" />
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {LogTreeView, Progress, ReportBasicInfo, StatisticTable} from '@/views/component/Report/components';
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import {StateType as DebugStateType} from "@/views/component/debug/store";
import bus from "@/utils/eventBus";
import {useI18n} from "vue-i18n";
import {getToken} from "@/utils/localToken";
import {Scenario} from "@/views/scenario/data";
import {message} from "ant-design-vue";
import {ProcessorInterface} from "@/utils/enum";
import {StateType as ReportStateType} from "@/views/report/store";

import {
  resetData,
  initData,
  progressStatus,
  progressValue,
  scenarioReports,
  statInfo,
  statisticData,
  updateExecLogs,
  updateExecResult,
  updateStatFromLog
} from '@/composables/useExecLogs';
import {momentUtc} from "@/utils/datetime";
import {CurrentUser} from "@/store/user";
import {notifyError, notifySuccess} from "@/utils/notify";

const {t} = useI18n();
const router = useRouter();

const store = useStore<{ Report: ReportStateType, Scenario: ScenarioStateType, Debug: DebugStateType, Global: GlobalStateType, Exec: ExecStatus, ProjectSetting, Environment,User }>();
const collapsed = computed<boolean>(() => store.state.Global.collapsed);
const nodeData = computed<any>(() => store.state.Scenario.nodeData);
const detailResult = computed<any>(() => store.state.Scenario.detailResult);
const currentUser:any = computed<CurrentUser>(()=> store.state.User.currentUser);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
const envList = computed(() => store.state.ProjectSetting.envList);
const scenarioId = computed(() => {
  return detailResult.value.id
});


const envName = computed(() => {
  const curEnv = envList.value.find((item: any) => item.id === currEnvId.value)
  return curEnv?.name || '暂无'
});

const reportId = ref('');
const show = ref(false)
const baseInfoList = computed(() => {
  if (!detailResult.value) return [];
  const curEnv = envList.value.find((item: any) => item.id === currEnvId.value)
  return [
    {value: detailResult?.value?.name || '暂无', label: '场景名称'},
    {value: momentUtc(new Date()) , label: '执行时间'},
    {value: curEnv?.name || '暂无', label: '执行环境'},
    {value: detailResult.value.creatorName || '暂无', label: '创建人'},
    {value: currentUser?.value?.name || '暂无', label: '执行人'},
    {value: detailResult.value.priority || '未设置', label: '优先级'},
  ]
});

// 每次重新渲染
const progressKey = ref(0);
const execStart = async () => {
  resetData();
  progressKey.value += 1;
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
  stopExec();
}

const stopExec = () => {
  const msg = {act: 'stop', execReq: {scenarioId: scenarioId.value}};
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg));
}

onMounted(async () => {
  progressStatus.value = 'in_progress';
  await execStart();
  bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
  bus.on(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);

})

onUnmounted(() => {
  execCancel();
  bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
  bus.off(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
})

const OnWebSocketMsg = (data: any) => {
  if (!data.msg) return;
  if (progressStatus.value === 'cancel') return;
  if (progressStatus.value === 'exception') return;

  const wsMsg = JSON.parse(data.msg);
  const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
  console.log('scenario wsMsg***', wsMsg);
  // 开始执行，初始化数据
  if (wsMsg.category == 'initialize') {
    // 重置数据, 重新初始化
    // initData();
    progressStatus.value = 'in_progress';
  }
  // 执行中
  else if (wsMsg.category == 'in_progress') {
    progressStatus.value = 'in_progress';
  }
  //  更新【场景】的执行结果
  else if (wsMsg.category == 'result' && log.scenarioId) {
    updateExecResult(log);
    reportId.value = log.id
  }
  // 更新【场景中每条编排】的执行记录
  else if (wsMsg.category === "processor" && log.scenarioId) {
    console.log('场景里每条编排的执行记录', log);
    updateExecLogs(log);
  }
  // 更新统计值
  else if (wsMsg.category === "stat") {
    updateStatFromLog(log);
  }
  else if (wsMsg.category === "exception") {
    progressStatus.value = 'exception';
    stopExec();
  }
  // 执行完毕
  else if (wsMsg.category == 'end') {
    progressStatus.value = 'end';
    show.value = true
    // refresh processor interface data in scenario if needed
    if (nodeData.value.processorType === ProcessorInterface.Interface) {
      store.dispatch('Debug/refreshInterfaceResultFromScenarioExec')
    }
  } else {
    console.log('wsMsg', wsMsg);
  }
}

// websocket 连接状态 查询
const onWebSocketConnStatusMsg = (data: any) => {
  if (!data.msg) {
    return;
  }
  const {conn}: any = JSON.parse(data.msg);
  progressStatus.value = conn === 'success' ? 'in_progress' : 'exception';
}

async function genReport() {
  const res = await store.dispatch('Scenario/genReport', {
    id: reportId.value,
  });
  if (res) {
    notifySuccess('生成报告成功');
  } else {
    notifyError('生成报告失败');
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
