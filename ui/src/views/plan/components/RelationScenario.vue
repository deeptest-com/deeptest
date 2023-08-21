<template>
    <a-modal
        class="associate-scenario-modal"
        title="关联测试场景1"
        :visible="associateModalVisible"
        :closable="true"
        @cancel="handleCancel"
        @ok="onOk"
        width="1000px">
        <ScenarioList
            :loading="loading"
            :list="scenarioLists"
            :pagination="pagination"
            :columns="columns"
            :show-scenario-operation="false"
            :selectedKeys="selectedScenarioIds"
            @refresh-list="getScenarioList"
            @select-row-keys="handleSelectRowKeys" />
    </a-modal>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, reactive, computed, watch } from 'vue';
import { useStore } from 'vuex';
import ScenarioList from './ScenarioList.vue';

import { StateType as PlanStateType } from '../store';

const props = defineProps<{
    associateModalVisible: Boolean,
}>();
const store = useStore<{ Plan: PlanStateType }>();
const emits = defineEmits(['onCancel', 'onOk']);
const scenarioLists = computed<any[]>(() => store.state.Plan.scenarios.list);
const currPlan = computed<any>(() => store.state.Plan.currPlan);
let pagination = computed<any>(() => store.state.Plan.scenarios.pagination);
let queryParams = reactive<any>({
  keywords: '', enabled: '1',
  planId: currPlan.value.id,
  page: pagination.value.current, pageSize: pagination.value.pageSize
});
const selectedScenarioIds = ref<number[]>([]);
const loading = ref<boolean>(false);

const columns: any[] = reactive([
    {
        title: '用例名称',
        dataIndex: 'name',
        width: 300,
        slots: { customRender: 'name' }
    },
    {
        title: '状态',
        width: 90,
        dataIndex: 'status',
        slots: { customRender: 'status' }
    },
    {
        title: '优先级',
        width: 110,
        dataIndex: 'priority',
    },
    {
        title: '最近更新',
        width: 180,
        slots: { customRender: 'updatedAt' },
        dataIndex: 'updatedAt',
    }
]);

function handleSelectRowKeys(value: any[]) {
    selectedScenarioIds.value = value;
}

function handleCancel() {
    selectedScenarioIds.value = [];
    emits('onCancel');
}

async function onOk() {
    console.log('selectScenarioIds: --', selectedScenarioIds);
    await store.dispatch('Plan/addScenario', { planId: currPlan.value.id, params: { scenarioIds: selectedScenarioIds.value } });
    selectedScenarioIds.value = [];
    emits('onOk');
}

async function getScenarioList(params) {
    loading.value = true;
    await store.dispatch('Plan/getScenarioList', {
        ...queryParams,
        ...params
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
