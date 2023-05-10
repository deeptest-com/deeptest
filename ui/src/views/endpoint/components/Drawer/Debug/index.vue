<template>
  <div id="debug-index" class="dp-splits-v">
    <div id="debug-top">
      <DebugMehod />
      <RequestInvocation
        :show-debug-data-url="false"
        :onSend="invokeInterface"
        :onSave="saveInterface">
      </RequestInvocation>
    </div> 
    <div id="debug-bottom">
      <div id="debug-content">
        <DebugBasicInfo />
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
import {onMounted, provide, ref, computed} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import { EnvironmentOutlined, HistoryOutlined } from '@ant-design/icons-vue';
import DebugInterface from './interface.vue';
import DebugMehod from './method.vue';
import DebugBasicInfo from './basicInfo.vue';
import RequestEnv from '@/views/component/debug/others/env/index.vue';
import RequestHistory from '@/views/component/debug/others/history/index.vue';
import RequestInvocation from '@/views/component/debug/request/Invocation.vue';

import {NotificationKeyCommon} from "@/utils/const";
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";
import {DebugInfo} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

provide('usedBy', UsedBy.InterfaceDebug)
const useForm = Form.useForm;
const debugInfo = computed<DebugInfo>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const {t} = useI18n();

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
    id: debugData.value.id,

    data: debugInfo.value
  }
  await store.dispatch('Debug/call', callData)
};

const saveInterface = async (data) => {
  console.log('saveInterface', data)

  const obj = Object.assign({}, data)
  delete obj.shareVars
  delete obj.envVars
  delete obj.globalEnvVars
  delete obj.globalParamVars

  const res = await store.dispatch('Debug/save', obj)
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