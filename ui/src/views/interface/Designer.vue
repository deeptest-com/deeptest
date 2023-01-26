<template>
  <div id="designer-main" class="dp-splits-v">
    <div id="design-content">
      <DesignInterface />
    </div>

    <div id="design-splitter" class="splitter"></div>

    <div id="design-right">
      <a-tabs v-model:activeKey="tabKey"
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
          <RequestEnv></RequestEnv>
        </a-tab-pane>

        <a-tab-pane key="history">
          <template #tab>
            <a-tooltip placement="left" overlayClassName="dp-tip-small">
              <template #title>历史</template>
              <HistoryOutlined/>
            </a-tooltip>
          </template>
          <RequestHistory></RequestHistory>
        </a-tab-pane>

        <a-tab-pane key="datapool">
          <template #tab>
            <a-tooltip placement="left" overlayClassName="dp-tip-small">
              <template #title>数据池</template>
              <TableOutlined />
            </a-tooltip>
          </template>
          <Datapool></Datapool>
        </a-tab-pane>

      </a-tabs>
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, provide, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import { HistoryOutlined, EnvironmentOutlined, TableOutlined } from '@ant-design/icons-vue';
import {useStore} from "vuex";

import {StateType} from "@/views/interface/store";
import {Interface} from "@/views/interface/data";
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";

import DesignInterface from './components/designer/Interface.vue';
import RequestEnv from './components/designer/others/env/index.vue';
import RequestHistory from './components/designer/others/history/index.vue';
import Datapool from './components/designer/others/datapool/index.vue';

provide('usedBy', UsedBy.interface)
const useForm = Form.useForm;

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();

const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const tabKey = ref('env')

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('designer-main',
      'design-content', 'design-splitter', 'design-right', 500, 300)
}

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
#designer-main {
  display: flex;
  height: 100%;

  #design-content {
    flex: 1;
    width: 0;
    height: 100%;
  }

  #design-right {
    width: 320px;
    height: 100%;
  }

  .switcher {
    position: fixed;
    right: 8px;
    bottom: 5px;
    cursor: pointer;
  }
}

</style>