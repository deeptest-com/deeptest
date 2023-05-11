<template>
    <a-modal
        class="associate-scenario-modal"
        title="关联测试场景"
        :visible="associateModalVisible"
        :closable="true"
        @cancel="handleCancel"
        @ok="onOk"
        width="1000px">
        <ScenarioList
            :loading="loading"
            :list="scenarioList"
            :pagination="pagination"
            :columns="columns"
            :show-scenario-operation="false"
            @refresh-list="getScenarioList"
            @select-row-keys="handleSelectRowKeys" />
    </a-modal>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, reactive, computed, watch, toRefs } from 'vue';
import { useStore } from 'vuex';
import ScenarioList from './PlanList.vue';

import { StateType as ScenarioStateType } from '@/views/scenario/store';
import { StateType as PlanStateType } from '../../store';
import { PaginationConfig, QueryParams } from '@/views/scenario/data';

const props = defineProps<{
    associateModalVisible: Boolean,
}>();
const store = useStore<{ Scenario: ScenarioStateType, Plan: PlanStateType }>();
const emits = defineEmits(['onCancel', 'onOk']);
const scenarioList = computed<any[]>(() => store.state.Scenario.listResult.list);
const currPlanId = computed(() => store.state.Plan.planId);
let pagination = computed<PaginationConfig>(() => store.state.Scenario.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});
let selectedScenarioIds: number[] = [];
const loading = ref<boolean>(false);

const columns: any[] = reactive([
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
        title: '最近更新',
        dataIndex: 'updatedAt',
    }
]);

function handleSelectRowKeys(value: any[]) {
    selectedScenarioIds = value;
}

function handleCancel() {
    emits('onCancel');
}

async function onOk() {
    console.log('selectScenarioIds: --', selectedScenarioIds);
    await store.dispatch('Plan/addScenario', { planId: currPlanId.value, params: { scenarioIds: selectedScenarioIds } });
    emits('onOk');
}

async function getScenarioList(params) {
    loading.value = true;
    await store.dispatch('Scenario/listScenario', {
        ...queryParams,
        params
    })
    loading.value = false;
}

watch(() => {
    return props.associateModalVisible;
}, (val) => {
    if (val) {
        getScenarioList(queryParams);
    }
}, { immediate: true })
</script>
<style scoped lang="less">
:deep(.ant-modal.associate-scenario-modal) {
    top: 50%;
    transform: translateY(-50%);
}

</style>
