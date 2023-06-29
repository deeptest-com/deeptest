<template>
  <div id="debug-form">{{debugInfo}}
    <div id="top-panel">
      <InterfaceRequest v-if="debugData.method"
                        :showRequestInvocation="false"
                        :showDebugDataUrl="false" />
    </div>

    <div id="design-splitter-v" :hidden="!debugData.method" />

    <div id="bottom-panel">
      <InterfaceResponse v-if="debugData.method" />
    </div>

    <VariableSelection/>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import {StateType as ProjectGlobal} from "@/store/project";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "../../endpoint/store";

import InterfaceRequest from './request/Index.vue';
import InterfaceResponse from './response/Index.vue';
import VariableSelection from './others/variable-replace/Selection.vue';

const {t} = useI18n();
const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectGlobal, Endpoint: Endpoint }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);

onMounted(() => {
  console.log('onMounted debug-interface')
})
onUnmounted(() => {
  console.log('onUnmounted debug-interface')
  store.dispatch('Debug/resetDataAndInvocations');
})

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
    width: 100%;
    padding: 0;
  }

  #bottom-panel {
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