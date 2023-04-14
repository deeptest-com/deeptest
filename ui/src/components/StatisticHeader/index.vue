<template>
  <div class="statistics">
    <a-row type="flex">
      <a-col flex="1 1 200px">
        <a-card title="统计数据">
          <div
            style="
              display: flex;
              justify-content: space-between;
              flex-wrap: wrap; /* 只要您把这个属性去掉,就不会自动换行了*/
            "
          >
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="项目成员(位)"
                  :value="11.28"
                  :precision="2"
                  suffix="%"
                  :value-style="{ color: '#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">13</div>
                    <arrow-up-outlined />
                  </template>
                </a-statistic>
                <a-statistic
                  title="项目成员(位)"
                  :value="11.28"
                  :precision="2"
                  suffix="%"
                  :value-style="{ color: '#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">1213</div>
                    <arrow-up-outlined />
                  </template>
                </a-statistic>
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="总接口数"
                  :value="9.3"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: '#cf1322', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">1203</div>
                    <arrow-down-outlined />
                  </template>
                </a-statistic>
                <a-statistic
                  title="总测试场景数（个）"
                  :value="9.3"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: '#cf1322', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">233</div>
                    <arrow-down-outlined />
                  </template>
                </a-statistic>
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="接口测试总体覆盖率（%）"
                  :value="11.28"
                  :precision="2"
                  suffix="%"
                  :value-style="{ color: '#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">80.8%</div>
                    <arrow-up-outlined />
                  </template>
                </a-statistic>
                <a-statistic
                  title="项目成员(位)"
                  :value="11.28"
                  :precision="2"
                  suffix="%"
                  :value-style="{ color: '#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">43</div>
                    <arrow-up-outlined />
                  </template>
                </a-statistic>
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="执行总次数"
                  :value="9.3"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: '#cf1322', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">9998</div>
                    <arrow-down-outlined />
                  </template>
                </a-statistic>
                <a-statistic
                  title="测试通过率"
                  :value="9.3"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: '#cf1322', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">237</div>
                    <arrow-down-outlined />
                  </template>
                </a-statistic>
              </div>
            </a-card>
          </div>
        </a-card>
      </a-col>
      <a-col class="pie" style="margin-left: 14px" flex="0 1 400px">
        <a-card title="发现缺陷分布">
          <Pie />
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, defineComponent } from "vue";
import { ArrowUpOutlined, ArrowDownOutlined } from "@ant-design/icons-vue";
import Pie from "@/components/Pie/index.vue";
import * as echarts from "echarts";

</script>
<style lang="less" scoped>
.statistics {
  padding:  16px;
  .card-content {
    display: flex;
    justify-content: space-between;
    &-num {
      font-weight: 400;
      font-size: 24px;

      color: rgba(0, 0, 0, 0.85);
    }
  }

  .pie {
    background: #fff;
    :deep(.ant-card) {
      height: 100%;
    }
    :deep(.ant-card-body) {
      height: 38vh !important;
      padding: 0 !important;
    }
  }
}
</style>
 

