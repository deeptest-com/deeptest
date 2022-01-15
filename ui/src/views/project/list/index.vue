<template>
  <div class="indexlayout-main-conent">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => setCreateFormVisible(true)">新建脚本</a-button>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr">
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
        >
          <template #name="{ text, record  }">
            <a :href="record.href" target="_blank">{{ text }}</a>
          </template>
          <template #status="{ record }">
            <a-tag v-if="record.disabled == 0" color="green">启用</a-tag>
            <a-tag v-else color="cyan">禁用</a-tag>
          </template>
          <template #action="{ record }">
            <a-button type="link" @click="() => editProject(record.id)">编辑
            </a-button>
            <a-button type="link" @click="() => deleteProject(record.id)">删除
            </a-button>
          </template>

        </a-table>
      </div>
    </a-card>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, reactive, Ref, ref} from "vue";
import {SelectTypes} from 'ant-design-vue/es/select';
import {PaginationConfig, QueryParams, Project} from '../data.d';
import {useStore} from "vuex";

import {StateType as ListStateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";

interface ListProjectPageSetupData {
  statusArr,
  queryParams,
  columns: any;
  list: ComputedRef<Project[]>;
  pagination: ComputedRef<PaginationConfig>;
  loading: Ref<boolean>;
  getList: (current: number) => Promise<void>;

  editProject: (id: number) => Promise<void>;
  deleteProject: (id: number) => void;

  onSearch: () => void;
}

export default defineComponent({
  name: 'ProjectListPage',
  components: {
  },
  setup(): ListProjectPageSetupData {
    const statusArr = ref<SelectTypes['options']>([
      {
        label: '所有状态',
        value: '1',
      },
      {
        label: '启用',
        value: '0',
      },
      {
        label: '禁用',
        value: '0',
      },
    ]);

    const router = useRouter();
    const store = useStore<{ ListProject: ListStateType }>();

    const list = computed<Project[]>(() => store.state.ListProject.queryResult.list);
    let pagination = computed<PaginationConfig>(() => store.state.ListProject.queryResult.pagination);
    let queryParams = reactive<QueryParams>({
      keywords: '', enabled: '1',
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
        width: 260,
        slots: {customRender: 'action'},
      },
    ];

    onMounted(() => {
      getList(1);
    })

    const loading = ref<boolean>(true);
    const getList = async (current: number): Promise<void> => {
      loading.value = true;

      await store.dispatch('project/queryProject', {
        keywords: queryParams.keywords,
        enabled: queryParams.enabled,
        pageSize: pagination.value.pageSize,
        page: current,
      });
      loading.value = false;
    }

    const editProject = async (id: number) => {
      console.log('editProject')
    }

    const deleteProject = (id: number) => {
      console.log('editProject')
    }

    const onSearch = debounce(() => {
      getList(1)
    }, 500);

    onMounted(() => {
      getList(1);
    })

    return {
      statusArr,
      queryParams,
      columns,
      list,
      pagination,
      loading,
      getList,

      editProject,
      deleteProject,

      onSearch,
    }
  }

})
</script>