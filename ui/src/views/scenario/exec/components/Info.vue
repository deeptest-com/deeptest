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
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import settings from "@/config/settings";
import {WebSocket} from "@/services/websocket";
import {WebSocketData} from "@/store/websoket";
import {WsMsg} from "@/types/data";

import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ScenarioStateType} from "../../store";
import bus from "@/utils/eventBus";
import Log from "./Log.vue"
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

const logMap = ref({} as any)
const logTreeData = ref({} as any)
const OnWebSocketMsg = (data: any) => {
  console.log('WebsocketMsgEvent in exec info')

  const wsMsg = JSON.parse(data.msg) as WsMsg
  if (wsMsg.category != '') {
    execResult.value.progressStatus = wsMsg.category
    return
  }

  const log = wsMsg.data
  console.log(1, log)
  logMap.value[log.id] = log

  if (log.parentId === 0) {
    logTreeData.value = log
  } else {
    if (!logMap.value[log.parentId]) logMap.value[log.parentId] = {}
    if (!logMap.value[log.parentId].logs) logMap.value[log.parentId].logs = []

    logMap.value[log.parentId].logs.push(log)
  }

  console.log(2, logTreeData)

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
  }
}

</style>