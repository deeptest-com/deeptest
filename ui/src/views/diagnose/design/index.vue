<template>
  <div class="diagnose-interface-design-main">
      <div id="diagnose-interface-debug-panel">
        <div class="tabs">
          <a-tabs type="editable-card" :hideAdd="true" v-model:activeKey="activeTabKey"
                  @edit="onTabEdit"
                  @change="changeTab">

            <a-tab-pane v-for="tab in interfaceTabs" :key="''+tab.id" :tab="getTitle(tab.title)">
              <template v-if="debugData?.method" >
                <DebugComp :onSaveDebugData="saveDiagnoseInterface"
                           :base-url-disabled="false" />
              </template>
            </a-tab-pane>

          </a-tabs>
        </div>
      </div>

      <div class="selection">
       <!-- <EnvSelection /> -->
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
import Invocation from '@/views/component/debug/request/Invocation.vue'

import DebugComp from '@/views/component/debug/index.vue';

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as DiagnoseInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";

provide('usedBy', UsedBy.DiagnoseDebug)

const store = useStore<{ Debug: Debug, DiagnoseInterface: DiagnoseInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const debugData = computed<any>(() => store.state.Debug.debugData);

const interfaceId = computed<any>(() => store.state.DiagnoseInterface.interfaceId);
const interfaceData = computed<any>(() => store.state.DiagnoseInterface.interfaceData);
const interfaceTabs = computed<any>(() => store.state.DiagnoseInterface.interfaceTabs);

const activeTabKey = ref('0')
const rightTabKey = ref('')

const changeTab = (key) => {
  console.log('changeTab', key)

  const found = interfaceTabs.value.find(function (item, index, arr) {
    return item.id === +key
  })

  store.dispatch('DiagnoseInterface/openInterfaceTab', found);
}

const usedBy = UsedBy.DiagnoseDebug
const loadDebugData = debounce(async () => {
  console.log('loadDebugData')

  store.dispatch('Debug/loadDataAndInvocations', {
    diagnoseInterfaceId: interfaceData.value.id,
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

const saveDiagnoseInterface = async (e) => {
  console.log('saveDiagnoseInterface')

    let data = JSON.parse(JSON.stringify(debugData.value))
    data = prepareDataForRequest(data)

    Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null})

    const res = await store.dispatch('DiagnoseInterface/saveDiagnoseDebugData', data)
    if (res === true) {
      notification.success({
        key: NotificationKeyCommon,
        message: `保存成功`,
      });
    } else {
      notification.success({
        key: NotificationKeyCommon,
        message: `保存失败`,
      });
    }
}

const onTabEdit = (key, action) => {
  console.log('onTabEdit', key, action)
  if (action === 'remove') {
    store.dispatch('DiagnoseInterface/removeInterfaceTab', +key);
  }
};

const getTitle = (title) => {
  const len = title.length
  if (len <= 12) return title

  return title.substr(0, 16) + '...' + title.substr(len-6, len);
};

const closeRightTab = () => {
  rightTabKey.value = ''
}

</script>

<style scoped lang="less">
.diagnose-interface-design-main {
  padding: 16px 0px 16px 16px;

  #diagnose-interface-debug-panel {
  }
}

</style>
