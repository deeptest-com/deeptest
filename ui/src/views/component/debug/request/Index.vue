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
import {prepareDataForRequest} from "@/views/interface1/service";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const invokeInterface = async () => {
  console.log('invokeInterface', debugData.value)

  // const reqData = prepareDataForRequest(Object.assign({usedBy}, debugData.value))
  // await saveInterface(reqData)

  const callData = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    id: debugData.value.id,

    data: {
      endpointId: debugData.value.endpointId,
      interfaceId: debugData.value.interfaceId,
      usedBy,
    }
  }
  await store.dispatch('Debug/invokeInterface', callData)
};

const saveInterface = async (data) => {
  console.log('saveInterface', data)
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

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>