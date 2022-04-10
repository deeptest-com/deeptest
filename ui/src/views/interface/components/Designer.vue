<template>
  <div class="designer-main">
    <div id="design-content" v-if="interfaceData.method">
      <div id="top-panel">
        <InterfaceRequest></InterfaceRequest>
      </div>

      <div id="design-splitter-v"></div>

      <div id="bottom-panel">
        <InterfaceResponse></InterfaceResponse>
      </div>
    </div>

    <div class="design-right">
      <a-tabs v-model:activeKey="tabKey" size="small" class="tabs-bar-center">
        <a-tab-pane key="history" tab="历史">
          <div class="history">

          </div>
        </a-tab-pane>
        <a-tab-pane key="env" tab="环境">
          <RequestEnv></RequestEnv>
        </a-tab-pane>
      </a-tabs>
    </div>
  </div>

</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {resizeHeight, resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType} from "@/views/interface/store";
import InterfaceRequest from './designer/request/Index.vue';
import InterfaceResponse from './designer/response/Index.vue';
import RequestEnv from './designer/others/env/index.vue';
import {Interface} from "@/views/interface/data";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
  },
  components: {
    InterfaceRequest, InterfaceResponse,
    RequestEnv,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const tabKey = ref('history')

    onMounted(() => {
      console.log('onMounted')

      resizeHeight('design-content', 'top-panel', 'design-splitter-v', 'bottom-panel',
          200, 100, 0)
    })

    return {
      interfaceData,
      tabKey,
    }
  }
})
</script>

<style lang="less">
.tabs-bar-center {
  .ant-tabs-nav-scroll {
    text-align: center;
  }
}
</style>

<style lang="less" scoped>
.designer-main {
  display: flex;
  height: 100%;

  #design-content {
    flex: 1;

    display: flex;
    flex-direction: column;
    position: relative;
    height: 100%;

    #top-panel {
      padding: 0;
      flex: 4;
      width: 100%;
    }

    #bottom-panel {
      flex: 6;
      padding: 4px;
      width: 100%;
      overflow: auto;
    }

    #design-splitter-v {
      width: 100%;
      height: 3px;
      background-color: #e6e9ec;
      cursor: ns-resize;

      &:hover {
        height: 3px;
        background-color: #D0D7DE;
      }

      &.active {
        height: 3px;
        background-color: #a9aeb4;
      }
    }
  }

  .design-right {
    width: 260px;

    border-left: solid 2px #e6e9ec;
  }
}

</style>