<template>
    <div class="report-table-filter">
        <div class="bulk-operation" v-if="showOperation">
            <a-tooltip placement="bottomLeft" color="#fff">
                <template #title>
                    <div class="actions">
                        <span class="action-item">导出</span>
                        <span class="action-item">删除</span>
                    </div>
                </template>
                <a-button type="primary">批量操作</a-button>
            </a-tooltip>
        </div>
        <a-form :model="formState" class="report-form">
            <div class="report-executor-selector">
                <a-form-item name="executor" label="执行人">
                    <a-select allowClear placeholder="请选择执行人" v-model:value="selectValue" :options="executorOptions"
                        style="width: 140px" @change="handleSelectChange" />
                </a-form-item>
            </div>
            <div class="report-excutime">
                <a-form-item name="executime" label="执行时间">
                    <a-range-picker :show-time="{ format: 'HH:mm' }" format="YYYY-MM-DD HH:mm"
                        :placeholder="['开始时间', '结束时间']" @ok="onRangeOk" @change="onRangeChange" />
                </a-form-item>
            </div>
            <div class="report-name">
                <a-input-search
                    v-model:value="executeName"
                    placeholder="请输入用例名称"
                    enter-button
                    @search="onSearch"
                />
            </div>
        </a-form>
    </div>
</template>
<script setup lang="ts">
import { ref, reactive, defineEmits, defineProps, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { StateType } from "../store";
import { StateType as ProjectStateType } from "@/store/project";
import { momentTimeStamp } from '@/utils/datetime';

defineProps({
    showOperation: {
        type: Boolean,
        default: true,
        required: false
    }
})

const store = useStore<{ Report: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const executorOptions = computed<any[]>(() => store.state.Report.members);

const selectValue = ref(null);
const executeName = ref('');
const formState = reactive({});
let executeTime = reactive<any>({ executeStartTime: '', executeEndTime: '' });

const emits = defineEmits(['handleFilter']);

const getMember = async (): Promise<void> => {
  await store.dispatch('Report/getMembers', currProject.value.id)
}

function onRangeOk(date: any) {
    console.log('current select executime ---', date);
    executeTime = { executeStartTime: momentTimeStamp(date[0]), executeEndTime: momentTimeStamp(date[1]) };
    refreshList({});
}

function onRangeChange(date: any) {
    if (date.length === 0) {
        executeTime = { executeStartTime: '', executeEndTime: '' };
        refreshList({});
    }
}

function handleSelectChange(value: any) {
    console.log('handleSelect---', value);
    refreshList({});
}

function refreshList(params) {
    emits('handleFilter', { keywords: executeName.value, createUserId: selectValue.value, ...executeTime, ...params })
}

const onSearch = (val: string) => {
  refreshList({});
};

watch(() => {
    return currProject.value;
}, (val: any) => {
    if (val.id) {
        getMember();
    }
}, {
    immediate: true
})
</script>
<style scoped lang="less">
.report-table-filter {
    width: 100%;
    display: flex;
    flex-direction: column;
    margin-bottom: 24px;

    .bulk-operation {
        margin-bottom: 24px;
    }

    .report-form {
        display: flex;
        align-items: center;
        flex: 1;
    }

    :deep(.ant-form-item) {
        margin: 0;
        margin-right: 20px;
    }
}

.actions {

    .action-item {
        display: block;
        width: 162px;
        height: 32px;
        background: #FFFFFF;
        font-weight: 400;
        font-size: 14px;
        line-height: 32px;
        text-align: center;
        color: rgba(0, 0, 0, 0.85);
        cursor: pointer;

        &:hover {
            background: #F5F5F5;
        }
    }

}
</style>
