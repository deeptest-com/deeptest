<template>
  <div class="project-main-list">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => edit(0)">新建项目</a-button>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr" class="status-select">
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
          <template #name="{ text  }">
            {{ text }}
          </template>
          <template #status="{ record }">
            <a-tag v-if="record.disabled == 0" color="green">启用</a-tag>
            <a-tag v-else color="cyan">禁用</a-tag>
          </template>
          <template #action="{ record }">
            <a-button type="link" @click="() => edit(record.id)">编辑
            </a-button>
            <a-button type="link" @click="() => remove(record.id)">删除
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

import {StateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal, notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";

export default defineComponent({
  name: 'ProjectListPage',
  components: {
  },
  setup() {
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
    const store = useStore<{ Project: StateType }>();

    const list = computed<Project[]>(() => store.state.Project.queryResult.list);
    let pagination = computed<PaginationConfig>(() => store.state.Project.queryResult.pagination);
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
      console.log('onMounted')
      getList(1);
    })

    const loading = ref<boolean>(true);
    const getList = async (current: number): Promise<void> => {
      loading.value = true;

      await store.dispatch('Project/queryProject', {
        keywords: queryParams.keywords,
        enabled: queryParams.enabled,
        pageSize: pagination.value.pageSize,
        page: current,
      });
      loading.value = false;
    }

    const edit = async (id: number) => {
      console.log('edit')
      router.push(`/project/edit/${id}`)
    }

    const remove = (id: number) => {
      console.log('remove')

      Modal.confirm({
        title: '删除项目',
        content: '确定删除指定的项目？',
        okText: '确认',
        cancelText: '取消',
        onOk: async () => {
          store.dispatch('Project/removeProject', id).then((res) => {
            console.log('res', res)
            if (res === true) {
              notification.success({
                key: NotificationKeyCommon,
                message: `删除成功`,
              });
            } else {
              notification.error({
                key: NotificationKeyCommon,
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

      edit,
      remove,

      onSearch,
    }
  }

})
</script>

<style lang="less" scoped>
.project-main-list {
  .status-select {
    width: 100px;
  }
}
</style>