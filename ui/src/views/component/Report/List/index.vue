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
                pagination.pageSize = size
                handleGetList({ page });
            },
        }" 
        :row-selection="{
            selectedRowKeys: selectedRowKeys,
            onChange: onSelectChange
        }" class="dp-table">
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
            <span class="report-planname" @click="handleQueryDetail(record)">{{ record.name }}</span>
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
</template>
<script setup lang="ts">
import { computed, ref, defineEmits, defineProps } from "vue";
import { useStore } from "vuex";
import { ColumnProps } from 'ant-design-vue/es/table/interface';
import { MoreOutlined } from "@ant-design/icons-vue";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType } from "../store";
import { PaginationConfig } from "../data";
import { momentUtc } from "@/utils/datetime";

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

const onSelectChange = (keys: Key[], rows: any) => {
    selectedRowKeys.value = [...keys];
}


const handleExport = (id: number) => {
    console.log('export')
}

const handleDelete = (id: number) => {
    console.log('remove')
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
}
</style>