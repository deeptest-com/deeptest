<template>
  <div class="scenario-detail-main">
    <a-card :bordered="false" :bodyStyle="{ paddingTop: '8px' }">
      <template #title>
        <div>报告详情 - {{ modelRef.name }}</div>
      </template>
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <div>
        <Log v-if="modelRef.logs" :logs="modelRef.logs"></Log>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { defineComponent, computed, ref, reactive, ComputedRef } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { useI18n } from "vue-i18n";
import { StateType } from "../store";
import Log from "@/views/scenario/exec/components/Log.vue"
import { Report } from "@/views/report/data";
const router = useRouter();
const { t } = useI18n();

const store = useStore<{ Report: StateType }>();
const modelRef = computed<Report>(() => store.state.Report.detailResult);

const get = async (id: number): Promise<void> => {
  await store.dispatch('Report/get', id);
}
const id = ref(+router.currentRoute.value.params.id)
get(id.value)

const back = (): void => {
  router.replace(`/report/index`)
}

</script>

<style lang="less" scoped>
.scenario-detail-main {
  .ant-card-body {
    padding-top: 0 !important;
  }
}
</style>
