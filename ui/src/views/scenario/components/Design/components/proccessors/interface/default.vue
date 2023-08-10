<template>
  <div class="processor_interface_default-main dp-relative  dp-processors-container">
    <ProcessorHeader/>
    <DebugComp :onSaveDebugData="saveScenarioInterface"
               :onSyncDebugData="syncDebugData" />
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";

import {computed, inject, onMounted, provide, ref, watch} from "vue";
import {notification} from 'ant-design-vue';
import {useStore} from "vuex";

import {UsedBy} from "@/utils/enum";
import {NotificationKeyCommon} from "@/utils/const";
import DebugComp from '@/views/component/debug/index.vue';
import ProcessorHeader from '../../common/ProcessorHeader.vue';

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import debounce from "lodash.debounce";

provide('usedBy', UsedBy.ScenarioDebug)

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const scenarioProcessorIdForDebug = computed<number>(() => store.state.Scenario.scenarioProcessorIdForDebug);

watch(scenarioProcessorIdForDebug,  () => {
  console.log('watch scenarioProcessorIdForDebug', scenarioProcessorIdForDebug.value)
  loadData()
}, {immediate: true, deep: true})

function loadData() {
  debounce(async () => {
    console.log('loadData', scenarioProcessorIdForDebug.value)

    if (scenarioProcessorIdForDebug.value === 0) return

   await store.dispatch('Debug/loadDataAndInvocations', {
      scenarioProcessorId: scenarioProcessorIdForDebug.value,
      usedBy: usedBy,
    });

  }, 300)()
}

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
  await store.dispatch('Scenario/syncDebugData')
};

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;
}
</style>
