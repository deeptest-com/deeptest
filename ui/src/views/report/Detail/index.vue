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
            <StatisticTable :scene="scene" :data="detailResult.statisticData" />
            <LogTreeView :treeData="detailResult.scenarioReports"/>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, computed } from 'vue';
import { useStore } from 'vuex';

import { ReportBasicInfo, StatisticTable, ScenarioCollapsePanel, EndpointCollapsePanel,LogTreeView } from '@/views/component/Report/components';


import { StateType as ReportStateType } from "../store";
import { StateType as PlanStateType } from '@/views/plan/store';

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
