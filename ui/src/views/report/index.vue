<template>
    <div class="report-container">
      <div class="report-table-filter">
          <TableFilter :formState="formState" @handle-filter="handleFilter" />
      </div>
      <div class="report-list">
          <List :loading="loading" :list="list" @get-list="getList" @query-detail="queryDetail"/>
      </div>
      <DetailDrawer
        :title="'测试报告详情'"
        :show-scenario-info="true"
        :scenario-expand-active="false"
        :drawer-visible="drawerVisible"
        :report-id="currPlanId"
        :scene="ReportDetailType.QueryDetail"
        @on-close="drawerVisible = false" />
    </div>
</template>
<script setup lang="ts">
import { computed, ref, watch, reactive } from "vue";
import { useStore } from "vuex";

import { TableFilter } from '@/views/component/Report/components';
import List from './List/index.vue';
import DetailDrawer from './Detail/index.vue';

import { StateType as ProjectStateType } from "@/store/project";
import { StateType } from "./store";
import { PaginationConfig } from "./data";
import { ReportDetailType } from "@/utils/enum";

const store = useStore<{ Report: StateType, ProjectGlobal: ProjectStateType }>();
// 全局选中的项目
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
// 报告列表信息
const list = computed<any>(() => store.state.Report.listResult.list);
// 分页数据
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);
// 初始查询参数
const queryParams = reactive({
  createUserId: null,
  executeStartTime: '',
  executeEndTime: '',
  keywords: '',
  page: 1,
  pageSize: pagination.value.pageSize || 20,
});

const loading = ref<boolean>(false);
const drawerVisible = ref<boolean>(false);
let formState = reactive({});
const currPlanId = ref(0);

const handleFilter = (params: any) => {
  formState = params;
  getList({});
}

const getList = async (params?: any): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/list', {
    ...queryParams,
    ...formState,
    ...params
  });
  loading.value = false;
};

const queryDetail = async (record: any) => {
    // console.log('查看报告详情：===', record);
    await store.dispatch('Report/initReportDetail');
    await store.dispatch('Report/get', record.id);
    currPlanId.value = record.id;
    drawerVisible.value = true;
};

watch(currProject, (val) => {
  if (val.id) {
    getList({ page: 1 });
  }
}, { immediate: true })
</script>
<style lang="less" scoped>
.report-container {
    margin: 16px;
    background: #ffffff;
    min-height: calc(100vh - 92px);
    overflow: hidden;
    padding: 20px;
    box-sizing: border-box;
}
</style>
