<template>
    <div class="debug-methods">
      <a-radio-group @change="changeMethod" v-model:value="selectedMethod" button-style="solid">
        <template v-for="method in requestMethodOpts" :key="method.value">
          <a-radio-button v-if="hasDefinedMethod(method.value)" class="has-defined"
                          :value="method.value"
                          :style="{ color: method.color }">
            {{ method.label }}
          </a-radio-button>
        </template>
      </a-radio-group>
    </div>
</template>
<script setup lang="ts">
import {computed, ref, inject, watch, onMounted} from 'vue';
import {useStore} from "vuex";

import {requestMethodOpts} from '@/config/constant';
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "@/views/endpoint/store";
import {UsedBy} from "@/utils/enum";

const usedBy = inject('usedBy') as UsedBy;
const store = useStore<{  Debug: Debug, Endpoint: Endpoint }>();
const interfaceDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);

const selectedMethod = ref('GET');

const changeMethod = async () => {
  console.log('changeMethod', selectedMethod.value)
  const endpointInterface = interfaceMethodToObjMap.value[selectedMethod.value]

  // sync with / to define page
  if (endpointInterface?.id) {
    await store.commit('Endpoint/setSelectedMethodDetail', endpointInterface);

    store.dispatch('Debug/loadDataAndInvocations', {
      endpointInterfaceId: endpointInterface.id,
      usedBy: usedBy,
    });

  } else {
    await store.commit('Endpoint/setSelectedMethodDetail', {});
  }
}

const initMethod = async () => {
  await store.dispatch('Endpoint/removeUnSavedMethods')
  if (interfaceDetail.value?.method) {
    selectedMethod.value = interfaceDetail.value?.method
  }
  await changeMethod()
}
initMethod()

function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}
</script>
<style scoped lang="less">
.debug-methods {
    .has-defined {
      box-shadow: none;
      background: rgb(245, 245, 245);
      border-color: rgb(217, 217, 217);
      &:before {
        display: none;
      }
      &.ant-radio-button-wrapper-checked {
        color: #FFF;
        background-color: #fff;
      }
    }
  }
</style>
