<template>
  <div id="request-main">
    <RequestConfig></RequestConfig>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {useStore} from "vuex";
import {UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import RequestConfig from './Config.vue';

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

// const invokeInterface = async () => {
//   console.log('invokeInterface', debugData.value)
//
//   const callData = {
//     serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
//     token: await getToken(),
//
//     data: debugData.value
//   }
//   await store.dispatch('Debug/call', callData)
// };
//
// const saveInterface = async (data) => {
//   console.log('saveInterface', data)
//
//   const obj = Object.assign({}, data)
//   delete obj.shareVars
//   delete obj.envVars
//   delete obj.globalEnvVars
//   delete obj.globalParamVars
//
//   const res = await store.dispatch('Debug/save', obj)
//   if (res === true) {
//     notification.success({
//       key: NotificationKeyCommon,
//       message: `保存成功`,
//     });
//   } else {
//     notification.success({
//       key: NotificationKeyCommon,
//       message: `保存失败`,
//     });
//   }
// };

// move to other files
// //保存场景调试数据
// const saveScenarioInterface = async (data) => {
//   const obj = Object.assign({}, data)
//   delete obj.shareVars
//   delete obj.envVars
//   delete obj.globalEnvVars
//   delete obj.globalParamVars
//
//   const res = await store.dispatch('Scenario/saveScenarioDebugData', obj)
//   if (res === true) {
//     notification.success({
//       key: NotificationKeyCommon,
//       message: `保存成功`,
//     });
//   } else {
//     notification.success({
//       key: NotificationKeyCommon,
//       message: `保存失败`,
//     });
//   }
// };
//
// //同步调试数据
// const syncDebugData = async () => {
//   await store.dispatch('Debug/loadDataAndInvocations', {
//     endpointInterfaceId: debugData.value.endpointInterfaceId,
//     scenarioProcessorId: 0,
//     usedBy: UsedBy.InterfaceDebug,
//   });
// };

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>
