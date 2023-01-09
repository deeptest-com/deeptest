<template>
  <div>
    <a-card :bordered="false">
      <div >
        <a-card :bordered="false" :bodyStyle="{paddingTop: '8px'}">
        <a-row>
          <a-col>测试报告--{{report.name}}</a-col>
          <a-col>
          <a-descriptions size="small" :column="4">
            <a-descriptions-item label="Count" color="green" > {{report.totalRequestNum}}</a-descriptions-item>
            <a-descriptions-item label="Passed" color="green"> {{report.passRequestNum}}</a-descriptions-item>
            <a-descriptions-item label="Failed" color="red">{{failCount}}</a-descriptions-item>
            <a-descriptions-item label="Excute Time">{{}}</a-descriptions-item>
          </a-descriptions>
          </a-col>
        </a-row>
            <!-- <a-button key="2" type="primary" display="display" style="float: right">全部收起</a-button>
            <a-button key="1" type="primary" display="true" style="float: right">全部展开</a-button> -->
          </a-card>
      </div>

      <a-card :bordered="false" :bodyStyle="{paddingTop: '8px'}">
      <div>
        <a-table :columns="columns" bordered :data-source="data" :pagination=false :rowKey="(record: any) => record.id">
        </a-table>
      </div>
    </a-card>
      <div class="scenario-detail-main" >
        <a-card :bordered="false" :bodyStyle="{paddingTop: '8px'}">
          <div>
            <Log v-if="report.logs" :logs="report.logs"></Log>
          </div>
        </a-card>
      </div>
    </a-card>
  </div>
</template>



<script setup lang="ts">
import {defineComponent, computed, ref, reactive, ComputedRef} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {BackTop} from "ant-design-vue";
import { useI18n } from "vue-i18n";
import {StateType} from "../store";
import {Report} from "@/views/reportNew/data";
import Log from "../Log.vue"
import { getAdditionalMeta } from "@intlify/core-base";
const router = useRouter();
const { t } = useI18n();


const store = useStore<{ Report: StateType }>();
//千万注意async/await的顺序问题
const report = computed<Report>(() => store.state.Report.detailResult);
const id = ref(+router.currentRoute.value.params.id)
const get = async (id: number): Promise<void> => {
  await store.dispatch('Report/get', id);
}
get(id.value)


const passCount = Number(store.state.Report.detailResult.passRequestNum)
const totalRequestNum = Number(store.state.Report.detailResult.totalRequestNum)
const failCount = totalRequestNum-passCount
const data = ref(store.state.Report.detailResult.logs[0].logs)

const columns = [
  {
    title: '序号',
    dataIndex: 'id',
    key:'rowIndex',
    width:80,
    align:"center",
    customRender: ({ index }: { index: number }) => {
            return `${index + 1}`
          }
  },
  {
    title: '接口名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '运行时间',
    dataIndex: 'runtime',
    width:120,
    align:"center",
    key: 'runtime',
  },
  {
    title: '执行结果',
    width:140,
    align:"center",
    key: 'resultStatus',
    dataIndex: 'resultStatus',
  },
];

</script>

<style lang="less" scoped>
.scenario-detail-main {
.ant-card-body {
  padding-top: 0 !important;
}
}
</style>
