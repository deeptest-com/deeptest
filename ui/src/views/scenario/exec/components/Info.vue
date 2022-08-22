<template>
  <div class="scenario-exec-info-main">
    <div class="scenario">
      <div class="header">
        <div class="title">
          {{execResult.name}}
        </div>

        <div class="progress">
          {{execResult.progressStatus ? t(execResult.progressStatus) : ''}}
        </div>

        <div class="status">
          {{execResult.resultStatus ? t(execResult.resultStatus) : ''}}
        </div>
        <div class="opt">
          <a-button v-if="execResult.progressStatus !== 'in_progress'" @click="execStart" type="link">开始执行</a-button>
          <a-button v-if="execResult.progressStatus === 'in_progress'" @click="execCancel" type="link">停止执行</a-button>
        </div>
      </div>

      <div class="logs">
        <Log v-if="logTreeData.logs" :logs="logTreeData.logs"></Log>
      </div>

      <div v-if="result.totalRequestNum > 0" class="result">
        <a-row>
          <a-col :span="4">开始时间</a-col>
          <a-col :span="4">{{ momentUtcDef(result.startTime) }}</a-col>
          <a-col :span="4">结束时间</a-col>
          <a-col :span="4">{{ momentUtcDef(result.endTime) }}</a-col>
          <a-col :span="4">耗时</a-col>
          <a-col :span="4">{{result.duration}}秒</a-col>
        </a-row>

        <a-row>
          <a-col :span="4">断言数</a-col>
          <a-col :span="4">{{result.totalAssertionNum}}</a-col>
          <a-col :span="4">通过数</a-col>
          <a-col :span="4">{{result.passAssertionNum}}</a-col>
          <a-col :span="4">失败数</a-col>
          <a-col :span="4">{{result.failAssertionNum}}</a-col>
        </a-row>

        <a-row>
          <a-col :span="4">请求数</a-col>
          <a-col :span="4">{{result.totalRequestNum}}</a-col>
          <a-col :span="4">成功数</a-col>
          <a-col :span="4">{{result.passRequestNum}}</a-col>
          <a-col :span="4">失败数</a-col>
          <a-col :span="4">{{result.failRequestNum}}</a-col>
        </a-row>

        <a-row>
          <a-col :span="4">接口数</a-col>
          <a-col :span="4">{{result.totalInterfaceNum}}</a-col>
          <a-col :span="4">成功数</a-col>
          <a-col :span="4">{{result.passInterfaceNum}}</a-col>
          <a-col :span="4">失败数</a-col>
          <a-col :span="4">{{result.failInterfaceNum}}</a-col>
        </a-row>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {WsMsg} from "@/types/data";

import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import bus from "@/utils/eventBus";
import Log from "./Log.vue"
import { momentUtcDef } from "@/utils/datetime";
import {useI18n} from "vue-i18n";
const { t } = useI18n();

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType, Global: GlobalStateType, Exec: ExecStatus; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const execResult = computed<any>(()=> store.state.Scenario.execResult);

const scenarioId = ref(+router.currentRoute.value.params.id)
store.dispatch('Scenario/loadExecResult', scenarioId.value);

const execStart = () => {
  console.log('execStart')
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', id: scenarioId.value}))
}

const execCancel = () => {
  console.log('execCancel')
  const msg = {act: 'stop', id: scenarioId.value}
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
}

onMounted(() => {
  console.log('onMounted exec info')
  bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
})

onUnmounted(() => {
  console.log('onUnmounted exec info')
  bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
})

const result = ref({} as any)
const logMap = ref({} as any)
const logTreeData = ref({} as any)
const OnWebSocketMsg = (data: any) => {
  console.log('WebsocketMsgEvent in exec info')

  const wsMsg = JSON.parse(data.msg) as WsMsg
  if (wsMsg.category == 'result') {
    result.value = wsMsg.data
    return

  } else if (wsMsg.category != '') {
    execResult.value.progressStatus = wsMsg.category
    if (wsMsg.category === 'in_progress') result.value = {}
    return

  }

  const log = wsMsg.data
  logMap.value[log.id] = log

  if (log.parentId === 0) {
    logTreeData.value = log
  } else {
    if (!logMap.value[log.parentId]) logMap.value[log.parentId] = {}
    if (!logMap.value[log.parentId].logs) logMap.value[log.parentId].logs = []

    logMap.value[log.parentId].logs.push(log)
  }

  // console.log(2, logTreeData)

  // const msgText = wsMsg.msg
  // store.dispatch('Scenario/updateExecResult', wsMsg.data).then(() => {
  //   console.log('===', msgText, execResult.value)
  // })

  // if ('isRunning' in wsMsg) {
  //   console.log(`change isRunning to ${wsMsg.isRunning}`)
  //   store.dispatch('Exec/setRunning', wsMsg.isRunning)
  // }
  //
  // if (item.info?.status === 'start') {
  //   const key = item.info.key + '-' + caseCount.value
  //   caseDetail.value[key] = logContentExpand.value
  // }
  //
  // item = genExecInfo(item, caseCount.value)
  // if (item.info && item.info.key && isInArray(item.info.status, ['pass', 'fail', 'skip'])) { // set case result
  //   store.dispatch('Result/list', {
  //     keywords: '',
  //     enabled: 1,
  //     pageSize: 10,
  //     page: 1
  //   });
  //   caseResult.value[item.info.key] = item.info.status
  // }
  //
  // wsMsg.out.push(item)
  // scroll('log-list')
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
        width: 200px;
        text-align: right;
      }
    }
    .logs {
      padding: 0px 12px;
    }
    .result {
      padding: 6px 12px;
      .ant-row {
        margin: 6px 0;
      }
    }
  }
}

</style>