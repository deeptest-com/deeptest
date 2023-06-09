<template>
  <div class="debug-main">
    <div id="debug-form">
      <div id="top-panel">
        <InterfaceRequest v-if="debugData.method" :show-reuqest-invocation="false" :show-debug-data-url="false"></InterfaceRequest>
      </div>

      <!-- <div id="design-splitter-v" :hidden="!debugData.method"></div> -->

      <div id="bottom-panel">
        <InterfaceResponse v-if="debugData.method"></InterfaceResponse>
      </div>

      <RequestVariable/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted } from "vue";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as ProjectGlobal} from "@/store/project";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "@/views/endpoint/store";

import InterfaceRequest from '@/views/component/debug/request/Index.vue';
import InterfaceResponse from '@/views/component/debug/response/Index.vue';
import RequestVariable from '@/components/Editor/RequestVariable.vue';

const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectGlobal, Endpoint: Endpoint }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

onMounted(() => {
  console.log('onMounted interface')
  window.addEventListener('resize', resizeHandler)
  resize()
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
