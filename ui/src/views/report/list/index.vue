<template>
  <div class="report-main">
    <a-card :bordered="false">
      <template #title>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.scenarioId" class="status-select" >
          <a-select-option value=""></a-select-option>
          <a-select-option v-for="item in scenarios" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
        </a-select>

        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                        placeholder="输入关键字搜索" style="width:270px;margin-left: 16px;"/>
      </template>

      <div>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {SelectTypes} from 'ant-design-vue/es/select';
import {useStore} from "vuex";

import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal} from "ant-design-vue";
import {StateType} from "@/views/report/store";
import {PaginationConfig, QueryParams, Report} from "@/views/report/data";

const scenarios = ref<any[]>([]);

const router = useRouter();
const store = useStore<{ Report: StateType }>();

const list = computed<Report[]>(() => store.state.Report.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Report.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', scenarioId: '',
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
    slots: {customRender: 'name'},
  },
  {
    title: '描述',
    dataIndex: 'desc',
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'status'},
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

const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;

  await store.dispatch('Report/listReport', {
    keywords: queryParams.keywords,
    scenarioId: queryParams.scenarioId,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
}

const view = (id: number) => {
  console.log('view')
  router.push(`/report/detail/${id}`)
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除项目',
    content: '确定删除？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Report/removeReport', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          message.success(`删除项目成功`)
          store.dispatch('Report/queryReport', id)
        } else {
          message.error(`删除项目失败`)
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

}
</style>