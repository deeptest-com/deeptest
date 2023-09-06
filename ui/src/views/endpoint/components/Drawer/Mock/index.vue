<template>
  <div class="endpoint-mock-main">
    <a-tabs type="card" v-model:activeKey="activeKey" class="tabs">
      <a-tab-pane key="expect" tab="期望">

      </a-tab-pane>
      <a-tab-pane key="script" tab="脚本">

      </a-tab-pane>
    </a-tabs>

    <div class="content">
      <ExclamationCircleOutlined />
      当您请求Mock接口时，会根据请求参数匹配的期望条件自动返回响应的结果，Mock请求地址：
      {{serverUrl}}/mocks/{{endpoint.id}}{{endpoint.path}}。
    </div>

    <div class="toolbar">
      是否开启<a-switch v-model:checked="advancedMockEnabled" class="switch" />
      <a-button @click="createExpect" type="primary" class="btn-create">新建期望</a-button>

    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, computed, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";
import {getUrls} from "@/utils/request";

const {t} = useI18n()

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);

const {serverUrl, agentUrl} = getUrls()
const activeKey = ref('expect')
const advancedMockEnabled = ref(true)

watch(() => endpoint.value.advancedMockDisabled, (newVal, oldVal) => {
  advancedMockEnabled.value = !endpoint.value.advancedMockDisabled
})

const createExpect = () => {
  console.log('createExpect')
}

</script>

<style lang="less" scoped>
.endpoint-mock-main {
  height: 100%;
  padding-top: 8px;
  position: relative;

  .tabs {
  }
  .content {

  }
  .toolbar {
    margin-top: -52px;
    text-align: right;

    .switch {
      margin-left: 6px;
    }
    .btn-create {
      margin-left: 30px;
    }
  }
}
</style>

