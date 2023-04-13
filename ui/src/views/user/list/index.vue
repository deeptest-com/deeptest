<template>
  <div class="user-main-list">
    <a-card :bordered="false">
      <template #extra>
        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.username"
                        placeholder="输入用户名搜索" style="width:270px;margin-left: 16px;"/>
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
          <template #action="{ record }">
            <a-button type="link" @click="() => remove(record.id)">
              <span>删除</span>
            </a-button>
          </template>
        </a-table>
      </div>
    </a-card>
  </div>
</template>
<script setup lang="ts">
import {PaginationConfig, QueryParams, User} from '../data.d';

import {computed, onMounted, reactive, ref} from "vue";
import {useStore} from "vuex";
import {StateType} from "@/views/user/store";
import {Modal, notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";

const router = useRouter();
const store = useStore<{ UserInternal: StateType }>();

const list = computed<User[]>(() => store.state.UserInternal.queryResult.list);

let pagination = computed<PaginationConfig>(() => store.state.UserInternal.queryResult.pagination);
let queryParams = reactive<QueryParams>({
  username: '',
  page: pagination.value.current,
  pageSize: pagination.value.pageSize
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
    title: '用户名',
    dataIndex: 'username',
  },
  {
    title: '姓名',
    dataIndex: 'name',
  },
  {
    title: '邮箱',
    dataIndex: 'email',
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

  await store.dispatch('UserInternal/queryUser', {
    username: queryParams.username,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
}

const onSearch = debounce(() => {
  getList(1)
}, 500);

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除用户',
    content: '确定删除指定的用户？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('UserInternal/removeUser', id).then((res) => {
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
</script>
<style lang="less" scoped>
</style>
