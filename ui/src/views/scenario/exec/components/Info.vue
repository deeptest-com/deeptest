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
          {{execData.status}}
        </div>
        <div class="opt">
          <a-button @click="exec" type="link">执行</a-button>
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
import {StateType as ScenarioStateType} from "../../store";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const execData = computed<any>(()=> store.state.Scenario.execData);

const scenarioId = ref(+router.currentRoute.value.params.id)

store.dispatch('Scenario/loadExecData', scenarioId.value);

const exec = () => {
  console.log('exec')
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execScenario', id: scenarioId.value}))
}

onMounted(() => {
  console.log('onMounted exec info')
})

onUnmounted(() => {
  console.log('onUnmounted')
})

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