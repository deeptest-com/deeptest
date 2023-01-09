<template>
  <div class="report-main">
    <a-card :bordered="false">
      <div>
        <a-page-header
            class="demo-page-header"
            style="border: 1px solid rgb(235, 237, 240)"
            title="测试报告-{{ modelRef.name }}"
            @back="() => $router.go(-1)"
        >
          <a-descriptions size="small" :column="1">
            <a-descriptions-item label="Created" >Lili Qu</a-descriptions-item>
            <a-descriptions-item label="Association">
              <a>421421</a>
            </a-descriptions-item>
            <a-descriptions-item label="Creation Time">2017-01-10</a-descriptions-item>
            <a-descriptions-item label="Effective Time">2017-10-10</a-descriptions-item>
            <a-descriptions-item label="Remarks">
              Gonghu Road, Xihu District, Hangzhou, Zhejiang, China
            </a-descriptions-item>
          </a-descriptions>
        </a-page-header>

            <a-button key="2" type="primary" display="display">全部收起</a-button>
            <a-button key="1" type="primary" display="true">全部展开</a-button>

      </div>
      <div class="scenario-detail-main">
        <a-table :columns="columns" :data-source="data">
          <template #name="{ text }">
            <a>{{ text }}</a>
          </template>
          <template #customTitle>
      <span>
        <smile-outlined />
        Name
      </span>
          </template>
          <template #tags="{ text: tags }">
      <span>
        <a-tag
            v-for="tag in tags"
            :key="tag"
            :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
        >
          {{ tag.toUpperCase() }}
        </a-tag>
      </span>
          </template>
        </a-table>
      </div>




      <div class="scenario-detail-main">
        <a-card :bordered="false" :bodyStyle="{paddingTop: '8px'}">
          <div>
            <Log v-if="modelRef.logs" :logs="modelRef.logs"></Log>
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
import { useI18n } from "vue-i18n";
import {StateType} from "../store";
import {Report} from "@/views/report/data";
const router = useRouter();
const { t } = useI18n();
import Log from "@/views/scenario/exec/components/Log.vue"

const store = useStore<{ Report: StateType }>();
const modelRef = computed<Report>(() => store.state.Report.detailResult);

const get = async (id: number): Promise<void> => {
  await store.dispatch('Report/get', id);
}
const id = ref(+router.currentRoute.value.params.id)
get(id.value)

const columns = [
  {
    dataIndex: 'name',
    key: 'name',
    slots: { title: 'customTitle', customRender: 'name' },
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address',
  },
  {
    title: 'Tags',
    key: 'tags',
    dataIndex: 'tags',
    slots: { customRender: 'tags' },
  },
  {
    title: 'Action',
    key: 'action',
    slots: { customRender: 'action' },
  },
];

const data = [
  {
    key: '1',
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '2',
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
    tags: ['loser'],
  },
  {
    key: '3',
    name: 'Joe Black',
    age: 32,
    address: 'Sidney No. 1 Lake Park',
    tags: ['cool', 'teacher'],
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
