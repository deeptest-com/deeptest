<template>
  <div id="designer-interface-main">
      <div id="top-panel">
        <InterfaceRequest v-if="interfaceData.id"></InterfaceRequest>
      </div>

      <div id="design-splitter-v" :hidden="!interfaceData.id"></div>

      <div id="bottom-panel">
        <InterfaceResponse v-if="interfaceData.id"></InterfaceResponse>
      </div>

    <RequestVariable />
  </div>
</template>

<script setup lang="ts">
import {
  computed,
  ComputedRef,
  provide,
  defineProps,
  onMounted,
  onUnmounted,
  PropType,
  Ref,
  ref,
  watch, inject
} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
import InterfaceRequest from './request/Index.vue';
import InterfaceResponse from './response/Index.vue';
import RequestVariable from '@/components/Editor/RequestVariable.vue';
import {Interface} from "@/views/interface/data";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Scenario: ScenarioStateType; Interface: InterfaceStateType }>();

const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.interface ? store.state.Interface.interfaceData : store.state.Scenario.interfaceData);

onMounted(() => {
  console.log('onMounted interface')
  resize()

  window.addEventListener('resize', resizeHandler)
})
onUnmounted(() => {
  console.log('onUnmounted interface')
  window.removeEventListener('resize', resizeHandler)
})

let id = interfaceData.value.id
watch(interfaceData, () => {
  console.log('watch interfaceData', interfaceData.value.id)

  if (interfaceData.value.id !== id) {
    store.dispatch('Interface/listValidExtractorVariableForInterface', usedBy)
  }
  id = interfaceData.value.id
}, {deep: true})

const resize = () => {
  resizeHeight('designer-interface-main', 'top-panel', 'design-splitter-v', 'bottom-panel',
      200, 100)
}

</script>

<style lang="less" scoped>
#designer-interface-main {
    flex: 1;

    flex-direction: column;
    position: relative;
    height: 100%;

    #top-panel {
      height: 50%;
      width: 100%;
      padding: 0;
    }

    #bottom-panel {
      height: 50%;
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

</style>