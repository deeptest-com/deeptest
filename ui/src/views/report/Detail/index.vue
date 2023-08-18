<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ detailResult.name || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo  :items="detailResult.basicInfoList" :scene="scene" :show-operation="true" />
            <StatisticTable  :data="statisticData" :value="statInfo"/>

            <LogTreeView :treeData="detailResult.scenarioReports"/>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, computed } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable,LogTreeView } from '@/views/component/Report/components';

import { StateType as ReportStateType } from "../store";
import { StateType as PlanStateType } from '@/views/plan/store';
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

const statInfo = computed(() => {
  const data = detailResult.value || {};
  if (data === null) return {};
  return {
    passAssertionNum: data.passAssertionNum,
    failAssertionNum: data.failAssertionNum,
    notTestNum: data.notTestNum,
  }
})
const statisticData = computed(() => {
  const data = detailResult.value || {};
  const {
    totalAssertionNum = 0,
    totalInterfaceNum = 0,
    totalProcessorNum = 0,
    totalRequestNum = 0,
    failAssertionNum = 0,
    failInterfaceNum = 0,
    failRequestNum = 0,
    finishProcessorNum = 0,
    passInterfaceNum = 0,
    passRequestNum = 0,
    passAssertionNum = 0,
    duration = 0,
  } = data;
  // 计算平均接口耗时
  let interfaceDuration = 0;
  const notTestNum = totalInterfaceNum - passInterfaceNum - failInterfaceNum;
  const passRate = getPercentStr(passAssertionNum, totalAssertionNum);
  const notPassRate = getPercentStr(failAssertionNum, totalAssertionNum);
  const notTestNumRate = getPercentStr(notTestNum, totalAssertionNum);
  // // 平均接口耗时
  const avgInterfaceDuration = getDivision(duration, totalRequestNum);
  return [
    {
      label: '通过',
      class: 'success',
      rate: passRate,
      value: `${passAssertionNum} 个`,
    },
    {
      label: '接口总耗时',
      value: `${duration} 毫秒`
    },
    {
      label: '失败',
      rate: notPassRate,
      value: `${failAssertionNum} 个`,
      class: 'fail',
    },
    {
      label: '平均接口耗时',
      value: `${avgInterfaceDuration} 毫秒`,
    },
    {
      label: '未测',
      value: ` ${notTestNum}个`,
      rate: notTestNumRate,
      class: 'notest',
    },
    {
      label: '检查点 (成功/失败)',
      value: `${totalAssertionNum} (${passAssertionNum}/${failAssertionNum})`,
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
