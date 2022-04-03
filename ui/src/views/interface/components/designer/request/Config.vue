<template>
  <div class="config-main">
    <a-tabs v-model:activeKey="activeKey" :animated="false">
      <a-tab-pane key="1" tab="查询参数">
        <RequestParameters v-if="activeKey === '1'"></RequestParameters>
      </a-tab-pane>

      <a-tab-pane key="2" tab="请求体">
        <RequestBody v-if="activeKey === '2'"></RequestBody>
      </a-tab-pane>

      <a-tab-pane key="3" tab="请求头">
        <RequestHeaders v-if="activeKey === '3'"></RequestHeaders>
      </a-tab-pane>

      <a-tab-pane key="4" tab="授权">
        <Authorization v-if="activeKey === '4'"></Authorization>
      </a-tab-pane>

<!--      <a-tab-pane key="5" tab="预处理">
        <PreRequestScript v-if="activeKey === '5'"></PreRequestScript>
      </a-tab-pane>

      <a-tab-pane key="6" tab="验证">
        <ValidationScript v-if="activeKey === '6'"></ValidationScript>
      </a-tab-pane>-->
    </a-tabs>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {Methods} from "@/views/interface/consts";

import RequestParameters from "./Config/Parameters.vue";
import RequestBody from "./Config/Body.vue";
import RequestHeaders from "./Config/Headers.vue";
import Authorization from "./Config/Authorization.vue";
import PreRequestScript from "./Config/PreRequestScript.vue";
import ValidationScript from "./Config/ValidationScript.vue";
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestConfig',
  props: {
  },
  components: {
    RequestParameters, RequestBody, RequestHeaders, Authorization,
    // PreRequestScript, ValidationScript,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const activeKey = ref('1');
    const methods = Methods;

    const selectMethod = (e) => {
      console.log('selectMethod', e)
    };
    const sendRequest = (e) => {
      console.log('sendRequest', e)
    };
    const clearAll = (e) => {
      console.log('clearAll', e)
    };
    const saveName = (e) => {
      console.log('saveName', e)
      e.preventDefault();
    };
    const copyLink = (e) => {
      console.log('copyLink', e)
    };
    const saveAs = (e) => {
      console.log('saveAs', e)
    };
    const none = (e) => {
      console.log('none', e)
      e.preventDefault()
    };


    return {
      interfaceData,
      activeKey,
      methods,
      selectMethod,
      sendRequest,
      clearAll,
      saveName,
      copyLink,
      saveAs,
      none,
    }
  }
})

</script>

<style lang="less">
.config-main {
  height: calc(100% - 32px);

  .ant-tabs-line {
    height: 100%;
    .ant-tabs-top-content {
      height: calc(100% - 61px);
    }
  }
}
</style>

<style lang="less" scoped>
.config-main {
  .ant-tabs-tabpane-active {
    height: 100%;
  }
}
</style>