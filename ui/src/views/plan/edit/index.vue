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
                    <a-tab-pane key="test-scenario" tab="测试场景">
                        <div style="padding-top: 20px">
                            <ScenarioList 
                                :list="planScenarioList" 
                                :show-scenario-operation="true" 
                                :columns="columns"
                                :loading="loading"
                                :pagination="scenarioPagination"
                                @refresh-list="getScenarioList" />
                        </div>
                    </a-tab-pane>
                    <a-tab-pane key="test-report" tab="测试报告" force-render>
                        <div style="padding-top: 20px">
                            <ReportList />
                        </div>
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
import ScenarioList from '../components/ScenarioList.vue';
import ReportList from '../components/ReportList.vue';

import { momentUtc } from '@/utils/datetime';
import { StateType as PlanStateType } from '../store';

const props = defineProps<{
    editDrawerVisible: Boolean
    tabActiveKey?: String
}>();

const store = useStore<{ Plan: PlanStateType }>();
const planDetail = computed(() => store.state.Plan.detailResult);
const planScenarioList = computed<any[]>(() => store.state.Plan.scenarioListResult.scenarioList);
const scenarioPagination = computed<any>(() => store.state.Plan.scenarioListResult.pagination);
const currPlanId = computed<number>(() => store.state.Plan.planId);
console.log(planScenarioList);
const emits = defineEmits(['onCancel']);
const activeKey = ref(props.tabActiveKey || 'test-scenario');
const loading = ref(false);

const columns: any[] = reactive([
    {
        title: '编号',
        dataIndex: 'serialNumber',
    },
    {
        title: '用例名称',
        dataIndex: 'name',
    },
    {
        title: '状态',
        dataIndex: 'status',
    },
    {
        title: '优先级',
        dataIndex: 'priority',
    },
    {
        title: '所属分类',
        dataIndex: 'categoryName',
    },
    {
        title: '创建人',
        dataIndex: 'createUserName',
    },
    {
        title: '最近更新',
        dataIndex: 'updatedAt',
    },
    {
        title: '操作',
        dataIndex: 'operation',
        slots: { customRender: 'operation' }
    },
]);

function onCancel() {
    emits('onCancel');
}

// 移除-关联-筛选时重新获取已关联的场景列表
async function getScenarioList(params: any) {
    loading.value = true;
    await store.dispatch('Plan/getScenarioList', { ...params, planId: currPlanId.value });
    loading.value = false;
}

watch(() => {
    return currPlanId.value;
}, (val) => {
    if (val) {
        getScenarioList({ planId: val });
    }
}, { immediate: true });

watch(() => {
    return props.tabActiveKey;
}, (val) => {
    console.log('props- tabActiveKey', val);
    activeKey.value = val || 'test-scenario';
}, { immediate: true });
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
