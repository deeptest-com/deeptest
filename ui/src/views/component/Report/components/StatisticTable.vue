<template>
  <div class="report-statistical-table">
    <template v-if="loading">
      <div class="statistical-loading">
        <div class="spinner"></div>
      </div>
    </template>
    <template v-else>
      <div class="statistical-main" ref="mainRef"></div>
    </template>
    <div class="statistical-info">
      <TextItem v-for="item in data"
                :key="item.label"
                class="statistical-info-item"
                :label-class-name="item.class === 'fail' ? 'failed' : 'success'"
                :label="item.label"
                :value="item.value"/>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ref, onMounted, watch, defineProps} from 'vue';
import * as echarts from 'echarts';
import TextItem from './TextItem.vue';

const props = defineProps<{
  value: string
  data: any[],
}>();

const mainRef = ref();
let myChart: any = null;
const loading = ref(false);

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
        {value: 0, name: '已完成'},
        {value: 0, name: '失败'},
        {value: 0, name: '未测试'},
      ]
    }
  ]
});

const option = ref({
  tooltip: {
    trigger: 'item',
    formatter: (params) => {
      return `${params.data.name}: ${params.data.value}`
    }
  },
  series: [
    {
      name: '执行详情',
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 2
      },
      label: {
        show: false,
        position: 'center'
      },
      color: ['#04C495', '#F63838', 'rgba(0, 0, 0, 0.28)'],
      emphasis: {
        label: {
          show: true,
          fontSize: 12,
          fontWeight: 'bold',
        }
      },
      labelLine: {
        show: false
      },
      data: [
        {value: 0, name: '通过'},
        {value: 0, name: '失败'},
        {value: 0, name: '未测'},
      ]
    }
  ]
})



function setChart() {
  if (!myChart) {
    myChart = echarts.init(mainRef.value);
  }
  myChart.setOption(option.value);
}

watch(() => {
  return props.value
}, (newVal:any) => {
  if(newVal) {
    setTimeout(() => {
      loading.value = false;
      option.value.series[0].data[0].value = newVal.passAssertionNum;
      option.value.series[0].data[1].value = newVal.failAssertionNum;
      option.value.series[0].data[2].value = newVal.notTestNum;
      setChart();
    }, 500);
  }
}, {
  immediate: true
});

</script>
<style scoped lang="less">
.report-statistical-table {
  height: 240px;
  background: #FFFFFF;
  border: 1px solid #E5E5E5;
  margin-bottom: 24px;
  display: flex;
  align-items: center;
  padding: 0 24px;
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
