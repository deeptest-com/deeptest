<template>
    <a-table
        :rowKey="(_record, index) => index"
        :columns="columns"
        :data-source="list"
        :loading="loading"
        :pagination="{
            ...pagination,
            onChange: (page) => {
                handleGetList({ page });
            },
            onShowSizeChange: (page, size) => {
                pagination.page = page
                pagination.pageSize = size
                handleGetList(pagination);
            },
        }"
        :scroll="{ x: 1240 }"
        class="dp-table">
        <template #serialNumber="{ record }">
            <span>{{ record.serialNumber }}</span>
        </template>
        <template #interfacePassRate="{ record }">
            <span>{{ record.interfacePassRate }}</span>
        </template>
        <template #createUserName="{ record }">
            <span>{{ record.createUserName }}</span>
        </template>
        <template #execPlan="{ record, column }">
            <ToolTipCell style="color: #447DFD;cursor: pointer;" :width="column.width"  @click="handleQueryDetail(record)" :text="record.name" />
        </template>
        <template #duration="{ record }">
            <span v-html="formatWithSeconds(record.duration)"></span>
        </template>
        <template #executionTime="{ record, column }">
            <ToolTipCell :width="column.width" :text="momentUtc(record.startTime)" />
        </template>

        <template #action="{ record }">
            <a-dropdown>
                <MoreOutlined />
                <template #overlay>
                    <a-menu>
<!--                        <a-menu-item key="1">-->
<!--                            <a class="operation-a" href="javascript:void (0)" @click="handleExport(record.id)">导出</a>-->
<!--                        </a-menu-item>-->
                        <a-menu-item key="2">
                            <a class="operation-a" href="javascript:void (0)" @click="handleQueryDetail(record)">查看报告</a>
                        </a-menu-item>
                        <a-menu-item key="3">
                            <a class="operation-a" href="javascript:void (0)" @click="handleDelete(record.id)">删除</a>
                        </a-menu-item>
                    </a-menu>
                </template>
            </a-dropdown>
        </template>
    </a-table>
</template>
<script setup lang="ts">
import {computed, ref, defineEmits, defineProps, createVNode} from "vue";
import { useStore } from "vuex";
import { ColumnProps } from 'ant-design-vue/es/table/interface';
import {message, Modal} from "ant-design-vue";
import {ExclamationCircleOutlined, MoreOutlined} from "@ant-design/icons-vue";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType } from "../store";
import { PaginationConfig } from "../data";
import { momentUtc, formatWithSeconds } from "@/utils/datetime";
import ToolTipCell from '@/components/Table/tooltipCell.vue';


defineProps({
    loading: {
        required: false,
        default: false,
        type: Boolean
    },
    list: {
        required: true,
        default: [],
    }
})
const emits = defineEmits(['queryDetail', 'getList']);

const store = useStore<{ Report: StateType, ProjectGlobal: ProjectStateType }>();
// 分页数据
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);
// 表格选中项
const selectedRowKeys = ref<Key[]>([]);

type Key = ColumnProps['key'];


const columns = [
    {
        title: '编号',
        dataIndex: 'serialNumber',
        slots: { customRender: 'serialNumber' },
        width: 120
    },
    {
        title: '测试计划',
        dataIndex: 'execPlan',
        width: 300,
        slots: { customRender: 'execPlan' },
    },
    {
        title: '测试通过率',
        dataIndex: 'interfacePassRate',
        width: 110,
    },
    {
        title: '执行人',
        dataIndex: 'createUserName',
        width: 110,
    },
    {
        title: '执行耗时',
        width: 120,
        dataIndex: 'duration',
        slots: { customRender: 'duration' },
    },
    {
        title: '执行时间',
        dataIndex: 'executionTime',
        width: 180,
        slots: { customRender: 'executionTime' },
    },
    {
        title: '操作',
        key: 'action',
        width: 80,
        fixed: 'right',
        slots: { customRender: 'action' },
    },
];

const onSelectChange = (keys: Key[], rows: any) => {
    selectedRowKeys.value = [...keys];
}


const handleExport = (id: number) => {
    console.log('export');
}

const handleDelete = async (id: number) => {
  Modal.confirm({
    title: () => '确定删除该报告吗？',
    icon: createVNode(ExclamationCircleOutlined),
    okText: () => '确定',
    okType: 'danger',
    cancelText: () => '取消',
    onOk: async () => {
      const res = store.dispatch('Report/remove', id);
      if (res) {
        message.success('删除成功');
      } else {
        message.error('删除失败');
      }
    },
  });
}

const handleQueryDetail = (record: any) => {
    emits('queryDetail', record);
}

const handleGetList = (params: any) => {
    emits('getList', params);
}

</script>
<style scoped lang="less">

.report-planname {
    cursor: pointer;
    color: #447DFD;
}
</style>
