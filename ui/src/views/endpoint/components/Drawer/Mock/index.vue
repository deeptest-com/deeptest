<template>
  <div class="endpoint-mock-main">
    <a-tabs type="card" v-model:activeKey="activeKey" class="tabs">
      <a-tab-pane key="expect" tab="期望" />
      <a-tab-pane key="script" tab="脚本" />
    </a-tabs>

    <div class="content">
      <EndpointMockExpect v-if="activeKey==='expect'" />

      <EndpointMockScript v-if="activeKey==='script'" />
    </div>

  </div>
</template>

<script setup lang="ts">
import {ref, computed, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";
import {getUrls} from "@/utils/request";
import EndpointMockExpect from './mock-expect.vue';
import EndpointMockScript from './mock-script.vue';

const {t} = useI18n()

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);

const activeKey = ref('expect')

</script>

<style lang="less" scoped>
.endpoint-mock-main {
  //height: calc(100vh - 100px);
  height: 100%;

  padding-top: 8px;
  position: relative;

  .content {
    height: calc(100% - 56px);
  }
}
</style>

