<template>
  <div class="response-renderer">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" :tab="title">
        <ResponseLensJson v-if="responseData.contentLang === 'json'"></ResponseLensJson>
        <ResponseLensXml v-if="responseData.contentLang === 'xml'"></ResponseLensXml>
        <ResponseLensHtml v-if="responseData.contentLang === 'html'"></ResponseLensHtml>
        <ResponseLensRaw v-if="responseData.contentLang === 'text'"></ResponseLensRaw>
        <ResponseLensImage v-if="responseData.contentLang === 'image'"></ResponseLensImage>
      </a-tab-pane>

      <a-tab-pane key="3" tab="响应头">
        <ResponseHeaders v-if="activeKey === '3'"></ResponseHeaders>
      </a-tab-pane>

      <a-tab-pane key="4">
        <ResponseExtract v-if="activeKey === '4'"></ResponseExtract>

        <template #tab>
          <a-badge v-if="extractorFail" dot><span class="link">提取器</span></a-badge>
          <span v-else>提取器</span>
        </template>

      </a-tab-pane>

      <a-tab-pane key="5">
        <ResponseCheck v-if="activeKey === '5'"></ResponseCheck>

        <template #tab>
          <a-badge v-if="checkpointFail" dot><span class="link">检查点</span></a-badge>
          <span v-else>检查点</span>
        </template>

      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import ResponseLensJson from "./Renderer/lenses/JSONLensRenderer.vue";
import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseExtract from "./Extractor.vue";
import ResponseCheck from "./Checkpoint.vue";

import ResponseLensXml from "@/views/interface1/components/designer/response/Renderer/lenses/XMLLensRenderer.vue";
import ResponseLensHtml from "@/views/interface1/components/designer/response/Renderer/lenses/HTMLLensRenderer.vue";
import ResponseLensImage from "@/views/interface1/components/designer/response/Renderer/lenses/ImageLensRenderer.vue";
import ResponseLensRaw from "@/views/interface1/components/designer/response/Renderer/lenses/RawLensRenderer.vue";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);
const extractorsData = computed<any>(() => store.state.Debug.extractorsData);
const checkpointsData = computed<any>(() => store.state.Debug.checkpointsData);

const title = computed(() => t(responseData.value.contentLang ? responseData.value.contentLang : 'empty'))
const activeKey = ref('1');

const extractorFail = computed(() => {
  for (let val of extractorsData.value) {
    if (val.result==='extractor_err') return true
  }
  return false
})

const checkpointFail = computed(() => {
  for (let val of checkpointsData.value) {
    if (val.resultStatus==='fail') return true
  }
  return false
})

// const listExtractor = () => {
//   usedBy === UsedBy.InterfaceDebug ? store.dispatch('Interface1/listExtractor') :
//       store.dispatch('Scenario/listExtractor')
// }
// listExtractor()
//
// const listCheckPoint = () => {
//   usedBy === UsedBy.InterfaceDebug ? store.dispatch('Interface1/listCheckpoint') :
//       store.dispatch('Scenario/listCheckpoint')
// }
// listCheckPoint()

</script>

<style lang="less">
.response-renderer {
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
  .link {
    color: #009688;
  }
</style>