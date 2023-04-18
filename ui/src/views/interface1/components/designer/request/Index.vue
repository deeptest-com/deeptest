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
import {computed, ComputedRef, defineComponent, inject, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import {StateType} from "@/views/interface1/store";
import RequestInvocation from './Invocation.vue';
import RequestConfig from './Config.vue';
import {Interface} from "@/views/interface1/data";
import {NotificationKeyCommon} from "@/utils/const";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
import {getToken} from "@/utils/localToken";
import {prepareDataForRequest} from "@/views/interface1/service";

const useForm = Form.useForm;

const usedBy = inject('usedBy') as UsedBy

const {t} = useI18n();
const store = useStore<{ Interface1: StateType; Scenario: ScenarioStateType }>();
const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.InterfaceDebug ? store.state.Interface1.interfaceData : store.state.Scenario.interfaceData);

const invokeInterface = async () => {
  console.log('invokeInterface', interfaceData.value)

  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    id: interfaceData.value.id,
    // usedBy: usedBy,
    data: prepareDataForRequest(Object.assign({usedBy}, interfaceData.value)),
  }

  usedBy === UsedBy.InterfaceDebug ? store.dispatch('Interface1/invokeInterface', data) :
    store.dispatch('Scenario/invokeInterface', data)
};

const saveInterface = (data) => {
  console.log('saveInterface', data)

  if (usedBy === UsedBy.InterfaceDebug) {
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