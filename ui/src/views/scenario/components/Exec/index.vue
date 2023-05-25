<template>
  <div class="scenario-exec-info-main">
    <ReportBasicInfo :items="baseInfoList || []" :showOperation="!!reportId"
                     :scene="ReportDetailType.ExecScenario"
                     @handleBtnClick="genReport"/>
    <StatisticTable :scene="ReportDetailType.ExecScenario" :data="statisticData"/>
    <Progress :exec-status="progressStatus" :percent="progressValue" @exec-cancel="execCancel"/>
    <EndpointCollapsePanel v-if="recordList.length > 0"
                           :recordList="recordList"/>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {WsMsg} from "@/types/data";
import {
  ReportBasicInfo,
  StatisticTable,
  ScenarioCollapsePanel,
  EndpointCollapsePanel,
  Progress
} from '@/views/component/Report/components';
import {ReportDetailType} from "@/utils/enum";
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import bus from "@/utils/eventBus";
import Log from "./Log.vue"
import {momentShort, momentUtc} from "@/utils/datetime";
import {useI18n} from "vue-i18n";
import {getToken} from "@/utils/localToken";
import {WsMsgCategory} from "@/utils/enum";
import {formatData} from "@/utils/formatData";
import {Scenario} from "@/views/scenario/data";
import {message} from "ant-design-vue";

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
const statisticData = ref({});
// const execResult = computed<any>(() => store.state.Scenario.execResult);
const progressValue = ref(10);
const recordList:any = ref([]);
const progressStatus = ref('in_progress');
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


const OnWebSocketMsg = (data: any) => {
  console.log('OnWebSocketMsg 832', data);
  if (!data.msg) return
  const wsMsg = JSON.parse(data.msg) as WsMsg;
  if (wsMsg.category) {
    if (wsMsg.category == WsMsgCategory.Result) { // update result
      const res = wsMsg.data;
      progressStatus.value = wsMsg.category;
      progressValue.value = 100;
      statisticData.value = {
        "duration": res.duration, //执行耗时（单位：s)
        "passScenarioNum": res.passScenarioNum || 0, //通过场景数
        "failScenarioNum": res.failScenarioNum || 0, //失败场景数
        "passInterfaceNum": res.passInterfaceNum || 0,
        "failInterfaceNum": res.failInterfaceNum || 0,
        "totalRequestNum": res.totalRequestNum || 0,
        "passRequestNum": res.passRequestNum || 0,
        "failRequestNum": res.failRequestNum || 0,
        "passAssertionNum": res.passAssertionNum || 0, //通过检查点数
        "failAssertionNum": res.failAssertionNum || 0, //失败检查点数
        "totalScenarioNum": data.totalScenarioNum || 0, //场景总数
        "totalInterfaceNum": data.totalInterfaceNum || 0, //接口总数
      }
      console.log('statisticData', statisticData.value);
      console.log('832res',res);
      reportId.value = res.id;
      recordList.value = formatData(res?.logs?.[0]?.logs || []);
    }
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
