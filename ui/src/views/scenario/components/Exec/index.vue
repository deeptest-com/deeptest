<template>
  <div class="scenario-exec-info-main">
    <ReportBasicInfo :items="baseInfoList || []" :scene="ReportDetailType.ExecScenario" />
    <StatisticTable :scene="ReportDetailType.ExecScenario" :data="execResult.statisticData" />
    <Progress :exec-status="execResult.progressStatus" :percent="execResult.progressValue || 10" @exec-cancel="execCancel" />
<!--    <template v-for="scenarioReportItem in execResult.scenarioReports" :key="scenarioReportItem.id">-->
<!--      <ScenarioCollapsePanel :show-scenario-info="showScenarioInfo"-->
<!--                             :expand-active="expandActive"-->
<!--                             :record="scenarioReportItem">-->
<!--        <template #endpointData>-->
<!--          <EndpointCollapsePanel v-if="scenarioReportItem.requestLogs.length > 0"-->
<!--                                 :recordList="scenarioReportItem.requestLogs" />-->
<!--        </template>-->
<!--      </ScenarioCollapsePanel>-->
<!--    </template>-->
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {WsMsg} from "@/types/data";
import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel, Progress } from '@/views/component/Report/Components';
import { ReportDetailType } from "@/utils/enum";
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import bus from "@/utils/eventBus";
import Log from "./Log.vue"
import {momentShort, momentUtc} from "@/utils/datetime";
import {useI18n} from "vue-i18n";
import {getToken} from "@/utils/localToken";
import {WsMsgCategory} from "@/utils/enum";
import {Scenario} from "@/views/scenario/data";
const { t } = useI18n();

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType, Global: GlobalStateType, Exec: ExecStatus,ProjectSetting,Environment }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
const envList = computed(() => store.state.ProjectSetting.envList);
// const environmentData = computed(() => store.state.Environment.environmentData);
const scenarioId = computed(() => {
  return detailResult.value.id
});
const baseInfoList = computed(() => {
  if (!detailResult.value) return [];
  console.log(envList.value)
  const curEnv = envList.value.find((item: any) => item.id === currEnvId.value)
  return [
    {value: detailResult.value.name, label: '场景名称'},
    {value: curEnv.name || '暂无', label: '执行环境'},
  ]
})
const execResult = computed<any>(() => store.state.Scenario.execResult);

// // 执行场景ID变化时，执行场景
// watch(() => {
//   scenarioId.value
// }, async (newVal:any, oldVal) => {
//   if(newVal){
//     await execStart();
//   }
// },{
//   immediate: true
// })

const execStart = async () => {
  console.log('execStart')
  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    scenarioId: scenarioId.value,
    environmentId: currEnvId.value,
  }
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', scenarioExecReq: data}))
}

const execCancel = () => {
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

const result = ref({} as any)
const logMap = ref({} as any)
const logTreeData = ref({} as any)
const OnWebSocketMsg = (data: any) => {

  if (!data.msg) return
  const wsMsg = JSON.parse(data.msg) as WsMsg;
  console.log('121212', wsMsg)

  // dealwith category
  if (wsMsg.category) {
    if (wsMsg.category == WsMsgCategory.Result) { // update result
      result.value = wsMsg.data
    } else { // update status
      execResult.value.progressStatus = wsMsg.category
      if (wsMsg.category === WsMsgCategory.InProgress) result.value = {}
    }
    return
  }
  const log = wsMsg.data
  logMap.value[log.id] = log
  if (log.parentId === 0) { // root
    logTreeData.value = log
    logTreeData.value.name = execResult.value.name
  } else {
    if (!logMap.value[log.parentId]) logMap.value[log.parentId] = {}
    if (!logMap.value[log.parentId].logs) logMap.value[log.parentId].logs = []
    logMap.value[log.parentId].logs.push(log)
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
