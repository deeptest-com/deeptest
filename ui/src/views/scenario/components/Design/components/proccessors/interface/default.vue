<template>
  <div class="processor_interface_default-main">
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
import {resizeWidth} from "@/utils/dom";
import {NotificationKeyCommon} from "@/utils/const";
import DebugComp from '@/views/component/debug/index.vue';

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import debounce from "lodash.debounce";

provide('usedBy', UsedBy.ScenarioDebug)

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const rightTabKey = ref('')

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const scenarioProcessorIdForDebug = computed<number>(() => store.state.Scenario.scenarioProcessorIdForDebug);

watch(scenarioProcessorIdForDebug, () => {
  console.log('watch scenarioProcessorIdForDebug', scenarioProcessorIdForDebug)
  loadData()
}, {deep: true})

const loadData = debounce(async () => {
  console.log('loadData', scenarioProcessorIdForDebug.value)
  if (scenarioProcessorIdForDebug.value === 0) {
    return
  }

  store.dispatch('Debug/loadDataAndInvocations', {
    scenarioProcessorId: scenarioProcessorIdForDebug.value,
    usedBy: usedBy,
  });

}, 300)
loadData()

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
  store.dispatch('Scenario/syncDebugData')
};

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('debug-index',
      'debug-content', 'debug-splitter', 'debug-right', 500, 38)
}

const closeRightTab = () => {
  rightTabKey.value = ''
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
    }

    #debug-right {
      width: 38px;
      height: 100%;
    }
  }
}
</style>

<style lang="less">
.processor_interface_default-main {

}
</style>
