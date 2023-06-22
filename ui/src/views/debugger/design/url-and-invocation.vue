<template>
  <div class="url-input-main">
    <div class="url">
      <a-input-group>
        <a-row type="flex" :gutter="0">
          <a-col flex="80px">
            <a-select class="select-env"
                      :options="methods"
                      v-model:value="method"
                      @change="changeMethod">
            </a-select>
          </a-col>

          <a-col flex="3">
            <a-input placeholder="站点地址"
                     v-model:value="debugData.baseUrl">
            </a-input>
          </a-col>

          <a-col flex="3">
            <!-- debug-url-3 -->
            <a-input class="uri" placeholder="请求路径"
                     :id="'testInterfaceUrl0'"
                     v-model:value="debugData.url"
                     v-contextmenu="e => onContextMenuShow(0, e)">
            </a-input>
          </a-col>

          <a-col flex="80px" class="send">
            <a-button type="primary" trigger="click" @click="send">
              <span>发送</span>
            </a-button>
          </a-col>

          <a-col flex="80px" class="save">
            <a-button @click="saveTestInterface" class="dp-bg-light">
              <SaveOutlined/>
              保存
            </a-button>
          </a-col>
        </a-row>
      </a-input-group>

    </div>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>
  </div>
</template>
<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {useStore} from "vuex";
import {notification} from "ant-design-vue";
import {SaveOutlined} from '@ant-design/icons-vue';

import {NotificationKeyCommon} from "@/utils/const";
import {getToken} from "@/utils/localToken";

import {StateType as DebugStateType} from "@/views/component/debug/store";
import {StateType as TestInterfaceStateType} from "@/views/debugger/store";
import {StateType as EndpointStateType} from "@/views/endpoint/store";
import {Methods} from "@/utils/enum";
import {getArrSelectItems} from "@/utils/comm";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {getContextMenuStyle2} from "@/utils/dom";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"

const store = useStore<{ TestInterface: TestInterfaceStateType, Debug: DebugStateType, Endpoint: EndpointStateType }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const serveServers: any = computed(() => store.state.TestInterface.serveServers);

const method = ref('GET')
const methods = getArrSelectItems(Methods)

const getEnvUrl = () => {
  if (!debugData.value || !serveServers.value) return

  serveServers.value?.forEach((item) => {
    if (debugData.value.serverId === item.id) {
      debugData.value.baseUrl = item.url
      return
    }
  })
}

watch((debugData), async (newVal) => {
  console.log('watch debugData', debugData?.value)
  getEnvUrl()
}, { immediate: true, deep: true })

watch((serveServers), async (newVal) => {
  console.log('watch serveServers', serveServers?.value)
  getEnvUrl()
}, { immediate: true, deep: true })

const changeMethod = (item) => {
  console.log('changeMethod', item)
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
const saveTestInterface = async (e) => {
  console.log('saveTestInterface')
  if (validateInfo()) {
    let data = JSON.parse(JSON.stringify(debugData.value))
    data = prepareDataForRequest(data)

    if (validateInfo()) {
      Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null})

      const res = await store.dispatch('TestInterface/saveTestDebugData', data)
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
  }
}

const validateInfo = () => {
  let msg = ''
  if (!debugData.value.url) {
    msg = '请求地址不能为空'
  }

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

const showContextMenu = ref(false)
const paramIndex = ref(-1)
let contextTarget = {} as any
const contextMenuStyle = ref({} as any)
const onContextMenuShow = (idx, e) => {
  console.log('onContextMenuShow', idx, e.target)
  if (!e) return

  contextMenuStyle.value = getContextMenuStyle2(e)
  contextTarget = e.target
  paramIndex.value = idx

  showContextMenu.value = true
}
const onMenuClick = (key) => {
  console.log('onMenuClick', key)

  if (key === 'use-variable') {
    bus.emit(settings.eventVariableSelectionStatus, {src: 'testInterfaceUrl', index: paramIndex.value, data: contextTarget});
  }

  showContextMenu.value = false
}

</script>

<style scoped lang="less">
.url-input-main {
  display: inline-block;
  overflow: hidden;
  width: 100%;

  .url {
    .select-env {
      width: 100%;
    }
  }
}
</style>
