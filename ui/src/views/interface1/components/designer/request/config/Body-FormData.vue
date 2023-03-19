<template>
  <div class="formdata-main">
    <div class="dp-param-grid">
      <div class="params">
        <a-row v-for="(item, idx) in interfaceData.bodyFormData" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-input v-model:value="item.name" @change="onFormDataChange(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col width="72px">
            <a-select
                v-model:value="item.type"
                @change="onFormDataChange(idx)"
                :bordered="false"
            >
              <a-select-option value="text">Text</a-select-option>
              <a-select-option value="file">File</a-select-option>
            </a-select>
          </a-col>

          <a-col flex="2" class="flow">
            <a-input v-if="item.type!=='file'" v-model:value="item.value" class="dp-bg-input-transparent" />
            <span v-if="item.type==='file'" class="filename">{{getFileName(item.value)}}</span>

            <a-button v-if="item.type==='file'" @click="selectFile(idx)">
              <UploadOutlined />
            </a-button>
          </a-col>

          <a-col flex="80px" class="dp-right dp-icon-btn-container">
            <a-tooltip v-if="!item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>禁用</template>
              <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
            </a-tooltip>

            <a-tooltip v-if="item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>启用</template>
              <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
            </a-tooltip>

            <a-tooltip @click="remove(idx)" overlayClassName="dp-tip-small">
              <template #title>移除</template>
              <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>

            <a-tooltip @click="insert(idx)" overlayClassName="dp-tip-small">
              <template #title>插入</template>
              <PlusOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>
          </a-col>
        </a-row>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">

import {computed, ComputedRef, defineComponent, inject, onMounted, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {
  CheckCircleOutlined, CloseCircleOutlined,  UploadOutlined, DeleteOutlined, PlusOutlined
} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface1/store";
import {Param, Interface, BodyFormDataItem} from "@/views/interface1/data";
import {notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import settings from "@/config/settings";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface1: StateType, Scenario: ScenarioStateType }>();
const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.interface ? store.state.Interface1.interfaceData : store.state.Scenario.interfaceData);

let uploadRef = ref()
const isElectron = ref(!!window.require)

const onFormDataChange = (idx) => {
  console.log('onFormDataChange', idx)
  if (interfaceData.value.bodyFormData.length <= idx + 1
      && (interfaceData.value.bodyFormData[idx].name !== '' || interfaceData.value.bodyFormData[idx].value !== '')) {
    interfaceData.value.bodyFormData.push({type: 'text'} as BodyFormDataItem)
  }
}

const add = () => {
  console.log('add')
  interfaceData.value.bodyFormData.push({type: 'text'} as BodyFormDataItem)
}
const removeAll = () => {
  console.log('removeAll', interfaceData.value.bodyFormData)
  interfaceData.value.bodyFormData = [{type: 'text'} as BodyFormDataItem]
}

const disable = (idx) => {
  console.log('enable', idx)
  interfaceData.value.bodyFormData[idx].disabled = !interfaceData.value.bodyFormData[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  interfaceData.value.bodyFormData.splice(idx, 1)
  add()
}
const insert = (idx) => {
  console.log('insert')
  interfaceData.value.bodyFormData.splice(idx + 1, 0, {type: 'text'} as BodyFormDataItem)
}

const selectedItemIndex = ref(0)

let ipcRenderer = undefined as any
if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
    console.log('from electron: ', data)
    interfaceData.value.bodyFormData[selectedItemIndex.value].value = data.filepath
  })
}

const selectFile = (index) => {
  console.log('selectFile', index)
  selectedItemIndex.value = index

  if (isElectron.value) {
    const data = {act: 'chooseFile'} as any
    ipcRenderer.send(settings.electronMsg, data)
  } else {
    notification.warn({
      key: NotificationKeyCommon,
      message: `请使用客户端上传文件`,
    });
  }
}

const getFileName = (path) => {
  if (!path) {
    return ''
  }
  return path.replace(/^.*[\\\\/]/, '')
}

onMounted(() => {
  console.log('onMounted', uploadRef.value)
})

</script>

<style lang="less" >
.formdata-main {
  height: 100%;
  overflow-y: auto;

  .flow {
    line-height: 32px;
    input {
      width: calc(100% - 46px)
    }
    .filename {
      padding: 0 10px;
    }
    .ant-btn {
      position: absolute;
      right: 0;
      z-index: 99;

      background: transparent;
      color: rgba(0, 0, 0, 0.65);
      border-color: #d9d9d9;
      &:hover, &:active {
        background: transparent;
        color: rgba(0, 0, 0, 0.65);
        border-color: #d9d9d9;
      }
    }
  }
}

</style>