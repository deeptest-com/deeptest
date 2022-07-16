<template>
  <div id="designer-interface-main">
      <div id="top-panel">
        <InterfaceRequest v-if="interfaceData.id"></InterfaceRequest>
      </div>

      <div id="design-splitter-v" :hidden="!interfaceData.id"></div>

      <div id="bottom-panel">
        <InterfaceResponse v-if="interfaceData.id"></InterfaceResponse>
      </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, onUnmounted, PropType, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler, resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType} from "@/views/interface/store";
import InterfaceRequest from './request/Index.vue';
import InterfaceResponse from './response/Index.vue';
import {Interface} from "@/views/interface/data";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import debounce from "lodash.debounce";

    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    onMounted(() => {
      console.log('onMounted interface')
      resize()

      window.addEventListener('resize', resizeHandler)
    })
    onUnmounted(() => {
      console.log('onUnmounted interface')
      window.removeEventListener('resize', resizeHandler)
    })

    watch(interfaceData, () => {
      console.log('watch interfaceData')
      // resize()
    }, {deep: true})

    const resize = () => {
      resizeHeight('designer-interface-main', 'top-panel', 'design-splitter-v', 'bottom-panel',
          200, 100, 50)
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