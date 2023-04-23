<template>
  <div class="debug-main">
    <div class="debug-methods">
      <a-radio-group @change="changeMethod" v-model:value="selectedMethod" button-style="solid">
        <template v-for="method in requestMethodOpts" :key="method.value">
          <a-radio-button
              v-if="hasDefinedMethod(method.value)"
              :value="method.value"
              class="has-defined">
            {{ method.label }}
          </a-radio-button>
        </template>
      </a-radio-group>
    </div>

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

    <!-- <div>{{currEndpointId}}</div>
        <hr>
        <div>{{currInterface}}</div>
        <hr>
        <div>{{debugData}}</div> -->

  </div>
</template>

<script setup lang="ts">
import {computed, inject, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";

import {requestMethodOpts} from '@/config/constant';
import {StateType as ProjectGlobal} from "@/store/project";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "../../endpoint/store";

import Path  from './path.vue';
import InterfaceRequest from './request/Index.vue';
import InterfaceResponse from './response/Index.vue';
import RequestVariable from '@/components/Editor/RequestVariable.vue';

import {UsedBy} from "@/utils/enum";
import {DebugInfo} from "@/views/component/debug/data";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectGlobal, Endpoint: Endpoint }>();

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
const selectedMethodDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);

const debugInfo = computed<DebugInfo>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);

const selectedMethod = ref(selectedMethodDetail.value?.method ? selectedMethodDetail.value?.method : 'GET');

const changeMethod = async () => {
  console.log('changeMethod', selectedMethod.value, interfaceMethodToObjMap)

  const endpointInterface = interfaceMethodToObjMap.value[selectedMethod.value]

  if (endpointInterface) {
    await store.commit('Endpoint/setSelectedMethodDetail', endpointInterface);
  } else {
    await store.commit('Endpoint/setSelectedMethodDetail', {});
  }

  await store.dispatch('Debug/loadData', {
    endpointInterfaceId: endpointInterface.id,
    // scenarioProcessorId: 0, // TODO: set in in scenario designer
    usedBy: usedBy,
  });

  store.dispatch('Debug/getLastInvocationResp', {
    endpointInterfaceId: debugInfo.value.endpointInterfaceId,
  })
  store.dispatch('Debug/listInvocation', {
    endpointInterfaceId: debugInfo.value.endpointInterfaceId,
  })
}
changeMethod()

function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

onMounted(() => {
  console.log('onMounted interface')
  resize()

  window.addEventListener('resize', resizeHandler)
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