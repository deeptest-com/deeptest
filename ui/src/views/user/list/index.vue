<template>
  <div class="user-main-list">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="成员">
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
      </a-tab-pane>
      <a-tab-pane key="2" tab="审批" force-render>
        <a-card :bordered="false">
          <a-table
            row-key="id"
            :columns="auditColumns"
            :data-source="auditLst.list"
            :loading="loading"
            :pagination="{
              ...auditLst.pagination,
              onChange: (page) => {
                getAudits(page);
              },
              onShowSizeChange: (page, size) => {
                pagination1.pageSize = size;
                getAudits(page);
              },
            }"
          >
            <template #status="{ text }">
              {{ text == 0 ? "待审批" : text == 1 ? "已同意" : "已拒绝" }}
            </template>
            <template #action="{ record }">
              <a-button
                type="link"
                style="padding: 0"
                @click="() => audit(record.id)"
                >审批</a-button
              >
            </template>
          </a-table>
        </a-card></a-tab-pane
      >
    </a-tabs>
  </div>

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
  <a-modal
    v-model:visible="auditModal"
    title="审批"
    @cancel="auditModal = false"
  >
    <p>是否通过该用户的审批？</p>

    <template #footer>
      <a-button key="back" danger @click="handleAudit(2)">拒绝</a-button>
      <a-button
        key="submit"
        type="primary"
        :loading="loading"
        @click="handleAudit(1)"
        >同意</a-button
      >
    </template>
  </a-modal>
</template>
<script setup lang="ts">
import { PaginationConfig, QueryParams, User } from "../data.d";

import { computed, onMounted, reactive, ref } from "vue";
import { useStore } from "vuex";
import { StateType } from "@/views/user/store";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import debounce from "lodash.debounce";
import { useRouter } from "vue-router";
import { getAuditList, doAudit } from "@/views/project/service";
import EditPage from "../edit/edit.vue";

const router = useRouter();
const store = useStore<{ UserInternal: StateType }>();

const list = computed<User[]>(() => store.state.UserInternal.queryResult.list);
let activeKey = ref("1");
let pagination = computed<PaginationConfig>(
  () => store.state.UserInternal.queryResult.pagination
);
let queryParams = reactive<QueryParams>({
  username: "",
  page: pagination.value.current,
  pageSize: pagination.value.pageSize,
});
const auditLst = ref({});
const auditModal = ref(false);
const auditId = ref(0);
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
const auditColumns = [
  {
    title: "申请人",
    dataIndex: "applyUserName",
    //  width: 150,
    // slots: { customRender: "username" },
  },
  {
    title: "申请加入项目",
    dataIndex: "projectName",
    //  width: 200,
    // slots: { customRender: "role" },
  },
  {
    title: "申请角色",
    dataIndex: "projectRoleName",
    //  width: 150,
    // slots: { customRender: "role" },
  },
  {
    title: "申请原因",
    dataIndex: "description",
    // width: 200,
    // slots: { customRender: "email" },
  },
  {
    title: "申请日期",
    dataIndex: "createdAt",
      width: 240,
    // slots: { customRender: "email" },
  },
  {
    title: "状态",
    dataIndex: "status",
     width: 100,
    slots: { customRender: "status" },
  },
  {
    title: "操作",
    key: "action",
    width: 50,
    slots: { customRender: "action" },
  },
];
onMounted(() => {
  getList(1);
  getRoles();
  getAudits(1);
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
const getAudits = (page: number) => {
  loading.value = true;

  getAuditList({
    pageSize: pagination1.value.pageSize,
    page: page,
  })
    .then((json) => {
      console.log("审批列表", json);
      if (json.code === 0) {
        // auditLst.value = json.data.result;
        auditLst.value = {
          current: 1,
          list: json.data.result || [],
          pagination: {
            ...pagination1.value,
            current: page,
            total: json.data.total || 0,
          },
        };
      }
    })
    .finally(() => {
      loading.value = false;
    });
};
const audit = (id: number) => {
  console.log("remove");
  auditModal.value = true;
  auditId.value = id;
};
const handleAudit = async (type: number) => {
  await doAudit({ id: auditId.value, status: type }).then((json) => {
    if (json.code === 0) {
      getAudits(1);
      notification.success({
        key: NotificationKeyCommon,
        message: `审批成功`,
      });
    } else {
      notification.error({
        key: NotificationKeyCommon,
        message: `审批失败`,
      });
    }
  });

  auditModal.value = false;
};
const visible = ref(false);
const currentUserId = ref(0);
const edit = (id: number) => {
  currentUserId.value = id;
  visible.value = true;
};

const handleOk = (e: MouseEvent) => {
  visible.value = false;
};

const closeModal = () => {
  visible.value = false;
};

const onSearch = debounce(() => {
  getList(1);
}, 500);

//角色列表
const getRoles = () => {
  store.dispatch("Project/getRoles");
  return;
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
<style lang="less" scoped>
.user-main-list {
  background: #fff;
}
</style>
