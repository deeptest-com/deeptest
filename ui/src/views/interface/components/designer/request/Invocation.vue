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
      <a-input v-model:value="interfaceData.url"  v-contextmenu="onContextMenuShow" class="dp-bg-light" />
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

    <div v-if="showContextMenu" :style="contextMenuStyle" class="context-menu">
      <div @click="onMenuClick('replaceVari')" class="item">替换变量</div>
      <div @click="onMenuClick('')" class="item">复制</div>
      <div @click="onMenuClick('')" class="item">关闭</div>
    </div>

    <RequestVariable
        v-if="requestVariableVisible"
        :interfaceId="interfaceData.id"
        :onFinish="requestVariableSelectFinish"
        :onCancel="requestVariableSelectCancel"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, ref, PropType, onMounted, onUnmounted, defineProps} from "vue";
import { notification, message } from 'ant-design-vue';
import { DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {Methods} from "@/utils/enum";
import {Interface} from "@/views/interface/data";
import {prepareDataForRequest} from "@/views/interface/service";
import {NotificationKeyCommon} from "@/utils/const"
import RequestVariable from "@/components/Editor/RequestVariable.vue";

const props = defineProps({
  onSend: {
    type: Function as PropType<(data) => void>,
    required: true
  },
  onSave: {
    type: Function as PropType<(data) => void>,
    required: true
  }
})

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const methods = Methods;

const selectMethod = (val) => {
  console.log('selectMethod', val.key)
  interfaceData.value.method = val.key
};

const sendRequest = (e) => {
  console.log('--- interface data', interfaceData.value)

  let data = JSON.parse(JSON.stringify(interfaceData.value))
  data = prepareDataForRequest(data)
  data.body = data.body.replaceAll('\n', '').replaceAll(' ', '')

  if (validateInfo()) {
    props.onSend(data)
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
  document.addEventListener("click", clearMenu)
})
onUnmounted(() => {
  document.removeEventListener("click", clearMenu)
})

const showContextMenu = ref(false)
let contextTarget = {} as any
const contextMenuStyle = ref({} as any)

const onContextMenuShow = (e, binding) => {
  console.log('onContextMenuShow', e, binding)
  contextMenuStyle.value.left = e.clientX + "px";
  contextMenuStyle.value.top = e.clientY - 60 > 6 ? e.clientY - 60 : 6  + "px";
  contextMenuStyle.value.maxHeight = "200px";

  contextTarget = e.target

  showContextMenu.value = true
}

const onMenuClick = (item) => {
  console.log('onMenuClick', item,
      contextTarget.value.substr(contextTarget.selectionStart, contextTarget.selectionEnd - contextTarget.selectionStart))
  showContextMenu.value = false

  requestVariableVisible.value = true
}

const requestVariableVisible = ref(false)

const requestVariableSelectFinish = (data) => {
  console.log('requestVariableSelectFinish', data)

  data.interfaceId = interfaceData.value.id
  data.projectId = interfaceData.value.projectId
  // store.dispatch('Interface/createExtractorOrUpdateResult', data).then((result) => {
  //   if (result) {
  //     requestVariableVisible.value = false
  //   }
  // })
}
const requestVariableSelectCancel = () => {
  console.log('requestVariableSelectCancel')
  requestVariableVisible.value = false
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

  .context-menu {
    position: fixed;
    padding: 6px 10px;
    border: 1px solid #dedfe1;
    background: #f0f2f5;
    z-index: 99;

    .item {
      margin: 5px 0;
      cursor: pointer;
    }
  }
}

</style>