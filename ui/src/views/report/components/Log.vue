<template>
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }" :visible="drawerVisible"
        @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>测试报告详情</div>
            </div>
        </template>
        <div class="drawer-content">
            <div class="report-basicinfo">
                <TextItem label="测试计划" :value="reportData.logName" />
                <TextItem label="执行时间" :value="reportData.logTime" />
                <TextItem label="执行环境" :value="reportData.logEnv" />
                <TextItem label="执行人" :value="reportData.logExecutor" />
                <a-button type="primary" class="report-export">导出报告</a-button>
            </div>
            <div class="report-statistical-table">
                <div class="statistical-main" ref="main"></div>
                <div class="statistical-info">
                    <div class="statistical-info-item">
                        <span class="label success">通过</span>
                        <span class="value">60% &nbsp; &nbsp; 8</span>
                    </div>
                    <div class="statistical-info-item">
                        <span class="label">总耗时</span>
                        <span class="value">
                            <span style="color: #04C495">4.9 &nbsp;</span>秒
                        </span>
                    </div>
                    <div class="statistical-info-item">
                        <span class="label failed">失败</span>
                        <span class="value">60% &nbsp; &nbsp; 8</span>
                    </div>
                    <div class="statistical-info-item">
                        <span class="label">平均接口请求耗时</span>
                        <span class="value"><span style="color: #04C495">4.9 &nbsp;</span>秒</span>
                    </div>
                    <div class="statistical-info-item">
                        <span class="label notest">未测</span>
                        <span class="value">60% &nbsp; &nbsp; 8</span>
                    </div>
                    <div class="statistical-info-item">
                        <span class="label">测试场景(成功/失败)</span>
                        <span class="value">(0/2)</span>
                    </div>
                </div>
            </div>
            <div class="report-list">
                <div class="report-item" v-for="logItem in reportData.logList" :key="logItem.id">
                    <div class="report-item-basic">
                        <div class="report-item-info report-item-name">{{ logItem.scenarioName }}</div>
                        <div class="report-item-info report-item-priority">{{ logItem.scenarioPriority }}</div>
                        <div class="report-item-info report-item-status">{{ logItem.scenarioStatus === 0 ? '已完成' : '未完成' }}
                        </div>
                        <div class="report-item-info report-item-rate">
                            <div class="report-progress"
                                :style="`background: linear-gradient(90deg, #04C495 ${logItem.scenarioProgress}%, #FF6963 0);`">
                            </div>
                            通过率 {{ logItem.scenarioProgress }}%
                        </div>
                        <div class="report-item-info report-item-operation" @click="handleExpand(logItem.id)">
                            {{ expandKey.includes(logItem.id + '') ? '收起' : '展开' }}
                            <template v-if="expandKey.includes(logItem.id + '')">
                                <up-outlined />
                            </template>
                            <template v-else>
                                <DownOutlined />
                            </template>

                        </div>
                    </div>
                    <div :class="['report-item-detail-list', expandKey.includes(logItem.id + '') ? 'active' : '']">
                        <a-table :showHeader="false" :columns="columns" :data-source="logItem.reponseList"
                            :pagination="false" :rowKey="(_record, index) => index">
                            <template #expandIcon="props">
                                <template v-if="props.expanded">
                                    <DownOutlined @click="$event => props.onExpand(props.record, $event)" />
                                </template>
                                <template v-else>
                                    <RightOutlined @click="$event => props.onExpand(props.record, $event)" />
                                </template>
                            </template>
                            <template #requestStatus="{ record }">
                                <div :class="['report-status', ClassMap[record.requestStatus]]">{{
                                    StatusMap[record.requestStatus] }}</div>
                            </template>
                            <template #requestMethod="{ record }">
                                <div :class="['report-method', ClassMap[record.requestStatus]]">
                                    {{ record.requestMethod }}
                                </div>
                            </template>
                            <template #requestCode="{ record }">
                                <div :class="['report-code', ClassMap[record.requestStatus]]">
                                    状态码: <span>{{ record.requestCode }}</span>
                                </div>
                            </template>
                            <template #requestTime="{ record }">
                                <div :class="['report-time', ClassMap[record.requestStatus]]">
                                    耗时: <span>{{ record.requestTime }}</span>
                                </div>
                            </template>
                            <template #requestType>
                                <div class="report-type">
                                    转单
                                </div>
                            </template>
                            <template #expandedRowRender="{ record }">
                                <div class="expand-wrapper">
                                    <div class="expand-content">
                                        <div class="expand-detail-info" v-for="info in record.requestInfo"
                                            :key="info.errorId">
                                            <div class="info-field"><exclamation-circle-outlined
                                                    style="color: #F63838;margin-right: 8px" />{{ info.errorField }}</div>
                                            <div class="info-tip">
                                                <span v-for="(tip, index) in info.errorTip" :key="index">{{ index + 1 }}. {{
                                                    tip
                                                }}</span>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </template>
                        </a-table>
                    </div>
                </div>
            </div>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, onMounted, watch } from 'vue';
