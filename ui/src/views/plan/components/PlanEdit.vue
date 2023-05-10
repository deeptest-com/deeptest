<template>
    <a-drawer :closable="true" :width="1000" :key="currPlanId" :visible="editDrawerVisible" @close="onCancel">
        <template #title>
        <div class="drawer-header">
            <div>编辑计划</div>
        </div>
        </template>
        <div class="drawer-content">
            <ConBoxTitle title="基本信息" backgroundStyle="background: #FBFBFB" />
            <div class="plan-basic-info">
                <TextItem label="负责人" :value="planDetail.adminName" />
                <TextItem label="创建时间" :value="momentUtc(planDetail.updatedAt)" />
                <TextItem label="最近更新" :value="momentUtc(planDetail.updatedAt)" />
                <TextItem label="最新执行通过率" labelStyle="width: 108px" :value="planDetail.testPassRate" />
                <TextItem label="执行次数" :value="planDetail.execTimes" />
                <TextItem label="执行环境" :value="planDetail.execEnv" />
            </div>
            <ConBoxTitle title="关联信息" backgroundStyle="background: #FBFBFB" />
            <div class="contract-wrapper">
                <a-tabs v-model:activeKey="activeKey">
                    <a-tab-pane key="1" tab="测试场景">
                        <ScenarioList />
                    </a-tab-pane>
                    <a-tab-pane key="2" tab="测试报告" force-render>
                        <ReportList />
                    </a-tab-pane>
                </a-tabs>
            </div> 
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, watch, reactive, computed } from 'vue';
import { useStore } from 'vuex';

import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import TextItem from '@/views/component/Report/TextItem.vue';
import ScenarioList from './ScenarioList.vue';
import ReportList from './ReportList.vue';

import { momentUtc } from '@/utils/datetime';
import { StateType as PlanStateType } from '../store';

const props = defineProps<{
    currPlanId: Number
    editDrawerVisible: Boolean
}>();

const store = useStore<{ Plan: PlanStateType }>();
const planDetail = computed(() => store.state.Plan.detailResult);
const planScenarioList = computed(() => store.state.Plan.scenarioList);
console.log(planScenarioList);
const emits = defineEmits(['onCancel']);
const activeKey = ref('1');

function onCancel() {
    emits('onCancel');
}

watch(() => {
    return props.currPlanId;
}, (val) => {
    if (val) {
        store.dispatch('Plan/getPlanDetail', val);
    }
})
</script>
<style scoped lang="less">
.plan-basic-info {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    position: relative;
    padding: 20px 0;
    .text-wrapper {
        width: 33%;
        height: 30px;
    }
}

.contract-wrapper {
    padding-top: 10px;
}
</style>
