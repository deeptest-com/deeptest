<template>
  <div class="response-renderer">
    <div class="left">
      <a-tabs v-model:activeKey="activeKey" class="dp-tabs-full-height">
        <a-tab-pane key="body" :tab="title" class="uppercase">
          <ResponseLensJson v-if="responseData.contentLang === 'json'" />
          <ResponseLensHtml v-else-if="responseData.contentLang === 'html'" />
          <ResponseLensXml v-else-if="responseData.contentLang === 'xml'" />
          <ResponseLensRaw v-else-if="responseData.contentLang === 'text'" />
          <ResponseLensImage v-else-if="isImage(responseData.contentType)" />
        </a-tab-pane>

        <a-tab-pane key="header" tab="响应头">
          <ResponseHeaders v-if="activeKey === 'header'" />
        </a-tab-pane>

        <a-tab-pane key="cookie" tab="Cookie">
          <ResponseCookies v-if="activeKey === 'cookie'" />
        </a-tab-pane>

        <a-tab-pane key="console" tab="控制台">
          <ResponseConsole v-if="activeKey === 'console'" />
        </a-tab-pane>

        <a-tab-pane key="info" tab="实际请求">
          <ResponseInfo v-if="activeKey === 'info'" />
          <!-- <template #tab>
            <a-badge v-if="extractorFail" dot><span class="link">提取器</span></a-badge>
            <span v-else>提取器</span>
          </template> -->
        </a-tab-pane>

      </a-tabs>
    </div>

    <div class="right">
      <ResponseResult />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseCookies from "./Renderer/Cookies.vue";
import ResponseConsole from "./Renderer/Console.vue";
import ResponseInfo from "./Renderer/Info.vue";

import ResponseResult from "./Renderer/Result.vue";
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

const isImage = (type) => {
  return type && type.indexOf('image') > -1
}

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
  display: flex;

  .left {
    height: 100%;
    flex: 1;
    .ant-tabs-line {
      height: 100%;
    }
    .link {
      color: #009688;
    }
  }
  .right {
    height: 100%;
    width: 360px;
  }


}
</style>

<style lang="less" scoped>

</style>