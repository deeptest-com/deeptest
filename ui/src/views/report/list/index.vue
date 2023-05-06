<template>
  <div class="report-main">
    <TableFilter :executor-options="members || []" @get-list="getList" />

    <div>
      <a-table 
        :rowKey="(_record, index) => index" 
        :columns="columns" 
        :data-source="list" 
        :loading="loading" 
        :pagination="{
          ...pagination,
          onChange: (page) => {
            getList({ page });
          },
          onShowSizeChange: (page, size) => {
            pagination.pageSize = size
            getList({ page });
          },
        }" 
        :row-selection="{
          selectedRowKeys: selectedRowKeys,
          onChange: onSelectChange
        }" 
        class="dp-table">
        <template #serialNumber="{ record }">
          <span style="cursor: pointer">{{ record.serialNumber }}</span>
        </template>
        <template #interfacePassRate="{ record }">
          <span>{{ record.interfacePassRate }}</span>
        </template>
        <template #createUserName="{ record }">
          <span>{{ record.createUserName }}</span>
        </template>
        <template #execPlan="{ record }">
          <span class="report-planname">{{ record.name }}</span>
        </template>
        <template #duration="{ record }">
          <span>{{ record.duration * 1000 }}ms</span>
        </template>
        <template #executionTime="{ record }">
          <span>{{ momentUtc(record.startTime) }}</span>
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

    <LogDetail :drawer-visible="reportDetailVisible" @on-close="reportDetailVisible = false" />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useStore } from "vuex";
import debounce from "lodash.debounce";
import { ColumnProps } from 'ant-design-vue/es/table/interface';
import { Modal, notification } from "ant-design-vue";
import { MoreOutlined } from "@ant-design/icons-vue";
import TableFilter from "../components/TableFilter.vue";
import LogDetail from "../components/Log.vue";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType } from "@/views/report/store";
import { PaginationConfig, Report } from "@/views/report/data";
import { NotificationKeyCommon } from "@/utils/const";
import { momentUtc } from "@/utils/datetime";

const store = useStore<{ Report: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const list = computed<Report[]>(() => store.state.Report.listResult.list);
const members = computed<Report[]>(() => store.state.Report.members);
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);

const selectedRowKeys = ref<Key[]>([]);
const reportDetailVisible = ref(false);
const queryParams = {
  currProjectId: 0,
  executeStartTime: '',
  executeEndTime: '',
  keywords: '',
  page: 1,
  pageSize: pagination.value.pageSize || 20,
};

type Key = ColumnProps['key'];


const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
    slots: { customRender: 'serialNumber' },
    width: 80
  },
  {
    title: '测试通过率',
    dataIndex: 'interfacePassRate',
    width: 120,
  },
  {
    title: '执行人',
    dataIndex: 'createUserName',
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
    dataIndex: 'duration',
    width: 200,
    slots: { customRender: 'duration' },
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

const loading = ref<boolean>(false);
const getListInfo = () => {
  getMember();
  getList({ page: 1 });
}
const getList = async (params: any): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/list', {
    ...queryParams,
    ...params
  });
  loading.value = false;
}

const getMember = async (): Promise<void> => {
  await store.dispatch('Report/getMembers', currProject.value.id)
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
  // getList(1)
}, 500);

watch(currProject, (val) => {
  console.log(val);
  if (val.id) {
    getListInfo();
  }
}, { immediate: true })

</script>

<style lang="less" scoped>
.report-main {
  .scenario-select {}
}
</style>