<template>
  <div class="project-members">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => invite(0)">邀请</a-button>
      </template>
      <template #extra>
        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                        placeholder="输入关键字搜索" style="width:270px;margin-left: 16px;"/>
      </template>

      <div>
        <a-table
            row-key="id"
            :columns="columns"
            :data-source="members.list"
            :loading="loading"
            :pagination="{
                ...pagination,
                onChange: (page) => {
                    getMembers(page);
                },
                onShowSizeChange: (page, size) => {
                    pagination.pageSize = size
                    getMembers(page);
                },
            }"
        >
          <template #username="{ text  }">
            {{ text }}
          </template>

          <template #email="{ text  }">
            {{ text }}
          </template>

          <template #action="{ record }">
            <a-button type="link" @click="() => remove(record.id)"
                      :disabled="currentUser.projectRoles[projectId] !== 'admin'">移除</a-button>
          </template>

        </a-table>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, Ref, ref} from "vue";
import {PaginationConfig, QueryParams, Project} from '../data.d';
import {useStore} from "vuex";

import {StateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {Modal, notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import {query, queryMembers, removeMember} from "@/views/project/service";
import {StateType as UserStateType} from "@/store/user";

const router = useRouter();
const store = useStore<{ Project: StateType, User: UserStateType }>();
const currentUser = computed<any>(()=> store.state.User.currentUser);

let pagination = computed<PaginationConfig>(() => store.state.Project.queryResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',
  page: pagination.value.current,
  pageSize: pagination.value.pageSize
});

const members = ref({});

const projectId = +router.currentRoute.value.params.id

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
    slots: {customRender: 'username'},
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    slots: {customRender: 'email'},
  },
  {
    title: '操作',
    key: 'action',
    width: 260,
    slots: {customRender: 'action'},
  },
];

const invite = () => {
  console.log('invite')
  router.push(`/project/invite/${projectId}`)
}

onMounted(() => {
  console.log('onMounted')
  getMembers(1);
})

const initState: StateType = {
  queryResult: {
    list: [],
    pagination: {
      total: 0,
      current: 1,
      pageSize: 10,
      showSizeChanger: true,
      showQuickJumper: true,
    },
  },
  detailResult: {} as Project,
  queryParams: {},
};

const loading = ref<boolean>(true);
const getMembers = (page: number) => {
  loading.value = true;

  queryMembers({
    id: projectId,
    keywords: queryParams.keywords,
    pageSize: pagination.value.pageSize,
    page: page,
  }).then((json) => {
    if (json.code === 0) {
      members.value = {
        ...initState.queryResult,
        list: json.data.result || [],
        pagination: {
          ...initState.queryResult.pagination,
          current: page,
          pageSize: pagination.value.pageSize,
          total: json.data.total || 0,
        },
      };
    }
  }).finally(() => {
    loading.value = false;
  })
}

const remove = (userId: number) => {
  console.log('remove')

  Modal.confirm({
    title: '移除成员',
    content: '确定移除指定的项目成员？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      removeMember(userId, projectId).then((json) => {
        if (json.code === 0) {
          getMembers(queryParams.page)

          notification.success({
            key: NotificationKeyCommon,
            message: `移除成功`,
          });
        } else {
          notification.error({
            key: NotificationKeyCommon,
            message: `移除失败`,
          });
        }
      })
    }
  });
}

const onSearch = debounce(() => {
  getMembers(1)
}, 500);

</script>

<style lang="less" scoped>
.project-members {

}
</style>