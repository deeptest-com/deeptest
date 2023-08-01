<template>
  <div id="debug-form">
    <div id="top-panel">
      <div v-if="serverConfig.demoTestSite" class="dp-red">
        您正在访问演示站点，所有的接口请求将被重定向到{{serverConfig.demoTestSite}}。
      </div>

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
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";

const {t} = useI18n();
const store = useStore<{  Debug: Debug, Endpoint: Endpoint, ProjectGlobal: ProjectGlobal, Global: GlobalStateType }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const serverConfig = computed<any>(() => store.state.Global.serverConfig);

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
  display: flex;
  flex-direction: column;
  height: calc(100% - 33px);

  padding: 0;
  position: relative;

  #top-panel {
    flex: 1;
    height: 0;

    width: 100%;
    padding: 0;
  }

  #bottom-panel {
    flex: 1;
    width: 100%;
    padding: 0;
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