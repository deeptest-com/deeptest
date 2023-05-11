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
                  v-if="type === 0"
                  title="入住项目(个)"
                  :value="card?.projectTotal"
                />

                <a-statistic
                  v-else
                  title="项目成员(位)"
                  :value="card?.userTotal"
                />
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="总接口数（个）"
                  :value="card?.interfaceHb"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.interfaceHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.interfaceTotal }}
                    </div>
                      <arrow-up-outlined   v-if="card?.interfaceHb >= 0"/>
                      <arrow-down-outlined v-else />
                  </template>
                </a-statistic>
                 <a-statistic
                  title="总测试场景数（个）"
                  :value="card?.scenarioHb"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.scenarioHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.scenarioTotal }}
                    </div>
                      <arrow-up-outlined   v-if="card?.scenarioHb >= 0"/>
                      <arrow-down-outlined v-else />
                  </template>
                </a-statistic>
           
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                  <a-statistic
                  title="接口测试总体覆盖率（%）"
                  :value="card?.coverageHb"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.coverageHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.coverage+"%" }}
                    </div>
                      <arrow-up-outlined   v-if="card?.coverageHb >= 0"/>
                      <arrow-down-outlined v-else />
                  </template>
                </a-statistic>
         
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="执行总次数"
                  :value="card?.execTotal"
                 
                />
                   <a-statistic
                  title="测试通过率（%）"
                  :value="card?.passRate+'%'"
                 
                />
               
              </div>
            </a-card>
          </div>
        </a-card>
      </a-col>
      <a-col class="pie" style="margin-left: 14px" flex="0 1 400px">
        <a-card title="发现缺陷分布">
          <Pie :params="pieData"/>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, watch, ref, defineProps } from "vue";
import { ArrowUpOutlined, ArrowDownOutlined } from "@ant-design/icons-vue";
import Pie from "@/components/Pie/index.vue";
import * as echarts from "echarts";
import { log } from "handsontable/helpers";
import { useStore } from "vuex";
import { StateType } from "@/views/home/store";
const store = useStore<{ Home: StateType }>();
const pieData = computed<any>(() => store.state.Home.pieData);
// 组件接收参数
const props = defineProps({
  // 请求API的参数
  params: { type: Object },
  type: { type: Number },
});
// console.log("staticstic params", props.params);
const card = ref<any>({});

// 监听项目数据变化
watch(
  () => {
    return props.params;
  },
  async (newVal: any) => {
    console.log("watch staticstic newVal", newVal);
    card.value = newVal.cardData;
  },
  {
    immediate: true,
  }
);
</script>
<style lang="less" scoped>
.statistics {
  // padding:  16px;
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
 

