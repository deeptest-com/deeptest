<template>
  <div class="invocation-main">
    <div class="toolbar">
      <div v-if="showMethodSelection" class="select-method">
        <a-select class="select-method"
                  v-model:value="debugData.method"
                  :disabled="usedBy === UsedBy.CaseDebug">
          <template v-for="method in Methods">
            <a-select-option v-if="hasDefinedMethod(method)"
                             :key="method"
                             :value="method">
              {{ method }}
            </a-select-option>
          </template>
        </a-select>
      </div>
      <div id="env-selector">
        <EnvSelector :show="showBaseUrl()" :server-id="serverId" @change="changeServer" :disabled="usedBy === UsedBy.ScenarioDebug" />
      </div>
      <div v-if="showBaseUrl()" class="base-url">
        <a-input placeholder="请输入地址"
                 v-model:value="debugData.baseUrl"
                 :disabled="baseUrlDisabled" />
      </div>

      <div class="url"
           :class="[isPathValid  ? '' :  'dp-field-error' ]">
        <a-tooltip placement="bottom" :visible="!isPathValid"  overlayClassName="dp-tip-small" :title="'请输入合法的路径,以http(s)开头'">
          <a-input placeholder="请输入路径"
                   v-model:value="debugData.url"
                   @change="pathUpdated"
                   :disabled="urlDisabled"
                   :title="urlDisabled ? '请在接口定义中修改' : ''"/>
        </a-tooltip>
      </div>

      <div class="send">
        <a-button type="primary" trigger="click"
                  @click="send"
                  :disabled="!isPathValid">
          <span>发送</span>
        </a-button>
      </div>

      <div class="save">
        <a-button trigger="click" class="dp-bg-light"
                  @click="save"
                  :disabled="!isPathValid">
          <icon-svg class="icon dp-icon-with-text" type="save" />
          保存
        </a-button>
      </div>

      <div v-if="usedBy === UsedBy.InterfaceDebug"
           :disabled="!isPathValid"
           class="save-as-case">
        <a-button trigger="click" @click="saveAsCase" class="dp-bg-light">
          另存为用例
        </a-button>
      </div>

      <div v-if="isShowSync"
           :disabled="!isPathValid"
           class="sync">
        <a-button trigger="click" @click="sync" class="dp-bg-light">
          <UndoOutlined/>
          同步
        </a-button>
      </div>
    </div>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onMounted, onUnmounted, PropType, ref, watch, Teleport} from "vue";
import {notification} from 'ant-design-vue';
import {UndoOutlined} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import IconSvg from "@/components/IconSvg";
import {Methods, ProcessorInterfaceSrc, UsedBy} from "@/utils/enum";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {NotificationKeyCommon} from "@/utils/const"

import {StateType as GlobalStateType} from "@/store/global";
import {StateType as DebugStateType} from "@/views/component/debug/store";
import {StateType as EndpointStateType} from "@/views/endpoint/store";

import {Endpoint} from "@/views/endpoint/data";
import useVariableReplace from "@/hooks/variable-replace";
import {getToken} from "@/utils/localToken";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import EnvSelector from "./config/EnvSelector.vue";
import {handlePathLinkParams} from "@/utils/dom";

const store = useStore<{ Debug: DebugStateType, Endpoint: EndpointStateType, Global: GlobalStateType }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const servers = computed<any[]>(() => store.state.Debug.serves);

const props = defineProps({
  onSave: {
    type: Function as PropType<(data) => void>,
    required: true
  },
  onSaveAsCase: {
    type: Function,
    required: false
  },
  onSync: {
    type: Function as PropType<() => void>,
    required: false
  },

  showMethodSelection: {
    type: Boolean,
    required: false,
    default: true
  },
  baseUrlDisabled: {
    type: Boolean,
    required: false,
    default: true
  },
  urlDisabled: {
    type: Boolean,
    required: false,
    default: false
  },
})
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const {showContextMenu, contextMenuStyle, onContextMenuShow, onMenuClick} = useVariableReplace('endpointInterfaceUrl')

const listServer = async (serveId) => {
  await store.dispatch('Debug/listServes', { serveId })
}

const showBaseUrl = () => {
  const notShow = debugData.value.usedBy === UsedBy.DiagnoseDebug
      || (debugData.value.usedBy === UsedBy.ScenarioDebug &&
                (debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Diagnose ||
                  debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Custom  ||
                  debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Curl
                  ))

  return !notShow
}

const isShowSync = computed(() => {
  const ret = usedBy === UsedBy.ScenarioDebug && (
      debugData.value.processorInterfaceSrc !== ProcessorInterfaceSrc.Custom  &&
      debugData.value.processorInterfaceSrc !== ProcessorInterfaceSrc.Curl)

  return ret
})

