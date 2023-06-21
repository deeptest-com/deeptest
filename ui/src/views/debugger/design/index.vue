<template>
  <div class="test-interface-design-main">
      <div v-if="debugData?.method" id="test-interface-debug-panel">
        <div id="test-interface-debug-content">
          <div class="tabs">
            <a-tabs v-model:activeKey="activeTabKey" @edit="onTabEdit" @change="changeTab" type="editable-card">
              <a-tab-pane v-for="tab in interfaceTabs" :key="''+tab.id" :tab="tab.title">

                <UrlAndInvocation />
                <DebugComp />

              </a-tab-pane>
            </a-tabs>
          </div>
        </div>

        <div id="test-interface-debug-splitter" class="splitter"></div>

        <div id="test-interface-debug-right">
          <a-tabs v-model:activeKey="rightTabKey"
                  tabPosition="right"
                  :tabBarGutter="0"
                  class="right-tab">

            <a-tab-pane key="env">
              <template #tab>
                <a-tooltip placement="left" overlayClassName="dp-tip-small">
                  <template #title>环境</template>
                  <EnvironmentOutlined/>
                </a-tooltip>
              </template>

              <RequestEnv v-if="rightTabKey==='env'"></RequestEnv>
            </a-tab-pane>

            <a-tab-pane key="history">
              <template #tab>
                <a-tooltip placement="left" overlayClassName="dp-tip-small">
                  <template #title>历史</template>
                  <HistoryOutlined/>
                </a-tooltip>
              </template>

              <RequestHistory v-if="rightTabKey==='history'"></RequestHistory>
            </a-tab-pane>
          </a-tabs>
        </div>
      </div>

      <div class="selection">
        <EnvSelection />
      </div>
  </div>
</template>

<script setup lang="ts">
import {computed, provide, ref, watch} from 'vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {UsedBy} from "@/utils/enum";
import { EnvironmentOutlined, HistoryOutlined } from '@ant-design/icons-vue';

import RequestEnv from '@/views/component/debug/others/env/index.vue';
import RequestHistory from '@/views/component/debug/others/history/index.vue';
import EnvSelection from './env-selection.vue'
import UrlAndInvocation from './url-and-invocation.vue'

import DebugComp from '@/views/component/debug/index.vue';

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as TestInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";

provide('usedBy', UsedBy.TestDebug)

const store = useStore<{ Debug: Debug, TestInterface: TestInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const debugData = computed<any>(() => store.state.Debug.debugData);

const interfaceId = computed<any>(() => store.state.TestInterface.interfaceId);
const interfaceData = computed<any>(() => store.state.TestInterface.interfaceData);
const interfaceTabs = computed<any>(() => store.state.TestInterface.interfaceTabs);

const activeTabKey = ref('0')
const rightTabKey = ref('env')

const changeTab = (key) => {
  console.log('changeTab', key)

  const found = interfaceTabs.value.find(function (item, index, arr) {
    return item.id === +key
  })

  store.dispatch('TestInterface/openInterfaceTab', found);
}

const usedBy = UsedBy.TestDebug
const loadDebugData = debounce(async () => {
  console.log('loadDebugData')

  store.dispatch('Debug/loadDataAndInvocations', {
    testInterfaceId: interfaceData.value.id,
    usedBy: usedBy,
  });
}, 300)

watch((interfaceData), async (newVal) => {
  console.log('watch interfaceData', interfaceData?.value)

  if (!interfaceData?.value) {
    store.dispatch('Debug/resetDataAndInvocations');
    return
  }

  loadDebugData()
  activeTabKey.value = ''+interfaceData.value.id
}, { immediate: true, deep: true })

const onTabEdit = (key, action) => {
  console.log('onTabEdit', key, action)
  if (action === 'remove') {
    store.dispatch('TestInterface/removeInterfaceTab', +key);
  }
};
</script>

<style lang="less">
.test-interface-design-main {
  .tabs {
    .ant-tabs-card {
      .ant-tabs-extra-content {
        margin-right: 160px;
      }
    }
  }
  .selection {
    position: absolute;
    top: 16px;
    right: 16px;
  }

  #test-interface-debug-right .right-tab {
    height: 100%;

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
}

</style>

<style scoped lang="less">
.test-interface-design-main {
  padding: 16px 0px 16px 16px;

  #test-interface-debug-panel {
    display: flex;

    #test-interface-debug-content {
      flex: 1;
      width: 0;
      height: 100%;
    }

    #test-interface-debug-right {
      margin-top: 50px;
      width: 260px;
      height: 100%;
    }

    .test-interface-debug-splitter {
      min-width: 20px;
    }
  }
}

</style>
