<template>
  <div class="endpoint-debug">
    <a-radio-group v-model:value="selectedMethod" button-style="solid">
      <a-radio-button
          :class="{'has-defined': hasDefinedMethod(method.value)}"
          :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
        {{ method.label }}
      </a-radio-button>
    </a-radio-group>

    <div>{{currEndpointId}}</div>
    <div>{{currInterface}}</div>
    <div>{{debugData}}</div>
  </div>
</template>

<script setup lang="ts">
import {
  ref,
  defineProps,
  computed, watch,
} from 'vue';
import {useStore} from "vuex";

import {
  requestMethodOpts,
  mediaTypesOpts,
  repCodeOpts,
  defaultCookieParams,
  defaultHeaderParams,
  defaultQueryParams,
  defaultPathParams,
  defaultEndpointDetail,
  defaultCodeResponse,
} from '@/config/constant';

import {StateType as Endpoint} from "../../../store";
import {StateType as Debug} from "@/store/debug";
import {StateType as ProjectGlobal} from "@/store/project";

const props = defineProps({});
const store = useStore<{ Debug, ProjectGlobal, Endpoint  }>();

console.log(store.state.ProjectGlobal)

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);

const currEndpointId = computed<number>(() => store.state.Debug.currEndpointId);
const currInterface = computed<number>(() => store.state.Debug.currInterface);
const debugData = computed<any>(() => store.state.Debug.debugData);

const selectedMethod = ref('GET');

watch(() => {
  return selectedMethod.value
}, (newVal, oldVal) => {
  console.log('watch selectedMethod in debug', newVal)

  store.dispatch('Debug/loadDebugData', {endpointId: endpointDetail.value.id});
}, {immediate: true});

function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

</script>

<style lang="less" scoped>
.endpoint-debug {
  .has-defined {
    color: #1890ff;
    &.ant-radio-button-wrapper-checked {
      color: #FFF;
    }
  }
}
</style>