import { DownOutlined, RightOutlined, UpOutlined, ExclamationCircleOutlined } from '@ant-design/icons-vue';
import TextItem from './TextItem.vue';
import * as echarts from 'echarts';

enum StatusMap {
    'success' = '通过',
    'expires' = '过期',
    'error' = '失败'
}

enum ClassMap {
    'success' = 'report-success',
    'expires' = 'report-expires',
    'error' = 'report-error'
}

defineProps<{
    drawerVisible: boolean
}>()

const emits = defineEmits(['onClose']);
const expandKey = ref<string[]>([]);
const main = ref();

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
                    requestStatus: 'success',
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

const columns = [
    { title: '请求状态', dataIndex: 'requestStatus', key: 'requestStatus', slots: { customRender: 'requestStatus' } },
    { title: '请求方法', dataIndex: 'requestMethod', key: 'requestMethod', slots: { customRender: 'requestMethod' } },
    { title: '请求url', dataIndex: 'requestUrl', key: 'requestUrl' },
    { title: '请求信息', dataIndex: 'requestData', key: 'requestData' },
    { title: '请求状态码', dataIndex: 'requestCode', key: 'requestCode', slots: { customRender: 'requestCode' } },
    { title: '请求耗时', dataIndex: 'requestTime', key: 'requestTime', slots: { customRender: 'requestTime' } },
    { title: '请求类型', dataIndex: 'requestMethod', key: 'requestMethod', slots: { customRender: 'requestType' } },
];

onMounted(() => {
    init();
})

function init() {
    if (!main.value) {
        return;
    }
    const myChart = echarts.init(main.value);
    const option: any = {
        color: ['#04C495', '#F63838', 'rgba(0, 0, 0, 0.28)'],
        series: [
            {
                name: 'Access From',
                type: 'pie',
                radius: ['40%', '70%'],
                avoidLabelOverlap: false,
                label: {
                    show: false,
                    position: 'center'
                },
                emphasis: {
                    label: {
                        show: true,
                        formatter: (params: any) => {
                            return [`{subTitle|${params.data.name}}`, `{title|${params.data.value}}`].join('\n')
                        },
                        rich: {
                            subTitle: {
                                fontSize: 12,
                                lineHeight: 18,
                                marginBottom: 10,
                                color: 'rgba(0, 0, 0, 0.85)'
                            },
                            title: {
                                fontSize: 24,
                                lineHeight: 29,
                                color: 'rgba(0, 0, 0, 0.85)'
                            }
                        }
                    },
                },
                labelLine: {
                    show: false
                },
                data: [
                    { value: 244, name: '已完成' },
                    { value: 100, name: '失败' },
                    { value: 20, name: '未测试' },
                ]
            }
        ]
    };

    option && myChart.setOption(option);
}

function onClose() {
    emits('onClose');
}

function handleExpand(id: any) {
    let lastExpandKeys = JSON.parse(JSON.stringify([...expandKey.value]));
    if (lastExpandKeys.includes(id + '')) {
        lastExpandKeys = lastExpandKeys.filter((item: string) => item !== id + '');
    } else {
        lastExpandKeys.push(id + '');
    }
    expandKey.value = lastExpandKeys;
}

watch(() => main.value, (val) => {
    if (val) {
        init();
    }
}, {
    immediate: true
})

</script>
<style scoped lang="less">
.report-drawer {
    :deep(.ant-drawer-header) {
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }
}

.drawer-content {
    .report-basicinfo {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        position: relative;
        margin-bottom: 24px;

        .text-wrapper {
            width: 40%;
            margin-bottom: 4px;
        }

        .report-export {
            position: absolute;
            right: 0;
            top: 0;
        }
    }

    .report-statistical-table {
        height: 240px;
        background: #FFFFFF;
        border: 1px solid #E5E5E5;
        margin-bottom: 24px;
        display: flex;
        align-items: center;
        padding: 0 118px 0 80px;
        box-sizing: border-box;

        .statistical-main {
            width: 214px;
            height: 214px;
            margin-right: 60px;
        }

        .statistical-info {
            display: flex;
            flex-wrap: wrap;
            flex: 1;

            .statistical-info-item {
                width: 50%;
                margin-bottom: 8px;
            }
        }
    }
}

