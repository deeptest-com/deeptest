<template>
  <div class="bearer-token-main author-content">
    <div class="params">
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Token</span>
        </a-col>

        <a-col flex="1">
          <a-input v-model:value="interfaceData.bearerToken.token" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
    </div>
    <div class="tips">
      <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
      <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent, inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ArrowRightOutlined, DeleteOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface1/store";
import {Interface} from "@/views/interface1/data";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

export default defineComponent({
  name: 'RequestAuthorBearerToken',
  components: {
    ArrowRightOutlined,
  },
  setup(props) {
    const usedBy = inject('usedBy') as UsedBy
    const {t} = useI18n();
    const store = useStore<{ Interface1: StateType, Scenario: ScenarioStateType }>();
    const interfaceData = computed<Interface>(
        () => usedBy === UsedBy.InterfaceDebug ? store.state.Interface1.interfaceData : store.state.Scenario.interfaceData);

    return {
      interfaceData,
    }
  }
})

</script>

<style lang="less" scoped>
.bearer-token-main {
}

</style>