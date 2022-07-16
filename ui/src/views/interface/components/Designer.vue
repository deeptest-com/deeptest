<template>
  <div class="designer-main">
    <div id="design-content">
      <DesignInterface />
    </div>

    <div v-if="interfaceData.id && showRightBar" class="design-right">
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
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import { HistoryOutlined, EnvironmentOutlined, LeftCircleOutlined, RightCircleOutlined } from '@ant-design/icons-vue';
import {useStore} from "vuex";

import {StateType} from "@/views/interface/store";
import DesignInterface from './designer/Interface.vue';
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
    DesignInterface,
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
    })

    watch(interfaceData, () => {
      console.log('watch interfaceData')
    }, {deep: true})

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
  height: 100%;

  .ant-tabs-left-content {
    padding-left: 0px;
  }
  .ant-tabs-right-content {
    padding-right: 0px;
    height: 100%;
    overflow-y: auto;
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
    height: 100%;
    width: 100%;
  }

  .design-right {
    width: 260px;
    height: 100%;
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