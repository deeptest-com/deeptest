<template>
    <div class="table-filter">
        <div class="left" v-if="showScenarioOperation">
            <a-button type="primary" @click="associateModalVisible = true">
                <template #icon><plus-outlined /></template>
                关联测试场景
            </a-button>
            <a-button type="default" :disabled="!selectedRowIds?.length" @click="handleRemove()">批量移除</a-button>
        </div>
        <div class="right">
            <a-form-item label="优先级">
                <a-select allowClear ref="select" v-model:value="formState.priority" style="width: 140px" :options="priorityOptions"
                    @change="handleChange" placeholder="请选择优先级"></a-select>
            </a-form-item>
            <a-form-item label="创建人">
                <a-select allowClear ref="select" v-model:value="formState.createUserId" style="width: 140px" :options="members"
                    @change="handleChange" placeholder="请选择创建人"></a-select>
            </a-form-item>
            <a-form-item>
                <a-input-search allowClear v-model:value="formState.keywords" placeholder="请输入需要搜索的用例名称" @search="handleChange" style="width: 220px" />
            </a-form-item>
        </div>
    </div>
    <a-table
        :row-selection="{
            selectedRowKeys: selectedRowKeys,
            onChange: onSelectChange
        }"
        :pagination="{
            ...pagination,
            showSizeChanger: false,
            onChange: (page) => {
                getList({ page });
            },
        }"
        row-key="id"
        :scroll="scroll"
        :loading="loading"
        :columns="columns"
        :data-source="list">
        <template #name="{ record, column }">
            <ToolTipCell :text="record.name" :width="column.width" />
        </template>
        <template #status="{ record }">
            <a-tag v-if="record.status" :color="planStatusColorMap.get(record.status)">{{ planStatusTextMap.get(record.status) }}</a-tag>
        </template>
        <template #updateAt="{ record, column }">
            <ToolTipCell :text="momentUtc(record.updateAt)" :width="column.width" />
        </template>
        <template #operation="{ record }">
            <a-button type="primary" @click="handleRemove(record)">
                移除
            </a-button>
        </template>
    </a-table>
    <RelationScenario
        :associate-modal-visible="associateModalVisible"
        @on-cancel="associateModalVisible = false"
        @on-ok="handleFinish"
    />
</template>
<script lang="ts" setup>
import { ref, reactive, defineProps, defineEmits, PropType, computed } from 'vue';
import { useStore } from 'vuex';
import { PlusOutlined } from '@ant-design/icons-vue';
import RelationScenario from './RelationScenario.vue';
import ToolTipCell from '@/components/Table/tooltipCell.vue';

import { StateType as PlanStateType } from '../store';
import { message, Modal } from 'ant-design-vue';
import { planStatusColorMap, planStatusTextMap } from '@/config/constant';
import { momentUtc } from '@/utils/datetime';

const props = defineProps({
    showScenarioOperation: {
        type: Boolean,
        default: true,
        required: false
    },
    list: {
        type: Array as PropType<any[]>,
        required: false
    },
    columns: {
        type: Array as PropType<any[]>
    },
    loading: {
        type: Boolean
    },
    pagination: {
        type: Object
    },
    planId: {
        type: Number,
        required: false
    },
    scroll: {
        type: Object,
        required: false,
    }
})

const emits = defineEmits(['selectRowKeys', 'refreshList']);
const store = useStore<{ Plan: PlanStateType }>();
const currPlan = computed<any>(() => store.state.Plan.currPlan);
const members = computed(() => store.state.Plan.members);
const associateModalVisible = ref(false);
const selectedRowKeys = ref<any[]>([]); // Check here to configure the default column
let selectedRowIds = reactive<any[]>([]);

const onSelectChange = (changableRowKeys: string[], rows: any) => {
    selectedRowKeys.value = changableRowKeys;
    selectedRowIds = rows.map((item: any) => item.id);
    emits('selectRowKeys', changableRowKeys);
};

const priorityOptions = ref<any>([
    {
        label: 'P0',
        value: 'P0'
    },
    {
        label: 'P1',
        value: 'P1'
    },
    {
        label: 'P2',
        value: 'P2'
    },
    {
        label: 'P3',
        value: 'P3'
    }
]);

const formState = reactive({ priority: null, createUserId: null, keywords: '' });

const getList = (params) => {
    console.log('changePage');
    emits('refreshList', { ...params, formState });
}

const handleChange = (value: string) => {
    console.log(`selected ${value}`);
    emits('refreshList', formState);
};

const handleRemove = async (record?: any) => {
    if (!record && selectedRowIds.length === 0) {
        message.warning('请先选择要删除的关联场景');
        return;
    }
    Modal.confirm({
        title: '确认要解除该测试场景的关联吗?',
        onOk: async () => {
            let scenarioIds: any[] = [];
            if (record && record.id) {
                scenarioIds.push(record.id);
            } else {
                scenarioIds = selectedRowIds;
            }
            const params = { scenarioIds };
            console.log('解除关联场景: --', params);
            await store.dispatch('Plan/removeScenario', { planId: currPlan.value.id, params });
            emits('refreshList', formState);
        }
    })
}

const handleFinish = async () => {
    associateModalVisible.value = false;
    emits('refreshList', formState);
}
</script>
<style scoped lang="less">
.table-filter {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-bottom: 20px;

    .left, .right {
        display: flex;
        align-items: center;

        :deep(.ant-row.ant-form-item), :deep(.ant-btn) {
            margin-right: 20px;
            margin-bottom: 0;

            &:last-child {
                margin: 0;
            }
        }
    }
}
</style>

