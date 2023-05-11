<template>
  <div  ref="main" style="width: 100%; height: 100%">
   
  </div>
  

</template>

<script lang="ts">
import { defineComponent, ref, onMounted, watch } from "vue";

import * as echarts from "echarts";
export default defineComponent({
  name: "Pie",
  props: {
    params: { type: Object },
  },
  setup(props) {
    const main = ref(); // 使用ref创建虚拟DOM引用，使用时用main.value
    let pieData = ref<any>({});
    onMounted(() => {
      // init();
    });
    // 监听项目数据变化
    watch(
      () => {
        return props.params;
      },
      async (newVal: any) => {
        console.log("watch pie newVal", newVal);
        pieData.value = newVal.pieData;
        // if (pieData.value?.total!=0) {
          init();
        // }
      },
      {
        immediate: false,
      }
    );
    function init() {
      // 基于准备好的dom，初始化echarts实例
      const myChart = echarts.init(main.value);
      let schoolData = [
        {
          name: "轻微",
          value: pieData.value?.minor || 0,
          color: "#447DFD",
        },
        {
          name: "致命",
          value: pieData.value?.deadly || 0,
          color: "#5344FD",
        },
        {
          name: "阻塞",
          value: pieData.value?.blocker || 0,
          color: "#26D1A1",
        },
        {
          name: "严重",
          value: pieData.value?.critical || 0,
          color: "#FF6963",
        },
        {
          name: "一般",
          value: pieData.value?.major || 0,
          color: "#FBC434",
        },
      ];
      schoolData = schoolData.filter((item) => {
        if (item.value != 0) {
          return item;
        }
      });
      // 指定图表的配置项和数据
      const option: any = {
        title: {
          text: "总计（个）", // 主标题
          x: "center",
          y: "42%",
          textStyle: {
            // 主标题样式
            fontSize: "12",
            color: "rgba(0, 0, 0, 0.65)",
          },

          subtext: pieData.value?.total || '0', // 副标题
          subtextStyle: {
            // 副标题样式
            color: "rgba(0, 0, 0, 0.85)",
            fontSize: 24,
            fontWeight: 500,
          },
        },
        tooltip: {
          trigger: "item",
          formatter: "<br/>{b} : {c} ({d}%)",
        },
        legend: {
          orient: "horizontal",
          itemGap: 0,
          itemHeight: 6,
          data: schoolData.map((a) => a.name),
          y: "bottom",
          x: "center",
          icon: "circle",
          textStyle: {
            color: "rgba(0, 0, 0, 0.46)",
            fontSize: "20px",
            fontWeight: 700,
          },
        },
        series: [
          {
            type: "pie",
            name: "缺陷分布",
            // radius: "55%",
            radius: ["40%", "70%"],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 20,
              borderColor: "#fff",
              borderWidth: schoolData.length > 1 ? 2 : 0,
            },
            emphasis: {
              label: {
                show: true,
                fontSize: 16,
                fontWeight: "bold",
              },
            },
            center: ["50%", "50%"],
            data: schoolData,
            labelLine: { show: false },
            label: {
              show: true,
              formatter: "{b} \n ({d}%)",
              // color: "#B1B9D3",
            },
            animationType: "scale",
            animationEasing: "exponentialInOut",
            animationDelay: function () {
              return Math.random() * 100;
            },
          },
        ],
        color: schoolData.map((item) => {
          return item.color;
        }),
  
      };
      // 赋值
      //   option.series = [
      //     {
      //       type: "pie",
      //       radius: "55%",
      //       center: ["50%", "30%"],
      //       // data: res.data.map((v) => {
      //       //   return { name: v.name, value: v.value }
      //       // })
      //       data: schoolData,

      //     },
      //   ];
      // 赋值
      // option.legend = [
      //   {
      //    data: schoolData.map((a) => a.name)
      //   }
      // ]
      //   // 赋值
      //   option.legend.data = schoolData.map((a) => a.name);

      // 使用刚指定的配置项和数据显示图表。
      myChart.setOption(option);
    }
    return {
      main,
      init,
    };
  },
});
</script>
 

