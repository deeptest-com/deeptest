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
      <a-descriptions :column="2">
        <a-descriptions-item :key="index" v-for="(item,index) in data" >
          <template #label>
            <span :class="['text-label', item.class]" >{{ item.label }}</span>
          </template>
          <span class="text-rate" v-if="item.rate">{{item.rate}}</span>
          <span class="text-value">{{item.value}}</span>
        </a-descriptions-item>
      </a-descriptions>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ref, onMounted, watch, defineProps} from 'vue';
import * as echarts from 'echarts';

const props = defineProps<{
  value: Object,
  data: any[],
}>();

const mainRef = ref();
let myChart: any = null;
const loading = ref(false);

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
        {value: 0, name: '通过接口'},
        {value: 0, name: '失败接口'},
        {value: 0, name: '未测接口'},
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
      option.value.series[0].data[0].value = newVal.interfacePass;
      option.value.series[0].data[1].value = newVal.interfaceFail;
      option.value.series[0].data[2].value = newVal.interfaceSkip;
      setChart();
    }, 500);
  }
}, {
  immediate: true
});

</script>
<style scoped lang="less">
.report-statistical-table {
  height: 200px;
  background: #FFFFFF;
  border: 1px solid #E5E5E5;
  margin-bottom: 24px;
  display: flex;
  align-items: center;
  padding: 0 16px;
  box-sizing: border-box;
  .statistical-main {
    width: 200px;
    height: 200px;
    margin-right: 32px;
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


.text-label {
  color: rgba(0, 0, 0, 0.85);
  margin-left: 8px;
  &:before {
    content: '';
    display: inline-block;
    margin-right: 7px;
    width: 6px;
    height: 6px;
    position: relative;
    top: -2px;
    border-radius: 50%;
  }

  &.success:before {
    background-color: #04C495;
  }

  &.fail:before {
    width: 6px;
    height: 6px;
    background-color: #F63838;
  }

  &.notest:before {
    width: 6px;
    height: 6px;
    background-color: rgba(0, 0, 0, 0.28);
  }
}

.text-value{
  display: inline-block;
  margin-left: 16px;
  width: 96px;
}
.text-rate{
  display: inline-block;
  width: 96px;
  margin-left: 16px;
}

</style>
