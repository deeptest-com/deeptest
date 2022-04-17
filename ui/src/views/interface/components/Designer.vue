<template>
  <div v-if="interfaceData.id" class="designer-main">
    <div id="design-content" v-if="interfaceData.method">
      <div id="top-panel">
        <InterfaceRequest></InterfaceRequest>
      </div>

      <div id="design-splitter-v"></div>

      <div id="bottom-panel">
        <InterfaceResponse></InterfaceResponse>
      </div>
    </div>

    <div v-if="showRightBar" class="design-right">
      <a-tabs v-model:activeKey="tabKey"
              tabPosition="right"
              :tabBarGutter="0"
              class="right-tab">
        <a-tab-pane key="history">
          <template #tab><HistoryOutlined /></template>
          <RequestHistory></RequestHistory>
        </a-tab-pane>

        <a-tab-pane key="env">
          <template #tab><EnvironmentOutlined /></template>
          <RequestEnv></RequestEnv>
        </a-tab-pane>
      </a-tabs>
    </div>

    <div class="switcher">
      <LeftCircleOutlined v-if="!showRightBar" @click="show" />
      <RightCircleOutlined v-if="showRightBar" @click="show" />
    </div>
  </div>

</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import { HistoryOutlined, EnvironmentOutlined, LeftCircleOutlined, RightCircleOutlined } from '@ant-design/icons-vue';
import {resizeHeight, resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType} from "@/views/interface/store";
import InterfaceRequest from './designer/request/Index.vue';
import InterfaceResponse from './designer/response/Index.vue';
import RequestEnv from './designer/others/env/index.vue';
import RequestHistory from './designer/others/history/index.vue';
import {Interface} from "@/views/interface/data";
import {getShowRightBar, setShowRightBar} from "@/utils/cache";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
  },
  components: {
    HistoryOutlined, EnvironmentOutlined, LeftCircleOutlined, RightCircleOutlined,
    InterfaceRequest, InterfaceResponse,
    RequestEnv, RequestHistory,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const tabKey = ref('history')

    const showRightBar = ref(false)
    getShowRightBar().then((val) => {
      showRightBar.value = val
    })

    const show = (e) => {
      console.log('show')
      showRightBar.value = !showRightBar.value
      setShowRightBar(showRightBar.value)
    }

    onMounted(() => {
      console.log('onMounted')

      resizeHeight('design-content', 'top-panel', 'design-splitter-v', 'bottom-panel',
          200, 100, 50)
    })

    return {
      interfaceData,
      tabKey,
      showRightBar,
      show,
    }
  }
})
</script>

<style lang="less">
.right-tab {
  .ant-tabs-left-content {
    padding-left: 0px;
  }
  .ant-tabs-right-content {
    padding-right: 0px;
  }
  .ant-tabs-nav-scroll {
    text-align: center;
  }
  .ant-tabs-tab {
    padding: 6px 10px !important;
    .anticon {
      margin-right: 2px !important;
    }
  }
  .ant-tabs-ink-bar {
    background-color: #d9d9d9 !important;
  }
}
</style>

<style lang="less" scoped>
.designer-main {
  display: flex;
  height: 100%;

  #design-content {
    flex: 1;

    flex-direction: column;
    position: relative;
    height: 100%;

    display: flex;
    #top-panel {
      height: 200px;
      padding: 0;
      width: 100%;
    }

    #bottom-panel {
      flex: 1;
      padding: 4px;
      width: 100%;
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

  .design-right {
    width: 260px;
    border-left: solid 1px #e6e9ec;
  }

  .switcher {
    position: fixed;
    right: 8px;
    bottom: 5px;
    cursor: pointer;
  }
}

</style>