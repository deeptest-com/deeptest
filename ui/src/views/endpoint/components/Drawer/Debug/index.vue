<template>
  <div class="endpoint-debug">
    <a-radio-group v-model:value="selectedMethod" button-style="solid">
      <template v-for="method in requestMethodOpts" :key="method.value">
        <a-radio-button
            v-if="hasDefinedMethod(method.value)"
            :value="method.value"
            class="has-defined">
          {{ method.label }}
        </a-radio-button>
      </template>
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
const currInterface = computed<any>(() => store.state.Debug.currInterface);
const debugData = computed<any>(() => store.state.Debug.debugData);

const selectedMethod = ref('GET');

watch(() => {
  return selectedMethod.value
}, (newVal, oldVal) => {
  console.log('watch selectedMethod in debug', newVal)

  store.dispatch('Debug/loadDebugData', {
    endpointId: currEndpointId.value, interfaceId: currInterface.value.id,
  });
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
