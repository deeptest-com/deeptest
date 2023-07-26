<template>
  <div class="endpoint-debug-cases-design-main">
    <div class="toolbar">
      <a-button type="link" trigger="click" @click="back">
        <span>返回用例列表</span>
      </a-button>
    </div>

    <div id="endpoint-debug-cases-design-panel">
      <div class="name">
        <EditAndShowField placeholder="请输入名称"
                          :custom-class="'custom-serve show-on-hover'"
                          :value="endpointCase.name"
                          @update="updateName" />
      </div>

      <DebugComp
          :topVal="'-36px'"
          :onSaveDebugData="saveCaseInterface"
          :urlDisabled="true"/>
    </div>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, provide, ref, watch} from 'vue';
import {useStore} from "vuex";
import debounce from "lodash.debounce";
import {UsedBy} from "@/utils/enum";

import DebugComp from '@/views/component/debug/index.vue';

import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as EndpointStateType} from '../../../store';
import {StateType as DiagnoseInterfaceStateType} from "@/views/diagnose/store";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import EditAndShowField from '@/components/EditAndShow/index.vue';

provide('usedBy', UsedBy.CaseDebug)
const usedBy = UsedBy.CaseDebug

const store = useStore<{ Debug: Debug, Endpoint: EndpointStateType, DiagnoseInterface: DiagnoseInterfaceStateType }>();
const endpointCase = computed<any>(() => store.state.Endpoint.caseDetail);
const debugData = computed<any>(() => store.state.Debug.debugData);

const props = defineProps({
  onBack: {
    type: Function,
    required: true,
  },
})

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

const saveCaseInterface = async (e) => {
  console.log('saveCaseInterface')

  let data = JSON.parse(JSON.stringify(debugData.value))
  data = prepareDataForRequest(data)

  Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null})

  const res = await store.dispatch('Endpoint/saveCaseDebugData', data)
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

const updateName = (val) => {
  endpointCase.value.name = val
  console.log('updateName', val, endpointCase.value)
  store.dispatch('Endpoint/updateCaseName', endpointCase.value)
}

const back = () => {
  console.log('back')
  props.onBack()
}

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
  padding: 0px 0px 16px 16px;
  position: relative;

  .toolbar {
    position: absolute;
    top: -36px;
    right: 110px;
    height: 50px;
    width: 120px;
  }

  #endpoint-debug-cases-design-panel {
    .name {
      width: 260px;
      line-height: 36px;
      font-weight: bold;
    }
  }
}

</style>
