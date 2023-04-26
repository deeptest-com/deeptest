<template>
  <div class="debug-main">
    <div id="debug-form">
      <div id="top-panel">
        <InterfaceRequest v-if="debugData.method"></InterfaceRequest>
      </div>

      <div id="design-splitter-v" :hidden="!debugData.method"></div>

      <div id="bottom-panel">
        <InterfaceResponse v-if="debugData.method"></InterfaceResponse>
      </div>

      <RequestVariable/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";

import InterfaceRequest from '@/views/component/debug/request/Index.vue';
import InterfaceResponse from '@/views/component/debug/response/Index.vue';
import RequestVariable from '@/components/Editor/RequestVariable.vue';

import {UsedBy} from "@/utils/enum";
import {DebugInfo} from "@/views/component/debug/data";
import debounce from "lodash.debounce";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug, Scenario: Scenario }>();

const endpointInterfaceId = computed<number>(() => store.state.Scenario.endpointInterfaceId);
const debugInfo = computed<DebugInfo>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);

const loadData = debounce(async () => {
  console.log('loadData', endpointInterfaceId.value)
  if (endpointInterfaceId.value === 0) {
    return
  }

  // await store.dispatch('Debug/loadData', {
  //   endpointInterfaceId: 0,
  //   scenarioProcessorId: 0,
  //   usedBy: usedBy,
  // });
  //
  // store.dispatch('Debug/getLastInvocationResp', {
  //   endpointInterfaceId: debugInfo.value.endpointInterfaceId,
  // })
  // store.dispatch('Debug/listInvocation', {
  //   endpointInterfaceId: debugInfo.value.endpointInterfaceId,
  // })
}, 300)
loadData()

watch(endpointInterfaceId, () => {
  console.log('watch endpointInterfaceId', endpointInterfaceId)
}, {deep: true})

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