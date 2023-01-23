<template>
  <div id="scenario-design-main">
    <div class="toolbar">
      <a-button @click="exec" type="link" class="dp-btn-group">执行</a-button>
      <a-button href="#/scenario/index" type="link" class="dp-btn-group">返回</a-button>
    </div>

    <div id="scenario-design-content" class="dp-splits-v">
      <div id="scenario-design-content-left">
        <Processor></Processor>
      </div>

      <div id="scenario-design-content-splitter" class="splitter"></div>

      <div id="scenario-design-content-right">
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

        </a-tabs>
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, provide, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import { HistoryOutlined, EnvironmentOutlined, LeftCircleOutlined, RightCircleOutlined } from '@ant-design/icons-vue';
import {useStore} from "vuex";
const router = useRouter();

import {StateType} from "@/views/interface/store";
import {resizeWidth} from "@/utils/dom";
import {useRouter} from "vue-router";
import {UsedBy} from "@/utils/enum";

import RequestEnv from '@/views/interface/components/designer/others/env/index.vue';
import RequestHistory from '@/views/interface/components/designer/others/history/index.vue';
import Processor from './components/Processor.vue';

provide('usedBy', UsedBy.scenario)
const useForm = Form.useForm;

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();

const id = ref(+router.currentRoute.value.params.id)
const tabKey = ref('env')

const exec = () => {
  router.push(`/scenario/exec/${id.value}`)
}

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('scenario-design-content',
      'scenario-design-content-left', 'scenario-design-content-splitter', 'scenario-design-content-right', 500, 300)
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
#scenario-design-main {
  height: 100%;

  .toolbar {
    height: 32px;
    padding-left: 12px;
    padding-right: 12px;
    text-align: right;
  }

  #scenario-design-content {
    height: calc(100% - 32px);
    width: 100%;

    #scenario-design-content-left {
      flex: 1;
      width: 0;
      height: 100%;
    }

    #scenario-design-content-right {
      width: 320px;
      height: 100%;
    }
  }
}

</style>