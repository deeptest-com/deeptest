<template>
    <div class="report-table-filter">
        <div class="bulk-operation">
            <a-tooltip placement="right">
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
                    <a-select placeholder="请选择执行人" v-model:value="selectValue" :options="executorOptions"
                        style="width: 140px" />
                </a-form-item>
            </div>
            <div class="report-excutime">
                <a-form-item name="executime" label="执行时间">
                    <a-range-picker :show-time="{ format: 'HH:mm' }" format="YYYY-MM-DD HH:mm"
                        :placeholder="['Start Time', 'End Time']" @ok="onRangeOk" />
                </a-form-item>
            </div>
            <div class="report-name">
                <a-input v-model:value="executeName" placeholder="请输入你需要搜索的用例名称">
                    <template #suffix>
                        <search-outlined />
                    </template>
                </a-input>
            </div>
        </a-form>
    </div>
</template>
<script setup lang="ts">
import { ref, defineProps, reactive } from 'vue';
import { SearchOutlined } from '@ant-design/icons-vue';

const selectValue = ref(null);
const executeName = ref('');
const formState = reactive({});
defineProps<{
    executorOptions: any[]
}>()

function onRangeOk(date: string) {
    console.log('current select executime ---', date);
}
</script>
<style scoped lang="less">
.report-table-filter {
    width: 100%;
    display: flex;
    align-items: center;
    margin-bottom: 24px;

    .bulk-operation {
        margin-right: 20px;


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

.action-item {
    display: block;
    width: 88px;
    height: 32px;
    text-align: center;
    line-height: 32px;
}
</style>