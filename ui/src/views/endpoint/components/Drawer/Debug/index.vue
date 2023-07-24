<template>
  <div id="endpoint-debug-index" class="dp-splits-v" v-if="endpointDetail?.id && endpointDetail?.interfaces[0]?.id">
    <div id="debug-top">
      <DebugMethod />
    </div>
    <div id="debug-bottom">
      <DebugComp :topVal="'10px'"
                  :onSaveDebugData="saveDebugInterface"
                 :showMethodSelection="false" />
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
</template>

<script setup lang="ts">
import {onMounted, provide, ref, computed,defineEmits} from "vue";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";

import {NotificationKeyCommon} from "@/utils/const";
import {UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Endpoint} from "@/views/endpoint/store";

import DebugMethod from './method.vue';
import DebugComp from '@/views/component/debug/index.vue';

const store = useStore<{  Debug: Debug,Endpoint:Endpoint,Global }>();
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);

provide('usedBy', UsedBy.InterfaceDebug)
const useForm = Form.useForm;
const debugData = computed<any>(() => store.state.Debug.debugData);
const {t} = useI18n();

const emit = defineEmits(['switchToDefineTab']);

const saveDebugInterface = async (data) => {
  console.log('saveDebugInterface', data)

  Object.assign(data, {shareVars: null, envVars: null, globalEnvVars: null, globalParamVars: null })
  store.commit("Global/setSpinning",true)
  const res = await store.dispatch('Debug/save', data).finally(()=>{
    store.commit("Global/setSpinning",false)
  })
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
  store.commit("Global/setSpinning",false)
};

</script>

<style lang="less" scoped>
#endpoint-debug-index {
  height: 100%;
  width: 100%;
  display: flex;
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
    flex: 1;
    //height: calc(100% - 46px);
    height: calc(100vh - 152px);
  }
}
</style>

