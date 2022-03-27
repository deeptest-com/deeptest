<template>
  <div class="response-renderer">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="JSON">
        <ResponseLensJson></ResponseLensJson>
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
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import ResponseLensJson from "./Renderer/lenses/JSONLensRenderer.vue";
import ResponseHeaders from "./Renderer/Headers.vue";
import ResponseResult from "./Renderer/Result.vue";

interface ResponseRendererSetupData {
  responseData: ComputedRef;
  activeKey: Ref<string>;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'ResponseRenderer',
  components: {
    ResponseHeaders, ResponseResult, ResponseLensJson,
  },
  setup(props): ResponseRendererSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseData = computed<any>(() => store.state.Interface.responseData);
    const activeKey = ref('1');

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      responseData,
      activeKey,
      doSomething,
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