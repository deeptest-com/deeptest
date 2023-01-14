<template>
  <div id="scenario-design-main">
    <div class="toolbar">
      <a-button @click="exec" type="link">执行场景</a-button>
    </div>

    <div id="main" class="main dp-splits-v">
      <div id="design-processor">
        <Processor></Processor>
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
import {Interface} from "@/views/interface/data";
import {resizeWidth} from "@/utils/dom";
import {useRouter} from "vue-router";
import RequestEnv from '@/views/interface/components/designer/others/env/index.vue';
import Processor from './Processor.vue';
import {UsedBy} from "@/utils/enum";

provide('usedBy', UsedBy.scenario)
const useForm = Form.useForm;

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();

const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

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
#scenario-design-main {
  height: 100%;

  .toolbar {
    height: 32px;
    text-align: right;
  }

  #main {
    height: calc(100% - 32px);
    width: 100%;

    #design-processor {
      flex: 1;
      width: 0;
      height: 100%;
    }

    #design-right {
      width: 320px;
      height: 100%;
    }
  }
}

</style>