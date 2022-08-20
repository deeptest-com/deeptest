<template>
  <div class="report-main">
    <a-card :bordered="false">
      <template #title>
        <div>测试报告</div>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.scenarioId" :dropdownMatchSelectWidth="false" class="scenario-select" >
          <a-select-option value="0">请选择场景</a-select-option>
          <a-select-option v-for="item in scenarios" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
        </a-select>

        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                        placeholder="输入关键字搜索" style="width:270px;margin-left: 16px;"/>
      </template>

      <div>
        <a-table
            row-key="id"
            :columns="columns"
            :data-source="list"
            :loading="loading"
            :pagination="{
                ...pagination,
                onChange: (page) => {
                    getList(page);
                },
                onShowSizeChange: (page, size) => {
                    pagination.pageSize = size
                    getList(page);
                },
            }"
            class="dp-table"
        >
          <template #name="{ text  }">
            {{ text }}
          </template>

          <template #execTime="{ record }">
            <span>{{ momentUtcDef(record.createdAt) }}</span>
          </template>

          <template #action="{ record }">
            <a-button type="link" @click="() => view(record.id)">查看</a-button>
            <a-button type="link" @click="() => remove(record.id)">删除</a-button>
          </template>

        </a-table>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {SelectTypes} from 'ant-design-vue/es/select';
import {useStore} from "vuex";

import debounce from "lodash.debounce";
import { momentUtcDef } from "@/utils/datetime";
import {useRouter} from "vue-router";
import {message, Modal} from "ant-design-vue";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
import {StateType} from "@/views/report/store";
import {PaginationConfig, QueryParams, Report} from "@/views/report/data";
import {Scenario} from "@/views/scenario/data";

const router = useRouter();
const store = useStore<{ Report: StateType, Scenario: ScenarioStateType }>();

const scenarios = computed<Scenario[]>(() => store.state.Scenario.listResult.list);
const list = computed<Report[]>(() => store.state.Report.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', scenarioId: '0',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});

const columns = [
  {
    title: '序号',
    dataIndex: 'index',
    width: 80,
    customRender: ({
                     text,
                     index
                   }: { text: any; index: number }) => (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: '名称',
    dataIndex: 'name',
    width: 300,
    slots: {customRender: 'name'},
  },
  {
    title: '执行时间',
    dataIndex: 'execTime',
    width: 200,
    slots: {customRender: 'execTime'},
  },
  {
    title: '操作',
    key: 'action',
    width: 200,
    slots: {customRender: 'action'},
  },
];

onMounted(() => {
  console.log('onMounted')
  getList(1);
})

store.dispatch('Scenario/listScenario', {});

const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/list', {
    keywords: queryParams.keywords,
    scenarioId: queryParams.scenarioId,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
}


const view = (id: number) => {
  console.log('view')
  router.push(`/report/${id}`)
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除报告',
    content: '确定删除指定的报告？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Report/remove', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          message.success(`删除报告成功`)
        } else {
          message.error(`删除报告失败`)
        }
      })
    }
  });
}

const onSearch = debounce(() => {
  getList(1)
}, 500);

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
.report-main {
  .scenario-select {

  }
}
</style>