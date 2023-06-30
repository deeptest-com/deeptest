<template>
  ={{endpointDetail?.interfaces[0].id}}=
  <div id="debug-index" class="dp-splits-v" v-if="endpointDetail?.interfaces?.length && endpointDetail?.interfaces[0].id">
    <div id="debug-top">
      <DebugMethod />

      <RequestInvocation
        :show-debug-data-url="false"
        :onSend="invokeInterface"
        :onSave="saveDebugInterface" />
    </div>

    <div id="debug-bottom">
      <div id="debug-content">
        <DebugUrlAndEnv />

        <DebugComp />
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
  <div v-else style="margin-top: 48px;">
    <a-empty
        image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
        :image-style="{height: '60px',}">
      <template #description>
      <span>
        您还未定义接口，请先定义接口才能使用调试功能
      </span>
      </template>
      <a-button type="primary" @click="emit('switchToDefineTab')">接口定义</a-button>
    </a-empty>
  </div>
</template>

<script setup lang="ts">
import {onMounted, provide, ref, computed,defineEmits} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import { EnvironmentOutlined, HistoryOutlined } from '@ant-design/icons-vue';

import {NotificationKeyCommon} from "@/utils/const";
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "@/views/endpoint/store";

import DebugMethod from './method.vue';
import DebugUrlAndEnv from './url-and-env.vue';
import RequestEnv from '@/views/component/debug/others/env/index.vue';
import RequestHistory from '@/views/component/debug/others/history/index.vue';
import RequestInvocation from '@/views/component/debug/request/Invocation.vue';

import DebugComp from '@/views/component/debug/index.vue';

const store = useStore<{  Debug: Debug,Endpoint:Endpoint }>();
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
provide('usedBy', UsedBy.InterfaceDebug)
const useForm = Form.useForm;
const debugData = computed<any>(() => store.state.Debug.debugData);
const {t} = useI18n();

const emit = defineEmits(['switchToDefineTab']);
const tabKey = ref('env')

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('debug-index',
      'debug-content', 'debug-splitter', 'debug-right', 500, 260)
}

const invokeInterface = async () => {
  console.log('invokeInterface', debugData.value)

  const callData = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),

    data: debugData.value
  }
  await store.dispatch('Debug/call', callData)
};

const saveDebugInterface = async (data) => {
  console.log('saveDebugInterface', data)

  Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null })

  const res = await store.dispatch('Debug/save', data)
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

</script>

<style lang="less">
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
</style>

<style lang="less" scoped>
#debug-index {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;

  #debug-top {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-right: 2px;
    margin-bottom: 20px;
  }

  #debug-bottom {
    display: flex;
  }

  #debug-content {
    flex: 1;
    width: 0;
    height: 100%;
  }

  #debug-right {
    width: 260px;
    height: 100%;
  }

  .splitter {
    min-width: 20px;
  }

  .switcher {
    position: fixed;
    right: 8px;
    bottom: 5px;
    cursor: pointer;
  }
}
</style>

