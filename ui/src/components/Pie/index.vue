<template>
  <div ref="main" style="width: 100%; height: 100%"></div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";

import * as echarts from "echarts";
export default defineComponent({
  name: "Pie",
  props: {
    value: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    const main = ref(); // 使用ref创建虚拟DOM引用，使用时用main.value
    onMounted(() => {
      init();
    });
    function init() {
      // 基于准备好的dom，初始化echarts实例
      const myChart = echarts.init(main.value);
      const schoolData = [
        {
          name: "轻微",
          value: 4253,
        },
        {
          name: "致命",
          value: 5691,
        },
        {
          name: "阻塞",
          value: 4536,
        },
        {
          name: "严重",
          value: 4369,
        },
        {
          name: "一般",
          value: 5124,
        },
      ];
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

          subtext: "226", // 副标题
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
            // radius: "55%",
            radius: ["40%", "70%"],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: "#fff",
              borderWidth: 2,
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
          },
        ],
        color: [
          "#447DFD",
          "#5344FD",
          "#68D079",
          "#E76D46",
          "#FBC434",
          "#75bedc",
        ],
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
 

