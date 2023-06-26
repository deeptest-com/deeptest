<template>
  <div class="processor_interface_default-main">
    <div id="debug-index" class="dp-splits-v">
      <div id="debug-content">
        <DebugUrlAndEnv />

        <RequestInvocation
            :showDebugDataUrl="true"
            :onSend="invokeInterface"
            :onSave="saveScenarioInterface"
            :onSync="syncDebugData" />

        <DebugInterface />
      </div>

      <div id="debug-splitter" class="splitter"></div>

      <div id="debug-right">
        <a-tabs v-model:activeKey="tabKey"
                tabPosition="right"
                :tabBarGutter="0"
                class="right-tab">

          <a-tab-pane key="env">
            <template #tab>
              <a-tooltip placement="left" overlayClassName="dp-tip-small">
                <template #title>环境</template>
                <EnvironmentOutlined/>
              </a-tooltip>
            </template>

            <RequestEnv v-if="tabKey==='env'"></RequestEnv>
          </a-tab-pane>

          <a-tab-pane key="history">
            <template #tab>
              <a-tooltip placement="left" overlayClassName="dp-tip-small">
                <template #title>历史</template>
                <HistoryOutlined/>
              </a-tooltip>
            </template>

            <RequestHistory v-if="tabKey==='history'"></RequestHistory>
          </a-tab-pane>

        </a-tabs>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";


import {computed, onMounted, provide, ref} from "vue";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import { EnvironmentOutlined, HistoryOutlined } from '@ant-design/icons-vue';

import {UsedBy} from "@/utils/enum";
import {resizeWidth} from "@/utils/dom";
import {getToken} from "@/utils/localToken";
import {NotificationKeyCommon} from "@/utils/const";

import RequestEnv from '@/views/component/debug/others/env/index.vue';
import RequestHistory from '@/views/component/debug/others/history/index.vue';

import RequestInvocation from '@/views/component/debug/request/Invocation.vue';
import DebugUrlAndEnv from './debug/url-and-env.vue';
import DebugInterface from './debug/interface.vue';

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";

provide('usedBy', UsedBy.ScenarioDebug)

const {t} = useI18n();

const tabKey = ref('env')

const store = useStore<{  Debug: Debug, Scenario: Scenario }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const invokeInterface = async () => {
  console.log('invokeInterface', debugData.value)

  const callData = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),

    data: debugData.value
  }
  await store.dispatch('Debug/call', callData)
};

const saveScenarioInterface = async (data) => {
  const obj = Object.assign({}, data)
  delete obj.shareVars
  delete obj.envVars
  delete obj.globalEnvVars
  delete obj.globalParamVars

  const res = await store.dispatch('Scenario/saveDebugData', obj)
  if (res === true) {
    notification.success({
      key: NotificationKeyCommon,
      message: `保存成功`,
    });
  } else {
    notification.success({
      key: NotificationKeyCommon,
      message: `保存失败`,
    });
  }
};

const syncDebugData = async () => {
  console.log('syncDebugData')
  await store.dispatch('Debug/syncDebugData');
};

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('debug-index',
      'debug-content', 'debug-splitter', 'debug-right', 500, 260)
}

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;

  #debug-index {
  display: flex;
  height: 100%;
  width: 100%;

  #debug-content {
    flex: 1;
    width: 0;
    height: 100%;

    .move-up {
      margin-top: -43px;
    }
  }

  #debug-right {
    width: 260px;
    height: 100%;
  }

  .switcher {
    position: fixed;
    right: 8px;
    bottom: 5px;
    cursor: pointer;
  }
}
}
</style>

<style lang="less">
.processor_interface_default-main {
  #debug-index #debug-right .right-tab {
    height: 100%;

    .ant-tabs-left-content {
      padding-left: 0px;
    }

    .ant-tabs-right-content {
      padding-right: 0px;
      height: 100%;

      .ant-tabs-tabpane {
        height: 100%;

        &.ant-tabs-tabpane-inactive {
          display: none;
        }
      }
    }

    .ant-tabs-nav-scroll {
      text-align: center;
    }

    .ant-tabs-tab {
      padding: 5px 10px !important;

      .anticon {
        margin-right: 2px !important;
      }
    }

    .ant-tabs-ink-bar {
      background-color: #d9d9d9 !important;
    }
  }
}
</style>
