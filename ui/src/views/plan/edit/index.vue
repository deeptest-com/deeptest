<template>
  <DrawerLayout :visible="editDrawerVisible" @close="onCancel" :stickyKey="stickyKey">
    <!-- 头部信息  -->
    <template #header>
      <div class="header-text">
        <!-- ::::todo 没有序列号，后端字段 -->
<!--        <span class="serialNumber">[{{ planDetail?.serialNumber }}]</span>-->
        <EditAndShowField :value="(currPlan && currPlan.name) || '暂无'" placeholder="输入计划名称" @update="handleUpdateName" />
      </div>
    </template>
    <!-- 基本信息 -->
    <template #basicInfo>
      <a-descriptions :title="null" size="small" :column="4">
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
    </template>
    <template #tabHeader>
      <div class="tab-header-items">
        <div class="tab-header-item"
             :class="{'active':tab.key === activeKey}" v-for="tab in tabsList"
             :key="tab.key"
             @click="changeTab(tab.key)">
          <span>{{ tab.label }}</span>
        </div>
      </div>
      <div class="tab-header-btns">
        <a-button class="plan-exec" type="primary" @click="handleEnvSelect">执行计划</a-button>
      </div>
    </template>
    <template #tabContent>
      <div class="tab-pane">
        <div style="padding-top: 20px" v-if="activeKey === 'test-scenario'">
          <ScenarioList
              :list="planScenarioList"
              :show-scenario-operation="true"
              :columns="columns"
              :scroll="{ x: 1240 }"
              :loading="loading"
              :pagination="scenarioPagination"
              @refresh-list="getScenarioList" />
        </div>
        <div style="padding-top: 20px"  v-if="activeKey === 'test-report'">
          <ReportList :show-report-list="activeKey === 'test-report'" />
        </div>
      </div>
    </template>
  </DrawerLayout>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, watch, reactive, computed } from 'vue';
import { useStore } from 'vuex';

import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import DrawerLayout from "@/views/component/DrawerLayout/index.vue";

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
        width: 300,
        slots: { customRender: 'name' }
    },
    {
        title: '状态',
        dataIndex: 'status',
        width: 80,
        slots: { customRender: 'status' }
    },
    {
        title: '优先级',
        width: 90,
        dataIndex: 'priority',
    },
    {
        title: '所属分类',
        width: 110,
        dataIndex: 'categoryName',
        ellipsis: true
    },
    {
        title: '创建人',
        width: 110,
        dataIndex: 'createUserName',
    },
    {
        title: '最近更新',
        dataIndex: 'updatedAt',
        width: 180,
        slots: { customRender: 'updateAt' }
    },
    {
        title: '操作',
        dataIndex: 'operation',
        width: 80,
        fixed: 'right',
        slots: { customRender: 'operation' },
    },
]);

const stickyKey = ref(0);

const tabsList = [
  {
    "key": "test-report",
    "label": "测试报告"
  },
  {
    "key": "test-scenario",
    "label": "测试场景"
  },
]
async function changeTab(value) {
  activeKey.value = value;
  stickyKey.value ++;
  emits('update:tabKey', value);
}

function onCancel() {
    emits('onCancel');
}

function handleEnvSelect() {
    emits('onSelectEnv',planDetail.value);
}

function handleChangeStatus(value) {
    console.log('changeStatus --', value);
    emits('onUpdate', { status: value });
}

function handleUpdateName(value) {
    emits('onUpdate', { name: value });
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

</style>
