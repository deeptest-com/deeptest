<template>
  <div class="endpoint-debug-index-wrapper">
    <div v-if="endpointDetail?.id && endpointDetail?.interfaces[0]?.id" id="endpoint-debug-index" class="debug-page-container-top dp-splits-v" >
      <div id="debug-top">
        <DebugMethod/>
      </div>
      <div id="debug-bottom">
        <DebugComp :topVal="'10px'"
                   :onSaveDebugData="saveDebugInterface"
                   :onSaveAsCase="saveAsCase"
                   :showMethodSelection="false"/>
      </div>
    </div>
    <div v-else style="margin-top: 48px;">
      <a-empty
          image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
          :image-style="{height: '60px',}">
        <template #description>
      <span>
        您还未定义接口，请先定义接口才能使用调试功能
      </span>
        </template>
        <a-button type="primary" @click="emit('switchToDefineTab')">接口定义</a-button>
      </a-empty>
    </div>

    <SaveAsCasePopup
        v-if="saveAsVisible"
        :visible="saveAsVisible"
        :model="saveAsModel"
        :onFinish="saveAsFinish"
        :onCancel="saveAsCancel"/>
  </div>
</template>

<script setup lang="ts">
import {computed, defineEmits, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";

import {NotificationKeyCommon} from "@/utils/const";
import {UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "@/views/endpoint/store";

import DebugMethod from './method.vue';
import DebugComp from '@/views/component/debug/index.vue';
import SaveAsCasePopup from "../Cases/edit.vue";

const store = useStore<{ Debug: Debug, Endpoint: Endpoint }>();
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);

provide('usedBy', UsedBy.InterfaceDebug)
const useForm = Form.useForm;
const {t} = useI18n();

const emit = defineEmits(['switchToDefineTab']);

const saveDebugInterface = async (data) => {
  console.log('saveDebugInterface', data)

  Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null})

  store.commit("Global/setSpinning",true)
  const res = await store.dispatch('Debug/save', data)
  store.commit("Global/setSpinning",false)

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
};

const saveAsVisible = ref(false)
const saveAsModel = ref({} as any)
const saveAsCase = () => {
  console.log('saveAsCase')
  saveAsVisible.value = true
  saveAsModel.value = {title: ''}
}
const saveAsFinish = async (model) => {
  console.log('saveAsFinish', model, debugData.value.url)

  const data = Object.assign({...model, debugData: debugData.value}, debugInfo.value)

  store.commit("Global/setSpinning",true)
  const res = await store.dispatch('Debug/saveAsCase', data)
  store.commit("Global/setSpinning",false)

  if (res === true) {
    saveAsVisible.value = false

    notification.success({
      key: NotificationKeyCommon,
      message: `另存为用例成功`,
    });
  } else {
    notification.success({
      key: NotificationKeyCommon,
      message: `另存为用例保存失败`,
    });
  }
}
const saveAsCancel = () => {
  console.log('saveAsVisible')
  saveAsVisible.value = false
}

</script>

<style lang="less" scoped>
.endpoint-debug-index-wrapper {
  height: 100%;

  #endpoint-debug-index {
    height: 100%;
    width: 100%;
    min-height: calc(100vh - 96px);
    flex-direction: column;

    #debug-top {
      display: flex;
      margin: 12px 0;
      height: 32px;
      //margin-bottom: 12px;
      align-items: center;
      justify-content: space-between;
      padding-right: 2px;
    }

    #debug-bottom {
      height: calc(100% - 46px);
      //height: calc(100vh - 152px);
    }
  }
}

</style>

