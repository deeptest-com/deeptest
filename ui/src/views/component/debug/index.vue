<template>
  <div id="debug-form">
    <div id="top-panel">
      <InterfaceRequest v-if="debugData.method"
                        :showRequestInvocation="false"
                        :showDebugDataUrl="false" />
    </div>

    <div id="design-splitter-v" :hidden="!debugData.method" />

    <div id="bottom-panel">
      <InterfaceResponse v-if="debugData.method" />
    </div>

    <RequestVariable/>
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

const {t} = useI18n();
const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectGlobal, Endpoint: Endpoint }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

// onMounted(() => {
//   console.log('onMounted interface')
//   window.addEventListener('resize', resizeHandler)
//   resize()
// })
// onUnmounted(() => {
//   console.log('onUnmounted interface')
//   window.removeEventListener('resize', resizeHandler)
// })
//
// const resize = () => {
//   resizeHeight('debug-form', 'top-panel', 'design-splitter-v', 'bottom-panel',
//       200, 360)
// }

</script>

<style lang="less" scoped>
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
</style>