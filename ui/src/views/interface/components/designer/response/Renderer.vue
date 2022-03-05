<template>
  <div class="main">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="JSON">
      </a-tab-pane>

      <a-tab-pane key="2" tab="原始内容"></a-tab-pane>
      <a-tab-pane key="3" tab="响应头"></a-tab-pane>
      <a-tab-pane key="4" tab="测试结果"></a-tab-pane>
    </a-tabs>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";

interface ResponseRendererSetupData {
  responseResult: ComputedRef;
  activeKey: Ref<string>;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'ResponseRenderer',
  components: {
  },
  setup(props): ResponseRendererSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseResult = computed<any>(() => store.state.Interface.responseResult);
    const activeKey = ref('1');

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      responseResult,
      activeKey,
      doSomething,
    }
  }
})

</script>

<style lang="less" scoped>
.main {

}

</style>