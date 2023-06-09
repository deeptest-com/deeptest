<template>
  <div id="request-main">
    <!-- 最新版本ui中 发送/保存按钮 调整为全局的保存，这里为了避免 scenario/design中引用出错，用这个条件判断先放着。后续再看是否移除掉 -->
    <template v-if="showRequestInvocation">
      <RequestInvocation
        :showDebugDataUrl="showDebugDataUrl"
        :onSend="invokeInterface"
        :onSave="saveInterface"
        :onSaveScenarioInterface="saveScenarioInterface"
        :onSync="syncDebugData"
      >
      </RequestInvocation>
    </template>

    <RequestConfig></RequestConfig>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted, defineProps} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import RequestInvocation from './Invocation.vue';
import RequestConfig from './Config.vue';
import {NotificationKeyCommon} from "@/utils/const";
import {UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";

defineProps({
  showRequestInvocation: {
    type: Boolean,
    required: false,
    default: true
  },
  showDebugDataUrl: {
    type: Boolean,
    required: false,
    default: true
  }
});

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import {DebugInfo} from "@/views/component/debug/data";
const store = useStore<{  Debug: Debug, Scenario: Scenario }>();

const debugInfo = computed<DebugInfo>(() => store.state.Debug.debugInfo);
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

//保存场景调试数据
const saveScenarioInterface = async (data) => {
  const obj = Object.assign({}, data)
  delete obj.shareVars
  delete obj.envVars
  delete obj.globalEnvVars
  delete obj.globalParamVars

  const res = await store.dispatch('Scenario/saveScenarioDebugData', obj)
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
  await store.dispatch('Debug/loadDataAndInvocations', {
    endpointInterfaceId: debugData.value.endpointInterfaceId,
    scenarioProcessorId: 0,
    usedBy: UsedBy.InterfaceDebug,
  });
};

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>
