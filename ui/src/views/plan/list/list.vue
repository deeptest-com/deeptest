<template>
  <div class="plan-list-main">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => create()">新建</a-button>
      </template>
      <template #extra>
        <a-select 
          allowClear
          @change="onSearch" 
          v-model:value="queryParams.status" 
          :options="planStatusOptions" 
          class="status-select" 
          style="width: 120px"
          placeholder="请选择状态">
        </a-select>
        <a-input-search 
          allowClear
          @change="onSearch" 
          @search="onSearch" 
          v-model:value="queryParams.keywords" 
          placeholder="输入关键字搜索"
          style="width:270px;margin-left: 16px;" />
      </template>

      <div>
        <a-table row-key="id" :columns="columns" :data-source="list" :loading="loading"
          :pagination="{
            ...pagination,
            onChange: (page) => {
              pagination.current = page;
              getList(page);
            },
            onShowSizeChange: (page, size) => {
              pagination.pageSize = size
              getList(page);
            },
          }" class="dp-table">

          <template #name="{ text, record }">
            <div class="plan-name">
              <EditAndShowField 
                :custom-class="'custom-endpoint show-on-hover'" 
                :value="text" 
                placeholder="请输入计划名称" 
                @edit="edit(record)"
                @update="(e: string) => updatePlan(e, record)" />
            </div>
          </template>
          <template #status="{ record }">
            <a-tag v-if="record.status" :color="planStatusColorMap.get(record.status)">{{ planStatusTextMap.get(record.status) }}</a-tag>
          </template>
          <template #updatedAt="{ record }">
            <span>{{ momentUtc(record.updatedAt) }}</span>
          </template>

          <template #action="{ record }">
            <a-dropdown>
              <MoreOutlined />
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="exec(record)">执行</a>
                  </a-menu-item>
                  <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="report(record.id)">测试报告</a>
                  </a-menu-item>
                  <!-- <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="exec(record.id)">执行</a>
                  </a-menu-item> -->
                  <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="clone(record.id)">克隆</a>
                  </a-menu-item>
                  <a-menu-item key="2">
                    <a class="operation-a" href="javascript:void (0)" @click="remove(record.id)">删除</a>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>

        </a-table>
      </div>
    </a-card>
  </div>

  <!-- 新建计划弹窗 -->
  <PlanCreate 
    :create-drawer-visible="createDrawerVisible"
    @on-cancel="createDrawerVisible = false" 
    @get-list="getList(1)"
  />
  <!-- 编辑计划抽屉 -->
  <PlanEdit 
    :tab-active-key="editTabActiveKey" 
    :edit-drawer-visible="editDrawerVisible" 
    @onExec="handleExec"
    @onUpdate="handleUpdate"
    @on-cancel="editDrawerVisible = false" />

  <!-- 执行计划抽屉 -->
  <ReportDetail 
    :drawer-visible="execReportVisible" 
    :title="execReportTitle" 
    :scenario-expand-active="true" 
    :show-scenario-info="true"
    :scene="ReportDetailType.ExecPlan"
    @on-close="execReportVisible = false"
  />
</template>

<script setup lang="ts">
import { computed, reactive, ref, watch } from "vue";
import { useStore } from "vuex";
import { message } from 'ant-design-vue';
import { MoreOutlined } from "@ant-design/icons-vue";
import { Modal } from "ant-design-vue";
import debounce from "lodash.debounce";

import PlanCreate from "../components/PlanCreate.vue";
import PlanEdit from "../edit/index.vue";
import EditAndShowField from "@/components/EditAndShow/index.vue";
import ReportDetail from "@/views/component/Report/Detail/Index.vue";

import { StateType as ProjectStateType } from "@/store/project";
import { PaginationConfig, Plan } from '../data.d';
import { StateType } from "../store";
import { momentUtc } from "@/utils/datetime";
import { planStatusColorMap, planStatusTextMap, planStatusOptions } from "@/config/constant";
import { ReportDetailType } from "@/utils/enum";
import settings from "@/config/settings";
import bus from "@/utils/eventBus";
import { useExec } from "../hooks/exec";

