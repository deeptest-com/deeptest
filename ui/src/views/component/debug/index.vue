<template>
  <div id="debug-index" class="dp-splits-v">
    <div id="debug-content">
      <Invocation :topVal="topVal"
      :showMethodSelection = "showMethodSelection"
                  :onSave="saveDebugData"
                  :onSync="syncDebugData" />
      <DebugConfig />
    </div>

    <div id="debug-splitter" class="splitter"></div>

    <div id="debug-right">
      <a-tabs tabPosition="right" class="right-tab"
              v-model:activeKey="rightTabKey"
              :tabBarGutter="0"
              @change="changeRightTab">

        <a-tab-pane key="env">
          <template #tab>
            <span id="env-tab">
              <a-tooltip placement="left" overlayClassName="dp-tip-small">
                <template #title>环境</template>
                <EnvironmentOutlined/>
              </a-tooltip>
            </span>
          </template>
        </a-tab-pane>

        <a-tab-pane key="history">
          <template #tab>
            <span id="his-tab">
              <a-tooltip placement="left" overlayClassName="dp-tip-small">
                <template #title>历史</template>
                <HistoryOutlined/>
              </a-tooltip>
            </span>
          </template>
        </a-tab-pane>

      </a-tabs>
    </div>

    <div v-if="rightTabKey==='env'"
         :style="posStyleEnv"
         class="right-float-tab dp-bg-white">
      <div class="dp-bg-light">
        <RequestEnv :onClose="closeRightTab" />
      </div>
    </div>
    <div v-if="rightTabKey==='history'"
         :style="posStyleHis"
         class="right-float-tab dp-bg-white">
      <div class="dp-bg-light">
        <RequestHistory :onClose="closeRightTab" />
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, PropType, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import {EnvironmentOutlined, HistoryOutlined} from '@ant-design/icons-vue';
import Invocation from '@/views/component/debug/request/Invocation.vue'
import RequestEnv from '@/views/component/debug/others/env/index.vue';
import RequestHistory from '@/views/component/debug/others/history/index.vue';
import DebugConfig  from './config.vue';

import {StateType as ProjectGlobal} from "@/store/project";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "../../endpoint/store";

import {StateType as GlobalStateType} from "@/store/global";
import {getRightTabPanelPosition, resizeWidth} from "@/utils/dom";

const {t} = useI18n();
const store = useStore<{  Debug: Debug, Endpoint: Endpoint, ProjectGlobal: ProjectGlobal, Global: GlobalStateType }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const props = defineProps({
  onSaveDebugData: {
    type: Function,
    required: true
  },
  onSyncDebugData: {
    type: Function,
    required: false
  },
  topVal: {
    type: String,
    required: true
  },
  showMethodSelection: {
    type: Boolean,
    required: true
  },
})

const rightTabKey = ref('')

const saveDebugData = async () => {
  props.onSaveDebugData()
};

const syncDebugData = async () => {
  if (props.onSyncDebugData)
    props.onSyncDebugData()
};

const posStyleEnv = ref({})
const posStyleHis = ref({})
onMounted(() => {
  console.log('onMounted in debug-index')
  resize()
})
onUnmounted(() => {
  console.log('onUnmounted in debug-index')
  store.dispatch('Debug/resetDataAndInvocations');
})

const resize = () => {
  resizeWidth('debug-index',
      'debug-content', 'debug-splitter', 'debug-right', 500, 38)
}

const changeRightTab = () => {
  console.log('changeRightTab')
  posStyleEnv.value = getRightTabPanelPosition('env-tab')
  posStyleHis.value = getRightTabPanelPosition('his-tab')
}

const closeRightTab = () => {
  rightTabKey.value = ''
}

</script>

<style lang="less">
#debug-index #debug-right .right-tab {
  //height: 100%;
  height: calc(100vh - 152px);
  .ant-tabs-left-content {
    padding-left: 0px;
  }

  .ant-tabs-right-content {
    padding-right: 0px;
    height: 100%;

    .ant-tabs-tabpane {
      height: 100%;

      &.ant-tabs-tabpane-inactive {
        display: none;
      }
    }
  }

  .ant-tabs-nav-scroll {
    text-align: center;
  }

  .ant-tabs-tab {
    padding: 5px 10px !important;

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
#debug-index {
  display: flex;
  height: 100%;
  width: 100%;

  #debug-content {
    flex: 1;
    width: 0;
    height: 100%;
  }

  #debug-right {
    width: 38px;
    height: 100%;
  }
  #debug-splitter {
    width: 1px;
    background-color: #f0f0f0;
  }
}
</style>
