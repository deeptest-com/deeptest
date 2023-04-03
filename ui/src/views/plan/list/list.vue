<template>
  <div class="plan-list-main">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => create()">新建</a-button>
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
                    getList(page, nodeDataCategory.id);
                },
                onShowSizeChange: (page, size) => {
                    pagination.pageSize = size
                    getList(page, nodeDataCategory.id);
                },
            }"
            class="dp-table">

          <template #name="{ text, record  }">
            {{ text }}
            <edit-outlined class="dp-primary" @click="edit(record)"/>
          </template>

          <template #status="{ record }">
            <a-tag v-if="record.disabled" color="green">禁用</a-tag>
            <a-tag v-else color="cyan">启用</a-tag>
          </template>

          <template #action="{ record }">
            <a-button type="link" @click="() => exec(record.id)">执行</a-button>
            <a-button type="link" @click="() => remove(record.id)">删除</a-button>
          </template>

        </a-table>

        <a-empty v-if="list.length === 0" :image="simpleImage" />
      </div>
    </a-card>
  </div>

  <a-drawer
      :closable="true"
      :width="1000"
      :key="currModel.id"
      :visible="createDrawerVisible"
  >
    <template #title>
      <div class="drawer-header">
        <div>新建计划</div>
      </div>
    </template>
    <div class="drawer-content">
      <PlanCreate
          :categoryId="nodeDataCategory.id"
          :onSaved="onSaved">
      </PlanCreate>
    </div>
  </a-drawer>

  <a-drawer
      :closable="true"
      :width="1000"
      :key="currModel.id"
      :visible="editDrawerVisible"
      @close="onEditFinish"
  >
    <template #title>
      <div class="drawer-header">
        <div>计划编辑</div>
      </div>
    </template>
    <div class="drawer-content">
      <PlanEdit
          :modelId="currModel.id"
          :categoryId="nodeDataCategory.id"
          :onFieldSaved="onFieldSaved">
      </PlanEdit>
    </div>
  </a-drawer>

</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref, UnwrapRef, watch} from "vue";
import { Empty } from 'ant-design-vue';
import {SelectTypes} from 'ant-design-vue/es/select';
import {PaginationConfig, QueryParams, Plan} from '../data.d';
import {useStore} from "vuex";

import debounce from "lodash.debounce";
import {StateType} from "../store";
import {useRouter} from "vue-router";
import {message, Modal, notification} from "ant-design-vue";
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import {StateType as ProjectStateType} from "@/store/project";

import PlanCreate from "../edit/create.vue";
import PlanEdit from "../edit/edit.vue";

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
const store = useStore<{ Plan: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const nodeDataCategory = computed<any>(()=> store.state.Plan.nodeDataCategory);

const list = computed<Plan[]>(() => store.state.Plan.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Plan.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});

const currModel = ref({})

watch(nodeDataCategory, () => {
  console.log('watch nodeDataCategory', nodeDataCategory.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

watch(currProject, () => {
  console.log('watch currProject', currProject.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

const loading = ref<boolean>(true);

const getList = debounce(async (current: number, categoryId: number): Promise<void> => {
  console.log('getList')

  loading.value = true;

  await store.dispatch('Plan/listPlan', {
    categoryId,
    keywords: queryParams.keywords,
    enabled: queryParams.enabled,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false
}, 600)

const exec = (id: number) => {
  console.log('exec')
  router.push(`/plan/exec/${id}`)
}

interface FormState {
  name: string;
  description: string;
  serveId?:string,

}

const createDrawerVisible = ref(false);
const create = () => {
  console.log('create')
  createDrawerVisible.value = true;
}
const onSaved = () => {
  console.log('onSaved')
  getList(pagination.value.current, nodeDataCategory.value.id)
  createDrawerVisible.value = false
}

const editDrawerVisible = ref(false);
const editFormState: UnwrapRef<FormState> = reactive({
  name: '',
  description: '',
  serveId:'',
});

const edit = (record) => {
  console.log('edit')
  currModel.value = record

  editDrawerVisible.value = true;
  editFormState.name = record.name;
  editFormState.description = record.description;
  editFormState.serveId = record.id;
}

const onFieldSaved = () => {
  console.log('onFieldSaved')
  getList(pagination.value.current, nodeDataCategory.value.id)
}
const onEditFinish = () => {
  console.log('onEditFinish')
  getList(pagination.value.current, nodeDataCategory.value.id)
  editDrawerVisible.value = false
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除计划',
    content: '确定删除指定的计划？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Plan/removePlan', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          getList(1, nodeDataCategory.value.id)

          notification.success({
            message: `删除成功`,
          });
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
  getList(1, nodeDataCategory.value.id)
}, 500);

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
})

</script>

<style lang="less" scoped>
.plan-list-main {

}
</style>