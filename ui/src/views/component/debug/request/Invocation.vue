<template>
  <div class="invocation-main">
    <div class="toolbar">
      <div v-if="showMethodSelection" class="select-method">
        <a-select class="select-method" v-model:value="debugData.method">
          <template v-for="method in Methods">
            <a-select-option v-if="hasDefinedMethod(method)"
                             :key="method"
                             :value="method">
              {{ method }}
            </a-select-option>
          </template>
        </a-select>
      </div>

      <div v-if="showBaseUrl()" class="base-url">
        <a-input placeholder="请输入地址"
                 v-model:value="debugData.baseUrl"
                 :disabled="baseUrlDisabled" />
      </div>

      <div class="url">
        <a-input placeholder="请输入路径"
                 v-model:value="debugData.url"
                 :disabled="urlDisabled"
                 :title="urlDisabled ? '请在接口定义中修改' : ''" />
      </div>

      <div class="send">
        <a-button type="primary" trigger="click" @click="send">
          <span>发送</span>
        </a-button>
      </div>

      <div class="save">
        <a-button trigger="click" @click="save" class="dp-bg-light">
          <icon-svg class="icon dp-icon-with-text" type="save" />
          保存
        </a-button>
      </div>

      <div v-if="usedBy === UsedBy.InterfaceDebug" class="save-as-case">
        <a-button trigger="click" @click="saveAsCase" class="dp-bg-light">
          另存为用例
        </a-button>
      </div>

      <div v-if="usedBy === UsedBy.ScenarioDebug" class="sync">
        <a-button trigger="click" @click="sync" class="dp-bg-light">
          <UndoOutlined/>
          同步
        </a-button>
      </div>
    </div>

    <!-- 选择环境 -->
    <Teleport to="body">
      <div v-if="showBaseUrl()" class="select-env-fixed" :style="{top: selectEnvTopPosition}">
        <a-select :value="serverId || null" @change="changeServer"
                  placeholder="请选择环境">
          <a-select-option v-for="(option, key) in servers" :key="key" :value="option.id">
            {{ option.description }}
          </a-select-option>
        </a-select>
      </div>
    </Teleport>

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

import {StateType as Debug} from "@/views/component/debug/store";
import {Endpoint} from "@/views/endpoint/data";
import useVariableReplace from "@/hooks/variable-replace";
import {getToken} from "@/utils/localToken";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"
import {serverList} from "@/views/project-settings/service";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const store = useStore<{ Debug: Debug, Endpoint,Global }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

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

const servers = ref([] as any[])
const listServer = async (serveId) => {
  servers.value = []
  if (!serveId) {
    return
  }

  const res = await serverList({
    serveId: serveId
  });
  if (res.code === 0) {
    servers.value = res.data
  }
  console.log('servers', servers)
}

const showBaseUrl = () => {
  console.log('showBaseUrl')

  const notShow = debugData.value.usedBy === UsedBy.DiagnoseDebug
      || (debugData.value.usedBy === UsedBy.ScenarioDebug &&
                (debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Diagnose ||
                  debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Custom  ||
                  debugData.value.processorInterfaceSrc === ProcessorInterfaceSrc.Curl
                  ))

  return !notShow
}

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

  if (usedBy !== UsedBy.DiagnoseDebug) {
    debugData.value.url = debugData?.value.url || endpointDetail.value?.path || ''
  }

}, {immediate: true, deep: true});

const serverId = computed(() => {
  return debugData?.value?.serverId || endpointDetail?.value?.serverId || servers.value[0]?.id || 0
});

function changeServer(id) {
  store.dispatch('Debug/changeServer', id)
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

const selectEnvTopPosition = ref('0px')
onMounted(() => {
  console.log('onMounted')
  selectEnvTopPosition.value = getSelectEnvTopPosition()
})
onUnmounted(() => {
  console.log('onUnmounted')
})

function hasDefinedMethod(method: string) {
  if (usedBy !== UsedBy.CaseDebug)
    return true

  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

const getSelectEnvTopPosition = () => {
  const elems = document.getElementsByClassName('invocation-main')
  if (elems.length === 0) return '0px'

  const rect = elems[0].getBoundingClientRect()
  if (!rect) return '0px'

  const top = rect?.top

  let val = 0
  if (usedBy === UsedBy.ScenarioDebug) {
    val = 50
  } else if (usedBy === UsedBy.DiagnoseDebug) {
    val = 40
  } else if (usedBy === UsedBy.InterfaceDebug) {
    val = 43
  } else if (usedBy === UsedBy.CaseDebug) {
    val = 34
  }

  return top - val + 'px'
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

<style lang="less">
.select-env-fixed { // related to body
  position: fixed;
  z-index: 99999;
  right: 22px;
  width: 120px;

  .ant-select {
    width: 100%;
  }
}
</style>

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
