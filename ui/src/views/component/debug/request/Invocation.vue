<template>
  <div class="invocation-main">
    <!-- 最新ui交互将调整这里的url显示。这里先用条件判断，保证接口管理/调试栏显示的是新ui，而scenario/design中仍保留 -->
    <div v-if="showDebugDataUrl" class="url">{{url}} - {{debugData.method}}</div>

    <div class="send">
      <a-dropdown-button type="primary" trigger="click" @click="sendRequest">
        <span>发送</span>

        <template #overlay>
          <a-menu>
            <a-menu-item @click="clearAll" key="clearAll">
              <UndoOutlined />
              全部清除
            </a-menu-item>
          </a-menu>
        </template>
        <template #icon><DownOutlined /></template>
      </a-dropdown-button>
    </div>

    <div v-if="usedBy===UsedBy.InterfaceDebug" class="save">
      <a-dropdown-button trigger="click" @click="save" class="dp-bg-light">
        <SaveOutlined />
        保存
        <template #overlay>
          <a-menu>
            <a-menu-item @click.prevent="none" key="copyLink" class="edit-name">
              <div class="dp-edit-interface-name">
                <div class="left">
                  <a-input @click.stop v-model:value="debugData.name" />
                </div>
                <div class="right">
                  <CheckOutlined @click.stop="saveName" class="save-button" />
                </div>
              </div>
            </a-menu-item>

            <a-menu-item @click="copyLink" key="copyLink">
              <LinkOutlined />
              复制链接
            </a-menu-item>

            <a-menu-item @click="saveAs" key="saveAs">
              <LinkOutlined />
              另存为
            </a-menu-item>
          </a-menu>
        </template>
        <template #icon><DownOutlined /></template>
      </a-dropdown-button>
    </div>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>

  </div>
</template>

<script setup lang="ts">
import {computed, ref, PropType, onMounted, onUnmounted, defineProps, inject, watch} from "vue";
import { notification, message } from 'ant-design-vue';
import { DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined, EditOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Methods, UsedBy} from "@/utils/enum";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {getContextMenuStyle, prepareDataForRequest} from "@/views/component/debug/service";
import {NotificationKeyCommon} from "@/utils/const"
import ContextMenu from "@/components/Editor/ContextMenu.vue"

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
import {addSepIfNeeded} from "@/utils/url";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const methods = Methods;

const props = defineProps({
  onSend: {
    type: Function as PropType<() => void>,
    required: true
  },
  onSave: {
    type: Function as PropType<(data) => void>,
    required: true
  },
  showDebugDataUrl: {
    type: Boolean,
    required: false,
    default: true
  }
})
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const url = computed(() => addSepIfNeeded(debugData.value.baseUrl) + debugData.value.url)

const selectMethod = (val) => {
  console.log('selectMethod', val.key)
  debugData.value.method = val.key
};

const sendRequest = (e) => {
  console.log('sendRequest', debugData.value)

  if (validateInfo()) {
    props.onSend()
  }
};

const save = (e) => {
  let data = JSON.parse(JSON.stringify(debugData.value))
  data = prepareDataForRequest(data)
  // console.log('-------', data.endpointInterfaceId)

  if (validateInfo()) {
    props.onSave(data)
  }
};

const saveName = (e) => {
  console.log('saveName', e)
  e.preventDefault();
};
const saveAs = (e) => {
  console.log('saveAs', e)
};

const copyLink = (e) => {
  console.log('copyLink', e)
};
const clearAll = (e) => {
  console.log('clearAll', e)
};
const none = (e) => {
  console.log('none', e)
  e.preventDefault()
};

const validateInfo = () => {
  let msg = ''
  if (!debugData.value.url) {
    msg = '请求地址不能为空'
  }
  // else if (!regxUrl.test(debugData.value.url)) {
  //   msg = '请求地址格式错误'
  // }

  if (msg) {
    notification.warn({
      key: NotificationKeyCommon,
      message: msg,
      placement: 'bottomLeft'
    });

    return false
  }

  return true
};

const clearMenu = () => {
  console.log('clearMenu')
  showContextMenu.value = false
}
onMounted(() => {
  console.log('onMounted')
})
onUnmounted(() => {
  console.log('onUnmounted')
})

const showContextMenu = ref(false)
let contextTarget = {} as any
const contextMenuStyle = ref({} as any)

const onMenuClick = (key) => {
  console.log('onMenuClick', key)

  if (key === 'use-variable') {
    bus.emit(settings.eventVariableSelectionStatus, {src: 'url', data: contextTarget});
  }
  showContextMenu.value = false
}

</script>

<style lang="less">
.dp-edit-interface-name {
  display: flex;
  .left {
    flex: 1;
  }
  .right {
    width: 30px;
    padding-left: 10px;
    .save-button {
      vertical-align: -5px
    }
  }
}
</style>

<style lang="less" scoped>
.invocation-main {
  display: flex;
  padding: 0;
  .url {
    flex: 1;
  }
  .send {
    width: 96px;
  }
  .save {
    width: 110px;
  }
}

</style>