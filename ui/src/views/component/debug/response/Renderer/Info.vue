<template>
  <div class="info-main">
    <div class="params">
      <ParamGrid title="查询参数" :list="requestData.queryParams" />
      <ParamGrid title="路径参数" :list="requestData.pathParams" />
      <ParamGrid title="请求头" :list="requestData.headers" />

      <ParamGrid v-if="requestData.bodyType==='multipart/form-data' && requestData.bodyFormData"
                 title="表单数据"
                 :list="requestData.bodyFormData" />

      <ParamGrid v-else-if="requestData.bodyType==='application/x-www-form-urlencoded' && requestData.bodyFormUrlencoded"
                 title="表单数据（UrlEncoded）"
                 :list="requestData.bodyFormUrlencoded" />

      <ParamContent v-else
                    title="请求体"
                    :content="requestData.body || ''" />

    </div>

    <div class="status">
      <ResponseMeta />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import ParamGrid from "../../comp/param-grid.vue";
import ParamContent from "../../comp/param-content.vue";

import ResponseMeta from "./Meta.vue";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const requestData = computed<any>(() => store.state.Debug.requestData);

</script>

<style lang="less" scoped>
.info-main {
  height: 100%;
  display: flex;

  .params {
    flex: 1;
    height: 100%;
    overflow-y: auto;
  }

  .status {
    height: 100%;
    overflow: hidden;
    width: 360px;
  }
}

</style>