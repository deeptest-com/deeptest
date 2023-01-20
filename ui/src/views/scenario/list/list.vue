<template>
  <div class="scenario-list-main">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => edit(0)">新建</a-button>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr" class="status-select">
        </a-select>
        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                        placeholder="输入关键字搜索" style="width:270px;margin-left: 16px;"/>
      </template>

      <div>
        <a-table
            v-if="list.length > 0"
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
          <template #status="{ record }">
            <a-tag v-if="record.disabled" color="green">禁用</a-tag>
            <a-tag v-else color="cyan">启用</a-tag>
          </template>

          <template #action="{ record }">
            <a-button type="link" @click="() => exec(record.id)">执行</a-button>

            <a-button type="link" @click="() => design(record.id)">设计</a-button>
            <a-button type="link" @click="() => edit(record.id)">编辑</a-button>
            <a-button type="link" @click="() => remove(record.id)">删除</a-button>
          </template>

        </a-table>

        <a-empty v-if="list.length === 0" :image="simpleImage" />
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref, watch} from "vue";
import { Empty } from 'ant-design-vue';
import {SelectTypes} from 'ant-design-vue/es/select';
import {PaginationConfig, QueryParams, Scenario} from '../data.d';
import {useStore} from "vuex";

import {StateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal, notification} from "ant-design-vue";
import {StateType as ProjectStateType} from "@/store/project";

const simpleImage = Empty.PRESENTED_IMAGE_SIMPLE

const statusArr = ref<SelectTypes['options']>([
  {
    label: '所有状态',
    value: '',
  },
  {
    label: '启用',
    value: '1',
  },
  {
    label: '禁用',
    value: '0',
  },
]);

const router = useRouter();
const store = useStore<{ Scenario: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const nodeDataCategory = computed<any>(()=> store.state.Scenario.nodeDataCategory);

const list = computed<Scenario[]>(() => store.state.Scenario.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Scenario.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});

watch(nodeDataCategory, () => {
  console.log('watch nodeDataCategory', nodeDataCategory.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

watch(currProject, () => {
  console.log('watch currProject', currProject.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

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
  getList(1, nodeDataCategory.value.id);
})

const loading = ref<boolean>(true);
const getList = async (current: number, categoryId: number): Promise<void> => {
  loading.value = true;

  await store.dispatch('Scenario/listScenario', {
    categoryId,
    keywords: queryParams.keywords,
    enabled: queryParams.enabled,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
}

const exec = (id: number) => {
  console.log('exec')
  router.push(`/scenario/exec/${id}`)
}

const design = (id: number) => {
  console.log('edit')
  router.push(`/scenario/design/${id}`)
}

const edit = (id: number) => {
  console.log('edit')
  router.push(`/scenario/edit/${id}`)
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除场景',
    content: '确定删除指定的场景？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Scenario/removeScenario', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          notification.success({
            message: `删除成功`,
          });
          store.dispatch('Scenario/queryScenario', id)
        } else {
          notification.error({
            message: `删除失败`,
          });
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
.scenario-list-main {

}
</style>