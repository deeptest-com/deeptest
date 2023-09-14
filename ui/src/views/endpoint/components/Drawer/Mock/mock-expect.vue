<template>
  <div class="endpoint-mock-expect-main">
    <ExclamationCircleOutlined />
    <span>
      当您请求Mock接口时，会根据请求参数匹配的期望条件自动返回响应的结果，Mock请求地址：
      {{getMockUrl(serverUrl)}}/mocks/{{endpoint.id}}{{endpoint.path}}。
    </span>

    <div class="toolbar">
      是否开启<a-switch @change="disable()" class="switch" v-model:checked="advancedMockEnabled" />
      <a-button @click="createExpect" type="primary" class="btn-create">新建期望</a-button>
    </div>

    <Expect />
  </div>
  <Detail v-if="open" @cancel="open = false" />
</template>

<script setup lang="ts">
import {ref, computed, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";
import {getUrls} from "@/utils/request";
import {disableAdvMock, disableScriptMock} from "@/views/endpoint/service";
import Expect from './expect/index.vue';
import Detail from './expect/detail.vue';

const {t} = useI18n()

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);

const open = ref(false);
const advancedMockEnabled = ref(true)
watch(() => endpoint.value.advancedMockDisabled, (newVal, oldVal) => {
  console.log('watch advancedMockEnabled', endpoint.value.advancedMockEnabled)
  advancedMockEnabled.value = !endpoint.value.advancedMockDisabled
}, {immediate: true})

const {serverUrl, agentUrl} = getUrls()
const getMockUrl = (serverUrl) => {
  return serverUrl.replace('/api/v1', '')
}

const createExpect = () => {
  open.value = true;
  store.commit('Endpoint/setMockExpectDetail', {});
}

const disable = () => {
  disableAdvMock(endpoint.value.id)
  endpoint.value.advancedMockDisabled = !endpoint.value.advancedMockDisabled
}

</script>

<style lang="less" scoped>
.endpoint-mock-expect-main {
  height: 100%;
  position: relative;

  .toolbar {
    position: absolute;
    top: -58px;
    right: 0px;

    text-align: right;

    .btn-create {
      margin-left: 30px;
    }
    .switch {
      margin-left: 16px;
    }
  }
}
</style>