.statistical-info-item {
    display: flex;
    align-items: center;
    &:nth-child(n) {
        .label {
            width: 48px;
            height: 22px;
            line-height: 22px;
            margin-right: 22px;
            display: flex;
            align-items: center;

            &:before {
                content: '';
                display: block;
                width: 6px;
                height: 6px;
                margin-right: 7px;
                border-radius: 50%;
            }

            &.success:before {
                background-color: #04C495;
            }

            &.failed:before {
                background-color: #F63838;
            }

            &.notest:before {
                background-color: rgba(0, 0, 0, 0.28);;
            }
        }
    }

    &:nth-child(2n) {
        display: flex;
        align-items: center;
        justify-content: space-between;

        .label {
            width: 147px;
        }
    }
}

.report-item {
    margin-bottom: 10px;

    .report-item-basic {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        box-sizing: border-box;
        padding: 14px 16px;
        height: 50px;
        background: #FFFFFF;
        border: 1px solid #E5E5E5;

        .report-item-info {
            padding-right: 50px;
            line-height: 22px;

            &:last-child {
                padding: 0;
            }
        }

        .report-item-priority {
            font-weight: bold;
        }

        .report-item-status {
            display: flex;
            align-items: center;

            &:before {
                content: '';
                display: block;
                width: 6px;
                height: 6px;
                border-radius: 50%;
                background-color: #04C495;
                margin-right: 10px;
            }
        }

        .report-item-name {
            width: 333px;
            font-weight: bold;
        }

        .report-item-rate {
            width: 292px;
            display: flex;
            align-items: center;
            font-family: 'PingFang SC';
            font-style: normal;
            font-weight: 400;
            font-size: 14px;
            color: rgba(0, 0, 0, 0.85);
            padding: 0;

            .report-progress {
                width: 180px;
                height: 6px;
                border-radius: 41px;
                margin-right: 16px;
            }
        }
    }
}

.report-item-detail-list {
    display: none;
    transition: all 0.5s ease-in-out;

    &.active {
        display: block;
    }

    :deep(.ant-table-tbody > tr:hover:not(.ant-table-expanded-row):not(.ant-table-row-selected) > td) {
        background-color: transparent;
    }

    :deep(.ant-table-tbody > tr) {
        padding: 0px 16px;
        height: 40px;
        box-sizing: border-box;
        background: #FBFBFB;
        /* 下划线 */
        box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
    }

    :deep(.ant-table-tbody .ant-table-expanded-row.ant-table-expanded-row-level-1 > td) {
        padding: 0;
    }

    .report-status {
        width: 36px;
        height: 20px;
        font-size: 12px;
        border-radius: 2px;
        text-align: center;
        line-height: 20px;

        &.report-success {
            background: #E6FFF4;
            color: #04C495;

        }

        &.report-error {
            background: #FFF2F0;
            ;
            color: #F63838;
        }

        &.report-expires {
            background: #FFF2F0;
            ;
            color: #F63838;
        }
    }

    .report-method {
        font-weight: bold;
        font-size: 14px;

        &.report-success {
            color: #04C495;
        }

        &.report-error,
        &.report-expires {
            color: #F63838;
        }
    }

    .report-time,
    .report-code {
        font-size: 14px;
        line-height: 22px;
        color: rgba(0, 0, 0, 0.85);

        span {
            color: #04C495;
        }
    }

    .report-type {
        font-size: 14px;
        line-height: 22px;
        color: #447DFD;
    }

    .expand-wrapper {
        padding: 9px 16px;
        column-span: 8;

        .expand-content {
            background-color: #fff;
            padding: 16px;

            .expand-detail-info {
                margin-bottom: 10px;

                &:last-child {
                    margin: 0;
                }
            }

            .info-field {
                font-weight: 400;
                font-size: 14px;
                line-height: 22px;
                color: rgba(0, 0, 0, 0.85);
                margin-bottom: 8px;
            }

            .info-tip {
                font-style: normal;
                font-weight: 400;
                font-size: 14px;
                line-height: 22px;
                color: rgba(0, 0, 0, 0.65);

                span {
                    display: block;
                }
            }
        }
    }

}
</style>