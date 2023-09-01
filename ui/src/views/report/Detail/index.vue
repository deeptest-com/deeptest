<template>
  <a-drawer 
    class="report-drawer" 
    wrapClassName="report-drawer-exec-history" 
    :closable="true" 
    :width="1000" 
    :visible="drawerVisible"
    @close="onClose">
    <template #title>
      <div class="drawer-header">
        <div>{{ detailResult.name || '测试报告详情' }}</div>
      </div>
    </template>
    <div class="report-exec-info-main">
      <ReportBasicInfo :items="detailResult.basicInfoList" :scene="scene" :show-operation="true"/>
      <StatisticTable :data="statisticData" :value="statInfo"/>
      <LogTreeView :treeData="scenarioReports"/>
    </div>
  </a-drawer>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, computed} from 'vue';
import {useStore} from 'vuex';

import {ReportBasicInfo, StatisticTable, LogTreeView} from '@/views/component/Report/components';

import {StateType as ReportStateType} from "../store";
import {StateType as PlanStateType} from '@/views/plan/store';
import {getDivision, getPercentStr} from "@/utils/number";

const props = defineProps<{
  drawerVisible: boolean
  title: string
  scenarioExpandActive: boolean
  showScenarioInfo: boolean
  scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
  reportId?: number
}>();

const emits = defineEmits(['onClose', 'execCancel']);


const store = useStore<{ Report: ReportStateType, Plan: PlanStateType }>();
const detailResult = computed<any>(() => store.state.Report.detailResult);
const expandActive = ref(props.scenarioExpandActive || false);

const scenarioReports = computed(() => {
  return detailResult.value.scenarioReports?.map((item) => {
    if (item?.logs?.length > 0) {
      return item.logs[0]
    }
  })
})
const statInfo = computed(() => {
  const data = JSON.parse(detailResult.value?.stat || '{}');
  return {
    interfacePass: data.interfacePass || 0,
    interfaceFail: data.interfaceFail || 0,
    interfaceSkip: data.interfaceSkip || 0,
  }
})
const statisticData = computed(() => {
  const data = JSON.parse(detailResult.value?.stat || '{}');
  const {
    checkpointFail= 0,
    checkpointPass= 0,
    interfaceCount= 0,
    interfaceDurationAverage= 0,
    interfaceDurationTotal= 0,
    interfaceFail= 0,
    interfacePass= 0,
    interfaceSkip= 0,
  } = data;
  const passRate = getPercentStr(interfacePass, interfaceCount);
  const notPassRate = getPercentStr(interfaceFail, interfaceCount);
  const notTestNumRate = getPercentStr(interfaceSkip, interfaceCount);
  return [
    {
      label: '通过接口',
      value: `${interfacePass} 个`,
      rate: passRate,
      class: 'success',
    },
    {
      label: '接口总耗时',
      value: `${interfaceDurationTotal} 毫秒`
    },
    {
      label: '失败接口',
      rate: notPassRate,
      value: `${interfaceFail} 个`,
      class: 'fail',
    },
    {
      label: '平均接口耗时',
      value: `${interfaceDurationAverage} 毫秒`,
    },
    {
      label: '未测接口',
      value: `${interfaceSkip}个`,
      rate: notTestNumRate,
      class: 'notest',
    },
    {
      label: '检查点 (成功/失败)',
      value: `${checkpointPass + checkpointFail} (${checkpointPass}/${checkpointFail})`,
    },
  ]

})

function onClose() {
  emits('onClose');
}

function execCancel() {
  emits('execCancel');
}

</script>
<style scoped lang="less">
.report-drawer {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>
