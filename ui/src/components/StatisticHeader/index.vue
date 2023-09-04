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
                  title="入住项目（个）"
                  :value="card?.projectTotal||0"

                />
                <a-statistic
                  v-else
                  title="项目成员（位）"
                  :value="card?.userTotal||0"
                />
              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="总接口数（个）"
                  :value="card?.interfaceHb||0"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.interfaceHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.interfaceTotal||0 }}
                    </div>

                       <span class="card-content-text">环比</span>
                      <arrow-up-outlined  class="card-content-up" v-if="card?.interfaceHb > 0"/>
                      <arrow-down-outlined class="card-content-down" v-if="card?.interfaceHb < 0"/>

                  </template>
                </a-statistic>
                 <a-statistic
                  title="总测试场景数（个）"
                  :value="card?.scenarioHb||0"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.scenarioHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.scenarioTotal ||0}}
                    </div>
                      <span class="card-content-text">环比</span>
                      <arrow-up-outlined  class="card-content-up"  v-if="card?.scenarioHb > 0"/>
                      <arrow-down-outlined class="card-content-down" v-if="card?.scenarioHb < 0" />
                  </template>
                </a-statistic>

              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                  <a-statistic
                  title="接口测试总体覆盖率（%）"
                  :value="card?.coverageHb||0"
                  :precision="2"
                  suffix="%"
                  class="demo-class"
                  :value-style="{ color: card?.coverageHb >= 0?'#cf1322':'#3f8600', fontSize: '18px' }"
                >
                  <template #prefix>
                    <div class="card-content-num">
                      {{ card?.coverage?card.coverage+"%":0+'%' }}
                    </div>
                    <span class="card-content-text">环比</span>
                      <arrow-up-outlined  class="card-content-up"   v-if="card?.coverageHb > 0"/>
                      <arrow-down-outlined class="card-content-down" v-if="card?.coverageHb < 0"  />
                  </template>
                </a-statistic>

              </div>
            </a-card>
            <a-card style="width: 49%; margin-bottom: 16px">
              <div class="card-content">
                <a-statistic
                  title="执行总次数"
                  :value="card?.execTotal||0"

                />
                   <a-statistic
                   :precision="2"
                   suffix="%"
                  title="测试通过率（%）"
                  :value="card?.passRate?card.passRate:0"

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
watch(() => {return props.params;}, async (newVal: any) => {
      console.log("watch staticstic newVal", newVal);
      card.value = newVal.cardData;
    }, {immediate: true});

</script>
<style lang="less" scoped>
.statistics {
  // padding:  16px;
  .card-content {
    display: flex;
    justify-content: space-between;
    :deep(.ant-statistic-content){
      font-size: 32px;
      font-weight: 400;
    }
    &-num {
      font-weight: 400;
      font-size: 32px;

      color: rgba(0, 0, 0, 0.85);
    }
    &-text{
      font-size: 12px;
      color: rgba(0, 0, 0, 0.46);;
    }
    &-up{
      font-size: 12px;
      border-radius:50%;
      background: rgba(255, 242, 240, 0.6);
      padding: 5px;
    }
     &-down{
      font-size: 12px;
      border-radius:50%;
      background: rgba(230, 255, 244, 0.6); padding: 5px;
    }
  }

  .pie {
    background: #fff;
    :deep(.ant-card) {
      height: 100%;
    }
    :deep(.ant-card-body) {
      // height: 38vh !important;
         height: calc(100% - 57px);
      padding: 0 !important;
    }
  }
}
</style>


