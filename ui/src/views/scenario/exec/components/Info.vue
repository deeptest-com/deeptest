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
          <a-button v-if="execResult.progressStatus !== 'in_progress'" @click="execStart" type="link">
            开始执行
          </a-button>

          <a-button v-if="execResult.progressStatus === 'in_progress'" @click="execCancel" type="link">
            停止执行
          </a-button>

          <a-button @click="design" type="link" class="dp-btn-group">设计</a-button>
          <a-button href="#/scenario/index" type="link" class="dp-btn-group">返回</a-button>
        </div>
      </div>

      <div class="logs">
        <Log v-if="logTreeData.logs" :logs="logTreeData.logs"></Log>
      </div>

      <div v-if="result.startTime" class="result">
        <a-row>
          <a-col :span="4">开始时间</a-col>
          <a-col :span="4">{{ momentShort(result.startTime) }}</a-col>
          <a-col :span="4">结束时间</a-col>
          <a-col :span="4">{{ momentShort(result.endTime) }}</a-col>
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
import { momentShort } from "@/utils/datetime";
import {useI18n} from "vue-i18n";
import {getToken} from "@/utils/localToken";
import {WsMsgCategory} from "@/utils/enum";
const { t } = useI18n();

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType, Global: GlobalStateType, Exec: ExecStatus; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);

const scenarioId = ref(+router.currentRoute.value.params.id)
store.dispatch('Scenario/loadExecResult', scenarioId.value);

const execStart = async () => {
  console.log('execStart')

  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    scenarioId: scenarioId.value,
  }

  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', scenarioExecReq: data}))
}

const execCancel = () => {
  console.log('execCancel')
  const msg = {act: 'stop', execReq: {scenarioId: scenarioId.value}}
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
}

const design = () => {
  console.log('design')
  router.push(`/scenario/design/${scenarioId.value}`)
}

onMounted(() => {
  console.log('onMounted exec info')
  bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
})

onUnmounted(() => {
  console.log('onUnmounted exec info')
  bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
})

const execResult = computed<any>(()=> store.state.Scenario.execResult);
const result = ref({} as any)
const logMap = ref({} as any)
const logTreeData = ref({} as any)
const OnWebSocketMsg = (data: any) => {
  console.log('--- WebsocketMsgEvent in exec info', data)

  if (!data.msg) return

  const wsMsg = JSON.parse(data.msg) as WsMsg

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