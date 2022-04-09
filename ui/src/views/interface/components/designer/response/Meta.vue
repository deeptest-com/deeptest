<template>
  <div class="response-meta">
    <a-row type="flex" :gutter="15">
      <a-col>
        <span>状态：{{ responseData.statusContent }}</span>
      </a-col>
      <a-col>
        <span>耗时：200毫秒</span>
      </a-col>
      <a-col>
        <span>大小：600字节</span>
      </a-col>
    </a-row>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";

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
    const store = useStore<{ Interface: StateType }>();
    const responseData = computed<any>(() => store.state.Interface.responseData);

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