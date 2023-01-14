<template>
  <div class="headers-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">响应头</a-col>
        <a-col flex="80px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>复制</template>
            <CopyOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in responseData.headers" :key="idx" type="flex" class="item">
        <a-col flex="1">
          <a-input v-model:value="item.name" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="80px">
          <a-input v-model:value="item.value" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, inject, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {
  QuestionCircleOutlined,
  DeleteOutlined,
  PlusOutlined,
  CheckCircleOutlined,
  CloseCircleOutlined,
  CopyOutlined
} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface, Response} from "@/views/interface/data";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface: StateType, Scenario: ScenarioStateType }>();
const responseData = computed<Response>(
    () =>  usedBy === UsedBy.interface ? store.state.Interface.responseData : store.state.Scenario.responseData);

const doSomething = (e) => {
  console.log('doSomething', e)
};

</script>

<style lang="less" scoped>
.headers-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .items {
    height: calc(100% - 28px);
    overflow-y: auto;
    .item {
      padding: 2px 3px;
      border-bottom: 1px solid #d9d9d9;

      .ant-col {
        border-right: 1px solid #d9d9d9;

        input {
          margin-top: 1px;
        }
      }
    }
  }

}

</style>