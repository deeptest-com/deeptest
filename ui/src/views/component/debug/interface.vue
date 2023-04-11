<template>
  <div class="debug-main">
    <Path/>

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
        <InterfaceRequest v-if="debugData.interfaceId"></InterfaceRequest>
      </div>

      <div id="design-splitter-v" :hidden="!debugData.interfaceId"></div>

      <div id="bottom-panel">
        <InterfaceResponse v-if="debugData.interfaceId"></InterfaceResponse>
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
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectGlobal, Endpoint: Endpoint }>();

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);

const currEndpointId = computed<number>(() => store.state.Debug.currEndpointId);
const currInterface = computed<any>(() => store.state.Debug.currInterface);
const debugData = computed<any>(() => store.state.Debug.debugData);

const selectedMethod = ref(currInterface.value?.method ? currInterface.value?.method : 'GET');

const changeMethod = () => {
  console.log('changeMethod', selectedMethod.value, interfaceMethodToObjMap)

  store.dispatch('Debug/setInterface', interfaceMethodToObjMap.value[selectedMethod.value]);
  store.dispatch('Debug/loadDebugData', {
    endpointId: currEndpointId.value, interfaceId: interfaceMethodToObjMap.value[selectedMethod.value].id,
  });
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
      200, 100)
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
    padding: 5px;

    flex-direction: column;
    position: relative;
    height: 100%;

    #top-panel {
      height: 50%;
      min-height: 200px;
      width: 100%;
      padding: 0;
    }

    #bottom-panel {
      height: 50%;
      min-height: 100px;
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