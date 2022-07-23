<template>
  <div class="scenario-exec-info-main">
    <div class="scenario">
      <div class="header">
        <div class="title">
          {{execData.name}}
        </div>

        <div class="progress">
          {{execData.progress}}
        </div>

        <div class="status">
          {{isRunning}}
        </div>
        <div class="opt">
          <a-button v-if="isRunning === 'false'" @click="execStart" type="link">开始执行</a-button>
          <a-button v-if="isRunning === 'true'" @click="execStop" type="link">停止执行</a-button>
        </div>
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

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType, Global: GlobalStateType, Exec: ExecStatus; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const execData = computed<any>(()=> store.state.Scenario.execData);
const isRunning = computed<string>(() => store.state.Exec.isRunning);

const scenarioId = ref(+router.currentRoute.value.params.id)

store.dispatch('Scenario/loadExecData', scenarioId.value);

const execStart = () => {
  console.log('execStart')
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', id: scenarioId.value}))
}

const execStop = () => {
  console.log('execStop')
  const msg = {act: 'stop'}
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

const OnWebSocketMsg = (data: any) => {
  console.log('WebsocketMsgEvent in exec info', data.msg)

  const wsMsg = JSON.parse(data.msg) as WsMsg

  if ('isRunning' in wsMsg) {
    console.log(`change isRunning to ${wsMsg.isRunning}`)
    store.dispatch('Exec/setRunning', wsMsg.isRunning)
  }

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
  }
}

</style>