<template>
  <div class="response-meta">
    <a-row type="flex" :gutter="15">
      <a-col>
        <span>状态：{{ responseData.statusContent }}</span>
      </a-col>
      <a-col>
        <span>耗时：{{ responseData.time }} 毫秒</span>
      </a-col>
      <a-col>
        <span>大小：{{ responseData.contentLength }} 字节</span>
      </a-col>
    </a-row>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface1/store";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

interface ResponseMetaSetupData {
  responseData: ComputedRef;

  doSomething: (e) => void;
}

export default defineComponent({
  name: 'ResponseMeta',
  components: {
  },
  setup(props): ResponseMetaSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface1: StateType; Scenario: ScenarioStateType }>();
    const responseData = computed<any>(
        () => UsedBy.interface ? store.state.Interface1.responseData : store.state.Scenario.responseData);

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      responseData,
      doSomething,
    }
  }
})

</script>

<style lang="less" scoped>
.response-meta {
  padding: 6px;
}

</style>