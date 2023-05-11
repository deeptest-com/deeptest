<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>{{ title || '测试报告详情' }}</div>
            </div>
        </template>
        <div class="drawer-content">
            <ReportBasicInfo :basic-info="{ logEnv: reportData.logEnv, logName: reportData.logName, logTime: reportData.logTime, logExecutor: reportData.logExecutor }" />
            <StatisticTable />
            <template v-for="logItem in reportData.logList" :key="logItem.id">
                <ScenarioCollapsePanel :show-scenario-info="showScenarioInfo" :expand-active="expandActive" :record="logItem">
                    <template #endpointData>
                        <EndpointCollapsePanel :recordList="logItem.reponseList" />
                    </template>
                </ScenarioCollapsePanel>
            </template>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref } from 'vue';
import ReportBasicInfo from './Components/BasicInfo.vue';
import StatisticTable from './Components/StatisticTable.vue';
import ScenarioCollapsePanel from './Components/ScenarioCollapsePanel.vue';
import EndpointCollapsePanel from './Components/EndpointCollapsePanel.vue';

const props = defineProps<{
    drawerVisible: boolean
    title: string
    scenarioExpandActive: boolean 
    showScenarioInfo: boolean
}>()

const emits = defineEmits(['onClose']);
const expandActive = ref(props.scenarioExpandActive || false);

const reportData = {
    logName: '测试计划名称11',
    logExecutor: '测试执行人',
    logEnv: '测试环境',
    logTime: '2023/04/24 10:20:20',
    logList: [
        {
            id: 1,
            scenarioName: '测试场景1',
            scenarioPriority: 'P1',
            scenarioStatus: 0,
            scenarioProgress: '66.66',
            reponseList: [
                {
                    requestId: 44444,
                    requestStatus: 'loading',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 122,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 55555,
                    requestStatus: 'success',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 112,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 66666,
                    requestStatus: 'error',
                    requestMethod: 'POST',
                    requestCode: '400',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 172,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                }
            ]
        },
        {
            id: 2,
            scenarioName: '测试场景2',
            scenarioPriority: 'P3',
            scenarioStatus: 0,
            scenarioProgress: '58',
            reponseList: [
                {
                    requestId: 11111,
                    requestStatus: 'success',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 142,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 2222,
                    requestStatus: 'error',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 145,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 33333,
                    requestStatus: 'expires',
                    requestMethod: 'POST',
                    requestCode: '400',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 146,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                }
            ]
        }
    ]
};

function onClose() {
    emits('onClose');
}

</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}
</style>