const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
  },
  {
    title: '摘要',
    dataIndex: 'name',
    slots: { customRender: 'name' },
    ellips: true
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: { customRender: 'status' },
  },
  {
    title: '测试通过率',
    dataIndex: 'testPassRate',
  },
  {
    title: '负责人',
    dataIndex: 'adminName',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
    slots: { customRender: 'updatedAt' },
  },
  {
    title: '操作',
    key: 'action',
    width: 80,
    slots: { customRender: 'action' },
  },
];

const store = useStore<{ Plan: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const nodeDataCategory = computed<any>(() => store.state.Plan.nodeDataCategory);
const currPlan = computed<any>(() => store.state.Plan.currPlan);

const list = computed<Plan[]>(() => store.state.Plan.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Plan.listResult.pagination);

const queryParams = reactive<any>({
  keywords: '', 
  status: null
});
const loading = ref<boolean>(false);
const createDrawerVisible = ref(false);
const editDrawerVisible = ref(false);
const editTabActiveKey = ref('test-scenario');
const execReportVisible = ref(false);
const execReportTitle = ref('');
const { execCancel, execStart, onWebSocketConnStatusMsg, OnWebSocketMsg } = useExec();

const getList = debounce(async (current: number): Promise<void> => {
  loading.value = true;
  await store.dispatch('Plan/listPlan', {
    ...queryParams,
    categoryId: nodeDataCategory.value.id,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false
}, 300);

const handleExec = () => {
  // editDrawerVisible.value = false;
  execReportTitle.value = currPlan.value && currPlan.value.name;
  execReportVisible.value = true;
};

const exec = async (record: any) => {
  await getCurrentPalnInfo(record);
  execReportTitle.value = record.name;
  execReportVisible.value = true;
};

const report = async (record: any) => {
  console.log('获取报告列表');
  editTabActiveKey.value = 'test-report';
  editDrawerVisible.value = true;
  getCurrentPalnInfo(record);
};

const clone = async (id: number) => {
  await store.dispatch('Plan/clonePlan', id);
};

const updatePlan = async (value: string, record: any) => {
  try {
    const { id, adminId, categoryId, testStage, desc, status } = record;
    await store.dispatch('Plan/savePlan', {
      id,
      adminId,
      categoryId,
      testStage,
      desc,
      status,
      name: value,
    });
  } catch(err) {
    console.log(err);
  }
};

const handleUpdate = async (params: any) => {
  try {
    const result = await store.dispatch('Plan/savePlan', {
      ...currPlan.value,
      ...params
    });
    if (result) {
      store.dispatch('Plan/getPlan', currPlan.value.id);
    } else {
      message.error('更新计划失败');
    }
  } catch(err) {
    console.log(err);
  }
}

const create = () => {
  console.log('create')
  createDrawerVisible.value = true;
};

const edit = async (record: any) => {
  editDrawerVisible.value = true;
  editTabActiveKey.value = 'test-scenario';
  getCurrentPalnInfo(record);
};

const getCurrentPalnInfo = async (record: any) => {
  const { id, adminId, categoryId, testStage, desc, status, name } = record;
  try {
    await store.dispatch('Plan/setCurrentPlan', { id, adminId, categoryId, testStage, desc, status, name });
    await store.dispatch('Plan/getPlan', record.id);
  } catch(err) {
    message.error('获取计划信息出错');
  }
};

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除计划',
    content: '确定删除指定的计划？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Plan/removePlan', id);
    }
  });
}

const onSearch = () => {
  getList(1);
};

watch(() => {
  return nodeDataCategory.value.id;
}, async (val) => {
  await getList(1);
}, { immediate: true, deep: true });

watch(() => {
  return currProject.value;
}, async (val) => {
  if (val.id) {
    await getList(1);
  }
}, { immediate: true });

watch(() => {
  return execReportVisible.value;
}, (val) => {
  if (val) {
    execStart();
    bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
    bus.on(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
  } else {
    execCancel();
    bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
    bus.off(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
  }
});


</script>

<style lang="less" scoped>
.operation-a {
  text-align: center;
  display: inline-block;
  width: 80px;
}

@media screen and (max-width: 1540px) {
  .plan-name {
    width: 180px;
  }
}
</style>
