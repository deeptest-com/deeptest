<template>
  <div class="main">
    <a-row type="flex" :gutter="15">
      <a-col>
        <span>状态：200.OK</span>
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
  responseResult: ComputedRef;

  doSomething: (e) => void;
}

export default defineComponent({
  name: 'ResponseMeta',
  components: {
  },
  setup(props): ResponseMetaSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseResult = computed<any>(() => store.state.Interface.modelResult);

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      responseResult,
      doSomething,
    }
  }
})

</script>

<style lang="less" scoped>
.main {
  padding: 6px;
}

</style>