<template>
  <div class="invocation-main">
    <div class="methods">
      <a-dropdown trigger="click">
        <template #overlay>
          <a-menu @click="selectMethod">
            <a-menu-item v-for="(item) in methods" :key="item">{{ item }}</a-menu-item>
          </a-menu>
        </template>
        <a-button class="dp-bg-light">
          <span class="curr-method">{{ interfaceData.method }}</span>
          <DownOutlined />
        </a-button>
      </a-dropdown>
    </div>
    <div class="url">
      <a-input v-model:value="interfaceData.url" v-contextmenu="onContextMenuShow" class="dp-bg-light" />
    </div>
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
    <div class="save">
      <a-dropdown-button trigger="click" @click="save" class="dp-bg-light">
        <SaveOutlined />
        保存
        <template #overlay>
          <a-menu>
            <a-menu-item @click.prevent="none" key="copyLink" class="edit-name">
              <div class="dp-edit-interface-name">
                <div class="left">
                  <a-input @click.stop v-model:value="interfaceData.name" />
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
import {computed, ref, PropType, onMounted, onUnmounted, defineProps, inject} from "vue";
import { notification, message } from 'ant-design-vue';
import { DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined, EditOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {Methods, UsedBy} from "@/utils/enum";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {Interface} from "@/views/interface/data";
import {getContextMenuStyle, prepareDataForRequest} from "@/views/interface/service";
import {NotificationKeyCommon} from "@/utils/const"
import ContextMenu from "@/components/Editor/ContextMenu.vue"
import {StateType as ScenarioStateType} from "@/views/scenario/store";

const props = defineProps({
  onSend: {
    type: Function as PropType<() => void>,
    required: true
  },
  onSave: {
    type: Function as PropType<(data) => void>,
    required: true
  }
})
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface: StateType, Scenario: ScenarioStateType }>();
const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.interface ? store.state.Interface.interfaceData : store.state.Scenario.interfaceData);

const methods = Methods;

const selectMethod = (val) => {
  console.log('selectMethod', val.key)
  interfaceData.value.method = val.key
};

const sendRequest = (e) => {
  console.log('--- interface data', interfaceData.value)

  if (validateInfo()) {
    props.onSend()
  }
};

const save = (e) => {
  let data = JSON.parse(JSON.stringify(interfaceData.value))
  console.log('save', data)
  data = prepareDataForRequest(data)
  console.log('save', data)

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
  if (!interfaceData.value.url) {
    msg = '请求地址不能为空'
  }
  // else if (!regxUrl.test(interfaceData.value.url)) {
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

const onContextMenuShow = (e) => {
  console.log('getContextMenuStyle', e.target)

  contextMenuStyle.value = getContextMenuStyle(e)
  contextTarget = e.target

  showContextMenu.value = true
}

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
  .methods {
    width: 116px;
    .curr-method {
      width: 65px;
    }
  }
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