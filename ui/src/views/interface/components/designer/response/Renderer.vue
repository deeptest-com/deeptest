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
        <ResponseHeaders></ResponseHeaders>
      </a-tab-pane>

      <a-tab-pane key="4">
        <ResponseExtractor></ResponseExtractor>

        <template #tab>
          <a-badge v-if="extractorFail" dot><span class="link">提取器</span></a-badge>
          <span v-else>提取器</span>
        </template>

      </a-tab-pane>

      <a-tab-pane key="5">
        <ResponseCheckpoint></ResponseCheckpoint>

        <template #tab>
          <a-badge v-if="checkpointFail" dot><span class="link">检查点</span></a-badge>
          <span v-else>检查点</span>
        </template>

      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, inject, PropType, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import { ExclamationOutlined } from '@ant-design/icons-vue';
import ResponseLensJson from "./Renderer/lenses/JSONLensRenderer.vue";
import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseExtractor from "./Extractor.vue";
import ResponseCheckpoint from "./Checkpoint.vue";
import ResponseLensXml from "@/views/interface/components/designer/response/Renderer/lenses/XMLLensRenderer.vue";
import ResponseLensHtml from "@/views/interface/components/designer/response/Renderer/lenses/HTMLLensRenderer.vue";
import ResponseLensImage from "@/views/interface/components/designer/response/Renderer/lenses/ImageLensRenderer.vue";
import ResponseLensRaw from "@/views/interface/components/designer/response/Renderer/lenses/RawLensRenderer.vue";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface: StateType, Scenario: ScenarioStateType }>();
const responseData = computed<any>(
    () => usedBy === UsedBy.interface ? store.state.Interface.responseData: store.state.Scenario.responseData);
const extractorsData = computed(
    () => usedBy === UsedBy.interface ? store.state.Interface.extractorsData: store.state.Scenario.extractorsData);
const checkpointsData = computed(
    () => usedBy === UsedBy.interface ? store.state.Interface.checkpointsData : store.state.Scenario.checkpointsData);

const title = ref(t('text'))

watch(responseData, () => {
  console.log('watch responseData')
  title.value = t(responseData.value.contentLang ? responseData.value.contentLang : 'empty')
}, {deep: true})

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

const listExtractor = () => {
  usedBy === UsedBy.interface ? store.dispatch('Interface/listExtractor', usedBy) :
      store.dispatch('Scenario/listExtractor', usedBy)
}
listExtractor()

const listCheckPoint = () => {
  usedBy === UsedBy.interface ? store.dispatch('Interface/listCheckpoint', usedBy) :
      store.dispatch('Scenario/listCheckpoint', usedBy)
}
listCheckPoint()

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