const getEnvUrl = () => {
  console.log('getEnvUrl', debugData?.value)

  if (!debugData.value || !servers.value) return

  servers.value?.forEach((item) => {
    if (debugData.value.serverId === item.id && debugData.value.baseUrl) {
      debugData.value.baseUrl = item.url
      return
    }
  })
}

const currServeId = ref(0)
const currServerId = ref(0)
watch(debugData, (newVal) => {
  console.log('watch debugData', debugData.value.serveId, debugData.value.currServerId)

  if (newVal.serveId != currServeId.value) {
    listServer(debugData.value.serveId)
    currServeId.value = debugData.value.serveId
  }

  if (currServerId.value != newVal.serverId) {
    getEnvUrl()
    currServerId.value = newVal.serverId
  }

  if (usedBy === UsedBy.InterfaceDebug || usedBy === UsedBy.CaseDebug) {
    debugData.value.url = debugData?.value.url || endpointDetail.value?.path || ''
  }

}, {immediate: true, deep: true});

const serverId = computed(() => {
  return store.state.Debug.currServe.environmentId || 0
});

function changeServer(id) {
  store.dispatch('Debug/changeServer', { serverId: id, requestEnvVars: false })
}

const send = async (e) => {
  const data = prepareDataForRequest(debugData.value)
  console.log('sendRequest', data)

  if (validateInfo()) {
    store.commit("Global/setSpinning",true)

    const callData = {
      serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
      token: await getToken(),
      data: data
    }
    await store.dispatch('Debug/call', callData).finally(()=>{
      store.commit("Global/setSpinning",false)
    })

    store.commit("Global/setSpinning",false)
  }
}

const save = (e) => {
  let data = JSON.parse(JSON.stringify(debugData.value))
  data = prepareDataForRequest(data)

  if (validateInfo()) {
     props.onSave(data)
  }

  bus.emit(settings.eventConditionSave, {});
}
const saveAsCase = () => {
  // console.log('saveAsCase', debugData.value.url)
  if (validateInfo() && props.onSaveAsCase) {
    props.onSaveAsCase()
  }
}

const sync = (e) => {
  if (validateInfo() && props.onSync) {
    props.onSync()
  }
};

const validateInfo = () => {
  let msg = ''
  if (usedBy !== UsedBy.DiagnoseDebug && !debugData.value.url) {
    msg = '请求地址不能为空'
  }
  // else if (!regxUrl.test(debugData.value.url)) {
  //   msg = '请求地址格式错误'
  // }

  if (msg) {
    notification.warn({
      key: NotificationKeyCommon,
      message: msg,
      placement: 'topRight'
    });

    return false
  }

  return true
};

onUnmounted(() => {
  console.log('onUnmounted')
})

function hasDefinedMethod(method: string) {
  if (usedBy !== UsedBy.CaseDebug) // selection not show for interface_debug
    return true

  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

function pathUpdated(e) {
  const path = e.target.value.trim();

  if (!validatePath()) {
    return
  }

  const ret = handlePathLinkParams(path, debugData.value?.pathParams)
  store.commit('Debug/setPathParams', ret)
}


const isPathValid = computed(() => {return validatePath()})
function validatePath() {
  const regx = /^https?:\/\/.+$/g;
  const isMatch = showBaseUrl() || regx.test(debugData.value?.url)

  return isMatch
}

// const showContextMenu = ref(false)
// const clearMenu = () => {
//   console.log('clearMenu')
//   showContextMenu.value = false
// }
//
// let contextTarget = {} as any
// const contextMenuStyle = ref({} as any)
//
// const onMenuClick = (key) => {
//   console.log('onMenuClick', key)
//
//   if (key === 'use-variable') {
//     bus.emit(settings.eventVariableSelectionStatus, {src: 'url', data: contextTarget});
//   }
//   showContextMenu.value = false
// }
</script>

<style lang="less" scoped>
.invocation-main {
  padding: 0;

  .toolbar {
    display: flex;

    .select-method {
      width: 100px;
    }

    .base-url {
      flex: 1;
    }

    .url {
      flex: 1;
      &.dp-field-error {
        border: 1px solid red !important;
      }

      input {
        &:focus {
          border: 1px solid #d9d9d9 !important;
          outline: none  !important;
          box-shadow: none  !important;
        }
      }
    }

    .send {
      margin-left: 8px;
      width: 66px;
    }

    .save {
      margin-left: 8px;
      width: 80px;
    }

    .save-as-case {
      margin-left: 8px;
      width: 102px;
    }

    .sync {
      margin-left: 8px;
      width: 80px;
    }

  }
}

</style>
