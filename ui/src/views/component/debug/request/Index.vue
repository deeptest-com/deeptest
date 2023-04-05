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

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const invokeInterface = async () => {
  console.log('invokeInterface', debugData.value)

  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    id: debugData.value.id,
    // usedBy: usedBy,
    data: prepareDataForRequest(Object.assign({usedBy}, debugData.value)),
  }

  usedBy === UsedBy.interface ? store.dispatch('Interface1/invokeInterface', data) :
    store.dispatch('Scenario/invokeInterface', data)
};

const saveInterface = (data) => {
  console.log('saveInterface', data)

  if (usedBy === UsedBy.interface) {
    store.dispatch('Interface1/saveInterface', data).then((res) => {
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
    })
  } else {
    store.dispatch('Scenario/saveInterface', data).then((res) => {
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
    })
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