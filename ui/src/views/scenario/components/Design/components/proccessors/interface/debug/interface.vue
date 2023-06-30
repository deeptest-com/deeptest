<template>
  <div class="debug-main">
    <DebugComp />
  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";
import debounce from "lodash.debounce";

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";

import {UsedBy} from "@/utils/enum";
import DebugComp from '@/views/component/debug/index.vue';

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug, Scenario: Scenario }>();

const scenarioProcessorIdForDebug = computed<number>(() => store.state.Scenario.scenarioProcessorIdForDebug);
// const endpointInterfaceIdForDebug = computed<number>(() => store.state.Scenario.endpointInterfaceIdForDebug);
const debugData = computed<any>(() => store.state.Debug.debugData);

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

onMounted(() => {
  console.log('onMounted interface')
  window.addEventListener('resize', resizeHandler)
  resize()
})
onUnmounted(() => {
  console.log('onUnmounted interface')
  window.removeEventListener('resize', resizeHandler)
})

const resize = () => {
  resizeHeight('debug-form', 'top-panel', 'design-splitter-v', 'bottom-panel',
      200, 360)
}
</script>

<style lang="less" scoped>
.debug-main {
  .debug-methods {
    .has-defined {
      color: #1890ff;
      &.ant-radio-button-wrapper-checked {
        color: #FFF;
      }
    }
  }

  #debug-form {
    flex: 1;
    padding: 5px 0;

    flex-direction: column;
    position: relative;
    height: 100%;
    max-height: 800px;

    #top-panel {
      height: 50%;
      min-height: 200px;
      width: 100%;
      padding: 0;
    }

    #bottom-panel {
      height: 360px;
      width: 100%;
      padding: 4px;
      overflow: auto;
    }

    #design-splitter-v {
      width: 100%;
      height: 2px;
      background-color: #e6e9ec;
      cursor: ns-resize;

      &:hover {
        height: 2px;
        background-color: #D0D7DE;
      }

      &.active {
        height: 2px;
        background-color: #a9aeb4;
      }
    }
  }
}

</style>
