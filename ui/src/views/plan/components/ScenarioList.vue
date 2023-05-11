<template>
    <div class="table-filter">
        <div class="left" v-if="showScenarioOperation">
            <a-button type="primary" @click="associateModalVisible = true">
                <template #icon><plus-outlined /></template>
                关联测试场景
            </a-button>
            <a-button type="default" @click="handleRemove">批量移除</a-button>
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
        :loading="loading"
        :columns="columns" 
        :data-source="list">
        <template #status="{ record }">
            {{ record.status }}
        </template>
        <template #operation="{ record }">
            <a-button type="primary" @click="handleRemove(record)"> 
                移除
            </a-button>
        </template>
    </a-table>
    <Associate 
        :associate-modal-visible="associateModalVisible" 
        @on-cancel="associateModalVisible = false" 
        @on-ok="handleFinish"
    />
</template>
<script lang="ts" setup>
import { ref, reactive, defineProps, defineEmits, PropType, computed } from 'vue';
import { useStore } from 'vuex';
import { PlusOutlined } from '@ant-design/icons-vue';
import Associate from './Associate.vue';

import { StateType as PlanStateType } from '../store';  
import { Modal } from 'ant-design-vue';

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
    }
})

const emits = defineEmits(['selectRowKeys', 'refreshList']);
const store = useStore<{ Plan: PlanStateType }>();
const planId = computed(() => store.state.Plan.planId);
const members = computed(() => store.state.Plan.members);
const associateModalVisible = ref(false);
const selectedRowKeys = ref<any[]>([]); // Check here to configure the default column
let selectedRowIds = reactive<number[]>([]);

const onSelectChange = (changableRowKeys: string[], rows: any) => {
    selectedRowKeys.value = changableRowKeys;
    selectedRowIds = rows.map(e => {
        return e.id;
    });
    emits('selectRowKeys', selectedRowIds);
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
            await store.dispatch('Plan/removeScenario', { planId: planId.value, params });
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
  
  