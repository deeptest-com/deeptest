<template>
  <div class="response-renderer">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" :tab="t(codeLang)">
        <ResponseLensJson v-if="codeLang === 'json'"></ResponseLensJson>
        <ResponseLensXml v-if="codeLang === 'xml'"></ResponseLensXml>
        <ResponseLensHtml v-if="codeLang === 'html'"></ResponseLensHtml>
        <ResponseLensImage v-if="codeLang === 'image'"></ResponseLensImage>
        <ResponseLensRaw v-if="codeLang === 'plaintext'"></ResponseLensRaw>
      </a-tab-pane>

      <a-tab-pane key="3" tab="响应头">
        <ResponseHeaders></ResponseHeaders>
      </a-tab-pane>

      <a-tab-pane key="4" tab="测试结果">
        <ResponseResult></ResponseResult>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import ResponseLensJson from "./Renderer/lenses/JSONLensRenderer.vue";
import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseResult from "./Renderer/Result.vue";
import {isInArray} from "@/utils/array";
import {getCodeLang} from "@/views/interface/service";
import ResponseLensXml from "@/views/interface/components/designer/response/Renderer/lenses/XMLLensRenderer.vue";
import ResponseLensHtml from "@/views/interface/components/designer/response/Renderer/lenses/HTMLLensRenderer.vue";
import ResponseLensImage from "@/views/interface/components/designer/response/Renderer/lenses/ImageLensRenderer.vue";
import ResponseLensRaw from "@/views/interface/components/designer/response/Renderer/lenses/RawLensRenderer.vue";

export default defineComponent({
  name: 'ResponseRenderer',
  components: {
    ResponseLensRaw,
    ResponseLensImage,
    ResponseLensHtml,
    ResponseLensXml,
    ResponseHeaders, ResponseResult, ResponseLensJson,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseData = computed<any>(() => store.state.Interface.responseData);
    const codeLang = ref('')

    codeLang.value = getCodeLang(responseData.value.contentLang)

    watch(responseData, () => {
      console.log('watch responseData')
      codeLang.value = getCodeLang(responseData.value.contentLang)
    }, {deep: true})

    const activeKey = ref('1');

    return {
      t,
      responseData,
      activeKey,
      codeLang,
    }
  }
})

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


</style>