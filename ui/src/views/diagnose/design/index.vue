<template>
  <a-spin :spinning="spinning">
  <div class="diagnose-interface-design-main">
      <div id="diagnose-interface-debug-panel">
        <a-tabs
          class="dp-tabs-full-height"
          type="editable-card"
          :hideAdd="true"
          :closable="true"
          v-if="interfaceTabs?.length"
          :activeKey="activeTabKey"
          @edit="onTabEdit"
          @change="changeTab">
          <a-tab-pane
            v-for="tab in interfaceTabs"
            :title="tab.title"
            :key="''+tab.id"
            class="dp-relative">
            <template #tab>
              <span :title="tab.title">{{ getTitle(tab.title) }}</span>
            </template>
            <template v-if="debugData?.method" >
              <DebugComp :onSaveDebugData="saveDiagnoseInterface"
                         :baseUrlDisabled="false" />
            </template>
          </a-tab-pane>
        </a-tabs>
        <div  v-else style="margin-top: 36px;">
          <a-empty  :description="'请先在左侧目录上选择需要调试的接口'"/>
        </div>
      </div>

      <div class="selection">
       <!-- <EnvSelection /> -->
      </div>
  </div>
  </a-spin>
</template>

<script setup lang="ts">
import {computed, provide, ref, watch} from 'vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {UsedBy} from "@/utils/enum";

import DebugComp from '@/views/component/debug/index.vue';

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as DiagnoseInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";
import {notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import {prepareDataForRequest} from "@/views/component/debug/service";
import openModal from "@/components/OpenModal/modal";
import {StateType as Debug} from "@/views/component/debug/store";
import ConfirmSave from "@/components/ConfirmSave/index.vue";
import {confirmToSave} from "@/utils/confirm";
import {notifySuccess} from "@/utils/notify";

provide('usedBy', UsedBy.DiagnoseDebug)

const store = useStore<{ Debug: Debug, DiagnoseInterface: DiagnoseInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType,Global }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const debugData = computed<any>(() => store.state.Debug.debugData);

const interfaceId = computed<any>(() => store.state.DiagnoseInterface.interfaceId);
const interfaceData = computed<any>(() => store.state.DiagnoseInterface.interfaceData);
const interfaceTabs = computed<any>(() => store.state.DiagnoseInterface.interfaceTabs);
const activeTabKey = ref('0')
const spinning = computed(()=> store.state.Global.spinning )

function changeTab(key) {
  console.log('changeTab', key)
  activeTabKey.value = key

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

// 切换项目时，需要判断正在调试的接口是否该项目下的接口，不是则需要清空 Tab list
watch(() => { return currProject.value.id },(newVal) => {
  if(newVal){
    store.commit('DiagnoseInterface/setInterfaceTabs',[])
  }
},{immediate:true})

const saveDiagnoseInterface = async (e) => {
  store.commit("Global/setSpinning",true)
  console.log('saveDiagnoseInterface')

    let data = JSON.parse(JSON.stringify(debugData.value))
    data = prepareDataForRequest(data)

    Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null})

    const res = await store.dispatch('DiagnoseInterface/saveDiagnoseDebugData', data).finally(()=> store.commit("Global/setSpinning",false))

  if (res === true) {
    notifySuccess(`保存成功`);
  } else {
    notifySuccess(`保存失败`);
  }
  store.commit("Global/setSpinning",false)
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

</script>

<style scoped lang="less">
.diagnose-interface-design-main {
  height: 100%;
  padding: 16px 0px 0 16px;

  #diagnose-interface-debug-panel {
    height: 100%;
  }
}

</style>
