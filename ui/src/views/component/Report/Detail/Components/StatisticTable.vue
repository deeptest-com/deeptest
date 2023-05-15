<template>
    <div class="report-statistical-table">
        <div class="statistical-main" ref="main"></div>
        <div class="statistical-info">
            <TextItem class="statistical-info-item" label-class-name="success" label="通过" :value="'60% &nbsp; &nbsp; 90'">
            </TextItem>
            <TextItem class="statistical-info-item" label="总耗时" label-style="width: 147px">
                <template #value>
                    <span style="color: #04C495">4.9 &nbsp;</span>秒
                </template>
            </TextItem>
            <TextItem class="statistical-info-item" label-class-name="failed" label="失败" :value="'60% &nbsp; &nbsp; 8'">
            </TextItem>
            <TextItem class="statistical-info-item" label="平均接口请求耗时" label-style="width: 147px">
                <template #value>
                    <span class="value"><span style="color: #04C495">4.9 &nbsp;</span>秒</span>
                </template>
            </TextItem>
            <TextItem class="statistical-info-item" label-class-name="notest" label="未测" :value="'60% &nbsp; &nbsp; 8'">
            </TextItem>
            <TextItem class="statistical-info-item" label="测试场景(成功/失败)" value="(0/2)" label-style="width: 147px"></TextItem>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, onMounted, watch, defineProps } from 'vue';
import * as echarts from 'echarts';
import TextItem from './TextItem.vue';

defineProps<{
    scene: string
}>();

const main = ref();
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


watch(() => main.value, (val) => {
    if (val) {
        init();
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
</style>