<template>
  <div class="scenario-exec-info-main">
    <ReportBasicInfo :showBtn="true"
                     :btnText="'另存为报告'"
                     :items="baseInfoList || []"
                     @handleBtnClick="genReport"/>
    <StatisticTable :data="statisticData" :value="statInfo"/>
    <LogTreeView :treeData="scenarioList" :isSingleScenario="true"/>
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
import {getDivision, getPercentStr} from "@/utils/number";
import {notifyError, notifySuccess} from "@/utils/notify";

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
    {value: reportsDetail.value.createUserName || '暂无', label: '创建人'},
    //  TODO，确定字段
    {value: reportsDetail.value.execUserName || '暂无', label: '执行人'},
    {value: reportsDetail.value.priority || '未设置', label: '优先级'},
  ]
})

const statInfo = computed(() => {
  const data = JSON.parse(reportsDetail.value?.stat || '{}');
  return {
    interfacePass: data.interfacePass || 0,
    interfaceFail: data.interfaceFail || 0,
    interfaceSkip: data.interfaceSkip || 0,
  }
})
const statisticData = computed(() => {
  const data = JSON.parse(reportsDetail.value?.stat || '{}');
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

/**
 * 适配场景执行报告日志数据
 * */
const scenarioList = computed(() => {
  return reportsDetail.value.logs;
})


async function genReport() {
  const res = await store.dispatch('Scenario/genReport', {
    id: reportsDetail.value.id,
  });
  if (res) {
    notifySuccess('生成报告成功');
  } else {
    notifyError('生成报告失败');
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
