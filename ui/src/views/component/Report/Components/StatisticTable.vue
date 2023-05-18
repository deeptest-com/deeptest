<template>
    <div class="report-statistical-table">
        <template v-if="loading">
            <div class="statistical-loading">
                <div class="spinner"></div>
            </div>
        </template>
        <template v-else>
            <div class="statistical-main" ref="main"></div>
        </template>
        <div class="statistical-info">
            <TextItem class="statistical-info-item" label-class-name="success" label="通过"
                :value="`${statiscalResult.passRate || '0.00%'} &nbsp; &nbsp;${statiscalResult.passNum || 0}`">
            </TextItem>
            <TextItem class="statistical-info-item" label="总耗时" label-style="width: 140px">
                <template #value>
                    <span style="color: #04C495" v-html="formatWithSeconds(statiscalResult.duration)"></span>
                </template>
            </TextItem>
            <TextItem class="statistical-info-item" label-class-name="failed" label="失败"
                :value="`${statiscalResult.failRate || '0.00%'} &nbsp; &nbsp;${statiscalResult.failNum || 0}`">
            </TextItem>
            <TextItem class="statistical-info-item" label="平均接口请求耗时" label-style="width: 140px">
                <template #value>
                    <span class="value"><span style="color: #04C495"
                            v-html="formatWithSeconds(statiscalResult.averageDuration)">
                        </span></span>
                </template>
            </TextItem>
            <TextItem class="statistical-info-item" label-class-name="notest" label="未测"
                :value="`${statiscalResult.notestRate || '0.00%'} &nbsp; &nbsp;${statiscalResult.notestNum || 0}`">
            </TextItem>
            <TextItem class="statistical-info-item" :label="labelMap"
                :value="`${statiscalResult.passNum || 0}/${statiscalResult.failNum || 0}`" label-style="width: 147px">
            </TextItem>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, onMounted, watch, defineProps } from 'vue';
import * as echarts from 'echarts';
import TextItem from './TextItem.vue';

import { ReportDetailType } from '@/utils/enum';
import { percentDef, formatWithSeconds } from '@/utils/datetime';

const props = defineProps<{
    scene: string
    data: any
}>();

const main = ref();
const myChart = ref<any>(null);
const statiscalResult = ref<any>({});
const loading = ref(false);
const labelMap = props.scene !== ReportDetailType.ExecScenario ? '测试场景(成功/失败)' : '断言数(成功/失败)';
const initOptions = ref<any>({
    tooltip: {
        trigger: 'item',
        formatter: (params) => {
            return `${params.data.name}: ${params.data.value}`
        }
    },
    color: ['#04C495', '#F63838', 'rgba(0, 0, 0, 0.28)'],
    series: [
        {
            name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            label: {
                position: 'center',
                show: true,
                color: '#1E7CE8',
                lineHeight: 16,
                fontSize: 12,
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
            labelLine: {
                show: false
            },
            data: [
                { value: 0, name: '已完成' },
                { value: 0, name: '失败' },
                { value: 0, name: '未测试' },
            ]
        }
    ]
});
onMounted(() => {
    init();
})

const transformWithUndefined = (num: number | undefined) => {
    return num || 0;
}

function init() {
    if (!main.value) {
        return;
    }
    myChart.value = echarts.init(main.value);
}

function initData(data: any) {
    if (Object.keys(data).length === 0) {
        loading.value = true;
        return;
    }
    const isExecScenario = props.scene === ReportDetailType.ExecScenario;
    const statiscalData: any = !isExecScenario ? {
        totalNum: transformWithUndefined(data.totalScenarioNum), //场景总数
        passNum: transformWithUndefined(data.passScenarioNum), //通过场景数
        failNum: transformWithUndefined(data.failScenarioNum), //失败场景数
        notestNum: transformWithUndefined(data.totalScenarioNum) - transformWithUndefined(data.passScenarioNum) - transformWithUndefined(data.failScenarioNum)
    } : {
        totalNum: transformWithUndefined(data.totalInterfaceNum), //接口总数
        passNum: transformWithUndefined(data.passInterfaceNum),
        failNum: transformWithUndefined(data.failInterfaceNum),
        notestNum: transformWithUndefined(data.totalInterfaceNum) - transformWithUndefined(data.passInterfaceNum) - transformWithUndefined(data.failInterfaceNum)
    };
    statiscalData.passRate = percentDef(statiscalData.passNum, statiscalData.totalNum);
    statiscalData.failRate = percentDef(statiscalData.failNum, statiscalData.totalNum);
    statiscalData.notestRate = percentDef(statiscalData.notestNum, statiscalData.totalNum);
    statiscalData.duration = data.duration;
    statiscalData.averageDuration = data.totalRequestNum ? data.duration / data.totalRequestNum : 0;
    const chartData = [{
        value: statiscalData.passNum,
        name: '通过'
    }, {
        value: statiscalData.failNum,
        name: '失败'
    }, {
        value: statiscalData.totalNum - statiscalData.passNum - statiscalData.failNum,
        name: '未测'
    }];
    initOptions.value.series[0].data = chartData;
    initOptions.value.series[0].label = {
        ...initOptions.value.series[0].label,
        formatter: () => {
            return [`{subTitle|通过}`, `{title|${statiscalData.passNum}}`].join('\n')
        },
    }
    statiscalResult.value = { ...statiscalData };
    if (!myChart.value) {
        myChart.value = echarts.init(main.value);
    }
    setTimeout(() => {
        loading.value = false;
        myChart.value.setOption({ ...initOptions.value });
    }, 500)
}


watch(() => main.value, (val) => {
    if (val) {
        init();
    }
}, {
    immediate: true
});

watch(() => {
    return [props.data, myChart.value];
}, val => {
    const [statiscalOriginalData, chartRef] = val;
    if (val[0] && chartRef) {
        initData(statiscalOriginalData);
    }
}, {
    immediate: true
})
</script>
<style scoped lang="less">
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

    .statistical-loading {
        width: 214px;
        height: 214px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-right: 60px;

        .spinner {
            animation: spinnerFour 1s linear infinite;
            border: solid 7px transparent;
            border-top: solid 7px #447DFD;
            border-radius: 100%;
            width: 100px;
            height: 100px;
        }
    }

    .statistical-info {
        display: flex;
        flex-wrap: wrap;
        flex: 1;

        .statistical-info-item {
            width: 50%;
            margin-bottom: 8px;
            display: flex;
            align-items: center;
        }
    }
}

@keyframes spinnerFour {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}
</style>