<template>
    <a-drawer :closable="true" :width="1000" :key="currPlan && currPlan.id" :visible="editDrawerVisible" @close="onCancel">
        <template #title>
            <div class="drawer-header" style="width: 360px">
                <EditAndShowField :value="(currPlan && currPlan.name) || '暂无'" placeholder="输入计划名称" @update="handleUpdateName" />
            </div>
        </template>
        <div class="drawer-content">
            <ConBoxTitle title="基本信息" backgroundStyle="background: #FBFBFB" />
            <div class="plan-basic-info">
                <a-descriptions :title="null" size="small">
                    <a-descriptions-item label="负责人">{{ planDetail.adminName }}</a-descriptions-item>
                    <a-descriptions-item label="创建时间">{{ momentUtc(planDetail.createdAt) }}</a-descriptions-item>
                    <a-descriptions-item label="最近更新">{{ momentUtc(planDetail.updatedAt) }}</a-descriptions-item>
                    <a-descriptions-item label="最新执行通过率">{{ planDetail.testPassRate }}</a-descriptions-item>
                    <a-descriptions-item label="执行次数">{{ planDetail.execTimes }}</a-descriptions-item>
                    <a-descriptions-item label="最近执行">{{ planDetail.execTime ? momentUtc(planDetail.execTime) : '' }}</a-descriptions-item>
                    <a-descriptions-item label="执行环境">{{ planDetail.execEnv }}</a-descriptions-item>
                    <a-descriptions-item label="状态">
                        <EditAndShowSelect
                            :label="planStatusTextMap.get((planDetail?.status || 'draft'))"
                            :value="planDetail.status"
                            :options="planStatusOptions"
                            @update="handleChangeStatus"/>
                    </a-descriptions-item>
                </a-descriptions>
                <a-button class="plan-exec" type="primary" @click="handleEnvSelect">执行计划</a-button>
            </div>
            <ConBoxTitle title="关联信息" backgroundStyle="background: #FBFBFB" />
            <div class="contract-wrapper">
                <a-tabs v-model:activeKey="activeKey" @change="handleTabChanged">
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
                            <ReportList :show-report-list="activeKey === 'test-report'" />
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
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import { ScenarioList, ReportList } from '../components';
import{ momentUtc } from '@/utils/datetime';
import { StateType as PlanStateType } from '../store';
import { planStatusOptions, planStatusTextMap } from '@/config/constant';

const props = defineProps<{
    editDrawerVisible: Boolean
    tabActiveKey?: String
}>();

const store = useStore<{ Plan: PlanStateType }>();
const planDetail = computed<any>(() => store.state.Plan.detailResult);
const planScenarioList = computed<any[]>(() => store.state.Plan.relationScenarios.scenarioList);
const scenarioPagination = computed<any>(() => store.state.Plan.relationScenarios.pagination);
const currPlan = computed<any>(() => store.state.Plan.currPlan);
const emits = defineEmits(['onCancel', 'onSelectEnv', 'onUpdate', 'update:tabKey']);
const activeKey = ref<string>('test-scenario');
const loading = ref(false);

const columns: any[] = reactive([
    {
        title: '编号',
        dataIndex: 'serialNumber',
    },
    {
        title: '用例名称',
        dataIndex: 'name',
        ellipsis: true
    },
    {
        title: '状态',
        dataIndex: 'status',
        slots: { customRender: 'status' }
    },
    {
        title: '优先级',
        dataIndex: 'priority',
    },
    {
        title: '所属分类',
        dataIndex: 'categoryName',
        ellipsis: true
    },
    {
        title: '创建人',
        dataIndex: 'createUserName',
    },
    {
        title: '最近更新',
        dataIndex: 'updatedAt',
        width: 160,
        slots: { customRender: 'updateAt' }
    },
    {
        title: '操作',
        dataIndex: 'operation',
        slots: { customRender: 'operation' },
    },
]);

function onCancel() {
    emits('onCancel');
}

function handleEnvSelect() {
    emits('onSelectEnv');
}

function handleChangeStatus(value) {
    console.log('changeStatus --', value);
    emits('onUpdate', { status: value });
}

function handleUpdateName(value) {
    emits('onUpdate', { name: value });
}

function handleTabChanged(val) {
    emits('update:tabKey', val);
}

// 移除-关联-筛选时重新获取已关联的场景列表
async function getScenarioList(params: any) {
    loading.value = true;
    await store.dispatch('Plan/getRelationScenarios', { ...params, planId: currPlan.value.id });
    loading.value = false;
}

watch([currPlan, () => props.editDrawerVisible], async (val: any) => {
    const [plan, visible] = val;
    if (plan && plan.id && visible) {
        await store.dispatch('Plan/getPlan', currPlan.value.id);
        getScenarioList({ planId: val.id });
    }
});

watch(() => props.tabActiveKey, (val: any) => {
    console.log('props- tabActiveKey', val);
    activeKey.value = val || 'test-scenario';
}, { deep: true });
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

    .plan-exec {
        position: absolute;
        top: -30px;
        right: -10px;
    }
}

.contract-wrapper {
    padding-top: 10px;
}
</style>
