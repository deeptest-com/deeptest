<template>
  <div class="report-main">
    <TableFilter :executor-options="[]" />

    <div>
      <a-table row-key="id" :columns="columns" :data-source="list" :loading="loading" :pagination="{
        ...pagination,
        onChange: (page) => {
          getList(page);
        },
        onShowSizeChange: (page, size) => {
          pagination.pageSize = size
          getList(page);
        },
      }" :row-selection="{
  selectedRowKeys: selectedRowKeys,
  onChange: onSelectChange
}" class="dp-table">
        <template #executor="{ record }">
          <span>{{ record.executor }}</span>
        </template>
        <template #execPlan="{ record }">
          <span class="report-planname">{{ record.name }}</span>
        </template>
        <template #executiveTime="{ record }">
          <span>{{ momentUtc(record.executiveTime) }}</span>
        </template>
        <template #executionTime="{ record }">
          <span>{{ record.executionTime }}</span>
        </template>

        <template #action="{ record }">
          <a-dropdown>
            <MoreOutlined />
            <template #overlay>
              <a-menu>
                <a-menu-item key="1">
                  <a class="operation-a" href="javascript:void (0)" @click="handleExport(record.id)">导出</a>
                </a-menu-item>
                <a-menu-item key="2">
                  <a class="operation-a" href="javascript:void (0)" @click="handleDelete(record.id)">删除</a>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </template>

      </a-table>
    </div>

    <a-drawer :visible="reportDetailVisible">
      <ReportDetail />
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useStore } from "vuex";

import debounce from "lodash.debounce";
import { ColumnProps } from 'ant-design-vue/es/table/interface';
import { momentUtc } from "@/utils/datetime";
import { Modal, notification } from "ant-design-vue";
import TableFilter from "../components/tableFilter.vue";
import ReportDetail from "../detail/index.vue";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ScenarioStateType } from "@/views/scenario/store";
import { StateType } from "@/views/report/store";
import { PaginationConfig, QueryParams, Report } from "@/views/report/data";
import { query } from "@/views/scenario/service";
import { NotificationKeyCommon } from "@/utils/const";

const store = useStore<{ Report: StateType, Scenario: ScenarioStateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const list = computed<Report[]>(() => store.state.Report.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);

const selectedRowKeys = ref<Key[]>([]);
let queryParams = reactive<QueryParams>({
  keywords: '', scenarioId: '0',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});
const reportDetailVisible = ref(false);

type Key = ColumnProps['key'];

watch(currProject, () => {
  getList(1);
}, { deep: false })

const columns = [
  {
    title: '编号',
    dataIndex: 'number',
    width: 80
  },
  {
    title: '测试通过率',
    dataIndex: 'execRate',
    width: 120,
  },
  {
    title: '执行人',
    dataIndex: 'executor',
    width: 80,
  },
  {
    title: '所属测试计划',
    dataIndex: 'execPlan',
    width: 200,
    slots: { customRender: 'execPlan' },
  },
  {
    title: '执行耗时',
    dataIndex: 'executiveTime',
    width: 200,
    slots: { customRender: 'executiveTime' },
  },
  {
    title: '执行时间',
    dataIndex: 'executionTime',
    width: 200,
    slots: { customRender: 'executionTime' },
  },
  {
    title: '操作',
    key: 'action',
    width: 80,
    slots: { customRender: 'action' },
  },
];

onMounted(() => {
  getList(1);
})

const scenarios = ref([] as any[])
query().then(json => {
  scenarios.value = json.data.result
})

const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/list', {
    keywords: queryParams.keywords,
    scenarioId: queryParams.scenarioId,
    pageSize: pagination.value.pageSize,
    page: current,
    order: 'desc',
  });
  loading.value = false;
}

const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
}


const handleExport = (id: number) => {
  console.log('export')
}

const handleDelete = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除报告',
    content: '确定删除指定的报告？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Report/remove', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          notification.success({
            key: NotificationKeyCommon,
            message: `删除成功`,
          });
        } else {
          notification.error({
            key: NotificationKeyCommon,
            message: `删除失败`,
          });
        }
      })
    }
  });
}

const onSearch = debounce(() => {
  getList(1)
}, 500);

</script>

<style lang="less" scoped>
.report-main {
  .scenario-select {}
}
</style>