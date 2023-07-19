<template>
  <div class="response-renderer">
    <a-tabs v-model:activeKey="activeKey" class="dp-tabs-full-height">
      <a-tab-pane key="body" :tab="title">
        <ResponseLensJson v-if="responseData.contentLang === 'json'"></ResponseLensJson>
        <ResponseLensXml v-if="responseData.contentLang === 'xml'"></ResponseLensXml>
        <ResponseLensHtml v-if="responseData.contentLang === 'html'"></ResponseLensHtml>
        <ResponseLensRaw v-if="responseData.contentLang === 'text'"></ResponseLensRaw>
        <ResponseLensImage v-if="responseData.contentLang === 'image'"></ResponseLensImage>
      </a-tab-pane>

      <a-tab-pane key="header" tab="响应头">
        <ResponseHeaders v-if="activeKey === 'header'"></ResponseHeaders>
      </a-tab-pane>

      <a-tab-pane key="cookie" tab="Cookie">
        <ResponseCookies v-if="activeKey === 'cookie'"></ResponseCookies>
      </a-tab-pane>

      <a-tab-pane key="info" tab="详情">
        <ResponseInfo v-if="activeKey === 'info'"></ResponseInfo>
        <!-- <template #tab>
          <a-badge v-if="extractorFail" dot><span class="link">提取器</span></a-badge>
          <span v-else>提取器</span>
        </template> -->
      </a-tab-pane>

    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseCookies from "./Renderer/Cookies.vue";
import ResponseInfo from "./Renderer/Info.vue";

import ResponseLensJson from "./Renderer/lenses/JSONLensRenderer.vue";
import ResponseLensXml from "@/views/component/debug/response/Renderer/lenses/XMLLensRenderer.vue";
import ResponseLensHtml from "@/views/component/debug/response/Renderer/lenses/HTMLLensRenderer.vue";
import ResponseLensImage from "@/views/component/debug/response/Renderer/lenses/ImageLensRenderer.vue";
import ResponseLensRaw from "@/views/component/debug/response/Renderer/lenses/RawLensRenderer.vue";
import {UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const title = computed(() => t(responseData.value.contentLang ? responseData.value.contentLang : 'empty'))
const activeKey = ref('body');

/* const extractorFail = computed(() => {
  for (let val of extractorsData.value) {
    if (val.result==='extractor_err') return true
  }
  return false
}) */

</script>

<style lang="less">
.response-renderer {
  height: 100%;

  .ant-tabs-line {
    height: 100%;
    .ant-tabs-top-content {
      height: calc(100% - 61px);
    }
  }
}
</style>

<style lang="less" scoped>
  .link {
    color: #009688;
  }
</style>