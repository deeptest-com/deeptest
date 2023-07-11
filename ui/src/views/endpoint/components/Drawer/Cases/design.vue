<template>
  <div class="endpoint-debug-cases-design-main">
    <div id="endpoint-debug-cases-design-panel">
      <div id="endpoint-debug-cases-design-content">
        <UrlAndInvocation />
        <DebugComp />
      </div>

      <div id="endpoint-debug-cases-design-splitter" class="splitter"></div>

      <div id="endpoint-debug-cases-design-right">
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
import EnvSelection from '@/views/diagnose/design/env-selection.vue'
import UrlAndInvocation from '@/views/diagnose/design/url-and-invocation.vue'

import DebugComp from '@/views/component/debug/index.vue';

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as EndpointStateType} from '../../../store';

provide('usedBy', UsedBy.CaseDebug)
const usedBy = UsedBy.CaseDebug

const store = useStore<{ Debug: Debug, Endpoint: EndpointStateType }>();
const endpointCase = computed<any>(() => store.state.Endpoint.caseDetail);
const debugData = computed<any>(() => store.state.Debug.debugData);

const rightTabKey = ref('env')

const loadDebugData = debounce(async () => {
  console.log('loadDebugData', endpointCase.value.id)

  store.dispatch('Debug/loadDataAndInvocations', {
    caseInterfaceId: endpointCase.value.id,
    usedBy: usedBy,
  });
}, 300)


watch(endpointCase, (newVal) => {
  if (!endpointCase.value) return

  console.log('watch endpointCase', endpointCase.value.id)
  loadDebugData()
}, {immediate: true, deep: true})

</script>

<style lang="less">
.endpoint-debug-cases-design-main {
  .selection {
    position: absolute;
    top: 16px;
    right: 16px;
  }

  #endpoint-debug-cases-design-right .right-tab {
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
.endpoint-debug-cases-design-main {
  padding: 16px 0px 16px 16px;

  #endpoint-debug-cases-design-panel {
    display: flex;

    #endpoint-debug-cases-design-content {
      flex: 1;
      width: 0;
      height: 100%;
    }

    #endpoint-debug-cases-design-right {
      margin-top: 50px;
      width: 260px;
      height: 100%;
    }

    .endpoint-debug-cases-design-splitter {
      min-width: 20px;
    }
  }
}

</style>
