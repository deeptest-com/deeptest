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
import {confirmToDo} from "@/utils/confirm";
import {scenarioTypeMapToText} from "@/views/scenario/components/Design/config";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {notifyError, notifySuccess} from "@/utils/notify";

provide('usedBy', UsedBy.ScenarioDebug)

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);
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

  const res = await store.dispatch('Scenario/saveDebugData', obj);

  if (res === true) {
    // 同步节点数据，更新节点名称或者 Method 方法
    // await store.dispatch('Scenario/getNode', nodeData.value);
    if(nodeData?.value?.method !== obj?.method && nodeData?.value?.processorID){
      store.commit('Scenario/setTreeDataMapItemProp', {
        id: nodeData.value.processorID,
        prop: 'method',
        value: obj.method,
      });
    }


    notifySuccess(`保存成功`);
  } else {
    notifyError(`保存失败`);
  }
};

const syncDebugData = async () => {
  console.log('syncDebugData')
  confirmToDo(`确定再次从` + getSrcName() + `？`, '已有数据将被删除。', async () => {
    await store.dispatch('Scenario/syncDebugData')
  })
};

function getSrcName() {
  const processorInterfaceSrc = nodeData.value?.processorInterfaceSrc;
  if (processorInterfaceSrc) {
    return scenarioTypeMapToText[processorInterfaceSrc] || '接口定义';
  }
  return scenarioTypeMapToText[nodeData.value?.processorType] || '接口定义';
}

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;

  .debug-index-wrapper {
    height: calc(100% - 40px);
  }
}
</style>
