<template>
  <div class="invocation-main">
    <div v-if="showDebugDataUrl" class="url"></div>

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
        同步
      </a-button>
    </div>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onMounted, onUnmounted, PropType, ref} from "vue";
import {notification} from 'ant-design-vue';
import {SaveOutlined} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Methods, UsedBy} from "@/utils/enum";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {prepareDataForRequest} from "@/views/component/debug/service";
import {NotificationKeyCommon} from "@/utils/const"
import ContextMenu from "@/components/Editor/ContextMenu.vue"
import {StateType as Debug} from "@/views/component/debug/store";
import {addSepIfNeeded} from "@/utils/url";

const store = useStore<{ Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const props = defineProps({
  onSend: {
    type: Function as PropType<() => void>,
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
  showDebugDataUrl: {
    type: Boolean,
    required: false,
    default: true
  }
})
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const send = (e) => {
  console.log('sendRequest', debugData.value)

  if (validateInfo()) {
    props.onSend()
  }
};

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
      placement: 'topRight'
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
    width: 90px;
  }

  .save {
    width: 90px;
  }

  .sync {
    width: 90px;
  }

  .save-scenario {
    width: 90px;
  }
}

</style>
