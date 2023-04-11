<template>
  <div id="debug-index" class="dp-splits-v">
    <div id="debug-content">
      <DesignInterface />
    </div>

    <div id="debug-splitter" class="splitter"></div>

    <div id="debug-right">
      <a-tabs v-model:activeKey="tabKey"
              tabPosition="right"
              :tabBarGutter="0"
              class="right-tab">

        <a-tab-pane key="history">
          <template #tab>
            <a-tooltip placement="left" overlayClassName="dp-tip-small">
              <template #title>历史</template>
              <HistoryOutlined/>
            </a-tooltip>
          </template>

          <RequestHistory v-if="tabKey==='history'"></RequestHistory>
        </a-tab-pane>

        <a-tab-pane key="datapool">
          <template #tab>
            <a-tooltip placement="left" overlayClassName="dp-tip-small">
              <template #title>数据池</template>
              <TableOutlined />
            </a-tooltip>
          </template>

          <Datapool v-if="tabKey==='datapool'"></Datapool>
        </a-tab-pane>

      </a-tabs>
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import { HistoryOutlined, TableOutlined } from '@ant-design/icons-vue';

import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";

import DesignInterface from './interface.vue';

import RequestHistory from './others/history/index.vue';
import Datapool from './others/datapool/index.vue';

provide('usedBy', UsedBy.interface)
const useForm = Form.useForm;

const {t} = useI18n();

const tabKey = ref('history')

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('debug-index',
      'debug-content', 'debug-splitter', 'debug-right', 500, 300)
}

</script>

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

<style lang="less">
#debug-index #debug-right .right-tab {
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