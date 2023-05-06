<template>
  <div id="request-main">
    <RequestInvocation
        :onSend="invokeInterface"
        :onSave="saveInterface">
    </RequestInvocation>

    <RequestConfig></RequestConfig>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import RequestInvocation from './Invocation.vue';
import RequestConfig from './Config.vue';
import {NotificationKeyCommon} from "@/utils/const";
import {UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
import {DebugInfo} from "@/views/component/debug/data";
const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<DebugInfo>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);

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

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>