<template>
  <div class="test-interface-design-main">
    <template v-if="debugData?.method">
      <div class="tabs">
        <a-tabs v-model:activeKey="activeTabKey" @edit="onTabEdit" @change="changeTab" type="editable-card">
          <a-tab-pane v-for="tab in interfaceTabs" :key="''+tab.id" :tab="tab.title">

            <UrlAndInvocation />
            <DebugComp />

          </a-tab-pane>
        </a-tabs>
      </div>

      <div class="selection">
        <EnvSelection />
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import {computed, provide, ref, watch} from 'vue';
import {useStore} from "vuex";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as TestInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";

import debounce from "lodash.debounce";
import {UsedBy} from "@/utils/enum";

import EnvSelection from './env-selection.vue'
import UrlAndInvocation from './url-and-invocation.vue'

import DebugComp from '@/views/component/debug/index.vue';

provide('usedBy', UsedBy.TestDebug)

const store = useStore<{ Debug: Debug, TestInterface: TestInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const debugData = computed<any>(() => store.state.Debug.debugData);

const interfaceId = computed<any>(() => store.state.TestInterface.interfaceId);
const interfaceData = computed<any>(() => store.state.TestInterface.interfaceData);
const interfaceTabs = computed<any>(() => store.state.TestInterface.interfaceTabs);

const activeTabKey = ref('0')

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
}
</style>

<style scoped lang="less">
.test-interface-design-main {
  padding: 16px;

  #debug-form {
    flex: 1;
    padding: 5px 0;

    flex-direction: column;
    position: relative;
    height: 100%;
    max-height: 800px;

    #top-panel {
      height: 50%;
      min-height: 200px;
      width: 100%;
      padding: 0;
    }

    #bottom-panel {
      height: 360px;
      width: 100%;
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
}

</style>
