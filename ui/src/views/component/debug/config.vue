<template>
  <div id="debug-form-main">
    <ContentPaneHorizontal>
      <template #top>
        <div id="debug-form-top">
          <div v-if="serverConfig.demoTestSite" class="dp-red">
            您正在访问演示站点，所有的接口请求将被重定向到{{serverConfig.demoTestSite}}。
          </div>

          <InterfaceRequest v-if="debugData.method"
                            :showRequestInvocation="false"
                            :showDebugDataUrl="false" />
        </div>
      </template>
      <template #bottom>
        <div id="debug-form-bottom">
          <InterfaceResponse v-if="debugData.method" />
        </div>
      </template>
    </ContentPaneHorizontal>
    <VariableSelection/>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import {resizeHeight} from "@/utils/dom";
import {StateType as ProjectGlobal} from "@/store/project";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "../../endpoint/store";
import ContentPaneHorizontal from "@/views/component/ContentPaneHorizontal/index.vue";
import InterfaceRequest from './request/Index.vue';
import InterfaceResponse from './response/Index.vue';
import VariableSelection from './others/variable-replace/Selection.vue';
import {StateType as GlobalStateType} from "@/store/global";

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
#debug-form-main {
  height: calc(100% - 33px);
  display: flex;
  flex-direction: column;
  padding: 0;
  position: relative;

  .container {
    margin-top: 0px !important;
    margin-bottom: 0px !important;
    height: 100%;
  }

  #debug-form-top {
    height: 100%;
    width: 100%;
    padding: 0;
  }

  #debug-form-bottom {
    width: 100%;
    height: 100%;
    padding: 0;
    overflow: auto;
  }

  #debug-form-splitter {
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
