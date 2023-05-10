<template>
    <div class="report-container">
        <div class="report-table-filter">
            <TableFilter @get-list="getList" />
        </div>
        <div class="report-list">
            <List :loading="loading" :list="list" @get-list="getList" @query-detail="queryDetail"/>
        </div>

        <DetailDrawer :title="'测试报告详情'" :show-scenario-info="true" :scenario-expand-active="false" :drawer-visible="drawerVisible" @on-close="drawerVisible = false" />
    </div>
</template>
<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { useStore } from "vuex";

import TableFilter from '@/views/component/Report/List/TableFilter.vue';
import List from '@/views/component/Report/List/index.vue';
import DetailDrawer from '@/views/component/Report/Detail/Index.vue';

import { StateType as ProjectStateType } from "@/store/project";
import { StateType } from "@/views/component/Report/store";
import { PaginationConfig, Report } from "@/views/component/Report/data";

const store = useStore<{ Report: StateType, ProjectGlobal: ProjectStateType }>();
// 全局选中的项目
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
// 报告列表信息
const list = computed<any>(() => store.state.Report.listResult.list);
// 分页数据
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);
// 初始查询参数
const queryParams = {
  currProjectId: 0,
  executeStartTime: '',
  executeEndTime: '',
  keywords: '',
  page: 1,
  pageSize: pagination.value.pageSize || 20,
};

const loading = ref<boolean>(false);
const drawerVisible = ref<boolean>(false);

const getList = async (params: any): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/list', {
    ...queryParams,
    ...params
  });
  loading.value = false;
};

const queryDetail = (record: any) => {
    console.log('查看报告详情：===', record);
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