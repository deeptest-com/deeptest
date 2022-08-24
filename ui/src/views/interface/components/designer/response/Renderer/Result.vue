<template>
  <div class="response-result-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>合计6个检查点，<span class="dp-color-fail">3</span>个失败。</span>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <div class="line">
        <span class="dp-color-fail dp-tip-small">
          <CloseOutlined class="dp-icon-btn dp-trans-80" />
        </span>
        <span>期待结果'200'，实际为'400'。</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { CloseOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface, Response} from "@/views/interface/data";

export default defineComponent({
  name: 'ResponseResult',
  components: {
    CloseOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseData = computed<Response>(() => store.state.Interface.responseData);

    return {
      responseData,
    }
  }
})

</script>

<style lang="less">
.response-result-main {
}
</style>

<style lang="less" scoped>
.response-result-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    padding: 6px;
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
    .line {
      font-size: 13px;
    }
  }
}
</style>