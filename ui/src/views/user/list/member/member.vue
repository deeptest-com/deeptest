<template>
    <a-card :bordered="false">
    <template #title>
      <a-button type="primary" @click="() => edit(0)">新建用户</a-button>
    </template>
    <template #extra>
      <a-input-search
        @change="onSearch"
        @search="onSearch"
        v-model:value="queryParams.username"
        placeholder="输入用户名搜索"
        style="width: 270px; margin-left: 16px"
      />
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
            pagination.pageSize = size;
            getList(page);
          },
        }"
      >
        <template #action="{ record }">
          <a-button type="link" @click="() => edit(record.id)">
            <span>编辑</span>
          </a-button>
          <a-button type="link" @click="() => remove(record.id)">
            <span>删除</span>
          </a-button>
        </template>
      </a-table>
    </div>
  </a-card>
  <a-modal
    v-model:visible="visible"
    @ok="handleOk"
    width="700px"
    :footer="null"
  >
    <EditPage
      :currentUserId="currentUserId"
      :getList="getList"
      :closeModal="closeModal"
    />
  </a-modal>
</template>

<script setup lang="ts">
import { PaginationConfig, QueryParams, User } from "../../data.d";
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useStore } from "vuex";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import debounce from "lodash.debounce";
import { useRouter } from "vue-router";
import EditPage from "../../edit/edit.vue";

const router = useRouter();
const store = useStore();

const list = computed<User[]>(() => store.state.UserInternal.queryResult.list);
const roles = ()=>{
  let rolesList = {}
  store.state.Project.roles.forEach((item:any)=>{
    rolesList[item.name] = item.displayName
  })
  return rolesList
}
let activeKey = ref("1");
let pagination = computed<PaginationConfig>(
  () => store.state.UserInternal.queryResult.pagination
);
let queryParams = reactive<QueryParams>({
  username: "",
  page: pagination.value.current,
  pageSize: pagination.value.pageSize,
});

const getRoleName = (val:any)=>{
  let rolesList = roles()
  return rolesList[val.text]
}

const columns = [
  {
    title: "序号",
    dataIndex: "index",
    align: "center",
    customRender: ({ text, index }: { text: any; index: number }) =>
      (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: "用户名",
    dataIndex: "username",
    align: "center",
  },
  {
    title: "姓名",
    dataIndex: "name",
    align: "center",
  },
  {
    title: "邮箱",
    dataIndex: "email",
    align: "center",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    slots: { customRender: "action" },
  },
];

onMounted(() => {
  getList(1);
  getRoles();
  // getAudits(1);
});

const pagination1 = ref({
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
});
const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;

  await store.dispatch("UserInternal/queryUser", {
    username: queryParams.username,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
};

const visible = ref(false);
const currentUserId = ref(0);
const edit = (id: number) => {
  currentUserId.value = id;
  visible.value = true;
};


const onSearch = debounce(() => {
  getList(1);
}, 500);

//角色列表
const getRoles = () => {
  store.dispatch("Project/getRoles");
  return;
};

const closeModal = () => {
  visible.value = false;
  getList(1);
};

const remove = (id: number) => {
  console.log("remove");

  Modal.confirm({
    title: "删除用户",
    content: "确定删除指定的用户？",
    okText: "确认",
    cancelText: "取消",
    onOk: async () => {
      store.dispatch("UserInternal/removeUser", id).then((res) => {
        console.log("res", res);
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
      });
    },
  });
};
</script>