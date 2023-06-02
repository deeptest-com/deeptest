<template>
  <div class="drawer-content">
    <ReportBasicInfo :showBtn="true"
                     :btnText="'生成报告'"
                     :items="baseInfoList || []"
                     @handleBtnClick="genReport"/>
    <StatisticTable :data="statisticData" :value="statInfo"/>
    <LogTreeView :treeData="scenarioList"/>
  </div>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, computed} from 'vue';
import {useStore} from 'vuex';

import {
  ReportBasicInfo,
  StatisticTable,
  LogTreeView
} from '@/views/component/Report/components';


import {PaginationConfig, Scenario} from "@/views/scenario/data";
import {momentUtc} from "@/utils/datetime"
import {message} from "ant-design-vue";
import {ReportDetailType} from '@/utils/enum';
import {formatData} from '@/utils/formatData';
import {getDivision, getPercentStr} from "@/utils/number";

const store = useStore<{ Scenario, ProjectGlobal, ServeGlobal, }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const scenariosReports = computed(() => store.state.Scenario.scenariosReports);
const reportsDetail: any = computed<PaginationConfig>(() => store.state.Scenario.reportsDetail);


const baseInfoList = computed(() => {
  if (!reportsDetail.value) return [];
  return [
    {value: reportsDetail.value.name, label: '场景名称'},
    {value: reportsDetail.value.startTime ? momentUtc(reportsDetail.value.startTime) : '暂无', label: '执行时间'},
    {value: reportsDetail.value.execEnv || '暂无', label: '执行环境'},
    {value: reportsDetail.value.createUserName || '暂无', label: '执行人'},
  ]
})

const statInfo = computed(() => {
  const data = reportsDetail.value;
  if (data === null) return {};
  return {
    passAssertionNum: data.passAssertionNum,
    failAssertionNum: data.failAssertionNum,
    notTestNum: data.notTestNum,
  }
})
const statisticData = computed(() => {
  const data = reportsDetail.value;
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
  const passRate = getPercentStr(passAssertionNum, totalAssertionNum);
  const notPassRate = getPercentStr(failAssertionNum, totalAssertionNum);
  const notTestNum = totalAssertionNum - passAssertionNum - failAssertionNum;
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

/**
 * 适配场景执行报告日志数据
 * */
const scenarioList = computed(() => {
  return [reportsDetail.value];
})


async function genReport() {
  const res = await store.dispatch('Scenario/genReport', {
    id: reportsDetail.value.id,
  });
  if (res) {
    message.success('生成报告成功');
  } else {
    message.error('生成报告失败');
  }
}

</script>
<style scoped lang="less">
.report-drawer {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>
