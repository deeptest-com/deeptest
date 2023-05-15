<template>
  <TableFilter :show-operation="false" @handle-filter="handleFilter"/>
  <a-table
      row-key="id"
      :loading="loading"
      :columns="columns"
      :data-source="scenariosReports"
      :pagination="{
            ...pagination,
            onChange: handlePageChanged
        }"
  >
    <template #createdAt="{ record }">
      <span>{{ momentUtc(record.createdAt) }}</span>
    </template>

  </a-table>
</template>
<script lang="ts" setup>
import {computed, reactive, ref, onMounted, watch} from 'vue';
import TableFilter from "@/views/component/Report/List/TableFilter.vue";
import {useStore} from "vuex";
import {momentUtc} from "@/utils/datetime";
import {Scenario} from "@/views/scenario/data";

const store = useStore<{ Scenario, ProjectGlobal, ServeGlobal }>();
const detailResult: any = computed<Scenario>(() => store.state.Scenario.detailResult);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const scenariosReports = computed(() => store.state.Scenario.scenariosReports);
const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
    slots: {customRender: 'serialNumber'},
    width: 120
  },
  {
    title: '测试通过率',
    dataIndex: 'interfacePassRate',
    width: 120,
  },
  {
    title: '执行耗时',
    dataIndex: 'duration',
    width: 80,
    slots: {customRender: 'duration'},
  },
  {
    title: '执行人',
    dataIndex: 'createUserName',
    width: 80,
  },
  {
    title: '执行时间',
    dataIndex: 'createdAt',
    width: 200,
    slots: {customRender: 'createdAt'},
  },
];
let formState = reactive({});
let pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

function handleFilter(params) {
  formState = params;
  refreshList({...formState});
}

function handlePageChanged(page) {
  pagination.current = page;
  refreshList({page, pageSize: pagination.pageSize});
}

const loading = ref(false);

async function refreshList(params: any) {
  loading.value = true;
  await store.dispatch('Scenario/getExecResultList', {
    data: {
      ...formState,
      ...params,
      scenarioId: detailResult.value.id,
    }
  });
  loading.value = false;
}

watch(() => {
  return detailResult.value.id
}, async (newVal) => {
  await refreshList({page: 1, pageSize: 10});
}, {
  immediate: true
})

</script>

