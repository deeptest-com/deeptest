<template>
  <div class="invocation-main">
    <div class="toolbar">
      <div v-if="showMethodSelection" class="select-method">
        <a-select class="select-method"
                  :options="methods"
                  v-model:value="debugData.method">
        </a-select>
      </div>

      <div class="base-url">
          <a-input placeholder="请输入地址"
                   :disabled="baseUrlDisabled"
                   v-model:value="debugData.baseUrl" />
      </div>

      <div class="url">
        <a-input placeholder="请输入路径"
                 v-model:value="debugData.url" />
      </div>

      <div class="send">
        <a-button type="primary" trigger="click" @click="send">
          <span>发送</span>
        </a-button>
      </div>

      <div class="save">
        <a-button trigger="click" @click="save" class="dp-bg-light">
          <SaveOutlined/>
          保存
        </a-button>
      </div>

      <div v-if="usedBy === UsedBy.ScenarioDebug" class="sync">
        <a-button trigger="click" @click="sync" class="dp-bg-light">
          <UndoOutlined />
          同步
        </a-button>
      </div>
    </div>

    <div class="select-env" :style="{top: topVal}">
      <a-select :value="serverId || null" @change="changeServer"
                placeholder="请选择环境" class="select-env">
        <a-select-option v-for="(option, key) in servers" :key="key" :value="option.id">
          {{ option.description }}
        </a-select-option>
      </a-select>
    </div>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onMounted, onUnmounted, PropType, ref, watch} from "vue";
import {notification} from 'ant-design-vue';
import {SaveOutlined, UndoOutlined} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Methods, UsedBy} from "@/utils/enum";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {NotificationKeyCommon} from "@/utils/const"

import {StateType as Debug} from "@/views/component/debug/store";
import {Endpoint} from "@/views/endpoint/data";
import useVariableReplace from "@/hooks/variable-replace";
import {getToken} from "@/utils/localToken";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"
import {serverList} from "@/views/project-settings/service";
import {getArrSelectItems} from "@/utils/comm";

const store = useStore<{ Debug: Debug, Endpoint }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

const props = defineProps({
  topVal: {
    type: String,
    required: true
  },
  onSave: {
    type: Function as PropType<(data) => void>,
    required: true
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
})
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const { showContextMenu, contextMenuStyle, onContextMenuShow, onMenuClick } = useVariableReplace('endpointInterfaceUrl')
const methods = getArrSelectItems(Methods)

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

  debugData.value.url = debugData?.value.url || endpointDetail.value?.path || ''

}, {immediate: true, deep: true});

const serverId = computed(() => {
  return debugData?.value?.serverId || endpointDetail?.value?.serverId || servers.value[0]?.id || 0
});

function changeServer(id) {
  store.dispatch('Debug/changeServer', id)
}

const send = async (e) => {
  console.log('sendRequest', debugData.value)

  if (validateInfo()) {
    const callData = {
      serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
      token: await getToken(),

      data: debugData.value
    }
    await store.dispatch('Debug/call', callData)
  }
}

const save = (e) => {
  let data = JSON.parse(JSON.stringify(debugData.value))
  data = prepareDataForRequest(data)

  if (validateInfo()) {
    props.onSave(data)
  }
};

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


onMounted(() => {
  console.log('onMounted')
})
onUnmounted(() => {
  console.log('onUnmounted')
})

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
.invocation-main {
  position: relative;
  .select-env {
    position: absolute;
    right: 0px;
    width: 120px;
    height: 36px;
    z-index: 9999;
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
      .ant-input[disabled] {
        color: rgba(0, 0, 0, 0.5);
      }
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

    .sync {
      margin-left: 8px;
      width: 80px;
    }

  }
}

</style>
